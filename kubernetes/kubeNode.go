/*
Copyright 2018 Sysdig.

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

package kubernetes

type KubeNode struct {
	Metadata KubeNodeMetadata `json:"metadata"`
	Spec     KubeNodeSpec     `json:"spec"`
	Status   KubeNodeStatus   `json:"status"`
}

type KubeNodeMetadata struct {
	Name              string `json:"name"`
	SelfLink          string `json:"selfLink"`
	Uid               string `json:"uid"`
	ResourceVersion   string `json:"resourceVersion"`
	CreationTimestamp string `json:"creationTimestamp"`
}

type KubeNodeSpec struct {
	PodCIDR    string `json:"podCIDR"`
	ExternalID string `json:"externalID"`
}

type KubeNodeStatus struct {
	Conditions []KubeNodeStatusConditions `json:"conditions"`
}

type KubeNodeStatusConditions struct {
	Type               string `json:"type"`
	Status             string `json:"status"`
	LastHeartbeatTime  string `json:"lastHeartbeatTime"`
	LastTransitionTime string `json:"lastTransitionTime"`
	Reason             string `json:"reason"`
	Message            string `json:"message"`
}
