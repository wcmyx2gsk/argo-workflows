package v1alpha1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func int32Ptr(i int32) *int32 {
	return &i
}

func TestTTLStrategyIsZero(t *testing.T) {
	t.Run("nil strategy", func(t *testing.T) {
		var s *TTLStrategy
		assert.True(t, s.IsZero())
	})

	t.Run("empty strategy", func(t *testing.T) {
		s := &TTLStrategy{}
		assert.True(t, s.IsZero())
	})

	t.Run("with SecondsAfterCompletion", func(t *testing.T) {
		s := &TTLStrategy{SecondsAfterCompletion: int32Ptr(60)}
		assert.False(t, s.IsZero())
	})

	t.Run("with SecondsAfterSuccess", func(t *testing.T) {
		s := &TTLStrategy{SecondsAfterSuccess: int32Ptr(120)}
		assert.False(t, s.IsZero())
	})

	t.Run("with SecondsAfterFailure", func(t *testing.T) {
		s := &TTLStrategy{SecondsAfterFailure: int32Ptr(30)}
		assert.False(t, s.IsZero())
	})

	// A strategy with zero value pointers (pointing to 0) should NOT be considered zero
	// because the user explicitly set the field, even if the value is 0.
	t.Run("with zero-value pointer fields", func(t *testing.T) {
		s := &TTLStrategy{SecondsAfterCompletion: int32Ptr(0)}
		assert.False(t, s.IsZero())
	})

	// All three fields set — also not zero
	t.Run("with all fields set", func(t *testing.T) {
		s := &TTLStrategy{
			SecondsAfterCompletion: int32Ptr(60),
			SecondsAfterSuccess:    int32Ptr(120),
			SecondsAfterFailure:    int32Ptr(30),
		}
		assert.False(t, s.IsZero())
	})

	// Only SecondsAfterFailure set — useful for retaining failed workflows longer for debugging
	t.Run("with only SecondsAfterFailure set for debug retention", func(t *testing.T) {
		s := &TTLStrategy{SecondsAfterFailure: int32Ptr(3600)}
		assert.False(t, s.IsZero())
	})
}

func TestTTLStrategyGetters(t *testing.T) {
	var nilStrategy *TTLStrategy
	assert.Nil(t, nilStrategy.GetSecondsAfterCompletion())
	assert.Nil(t, nilStrategy.GetSecondsAfterSuccess())
	assert.Nil(t, nilStrategy.GetSecondsAfterFailure())

	s := &TTLStrategy{
		SecondsAfterCompletion: int32Ptr(60),
		SecondsAfterSuccess:    int32Ptr(120),
		SecondsAfterFailure:    int32Ptr(30),
	}
	assert.Equal(t, int32(60), *s.GetSecondsAfterCompletion())
	assert.Equal(t, int32(120), *s.GetSecondsAfterSuccess())
	assert.Equal(t, int32(30), *s.GetSecondsAfterFailure())
}
