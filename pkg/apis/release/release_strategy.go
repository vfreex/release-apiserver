
/*
Copyright 2020 The OpenShift Release APIServer Authors.

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


package release

import (
	"context"
	"regexp"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/klog"
)

// Validate checks that an instance of Release is well formed
func (ReleaseStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*Release)
	klog.V(5).Infof("Validating fields for Release %s", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	if matched, _ := regexp.MatchString(`[\w\-]+`, o.Spec.ReleaseStreamName); !matched {
		errors = append(errors, field.Invalid(field.NewPath("spec", "releaseStreamName"),
			o.Spec.ReleaseStreamName, "releaseStreamName may not be valid."))
	}
	return errors
}
