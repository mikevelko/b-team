package testutils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/pw-software-engineering/b-team/server/pkg/auth"
	"github.com/pw-software-engineering/b-team/server/pkg/bookly"

	"github.com/pw-software-engineering/b-team/server/pkg/httputils"
	"github.com/stretchr/testify/require"
)

// RequestOpt is optional argument for JSONRequest, which lets you e.g. add session header to request
type RequestOpt func(req *http.Request)

// WithSessionHeader func
func WithSessionHeader(s bookly.Session) RequestOpt {
	sessionJSON, _ := json.Marshal(s)
	return func(req *http.Request) {
		req.Header.Set(auth.HeaderXSession, string(sessionJSON))
	}
}

// JSONRequest initialize JSON request for testing
func JSONRequest(t *testing.T, method string, url string, marshallable interface{}, opts ...RequestOpt) *http.Request {
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(marshallable)
	require.NoError(t, err)
	req, err := http.NewRequest(method, url, &body)
	req.Header.Add("Content-Type", "application/json")
	require.NoError(t, err)
	for _, opt := range opts {
		opt(req)
	}

	return req
}

// ErrRespFromBody tries to parse response body to httputils.ErrorResponse
func ErrRespFromBody(t *testing.T, body io.Reader) httputils.ErrorResponse {
	var resp httputils.ErrorResponse
	require.NoError(t, json.NewDecoder(body).Decode(&resp))
	return resp
}
