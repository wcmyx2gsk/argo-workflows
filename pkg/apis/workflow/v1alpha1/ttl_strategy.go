package v1alpha1

// TTLStrategy is the strategy for the time to live depending on if the workflow succeeded or failed.
// Note: SecondsAfterCompletion takes precedence if multiple fields are set.
// Personal note: SecondsAfterSuccess and SecondsAfterFailure are only evaluated when
// SecondsAfterCompletion is not set. Keep this in mind when configuring cleanup policies.
// Personal note: I typically set SecondsAfterFailure to a higher value (e.g. 7 days = 604800s) so
// failed workflows stick around long enough for debugging before being cleaned up.
// Personal note: For success, 1 day (86400s) is usually sufficient in my workflows.
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

// HasSuccessOrFailureTTL returns true if either success or failure TTL is set,
// which allows more granular cleanup behavior than SecondsAfterCompletion alone.
// Note: this is useful for keeping failed workflows around longer for debugging.
func (t *TTLStrategy) HasSuccessOrFailureTTL() bool {
	if t == nil {
		return false
	}
	return t.SecondsAfterSuccess != nil || t.SecondsAfterFailure != nil
}
