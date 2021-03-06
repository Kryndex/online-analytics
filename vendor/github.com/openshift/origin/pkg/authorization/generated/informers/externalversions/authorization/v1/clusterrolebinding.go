// This file was automatically generated by informer-gen

package v1

import (
	authorization_v1 "github.com/openshift/origin/pkg/authorization/apis/authorization/v1"
	clientset "github.com/openshift/origin/pkg/authorization/generated/clientset"
	internalinterfaces "github.com/openshift/origin/pkg/authorization/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/origin/pkg/authorization/generated/listers/authorization/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ClusterRoleBindingInformer provides access to a shared informer and lister for
// ClusterRoleBindings.
type ClusterRoleBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClusterRoleBindingLister
}

type clusterRoleBindingInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newClusterRoleBindingInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.AuthorizationV1().ClusterRoleBindings().List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.AuthorizationV1().ClusterRoleBindings().Watch(options)
			},
		},
		&authorization_v1.ClusterRoleBinding{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *clusterRoleBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&authorization_v1.ClusterRoleBinding{}, newClusterRoleBindingInformer)
}

func (f *clusterRoleBindingInformer) Lister() v1.ClusterRoleBindingLister {
	return v1.NewClusterRoleBindingLister(f.Informer().GetIndexer())
}
