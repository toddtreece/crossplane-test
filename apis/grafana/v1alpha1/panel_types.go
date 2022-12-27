/*
Copyright 2022 The Crossplane Authors.

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

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// PanelParameters are the configurable fields of a Panel.
type PanelParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// PanelObservation are the observable fields of a Panel.
type PanelObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A PanelSpec defines the desired state of a Panel.
type PanelSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       PanelParameters `json:"forProvider"`
}

// A PanelStatus represents the observed state of a Panel.
type PanelStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          PanelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Panel is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,toddtest}
type Panel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PanelSpec   `json:"spec"`
	Status PanelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PanelList contains a list of Panel
type PanelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Panel `json:"items"`
}

// Panel type metadata.
var (
	PanelKind             = reflect.TypeOf(Panel{}).Name()
	PanelGroupKind        = schema.GroupKind{Group: Group, Kind: PanelKind}.String()
	PanelKindAPIVersion   = PanelKind + "." + SchemeGroupVersion.String()
	PanelGroupVersionKind = SchemeGroupVersion.WithKind(PanelKind)
)

func init() {
	SchemeBuilder.Register(&Panel{}, &PanelList{})
}
