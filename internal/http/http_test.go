package http

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func Test(t *testing.T) {
	logger := zap.NewExample()
	defer logger.Sync()

	httpServerIsRunning := new(sync.WaitGroup)
	httpServerIsRunning.Add(1)
	isReady := &atomic.Value{}
	isReady.Store(false)

	// Launch an HTTP server
	go func() {
		server := NewServer(logger, "12345")
		httpServerIsRunning.Done()
		server.Serve()
	}()

	httpServerIsRunning.Wait()
	time.Sleep(30 * time.Millisecond) // makes sure the serve.Serve() has had time to execute

	// Test the HTTP server

	// The server should be healthy
	resp, err := http.Get(fmt.Sprintf("http://localhost:12345%s", livenessPath))
	require.Nil(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
