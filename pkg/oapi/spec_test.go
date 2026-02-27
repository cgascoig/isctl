package oapi

import (
	"sync"
	"testing"

	"github.com/icza/dyno"
	"github.com/stretchr/testify/assert"
)

func TestLazyLoadSpec(t *testing.T) {
	s := lazyLoadSpec()
	assert.NotNil(t, s)

	oapi, err := dyno.Get(s, "openapi")
	assert.NoError(t, err)
	assert.Equal(t, "3.0.2", oapi)
}

func TestLazyLoadSpecConcurrency(t *testing.T) {
	const goroutines = 10
	var wg sync.WaitGroup
	results := make(chan map[string]any, goroutines)

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s := lazyLoadSpec()
			results <- s
		}()
	}

	wg.Wait()
	close(results)

	// Check all returned the same pointer
	var firstResult map[string]any
	for s := range results {
		if firstResult == nil {
			firstResult = s
		} else {
			assert.Equal(t, firstResult, s, "concurrent lazyLoadSpec() calls should return same map")
		}
	}
}
