/*
Copyright 2021 The Kubernetes Authors.

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

package util

import (
	operatorv1 "sigs.k8s.io/cluster-api-operator/api/v1alpha2"
	"sigs.k8s.io/cluster-api-operator/internal/controller/genericprovider"
	clusterctlv1 "sigs.k8s.io/cluster-api/cmd/clusterctl/api/v1alpha3"
)

func IsCoreProvider(p genericprovider.GenericProvider) bool {
	_, ok := p.(*operatorv1.CoreProvider)
	return ok
}

// ClusterctlProviderType returns the provider type from the genericProvider.
func ClusterctlProviderType(genericProvider operatorv1.GenericProvider) clusterctlv1.ProviderType {
	switch genericProvider.(type) {
	case *operatorv1.CoreProvider:
		return clusterctlv1.CoreProviderType
	case *operatorv1.ControlPlaneProvider:
		return clusterctlv1.ControlPlaneProviderType
	case *operatorv1.InfrastructureProvider:
		return clusterctlv1.InfrastructureProviderType
	case *operatorv1.BootstrapProvider:
		return clusterctlv1.BootstrapProviderType
	case *operatorv1.AddonProvider:
		return clusterctlv1.AddonProviderType
	case *operatorv1.IPAMProvider:
		return clusterctlv1.IPAMProviderType
	}

	return clusterctlv1.ProviderTypeUnknown
}
