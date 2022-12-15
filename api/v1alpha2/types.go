// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package v1alpha2

// +k8s:conversion-gen=github.com/siderolabs/cluster-api-bootstrap-provider-talos/api/v1alpha3

// TalosConfigTemplateResource defines the Template structure
type TalosConfigTemplateResource struct {
	Spec TalosConfigSpec `json:"spec,omitempty"`
}
