package testutils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/pw-software-engineering/b-team/server/pkg/httputils"
	"github.com/stretchr/testify/require"
)

// JSONRequest initialize JSON request for testing
func JSONRequest(t *testing.T, method string, url string, marshallable interface{}) *http.Request {
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(marshallable)
	require.NoError(t, err)
	req, err := http.NewRequest(method, url, &body)
	req.Header.Add("Content-Type", "application/json")
	require.NoError(t, err)
	return req
}

// ErrRespFromBody tries to parse response body to httputils.ErrorResponse
func ErrRespFromBody(t *testing.T, body io.Reader) httputils.ErrorResponse {
	var resp httputils.ErrorResponse
	require.NoError(t, json.NewDecoder(body).Decode(&resp))
	return resp
}
