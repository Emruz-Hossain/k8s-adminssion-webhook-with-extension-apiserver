package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Kubecar struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KubecarSpec `json:"spec"`
}

type KubecarSpec struct {
	AccidentCount int32 `json:"accident_count"`
	TrafficRuleViolationCount int32 `json:"traffic_rule_violation_count"`
	DrivingSkillPoint int32 `json:"driving_skill_point"`
}


// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KubecarList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	
	Items []Kubecar `json:"items"`
}