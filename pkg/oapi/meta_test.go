package oapi

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMeta(t *testing.T) {
	m, err := GetMeta()
	assert.NoError(t, err)
	assert.NotNil(t, m)
	assert.Greater(t, len(*m), 0, "metadata should not be empty")
}

func TestGetMetaReturnsConsistentData(t *testing.T) {
	m1, err1 := GetMeta()
	assert.NoError(t, err1)

	m2, err2 := GetMeta()
	assert.NoError(t, err2)

	// Both calls should return the same pointer (lazy loaded once)
	assert.Equal(t, m1, m2, "GetMeta() should return same pointer on subsequent calls")
}

func TestGetMetaContainsExpectedData(t *testing.T) {
	m, err := GetMeta()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(*m), 1, "expected at least 1 metadata entry")

	// Check first entry has expected fields
	first := (*m)[0]
	assert.NotEmpty(t, first.ClassId)
	assert.NotEmpty(t, first.Name)
	assert.NotEmpty(t, first.ObjectType)
}

func TestGetMetaConcurrency(t *testing.T) {
	const goroutines = 10
	var wg sync.WaitGroup
	results := make(chan *Meta, goroutines)
	errs := make(chan error, goroutines)

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m, err := GetMeta()
			results <- m
			errs <- err
		}()
	}

	wg.Wait()
	close(results)
	close(errs)

	// Check all goroutines succeeded
	for err := range errs {
		assert.NoError(t, err)
	}

	// Check all returned the same pointer
	var firstResult *Meta
	for m := range results {
		if firstResult == nil {
			firstResult = m
		} else {
			assert.Equal(t, firstResult, m, "concurrent GetMeta() calls should return same pointer")
		}
	}
}

func TestClassMetaStructure(t *testing.T) {
	m, err := GetMeta()
	assert.NoError(t, err)

	for i, cm := range *m {
		assert.NotEmpty(t, cm.ClassId, "entry %d: ClassId should not be empty", i)
		assert.NotEmpty(t, cm.Name, "entry %d: Name should not be empty", i)
		assert.NotEmpty(t, cm.ObjectType, "entry %d: ObjectType should not be empty", i)
		assert.NotEmpty(t, cm.MetaType, "entry %d: MetaType should not be empty", i)
	}
}

func TestPropertyMetaStructure(t *testing.T) {
	m, err := GetMeta()
	assert.NoError(t, err)

	// Find an entry with properties
	for _, cm := range *m {
		if len(cm.Properties) > 0 {
			for j, pm := range cm.Properties {
				assert.NotEmpty(t, pm.Name, "property %d in class %s: Name should not be empty", j, cm.Name)
				assert.NotEmpty(t, pm.Type, "property %d in class %s: Type should not be empty", j, cm.Name)
				assert.NotEmpty(t, pm.ApiAccess, "property %d in class %s: ApiAccess should not be empty", j, cm.Name)
			}
			return // Tested one class with properties
		}
	}
}

func TestRelationshipMetaStructure(t *testing.T) {
	m, err := GetMeta()
	assert.NoError(t, err)

	// Find an entry with relationships
	for _, cm := range *m {
		if len(cm.Relationships) > 0 {
			for j, rm := range cm.Relationships {
				assert.NotEmpty(t, rm.Name, "relationship %d in class %s: Name should not be empty", j, cm.Name)
				assert.NotEmpty(t, rm.Type, "relationship %d in class %s: Type should not be empty", j, cm.Name)
			}
			return // Tested one class with relationships
		}
	}
}

func TestGetIdentityConstraints(t *testing.T) {
	m, err := GetMeta()
	assert.NoError(t, err)

	assert.Equal(t, []string{"VlanId",
		"EthNetworkPolicy"}, m.GetIdentityConstraints("fabric.Vlan"))
}
