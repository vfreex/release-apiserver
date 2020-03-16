package proxy

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"sigs.k8s.io/apiserver-builder-alpha/pkg/builders"
)

// +k8s:deepcopy-gen=false
type KojiImageBuildREST struct {
	*genericregistry.Store
}


func (r *KojiImageBuildREST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	return nil, nil
}

// Get retrieves the object from the storage. It is required to support Patch.
func (r *KojiImageBuildREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	return &KojiImageBuild{}, nil
}

// Update alters the status subset of an object.
func (r *KojiImageBuildREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	return nil, false, nil
}

func (r *KojiImageBuildREST) New() runtime.Object {
	return &KojiImageBuild{}
}

func (r *KojiImageBuildREST) NamespaceScoped() bool {
	return true
}

func (r *KojiImageBuildREST) ShortNames() []string {
	return []string{"ki"}
}

func (r *KojiImageBuildREST) Categories() []string {
	return []string{""}
}


func NewKojiImageBuildREST(optsGetter generic.RESTOptionsGetter) rest.Storage {
	groupResource := schema.GroupResource{
		Group:    "miskatonic.k8s.io",
		Resource: "students",
	}
	strategy := &KojiImageBuildStrategy{builders.StorageStrategySingleton}
	store := &genericregistry.Store{
		NewFunc:                  func() runtime.Object { return &KojiImageBuild{} },
		NewListFunc:              func() runtime.Object { return &KojiImageBuildList{} },
		DefaultQualifiedResource: groupResource,

		CreateStrategy: strategy, // TODO: specify create strategy
		UpdateStrategy: strategy, // TODO: specify update strategy
		DeleteStrategy: strategy, // TODO: specify delete strategy
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter}
	if err := store.CompleteWithOptions(options); err != nil {
		panic(err) // TODO: Propagate error up
	}
	return &KojiImageBuildREST{store}
}

