// Copyright 2018 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// SchemeGroupVersion is group version used to register these objects.
var SchemeGroupVersion = GroupVersion

// Resource takes an unqualified resource and returns a Group-qualified GroupResource.
func Resource(resource string) schema.GroupResource {
	return GroupVersion.WithResource(resource).GroupResource()
}

func init() {
	SchemeBuilder.Register(&TFJob{}, &TFJobList{})
	SchemeBuilder.Register(&PyTorchJob{}, &PyTorchJobList{})
	SchemeBuilder.Register(&MPIJob{}, &MPIJobList{})
	SchemeBuilder.Register(&MXJob{}, &MXJobList{})
	SchemeBuilder.Register(&XGBoostJob{}, &XGBoostJobList{})

	SchemeBuilder.SchemeBuilder.Register(addDefaultingFuncs)
}

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}
