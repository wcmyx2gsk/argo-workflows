package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// WorkflowPhase is a label for the condition of a workflow at the current time.
type WorkflowPhase string

// NodePhase is a label for the condition of a node at the current time.
type NodePhase string

const (
	// WorkflowPending means the workflow has been accepted by the system, but not yet started.
	WorkflowPending WorkflowPhase = "Pending"
	// WorkflowRunning means the workflow is currently being executed.
	WorkflowRunning WorkflowPhase = "Running"
	// WorkflowSucceeded means the workflow has been successfully completed.
	WorkflowSucceeded WorkflowPhase = "Succeeded"
	// WorkflowFailed means the workflow was not successful.
	WorkflowFailed WorkflowPhase = "Failed"
	// WorkflowError means the workflow had an error.
	WorkflowError WorkflowPhase = "Error"
)

// Workflow is the definition of a workflow resource.
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Workflow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata" protobuf:"bytes,1,opt,name=metadata"`
	Spec              WorkflowSpec   `json:"spec" protobuf:"bytes,2,opt,name=spec"`
	Status            WorkflowStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// WorkflowSpec is the spec of a Workflow.
type WorkflowSpec struct {
	// Entrypoint is the name of the template to use as the starting point of the workflow.
	Entrypoint string `json:"entrypoint,omitempty" protobuf:"bytes,1,opt,name=entrypoint"`

	// TTLStrategy defines the strategy for the time to live of a workflow.
	TTLStrategy *TTLStrategy `json:"ttlStrategy,omitempty" protobuf:"bytes,2,opt,name=ttlStrategy"`

	// ActiveDeadlineSeconds is an optional duration in seconds relative to the
	// workflow start time which the workflow is allowed to run before the
	// controller terminates the workflow.
	ActiveDeadlineSeconds *int64 `json:"activeDeadlineSeconds,omitempty" protobuf:"varint,3,opt,name=activeDeadlineSeconds"`
}

// WorkflowStatus contains overall status information about a workflow.
type WorkflowStatus struct {
	// Phase is the current phase of the workflow.
	Phase WorkflowPhase `json:"phase,omitempty" protobuf:"bytes,1,opt,name=phase,casttype=WorkflowPhase"`

	// StartedAt is the time the workflow was started.
	StartedAt metav1.Time `json:"startedAt,omitempty" protobuf:"bytes,2,opt,name=startedAt"`

	// FinishedAt is the time the workflow finished.
	FinishedAt metav1.Time `json:"finishedAt,omitempty" protobuf:"bytes,3,opt,name=finishedAt"`

	// Message is a human-readable message indicating details about why the workflow is in this condition.
	Message string `json:"message,omitempty" protobuf:"bytes,4,opt,name=message"`
}

// WorkflowList is list of Workflow resources.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type WorkflowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Workflow `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// DeepCopyObject implements the runtime.Object interface.
func (w *Workflow) DeepCopyObject() interface{} {
	return w.DeepCopy()
}

// DeepCopy returns a deep copy of the Workflow.
func (w *Workflow) DeepCopy() *Workflow {
	if w == nil {
		return nil
	}
	out := new(Workflow)
	*out = *w
	return out
}
