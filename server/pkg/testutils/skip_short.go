package testutils

import "testing"

// SetIntegration will skip given test if go test is run with -short flag
func SetIntegration(t *testing.T) {
    if testing.Short() {
        t.Skip("integration test are omitted in short mode")
    }
}
