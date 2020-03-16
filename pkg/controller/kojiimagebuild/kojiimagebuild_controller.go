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

package kojiimagebuild

import (
	"context"

	proxyv1alpha1 "github.com/vfreex/release-apiserver/pkg/apis/proxy/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new KojiImageBuild Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileKojiImageBuild{Client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("kojiimagebuild-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to KojiImageBuild
	err = c.Watch(&source.Kind{Type: &proxyv1alpha1.KojiImageBuild{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create
	// Uncomment watch a Deployment created by KojiImageBuild - change this for objects you create
	// err = c.Watch(&source.Kind{Type: &appsv1.Deployment{}}, &handler.EnqueueRequestForOwner{
	//	IsController: true,
	//	OwnerType:    &proxyv1alpha1.KojiImageBuild{},
	// })
	// if err != nil {
	//	return err
	// }

	return nil
}

var _ reconcile.Reconciler = &ReconcileKojiImageBuild{}

// ReconcileKojiImageBuild reconciles a KojiImageBuild object
type ReconcileKojiImageBuild struct {
	client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a KojiImageBuild object and makes changes based on the state read
// and what is in the KojiImageBuild.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  The scaffolding writes
// a Deployment as an example
// +kubebuilder:rbac:groups=proxy.art.openshift.io,resources=kojiimagebuilds,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=proxy.art.openshift.io,resources=kojiimagebuilds/status,verbs=get;update;patch
func (r *ReconcileKojiImageBuild) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	// Fetch the KojiImageBuild instance
	instance := &proxyv1alpha1.KojiImageBuild{}
	err := r.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}
