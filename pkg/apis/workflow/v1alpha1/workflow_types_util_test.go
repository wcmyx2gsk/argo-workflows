package v1alpha1

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func newTestWorkflow(phase WorkflowPhase, finishedAt metav1.Time) *Workflow {
	return &Workflow{
		ObjectMeta: metav1.ObjectMeta{
			CreationTimestamp: metav1.NewTime(time.Now().Add(-5 * time.Minute)),
		},
		Status: WorkflowStatus{
			Phase:      phase,
			FinishedAt: finishedAt,
		},
	}
}

func TestIsCompleted(t *testing.T) {
	assert.True(t, newTestWorkflow(WorkflowSucceeded, metav1.Now()).IsCompleted())
	assert.True(t, newTestWorkflow(WorkflowFailed, metav1.Now()).IsCompleted())
	// WorkflowError should also be considered completed
	assert.True(t, newTestWorkflow(WorkflowError, metav1.Now()).IsCompleted())
	assert.False(t, newTestWorkflow(WorkflowRunning, metav1.Time{}).IsCompleted())
	assert.False(t, newTestWorkflow(WorkflowPending, metav1.Time{}).IsCompleted())
	// WorkflowUnknown phase should not be considered completed
	assert.False(t, newTestWorkflow(WorkflowUnknown, metav1.Time{}).IsCompleted())
}

func TestIsRunning(t *testing.T) {
	assert.True(t, newTestWorkflow(WorkflowRunning, metav1.Time{}).IsRunning())
	assert.False(t, newTestWorkflow(WorkflowSucceeded, metav1.Now()).IsRunning())
	// A pending workflow should not be considered running
	assert.False(t, newTestWorkflow(WorkflowPending, metav1.Time{}).IsRunning())
}

func TestIsFailed(t *testing.T) {
	assert.True(t, newTestWorkflow(WorkflowFailed, metav1.Now()).IsFailed())
	assert.True(t, newTestWorkflow(WorkflowError, metav1.Now()).IsFailed())
	assert.False(t, newTestWorkflow(WorkflowSucceeded, metav1.Now()).IsFailed())
}

func TestHasExpired(t *testing.T) {
	secondsAfterSuccess := int32(10)
	finishedAt := metav1.NewTime(time.Now().Add(-1 * time.Minute))

	w := newTestWorkflow(WorkflowSucceeded, finishedAt)
	w.Spec.TTLStrategy = &TTLStrategy{
		SecondsAfterSuccess: &secondsAfterSuccess,
	}
	assert.True(t, w.HasExpired(), "workflow should have expired after success TTL")

	secondsAfterSuccess = int32(3600)
	w.Spec.TTLStrategy.SecondsAfterSuccess = &secondsAfterSuccess
	assert.False(t, w.HasExpired(), "workflow should not have expired with long TTL")

	wNoTTL := newTestWorkflow(WorkflowSucceeded, finishedAt)
	assert.False(t, wNoTTL.HasExpired(), "workflow without TTL should not expire")
}

func TestGetDuration(t *testing.T) {
	finishedAt := metav1.NewTime(time.Now().Add(-2 * time.Minute))
	w := newTestWorkflow(WorkflowSucceeded, finishedAt)
	duration := w.GetDuration()
	assert.Greater(t, duration.Seconds(), float64(0))
	// duration should be approximately 2 minutes (within a reasonable margin)
	// using 300s (5 min) as upper bound since creation timestamp is also -5 min
	assert.Less(t, duration.Seconds(), float64(300))
	// sanity check: duration should be at least 1 minute
	assert.GreaterOrEqual(t, duration.Seconds(), float64(60))
}
