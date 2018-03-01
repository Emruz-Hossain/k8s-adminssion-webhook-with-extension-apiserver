package util

import (
	"encoding/json"

	kc_v1alpha1 "github.com/emruz-hossain/k8s-admission-webhook-with-extension-apiserver/apis/kubecar/v1alpha1"
	"github.com/emruz-hossain/k8s-admission-webhook-with-extension-apiserver/client/clientset/versioned/typed/kubecar/v1alpha1"

	"github.com/appscode/kutil"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/jsonmergepatch"
)

func PatchKubecar(c v1alpha1.KubecarV1alpha1Interface, cur *kc_v1alpha1.Kubecar, transform func(kubecar *kc_v1alpha1.Kubecar) *kc_v1alpha1.Kubecar) (*kc_v1alpha1.Kubecar, kutil.VerbType, error) {
	curJson, err := json.Marshal(cur)
	if err != nil {
		return nil, kutil.VerbUnchanged, err
	}

	modJson, err := json.Marshal(transform(cur.DeepCopy()))
	if err != nil {
		return nil, kutil.VerbUnchanged, err
	}

	patch, err := jsonmergepatch.CreateThreeWayJSONMergePatch(curJson, modJson, curJson)
	if err != nil {
		return nil, kutil.VerbUnchanged, err
	}
	if len(patch) == 0 || string(patch) == "{}" {
		return cur, kutil.VerbUnchanged, nil
	}
	out, err := c.Kubecars(cur.Namespace).Patch(cur.Name, types.MergePatchType, patch)
	return out, kutil.VerbPatched, err
}
