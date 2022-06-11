// Copyright 2021 The Kubeflow Authors
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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/kubeflow/training-operator/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// MPIJobs returns a MPIJobInformer.
	MPIJobs() MPIJobInformer
	// MXJobs returns a MXJobInformer.
	MXJobs() MXJobInformer
	// PyTorchJobs returns a PyTorchJobInformer.
	PyTorchJobs() PyTorchJobInformer
	// TFJobs returns a TFJobInformer.
	TFJobs() TFJobInformer
	// XGBoostJobs returns a XGBoostJobInformer.
	XGBoostJobs() XGBoostJobInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// MPIJobs returns a MPIJobInformer.
func (v *version) MPIJobs() MPIJobInformer {
	return &mPIJobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MXJobs returns a MXJobInformer.
func (v *version) MXJobs() MXJobInformer {
	return &mXJobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PyTorchJobs returns a PyTorchJobInformer.
func (v *version) PyTorchJobs() PyTorchJobInformer {
	return &pyTorchJobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TFJobs returns a TFJobInformer.
func (v *version) TFJobs() TFJobInformer {
	return &tFJobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// XGBoostJobs returns a XGBoostJobInformer.
func (v *version) XGBoostJobs() XGBoostJobInformer {
	return &xGBoostJobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
