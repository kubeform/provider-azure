// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ActiveDirectoryAdministratorInitParameters struct {

	// The login name of the principal to set as the server administrator
	Login *string `json:"login,omitempty" tf:"login,omitempty"`

	// The ID of the principal to set as the server administrator. For a managed identity this should be the Client ID of the identity.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The name of the resource group for the MySQL server. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The name of the MySQL Server on which to set the administrator. Changing this forces a new resource to be created.
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// The Azure Tenant ID
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type ActiveDirectoryAdministratorObservation struct {

	// The ID of the MySQL Active Directory Administrator.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The login name of the principal to set as the server administrator
	Login *string `json:"login,omitempty" tf:"login,omitempty"`

	// The ID of the principal to set as the server administrator. For a managed identity this should be the Client ID of the identity.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The name of the resource group for the MySQL server. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The name of the MySQL Server on which to set the administrator. Changing this forces a new resource to be created.
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// The Azure Tenant ID
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type ActiveDirectoryAdministratorParameters struct {

	// The login name of the principal to set as the server administrator
	// +kubebuilder:validation:Optional
	Login *string `json:"login,omitempty" tf:"login,omitempty"`

	// The ID of the principal to set as the server administrator. For a managed identity this should be the Client ID of the identity.
	// +kubebuilder:validation:Optional
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The name of the resource group for the MySQL server. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The name of the MySQL Server on which to set the administrator. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// The Azure Tenant ID
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

// ActiveDirectoryAdministratorSpec defines the desired state of ActiveDirectoryAdministrator
type ActiveDirectoryAdministratorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ActiveDirectoryAdministratorParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ActiveDirectoryAdministratorInitParameters `json:"initProvider,omitempty"`
}

// ActiveDirectoryAdministratorStatus defines the observed state of ActiveDirectoryAdministrator.
type ActiveDirectoryAdministratorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ActiveDirectoryAdministratorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ActiveDirectoryAdministrator is the Schema for the ActiveDirectoryAdministrators API. Manages an Active Directory administrator on a MySQL server
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ActiveDirectoryAdministrator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.login) || (has(self.initProvider) && has(self.initProvider.login))",message="spec.forProvider.login is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.objectId) || (has(self.initProvider) && has(self.initProvider.objectId))",message="spec.forProvider.objectId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resourceGroupName) || (has(self.initProvider) && has(self.initProvider.resourceGroupName))",message="spec.forProvider.resourceGroupName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serverName) || (has(self.initProvider) && has(self.initProvider.serverName))",message="spec.forProvider.serverName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tenantId) || (has(self.initProvider) && has(self.initProvider.tenantId))",message="spec.forProvider.tenantId is a required parameter"
	Spec   ActiveDirectoryAdministratorSpec   `json:"spec"`
	Status ActiveDirectoryAdministratorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ActiveDirectoryAdministratorList contains a list of ActiveDirectoryAdministrators
type ActiveDirectoryAdministratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ActiveDirectoryAdministrator `json:"items"`
}

// Repository type metadata.
var (
	ActiveDirectoryAdministrator_Kind             = "ActiveDirectoryAdministrator"
	ActiveDirectoryAdministrator_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ActiveDirectoryAdministrator_Kind}.String()
	ActiveDirectoryAdministrator_KindAPIVersion   = ActiveDirectoryAdministrator_Kind + "." + CRDGroupVersion.String()
	ActiveDirectoryAdministrator_GroupVersionKind = CRDGroupVersion.WithKind(ActiveDirectoryAdministrator_Kind)
)

func init() {
	SchemeBuilder.Register(&ActiveDirectoryAdministrator{}, &ActiveDirectoryAdministratorList{})
}
