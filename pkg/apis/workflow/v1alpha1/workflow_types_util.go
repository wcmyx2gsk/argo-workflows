package v1alpha1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetDuration returns the duration of the workflow if completed, otherwise the time since creation.
func (w *Workflow) GetDuration() time.Duration {
	if w.Status.FinishedAt.IsZero() {
		return time.Since(w.CreationTimestamp.Time)
	}
	return w.Status.FinishedAt.Sub(w.CreationTimestamp.Time)
}

// IsCompleted returns true if the workflow has a completed phase.
func (w *Workflow) IsCompleted() bool {
	return w.Status.Phase.Completed()
}

// IsRunning returns true if the workflow is currently running.
func (w *Workflow) IsRunning() bool {
	return w.Status.Phase == WorkflowRunning
}

// IsFailed returns true if the workflow has failed.
func (w *Workflow) IsFailed() bool {
	return w.Status.Phase == WorkflowFailed || w.Status.Phase == WorkflowError
}

// IsSucceeded returns true if the workflow succeeded.
func (w *Workflow) IsSucceeded() bool {
	return w.Status.Phase == WorkflowSucceeded
}

// GetTTLStrategy returns the effective TTL strategy for the workflow.
func (w *Workflow) GetTTLStrategy() *TTLStrategy {
	if w.Spec.TTLStrategy != nil {
		return w.Spec.TTLStrategy
	}
	return nil
}

// HasExpired returns true if the workflow has exceeded its TTL.
func (w *Workflow) HasExpired() bool {
	ttl := w.GetTTLStrategy()
	if ttl == nil || !w.IsCompleted() {
		return false
	}
	finishedAt := w.Status.FinishedAt
	if finishedAt.IsZero() {
		return false
	}
	if w.IsSucceeded() && ttl.SecondsAfterSuccess != nil {
		return metav1.Now().After(finishedAt.Add(time.Duration(*ttl.SecondsAfterSuccess) * time.Second))
	}
	if w.IsFailed() && ttl.SecondsAfterFailure != nil {
		return metav1.Now().After(finishedAt.Add(time.Duration(*ttl.SecondsAfterFailure) * time.Second))
	}
	if ttl.SecondsAfterCompletion != nil {
		return metav1.Now().After(finishedAt.Add(time.Duration(*ttl.SecondsAfterCompletion) * time.Second))
	}
	return false
}
