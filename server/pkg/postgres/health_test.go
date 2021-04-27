package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_healthCheckSuccess(t *testing.T) {
	pool, cleanup, err := NewPool(conf)
	require.NoError(t, err)
	t.Cleanup(cleanup)
	check := newCheck(pool)

	_, err = check.Execute()
	assert.NoError(t, err)
}

func Test_healthCheckFail(t *testing.T) {
	pool, cleanup, err := NewPool(conf)
	require.NoError(t, err)

	cleanup() // clean up all connections
	check := newCheck(pool)

	detail, err := check.Execute()
	assert.Error(t, err)
	assert.Contains(t, detail, conf.Host)
}
