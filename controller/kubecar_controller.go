package controller

import (
	"fmt"
	"time"

	"github.com/appscode/go/log"
	"github.com/appscode/go/runtime"
	"github.com/appscode/kutil/meta"
	"github.com/appscode/kutil/tools/queue"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	kc_v1alpha1 "github.com/emruz-hossain/k8s-admission-webhook-with-extension-apiserver/apis/kubecar/v1alpha1"
	clientset "github.com/emruz-hossain/k8s-admission-webhook-with-extension-apiserver/client/clientset/versioned"
	kubecar_informer "github.com/emruz-hossain/k8s-admission-webhook-with-extension-apiserver/client/informers/externalversions"
	kubecar_lister "github.com/emruz-hossain/k8s-admission-webhook-with-extension-apiserver/client/listers/kubecar/v1alpha1"
	kc_util "github.com/emruz-hossain/k8s-admission-webhook-with-extension-apiserver/util"
)

type Options struct {
	MaxNumReques int
	ResyncPeriod time.Duration
	NumThreads   int
}

type KubecarController struct {
	opt Options

	k8sClient     kubernetes.Interface
	kubecarClient clientset.Interface

	kubecarInformerFactory kubecar_informer.SharedInformerFactory
	kubecarLister          kubecar_lister.KubecarLister
	kubecarQueue           *queue.Worker
	kubecarInformer        cache.SharedIndexInformer
}

func NewOptions() *Options {
	return &Options{
		MaxNumReques: 5,
		ResyncPeriod: time.Minute * 5,
		NumThreads:   1,
	}
}

func NewKubecarController(k8sClient kubernetes.Interface, kubecarClient clientset.Interface, opt Options) *KubecarController {
	cntrl := &KubecarController{
		opt: opt,

		k8sClient:     k8sClient,
		kubecarClient: kubecarClient,

		kubecarInformerFactory: kubecar_informer.NewSharedInformerFactory(kubecarClient, opt.ResyncPeriod),
	}

	cntrl.initKubecarWatcher()

	return cntrl

}

func (c *KubecarController) initKubecarWatcher() {
	c.kubecarInformer = c.kubecarInformerFactory.Kubecar().V1alpha1().Kubecars().Informer()
	c.kubecarLister = c.kubecarInformerFactory.Kubecar().V1alpha1().Kubecars().Lister()
	c.kubecarQueue = queue.New("Kubecar", c.opt.MaxNumReques, c.opt.NumThreads, c.runKubecarEventProcessor)
	c.kubecarInformer.AddEventHandler(queue.DefaultEventHandler(c.kubecarQueue.GetQueue()))

}

func (c *KubecarController) runKubecarEventProcessor(key string) error {
	obj, exist, err := c.kubecarInformer.GetIndexer().GetByKey(key)
	if err != nil {
		log.Errorln("Fetching object with key", key, "failed. Reason:", err.Error())
		return err
	}

	if !exist {
		log.Warning("Kubecar ", key, "doesnot exit anymore.")
	} else {
		log.Infoln("Add/Update/Sync happed for ", key)
		kc := obj.(*kc_v1alpha1.Kubecar)

		// Every accident reduce 10 driving skill points and every traffic rules violation reduce 2 driving skill points.
		// If driving skill point is 0 then the car is forbidden to run on street. Hence, it must be removed.
		if 100-(kc.Spec.AccidentCount*10+kc.Spec.TrafficRuleViolationCount*2) <= 0 {
			err = c.kubecarClient.KubecarV1alpha1().Kubecars(kc.Namespace).Delete(kc.Name, meta.DeleteInBackground())
			return err
		}

		if 100-(kc.Spec.AccidentCount*10+kc.Spec.TrafficRuleViolationCount*2) != kc.Spec.DrivingSkillPoint {
			kc_util.PatchKubecar(c.kubecarClient.KubecarV1alpha1(), kc, func(in *kc_v1alpha1.Kubecar) *kc_v1alpha1.Kubecar {
				in.Spec.DrivingSkillPoint = in.Spec.AccidentCount*10 + in.Spec.TrafficRuleViolationCount*2
				return in
			})
		}

	}
	return nil
}

func (c *KubecarController) Run(stopCh <-chan struct{}) {
	defer runtime.HandleCrash()

	c.kubecarInformerFactory.Start(stopCh)

	for _, v := range c.kubecarInformerFactory.WaitForCacheSync(stopCh) {
		if !v {
			runtime.HandleError(fmt.Errorf("timed out. waiting fro caches to sync."))
			return
		}
	}
	c.kubecarQueue.Run(stopCh)

	<-stopCh
	log.Infoln("Stoppig kubecar controller")

	return
}
