/*
Copyright 2023 The Nephio Authors.

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

package bootstrap

import (
	"strings"

	"github.com/henderiw-nephio/bootstrap-controller/pkg/applicator"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ClusterType string

const (
	ClusterTypeCapi         ClusterType = "capi"
	ClusterTypeNoKubeConfig ClusterType = "none"
)

func getClusterType(secret *corev1.Secret) ClusterType {
	switch string(secret.Type) {
	case "cluster.x-k8s.io/secret":
		if strings.Contains(secret.GetName(), "kubeconfig") {
			return ClusterTypeCapi
		}
	}
	return ClusterTypeNoKubeConfig
}

func getCapiClusterClient(secret *corev1.Secret) (applicator.APIPatchingApplicator, error) {
	//provide a restconfig from the secret value
	config, err := clientcmd.RESTConfigFromKubeConfig(secret.Data["value"])
	if err != nil {
		return applicator.APIPatchingApplicator{}, err
	}
	// build a cluster client from the kube rest config
	clClient, err := client.New(config, client.Options{})
	if err != nil {
		return applicator.APIPatchingApplicator{}, err
	}
	return applicator.NewAPIPatchingApplicator(clClient), nil
}
