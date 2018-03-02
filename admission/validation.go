package admission

import (
	"fmt"
	"net/http"

	"encoding/json"

	api "github.com/emruz-hossain/k8s-admission-webhook-with-extension-apiserver/apis/kubecar/v1alpha1"

	clientset "github.com/emruz-hossain/k8s-admission-webhook-with-extension-apiserver/client/clientset/versioned"

	admission "k8s.io/api/admission/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
)

type KubecarValidator struct {
	kubecarClient clientset.Interface
}

func (*KubecarValidator) ValidatingResource() (plural schema.GroupVersionResource, singular string) {
	return schema.GroupVersionResource{
			Group:    "admission.kubecar.emruz.com",
			Version:  "v1alpha1",
			Resource: "validations",
		},
		"validation"
}

func (f *KubecarValidator) Validate(req *admission.AdmissionRequest) *admission.AdmissionResponse {
	fmt.Println("KubecarValidator: " + req.Operation)

	if req.Operation == admission.Delete {
		return &admission.AdmissionResponse{Allowed: true}
	}

	if req.Operation == admission.Update {
		obj := &api.Kubecar{}
		if err := json.Unmarshal(req.Object.Raw, obj); err != nil {
			return &admission.AdmissionResponse{
				Allowed: false,
				Result: &metav1.Status{
					Status:  metav1.StatusFailure, Code: http.StatusBadRequest, Reason: metav1.StatusReasonBadRequest,
					Message: "invalid kubecar object",
				},
			}
		}

		oldObj := &api.Kubecar{}
		if err := json.Unmarshal(req.OldObject.Raw, oldObj); err != nil {
			return &admission.AdmissionResponse{
				Allowed: false,
				Result: &metav1.Status{
					Status:  metav1.StatusFailure, Code: http.StatusBadRequest, Reason: metav1.StatusReasonBadRequest,
					Message: "invalid old kubecar object",
				},
			}
		}

		// deny update if Accident count or Traffic rules violation is negative or Driving Skill does not match
		if obj.Spec.AccidentCount < 0 || obj.Spec.TrafficRuleViolationCount < 0 || obj.Spec.DrivingSkillPoint > 100 {
			return &admission.AdmissionResponse{
				Allowed: false,
				Result: &metav1.Status{
					Status:  metav1.StatusFailure, Code: http.StatusBadRequest, Reason: metav1.StatusReasonBadRequest,
					Message: "invalid specification",
				},
			}
		}
	}

	return &admission.AdmissionResponse{Allowed: true}
}

func (*KubecarValidator) Initialize(kubeClientConfig *rest.Config, stopCh <-chan struct{}) error {
	fmt.Println("KubecarValidator: Initialize")
	return nil
}
