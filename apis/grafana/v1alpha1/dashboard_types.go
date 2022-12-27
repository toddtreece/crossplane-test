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

// DashboardParameters are the configurable fields of a Dashboard.
type DashboardParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// DashboardObservation are the observable fields of a Dashboard.
type DashboardObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A DashboardSpec defines the desired state of a Dashboard.
type DashboardSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DashboardParameters `json:"forProvider"`
}

// A DashboardStatus represents the observed state of a Dashboard.
type DashboardStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DashboardObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Dashboard is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,toddtest}
type Dashboard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DashboardSpec   `json:"spec"`
	Status DashboardStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DashboardList contains a list of Dashboard
type DashboardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dashboard `json:"items"`
}

// Dashboard type metadata.
var (
	DashboardKind             = reflect.TypeOf(Dashboard{}).Name()
	DashboardGroupKind        = schema.GroupKind{Group: Group, Kind: DashboardKind}.String()
	DashboardKindAPIVersion   = DashboardKind + "." + SchemeGroupVersion.String()
	DashboardGroupVersionKind = SchemeGroupVersion.WithKind(DashboardKind)
)

func init() {
	SchemeBuilder.Register(&Dashboard{}, &DashboardList{})
}
