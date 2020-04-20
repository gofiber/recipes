package main

import (
	"testing"
)

func TestLoadingX509CertFromPKCS12Store(t *testing.T) {
	testCases := []struct {
		path     string
		password string
		expected string
	}{
		{"./security/server.p12", "changeit", "app.fiber.local"},
	}

	for _, tC := range testCases {
		tlsCert, error := initTLSConfig(tC.path, tC.password)

		if error != nil || tlsCert == nil {
			t.Errorf("%s: Failed to get TLS config object. Path: '%s', Password length: %v", t.Name(), tC.path, len(tC.password))
		}

		if tlsCert.Leaf.Subject.CommonName != tC.expected {
			t.Errorf("Unable to match expected CommonName outcome of '%s'. Found '%s'", tC.expected, tlsCert.Leaf.Subject.CommonName)
		}
	}
}
