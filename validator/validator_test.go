package main

import (
	"encoding/json"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

const (
	maxDeploymentIDLength = 100
	maxExpiryMonths       = 6
	minBatchSize          = 2
	maxBatchSize          = 20
)

// PVTIssuer represents the structure of each entry in pvt-issuers.json
type PVTIssuer struct {
	Name         string `json:"name"`
	Contact      string `json:"contact"`
	Endpoint     string `json:"endpoint"`
	DeploymentID string `json:"deploymentID"`
	BatchSize    int    `json:"batchSize"`
	Expiry       string `json:"expiry"`
}

// assertSecureURL verifies that a URL conforms to security requirements.
func assertSecureURL(t *testing.T, rawURL, label string) *url.URL {
	t.Helper()
	parsed, err := url.Parse(rawURL)
	if err != nil {
		t.Fatalf("%s URL %q is invalid: %v", label, rawURL, err)
	}
	if parsed.Scheme != "https" || parsed.Host == "" {
		t.Errorf("%s URL %q must be a valid HTTPS URL", label, rawURL)
	}
	if parsed.Opaque != "" {
		t.Errorf("%s URL %q must not be opaque", label, rawURL)
	}
	if parsed.OmitHost {
		t.Errorf("%s URL %q must not omit host", label, rawURL)
	}
	if parsed.User != nil {
		t.Errorf("%s URL %q must not contain user info", label, rawURL)
	}
	if parsed.RawQuery != "" || parsed.ForceQuery {
		t.Errorf("%s URL %q must not contain queries", label, rawURL)
	}
	if parsed.RawFragment != "" {
		t.Errorf("%s URL %q must not contain fragments", label, rawURL)
	}
	return parsed
}

func TestValidatePVTIssuers(t *testing.T) {
	data, err := os.ReadFile("../pvt-issuers.json")
	if err != nil {
		t.Fatalf("Failed to read pvt-issuers.json: %v", err)
	}

	var issuers map[string]PVTIssuer
	if err := json.Unmarshal(data, &issuers); err != nil {
		t.Fatalf("pvt-issuers.json is not valid JSON: %v", err)
	}

	for origin, issuer := range issuers {
		t.Run(origin, func(t *testing.T) {
			// 1. Validate Origin URL
			parsedOrigin := assertSecureURL(t, origin, "Origin")

			// 2. Validate Name
			if strings.TrimSpace(issuer.Name) == "" {
				t.Error("Name cannot be empty")
			}

			// 3. Validate Contact
			if !strings.Contains(issuer.Contact, "@") {
				t.Errorf("Contact %q must be a valid email address", issuer.Contact)
			}

			// 4. Validate Endpoint URL
			parsedEndpoint := assertSecureURL(t, issuer.Endpoint, "Endpoint")

			// 5. Compare Origin and Endpoint
			if parsedOrigin != nil && parsedEndpoint != nil {
				if parsedEndpoint.Scheme != parsedOrigin.Scheme {
					t.Errorf("Endpoint scheme %q does not match origin scheme %q", parsedEndpoint.Scheme, parsedOrigin.Scheme)
				}
				if parsedEndpoint.Host != parsedOrigin.Host {
					t.Errorf("Endpoint host %q does not match origin host %q", parsedEndpoint.Host, parsedOrigin.Host)
				}
			}

			// 6. Validate DeploymentID
			if strings.TrimSpace(issuer.DeploymentID) == "" {
				t.Error("deplotmentID cannot be empty")
			} else if len(issuer.DeploymentID) > maxDeploymentIDLength {
				t.Errorf("deplotmentID %q must be at most %d characters, got %d", issuer.DeploymentID, maxDeploymentIDLength, len(issuer.DeploymentID))
			}

			// 7. Validate BatchSize
			if issuer.BatchSize < minBatchSize || issuer.BatchSize > maxBatchSize {
				t.Errorf("batchSize %d must be between %d and %d", issuer.BatchSize, minBatchSize, maxBatchSize)
			}

			// 8. Validate Expiry
			expiryUnix, err := strconv.ParseInt(issuer.Expiry, 10, 64)
			if err != nil {
				t.Errorf("Expiry %q is not a valid Unix timestamp", issuer.Expiry)
			} else {
				expiryTime := time.Unix(expiryUnix, 0)
				now := time.Now()
				if expiryTime.Before(now) {
					t.Errorf("Issuer config has expired (Expiry: %s)", expiryTime.Format(time.RFC3339))
				}
				maxExpiry := now.AddDate(0, maxExpiryMonths, 0)
				if expiryTime.After(maxExpiry) {
					t.Errorf("Expiry %s is too far in the future (max %d months from now)", expiryTime.Format(time.RFC3339), maxExpiryMonths)
				}
			}
		})
	}
}
