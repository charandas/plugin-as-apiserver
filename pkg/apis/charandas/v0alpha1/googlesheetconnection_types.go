/*
Copyright 2023.

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

package v0alpha1

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource/resourcestrategy"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleSheetConnection
// +k8s:openapi-gen=true
type GoogleSheetConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GoogleSheetConnectionSpec   `json:"spec,omitempty"`
	Status GoogleSheetConnectionStatus `json:"status,omitempty"`
}

// GoogleSheetConnectionList
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type GoogleSheetConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []GoogleSheetConnection `json:"items"`
}

// GoogleSheetConnectionSpec defines the desired state of GoogleSheetConnection
type GoogleSheetConnectionSpec struct {
}

var _ resource.Object = &GoogleSheetConnection{}
var _ resourcestrategy.Validater = &GoogleSheetConnection{}

func (in *GoogleSheetConnection) GetObjectMeta() *metav1.ObjectMeta {
	return &in.ObjectMeta
}

func (in *GoogleSheetConnection) NamespaceScoped() bool {
	return false
}

func (in *GoogleSheetConnection) New() runtime.Object {
	return &GoogleSheetConnection{}
}

func (in *GoogleSheetConnection) NewList() runtime.Object {
	return &GoogleSheetConnectionList{}
}

func (in *GoogleSheetConnection) GetGroupVersionResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "charandas.example.com",
		Version:  "v0alpha1",
		Resource: "googlesheetconnections",
	}
}

func (in *GoogleSheetConnection) IsStorageVersion() bool {
	return true
}

func (in *GoogleSheetConnection) Validate(ctx context.Context) field.ErrorList {
	// TODO(user): Modify it, adding your API validation here.
	return nil
}

var _ resource.ObjectList = &GoogleSheetConnectionList{}

func (in *GoogleSheetConnectionList) GetListMeta() *metav1.ListMeta {
	return &in.ListMeta
}

// GoogleSheetConnectionStatus defines the observed state of GoogleSheetConnection
type GoogleSheetConnectionStatus struct {
}

func (in GoogleSheetConnectionStatus) SubResourceName() string {
	return "status"
}

// GoogleSheetConnection implements ObjectWithStatusSubResource interface.
var _ resource.ObjectWithStatusSubResource = &GoogleSheetConnection{}

func (in *GoogleSheetConnection) GetStatus() resource.StatusSubResource {
	return in.Status
}

// GoogleSheetConnectionStatus{} implements StatusSubResource interface.
var _ resource.StatusSubResource = &GoogleSheetConnectionStatus{}

func (in GoogleSheetConnectionStatus) CopyTo(parent resource.ObjectWithStatusSubResource) {
	parent.(*GoogleSheetConnection).Status = in
}

var _ resource.ObjectWithArbitrarySubResource = &GoogleSheetConnection{}

func (in *GoogleSheetConnection) GetArbitrarySubResources() []resource.ArbitrarySubResource {
	return []resource.ArbitrarySubResource{}
}
