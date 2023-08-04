package main

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
)

var _ rest.Getter = (*SubresourceStreamerREST)(nil)

type SubresourceStreamerREST struct {
	store *genericregistry.Store
}

func NewStaticQuerySubresourceStreamerREST(resource schema.GroupResource) {
}

func NewSubresourceStreamerREST(resource schema.GroupResource, singularResource schema.GroupResource, strategy streamerStrategy, optsGetter generic.RESTOptionsGetter, tableConvertor rest.TableConvertor) *SubresourceStreamerREST {
	var storage SubresourceStreamerREST
	store := &genericregistry.Store{
		NewFunc:     func() runtime.Object { return &SubresourceStreamer{} },
		NewListFunc: func() runtime.Object { return &SubresourceStreamer{} },

		DefaultQualifiedResource:  resource,
		SingularQualifiedResource: singularResource,

		CreateStrategy:      strategy,
		UpdateStrategy:      strategy,
		DeleteStrategy:      strategy,
		ResetFieldsStrategy: strategy,

		TableConvertor: tableConvertor,
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter}
	if err := store.CompleteWithOptions(options); err != nil {
		panic(err) // TODO: Propagate error up
	}
	storage.store = store
	return &storage

}

func (r *SubresourceStreamerREST) New() runtime.Object {
	return r.store.New()
}

func (r *SubresourceStreamerREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	o, err := r.store.Get(ctx, name, options)
	if err != nil {
		return nil, err
	}
	return o, nil
}
