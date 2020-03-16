/*
Copyright 2020 The OCP Release APIServer Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package proxy

import (
	"context"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"
	"net/http"
	"regexp"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ rest.CreaterUpdater = &KojiImageBuildPackageREST{}
var _ rest.Patcher = &KojiImageBuildPackageREST{}

// +k8s:deepcopy-gen=false
type KojiImageBuildPackageREST struct {
	Registry KojiImageBuildRegistry
}

func (r *KojiImageBuildPackageREST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	sub := obj.(*KojiImageBuildPackage)
	rec, err := r.Registry.GetKojiImageBuild(ctx, sub.Name, &metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	// Modify rec in someway before writing it back to storage

	r.Registry.UpdateKojiImageBuild(ctx, rec)
	return rec, nil
}

// Get retrieves the object from the storage. It is required to support Patch.
func (r *KojiImageBuildPackageREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	//build, err := r.Registry.GetKojiImageBuild(ctx, name, &metav1.GetOptions{})
	//	//if err != nil {
	//	//	return nil, err
	//	//}
	vsplit := strings.Split(name, "-")
	n := strings.Join(vsplit[0:len(vsplit) - 2], "-")
	v := vsplit[len(vsplit) - 2]
	rel := vsplit[len(vsplit) - 1]

	obj := &KojiImageBuildPackage{}
	obj.Name = name + "-packagelist"
	obj.Annotations = make(map[string]string)
	//obj.Annotations["koji-build-version"] = "0.0.0"
	//obj.Annotations["koji-build-release"] = "1"
	obj.Status.Packages = make(map[string][]string)
	//obj.Status.Packages["x86_64"]

	// http://download.eng.bos.redhat.com/brewroot/packages/openshift-jenkins-2-container/v4.4.0/202003160957/data/logs/x86_64.log
	resp, err := http.Get("http://download.eng.bos.redhat.com/brewroot/packages/" +
		n + "/" + v + "/" + rel + "/data/logs/x86_64.log")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyStr := string(body)

	re, err := regexp.Compile(`Installing : ([a-zA-Z0-9\-\.]+)\s+`)
	if err != nil {
		return nil, err
	}
	t := re.FindAllStringSubmatch(bodyStr, -1)
	for _, pkg := range t {
		obj.Status.Packages["x86_64"] = append(obj.Status.Packages["x86_64"], pkg[1])
	}

	return obj, nil
}

// Update alters the status subset of an object.
func (r *KojiImageBuildPackageREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	return nil, false, nil
}

func (r *KojiImageBuildPackageREST) New() runtime.Object {
	return &KojiImageBuildPackage{}
}
