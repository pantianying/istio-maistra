// Code generated by xns-informer-gen. DO NOT EDIT.

package v1alpha3

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	informers "istio.io/client-go/pkg/informers/externalversions/networking/v1alpha3"
	listers "istio.io/client-go/pkg/listers/networking/v1alpha3"
	"k8s.io/client-go/tools/cache"
)

type workloadGroupInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.WorkloadGroupInformer = &workloadGroupInformer{}

func NewWorkloadGroupInformer(f xnsinformers.SharedInformerFactory) informers.WorkloadGroupInformer {
	resource := v1alpha3.SchemeGroupVersion.WithResource("workloadgroups")
	converter := xnsinformers.NewListWatchConverter(
		f.GetScheme(),
		&v1alpha3.WorkloadGroup{},
		&v1alpha3.WorkloadGroupList{},
	)

	informer := f.ForResource(resource, xnsinformers.ResourceOptions{
		ClusterScoped:      false,
		ListWatchConverter: converter,
	})

	return &workloadGroupInformer{informer: informer.Informer()}
}

func (i *workloadGroupInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *workloadGroupInformer) Lister() listers.WorkloadGroupLister {
	return listers.NewWorkloadGroupLister(i.informer.GetIndexer())
}