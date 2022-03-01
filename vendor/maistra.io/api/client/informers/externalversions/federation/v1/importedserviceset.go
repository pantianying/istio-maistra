// Copyright Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	internalinterfaces "maistra.io/api/client/informers/externalversions/internalinterfaces"
	v1 "maistra.io/api/client/listers/federation/v1"
	versioned "maistra.io/api/client/versioned"
	federationv1 "maistra.io/api/federation/v1"
)

// ImportedServiceSetInformer provides access to a shared informer and lister for
// ImportedServiceSets.
type ImportedServiceSetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ImportedServiceSetLister
}

type importedServiceSetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewImportedServiceSetInformer constructs a new informer for ImportedServiceSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewImportedServiceSetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredImportedServiceSetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredImportedServiceSetInformer constructs a new informer for ImportedServiceSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredImportedServiceSetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FederationV1().ImportedServiceSets(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FederationV1().ImportedServiceSets(namespace).Watch(context.TODO(), options)
			},
		},
		&federationv1.ImportedServiceSet{},
		resyncPeriod,
		indexers,
	)
}

func (f *importedServiceSetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredImportedServiceSetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *importedServiceSetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&federationv1.ImportedServiceSet{}, f.defaultInformer)
}

func (f *importedServiceSetInformer) Lister() v1.ImportedServiceSetLister {
	return v1.NewImportedServiceSetLister(f.Informer().GetIndexer())
}