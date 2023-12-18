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




type MSSQLServerMicrosoftSupportAuditingPolicyInitParameters struct {


// Whether to enable the extended auditing policy. Possible values are true and false. Defaults to true.
Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

// Enable audit events to Azure Monitor? Defaults to true.
LogMonitoringEnabled *bool `json:"logMonitoringEnabled,omitempty" tf:"log_monitoring_enabled,omitempty"`

// The number of days to retain logs for in the storage account. Defaults to 0.
RetentionInDays *float64 `json:"retentionInDays,omitempty" tf:"retention_in_days,omitempty"`

// Is storage_account_access_key value the storage's secondary key?
StorageAccountAccessKeyIsSecondary *bool `json:"storageAccountAccessKeyIsSecondary,omitempty" tf:"storage_account_access_key_is_secondary,omitempty"`
}


type MSSQLServerMicrosoftSupportAuditingPolicyObservation struct {


// The ID of the SQL database to set the extended auditing policy. Changing this forces a new resource to be created.
DatabaseID *string `json:"databaseId,omitempty" tf:"database_id,omitempty"`

// Whether to enable the extended auditing policy. Possible values are true and false. Defaults to true.
Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

// The ID of the MS SQL Database Extended Auditing Policy.
ID *string `json:"id,omitempty" tf:"id,omitempty"`

// Enable audit events to Azure Monitor? Defaults to true.
LogMonitoringEnabled *bool `json:"logMonitoringEnabled,omitempty" tf:"log_monitoring_enabled,omitempty"`

// The number of days to retain logs for in the storage account. Defaults to 0.
RetentionInDays *float64 `json:"retentionInDays,omitempty" tf:"retention_in_days,omitempty"`

// Is storage_account_access_key value the storage's secondary key?
StorageAccountAccessKeyIsSecondary *bool `json:"storageAccountAccessKeyIsSecondary,omitempty" tf:"storage_account_access_key_is_secondary,omitempty"`

// The blob storage endpoint (e.g. https://example.blob.core.windows.net). This blob storage will hold all extended auditing logs.
StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`
}


type MSSQLServerMicrosoftSupportAuditingPolicyParameters struct {


// The ID of the SQL database to set the extended auditing policy. Changing this forces a new resource to be created.
// +crossplane:generate:reference:type=kubedb.dev/provider-azure/apis/sql/v1alpha1.MSSQLDatabase
// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
// +kubebuilder:validation:Optional
DatabaseID *string `json:"databaseId,omitempty" tf:"database_id,omitempty"`

// Reference to a MSSQLDatabase in sql to populate databaseId.
// +kubebuilder:validation:Optional
DatabaseIDRef *v1.Reference `json:"databaseIdRef,omitempty" tf:"-"`

// Selector for a MSSQLDatabase in sql to populate databaseId.
// +kubebuilder:validation:Optional
DatabaseIDSelector *v1.Selector `json:"databaseIdSelector,omitempty" tf:"-"`

// Whether to enable the extended auditing policy. Possible values are true and false. Defaults to true.
// +kubebuilder:validation:Optional
Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

// Enable audit events to Azure Monitor? Defaults to true.
// +kubebuilder:validation:Optional
LogMonitoringEnabled *bool `json:"logMonitoringEnabled,omitempty" tf:"log_monitoring_enabled,omitempty"`

// The number of days to retain logs for in the storage account. Defaults to 0.
// +kubebuilder:validation:Optional
RetentionInDays *float64 `json:"retentionInDays,omitempty" tf:"retention_in_days,omitempty"`

// Is storage_account_access_key value the storage's secondary key?
// +kubebuilder:validation:Optional
StorageAccountAccessKeyIsSecondary *bool `json:"storageAccountAccessKeyIsSecondary,omitempty" tf:"storage_account_access_key_is_secondary,omitempty"`

// The access key to use for the auditing storage account.
// +kubebuilder:validation:Optional
StorageAccountAccessKeySecretRef *v1.SecretKeySelector `json:"storageAccountAccessKeySecretRef,omitempty" tf:"-"`

// The blob storage endpoint (e.g. https://example.blob.core.windows.net). This blob storage will hold all extended auditing logs.
// +crossplane:generate:reference:type=kubedb.dev/provider-azure/apis/storage/v1alpha1.Account
// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("primary_blob_endpoint",true)
// +kubebuilder:validation:Optional
StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`

// Reference to a Account in storage to populate storageEndpoint.
// +kubebuilder:validation:Optional
StorageEndpointRef *v1.Reference `json:"storageEndpointRef,omitempty" tf:"-"`

// Selector for a Account in storage to populate storageEndpoint.
// +kubebuilder:validation:Optional
StorageEndpointSelector *v1.Selector `json:"storageEndpointSelector,omitempty" tf:"-"`
}

// MSSQLServerMicrosoftSupportAuditingPolicySpec defines the desired state of MSSQLServerMicrosoftSupportAuditingPolicy
type MSSQLServerMicrosoftSupportAuditingPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider       MSSQLServerMicrosoftSupportAuditingPolicyParameters `json:"forProvider"`
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
	InitProvider       MSSQLServerMicrosoftSupportAuditingPolicyInitParameters `json:"initProvider,omitempty"`
}

// MSSQLServerMicrosoftSupportAuditingPolicyStatus defines the observed state of MSSQLServerMicrosoftSupportAuditingPolicy.
type MSSQLServerMicrosoftSupportAuditingPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider          MSSQLServerMicrosoftSupportAuditingPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLServerMicrosoftSupportAuditingPolicy is the Schema for the MSSQLServerMicrosoftSupportAuditingPolicys API. Manages a MS SQL Database Extended Auditing Policy.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MSSQLServerMicrosoftSupportAuditingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MSSQLServerMicrosoftSupportAuditingPolicySpec   `json:"spec"`
	Status            MSSQLServerMicrosoftSupportAuditingPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLServerMicrosoftSupportAuditingPolicyList contains a list of MSSQLServerMicrosoftSupportAuditingPolicys
type MSSQLServerMicrosoftSupportAuditingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSSQLServerMicrosoftSupportAuditingPolicy `json:"items"`
}

// Repository type metadata.
var (
	MSSQLServerMicrosoftSupportAuditingPolicy_Kind             = "MSSQLServerMicrosoftSupportAuditingPolicy"
	MSSQLServerMicrosoftSupportAuditingPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MSSQLServerMicrosoftSupportAuditingPolicy_Kind}.String()
	MSSQLServerMicrosoftSupportAuditingPolicy_KindAPIVersion   = MSSQLServerMicrosoftSupportAuditingPolicy_Kind + "." + CRDGroupVersion.String()
	MSSQLServerMicrosoftSupportAuditingPolicy_GroupVersionKind = CRDGroupVersion.WithKind(MSSQLServerMicrosoftSupportAuditingPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&MSSQLServerMicrosoftSupportAuditingPolicy{}, &MSSQLServerMicrosoftSupportAuditingPolicyList{})
}
