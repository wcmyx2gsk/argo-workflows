package v1alpha1

// TTLStrategy is the strategy for the time to live depending on if the workflow succeeded or failed.
type TTLStrategy struct {
	// SecondsAfterCompletion is the number of seconds to live after completion
	// +optional
	SecondsAfterCompletion *int32 `json:"secondsAfterCompletion,omitempty" protobuf:"varint,1,opt,name=secondsAfterCompletion"`
	// SecondsAfterSuccess is the number of seconds to live after success
	// +optional
	SecondsAfterSuccess *int32 `json:"secondsAfterSuccess,omitempty" protobuf:"varint,2,opt,name=secondsAfterSuccess"`
	// SecondsAfterFailure is the number of seconds to live after failure
	// +optional
	SecondsAfterFailure *int32 `json:"secondsAfterFailure,omitempty" protobuf:"varint,3,opt,name=secondsAfterFailure"`
}

// GetSecondsAfterCompletion returns the seconds after completion or nil.
func (t *TTLStrategy) GetSecondsAfterCompletion() *int32 {
	if t == nil {
		return nil
	}
	return t.SecondsAfterCompletion
}

// GetSecondsAfterSuccess returns the seconds after success or nil.
func (t *TTLStrategy) GetSecondsAfterSuccess() *int32 {
	if t == nil {
		return nil
	}
	return t.SecondsAfterSuccess
}

// GetSecondsAfterFailure returns the seconds after failure or nil.
func (t *TTLStrategy) GetSecondsAfterFailure() *int32 {
	if t == nil {
		return nil
	}
	return t.SecondsAfterFailure
}

// IsZero returns true if no TTL values are set.
func (t *TTLStrategy) IsZero() bool {
	if t == nil {
		return true
	}
	return t.SecondsAfterCompletion == nil &&
		t.SecondsAfterSuccess == nil &&
		t.SecondsAfterFailure == nil
}
