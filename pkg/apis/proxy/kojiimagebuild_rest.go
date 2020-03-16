package proxy

import (
	"context"
	"crypto/tls"
	"github.com/kolo/xmlrpc"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"net/http"
	"sigs.k8s.io/apiserver-builder-alpha/pkg/builders"
	"strconv"
)

// +k8s:deepcopy-gen=false
type KojiImageBuildREST struct {
	*genericregistry.Store
	//Registry KojiImageBuildRegistry
}

func (r *KojiImageBuildREST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	//ki := obj.(*KojiImageBuild)
	//ki.Status.KojiState = "Unknown"
	//ki.Spec.Name = "ose-something-container"
	//ki.Spec.Version = "0.0.0"
	//ki.Spec.Release = "202003160000"
	//r.Store.Create(ctx, ki, createValidation, options)
	//return ki, nil
	return nil, nil
}

// Get retrieves the object from the storage. It is required to support Patch.
func (r *KojiImageBuildREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	//ki, err := r.Store.Get(ctx, name, options)
	//if err != nil {
	//	t := time.Now().UTC()
	//	ts := t.Format("20060102150405")
	//	ki = &KojiImageBuild{Spec: KojiImageBuildSpec{Name: name, Version: "0.0.0", Release: ts}}
	//	return ki, nil
	//}

	rpc, _ := xmlrpc.NewClient("https://brewhub.engineering.redhat.com/brewhub",
		&http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		})
	var buildinfo struct {
		BuildId      int     `xmlrpc:"build_id"`
		OwnerName    string  `xmlrpc:"owner_name"`
		PackageName  string  `xmlrpc:"package_name"`
		State        BuildState     `xmlrpc:"state"`
		Nvr          string  `xmlrpc:"nvr"`
		Version      string  `xmlrpc:"version"`
		Release      string  `xmlrpc:"release"`
		Epoch        string  `xmlrpc:"epoch"`
		CreationTs   float64 `xmlrpc:"creation_ts"`
		StartTs      float64 `xmlrpc:"start_ts"`
		CompletionTs float64 `xmlrpc:"completion_ts"`
		Source       string  `xmlrpc:"source"`

		Extra struct {
			Source struct {
				OriginalUrl string `xmlrpc:"original_url"`
			} `xmlrpc:"source"`
			Image struct {
				Index struct {
					UniqueTags []string `xmlrpc:"unique_tags"`
					Pull       []string `xmlrpc:"pull"`
					Digests    struct {
						ManifestList string `xmlrpc:"application/vnd.docker.distribution.manifest.list.v2+json"`
					} `xmlrpc:"digests"`
				} `xmlrpc:"index"`
			} `xmlrpc:"image"`
		} `xmlrpc:"extra"`
	}
	if err := rpc.Call("getBuild", name, &buildinfo); err != nil {
		return nil, err
	}
	obj := &KojiImageBuild{}
	obj.Name = name
	obj.Spec.Name = buildinfo.Nvr
	obj.Spec.Version = buildinfo.Version
	obj.Spec.Release = buildinfo.Release
	obj.Labels = make(map[string]string)
	obj.Labels["release-channel.art.openshift.io"] = "ocp-4.4"
	obj.Labels["release.art.openshift.io"] = "4.4.0-rc.0"
	obj.Status.KojiState = buildinfo.State.String()
	obj.Status.KojiBuildLink = "https://brewweb.engineering.redhat.com/brew/buildinfo?buildID=" + strconv.Itoa(buildinfo.BuildId)
	return obj, nil
}

// Update alters the status subset of an object.
func (r *KojiImageBuildREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {

	//ki, b, err := r.Store.Update(ctx, name, objInfo, createValidation, updateValidation, forceAllowCreate, options)
	//objInfo.UpdatedObject()
	return nil, false, nil

	//obj := &KojiImageBuild{}
	//obj.Spec.Name = name
	//obj.Spec.Version = "4.4.0"
	//obj.Spec.Release = time.Now().UTC().Format("20060102150405")
	//obj.Name = obj.Spec.Name + "-" + obj.Spec.Version + "-" + obj.Spec.Release
	//obj.Labels = make(map[string]string)
	//obj.Labels["release-channel.art.openshift.io"] = "ocp-4.4"
	//obj.Status.KojiState = "Patched"
	//obj.Status.KojiBuildLink = "https://brewweb.engineering.redhat.com/foo/"
	//return obj, true, nil
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
		Group:    "proxy.art.openshift.io",
		Resource: "kojiimagebuilds",
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


// BUILD_STATES
// https://pagure.io/koji/blob/8172bff0535f364eab1a2e7203e494526d7a425d/f/koji/__init__.py#_196
type BuildState int

const (
	BUILDING BuildState = iota
	COMPLETE
	DELETED
	FAILED
	CANCELED
)

func (s BuildState) String() string {
	return [...]string{"Building", "Complete", "Deleted", "Canceled"}[s]
}