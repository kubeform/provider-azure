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

type MSSQLFailoverGroupInitParameters struct {

	// A partner_server block as defined below.
	PartnerServer []PartnerServerInitParameters `json:"partnerServer,omitempty" tf:"partner_server,omitempty"`

	// A read_write_endpoint_failover_policy block as defined below.
	ReadWriteEndpointFailoverPolicy []ReadWriteEndpointFailoverPolicyInitParameters `json:"readWriteEndpointFailoverPolicy,omitempty" tf:"read_write_endpoint_failover_policy,omitempty"`

	// Whether failover is enabled for the readonly endpoint. Defaults to false.
	ReadonlyEndpointFailoverPolicyEnabled *bool `json:"readonlyEndpointFailoverPolicyEnabled,omitempty" tf:"readonly_endpoint_failover_policy_enabled,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MSSQLFailoverGroupObservation struct {

	// A set of database names to include in the failover group.
	Databases []*string `json:"databases,omitempty" tf:"databases,omitempty"`

	// The ID of the Failover Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A partner_server block as defined below.
	PartnerServer []PartnerServerObservation `json:"partnerServer,omitempty" tf:"partner_server,omitempty"`

	// A read_write_endpoint_failover_policy block as defined below.
	ReadWriteEndpointFailoverPolicy []ReadWriteEndpointFailoverPolicyObservation `json:"readWriteEndpointFailoverPolicy,omitempty" tf:"read_write_endpoint_failover_policy,omitempty"`

	// Whether failover is enabled for the readonly endpoint. Defaults to false.
	ReadonlyEndpointFailoverPolicyEnabled *bool `json:"readonlyEndpointFailoverPolicyEnabled,omitempty" tf:"readonly_endpoint_failover_policy_enabled,omitempty"`

	// The ID of the primary SQL Server on which to create the failover group. Changing this forces a new resource to be created.
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MSSQLFailoverGroupParameters struct {

	// A set of database names to include in the failover group.
	// +crossplane:generate:reference:type=MSSQLDatabase
	// +crossplane:generate:reference:extractor=kubedb.dev/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Databases []*string `json:"databases,omitempty" tf:"databases,omitempty"`

	// References to MSSQLDatabase to populate databases.
	// +kubebuilder:validation:Optional
	DatabasesRefs []v1.Reference `json:"databasesRefs,omitempty" tf:"-"`

	// Selector for a list of MSSQLDatabase to populate databases.
	// +kubebuilder:validation:Optional
	DatabasesSelector *v1.Selector `json:"databasesSelector,omitempty" tf:"-"`

	// A partner_server block as defined below.
	// +kubebuilder:validation:Optional
	PartnerServer []PartnerServerParameters `json:"partnerServer,omitempty" tf:"partner_server,omitempty"`

	// A read_write_endpoint_failover_policy block as defined below.
	// +kubebuilder:validation:Optional
	ReadWriteEndpointFailoverPolicy []ReadWriteEndpointFailoverPolicyParameters `json:"readWriteEndpointFailoverPolicy,omitempty" tf:"read_write_endpoint_failover_policy,omitempty"`

	// Whether failover is enabled for the readonly endpoint. Defaults to false.
	// +kubebuilder:validation:Optional
	ReadonlyEndpointFailoverPolicyEnabled *bool `json:"readonlyEndpointFailoverPolicyEnabled,omitempty" tf:"readonly_endpoint_failover_policy_enabled,omitempty"`

	// The ID of the primary SQL Server on which to create the failover group. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=MSSQLServer
	// +crossplane:generate:reference:extractor=kubedb.dev/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a MSSQLServer to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a MSSQLServer to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PartnerServerInitParameters struct {
}

type PartnerServerObservation struct {

	// The ID of a partner SQL server to include in the failover group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The location of the partner server.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The replication role of the partner server. Possible values include Primary or Secondary.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type PartnerServerParameters struct {

	// The ID of a partner SQL server to include in the failover group.
	// +crossplane:generate:reference:type=MSSQLServer
	// +crossplane:generate:reference:extractor=kubedb.dev/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a MSSQLServer to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a MSSQLServer to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`
}

type ReadWriteEndpointFailoverPolicyInitParameters struct {

	// The grace period in minutes, before failover with data loss is attempted for the read-write endpoint. Required when mode is Automatic.
	GraceMinutes *float64 `json:"graceMinutes,omitempty" tf:"grace_minutes,omitempty"`

	// The failover policy of the read-write endpoint for the failover group. Possible values are Automatic or Manual.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type ReadWriteEndpointFailoverPolicyObservation struct {

	// The grace period in minutes, before failover with data loss is attempted for the read-write endpoint. Required when mode is Automatic.
	GraceMinutes *float64 `json:"graceMinutes,omitempty" tf:"grace_minutes,omitempty"`

	// The failover policy of the read-write endpoint for the failover group. Possible values are Automatic or Manual.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type ReadWriteEndpointFailoverPolicyParameters struct {

	// The grace period in minutes, before failover with data loss is attempted for the read-write endpoint. Required when mode is Automatic.
	// +kubebuilder:validation:Optional
	GraceMinutes *float64 `json:"graceMinutes,omitempty" tf:"grace_minutes,omitempty"`

	// The failover policy of the read-write endpoint for the failover group. Possible values are Automatic or Manual.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode" tf:"mode,omitempty"`
}

// MSSQLFailoverGroupSpec defines the desired state of MSSQLFailoverGroup
type MSSQLFailoverGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MSSQLFailoverGroupParameters `json:"forProvider"`
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
	InitProvider MSSQLFailoverGroupInitParameters `json:"initProvider,omitempty"`
}

// MSSQLFailoverGroupStatus defines the observed state of MSSQLFailoverGroup.
type MSSQLFailoverGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MSSQLFailoverGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLFailoverGroup is the Schema for the MSSQLFailoverGroups API. Manages a Microsoft Azure SQL Failover Group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MSSQLFailoverGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.partnerServer) || (has(self.initProvider) && has(self.initProvider.partnerServer))",message="spec.forProvider.partnerServer is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.readWriteEndpointFailoverPolicy) || (has(self.initProvider) && has(self.initProvider.readWriteEndpointFailoverPolicy))",message="spec.forProvider.readWriteEndpointFailoverPolicy is a required parameter"
	Spec   MSSQLFailoverGroupSpec   `json:"spec"`
	Status MSSQLFailoverGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLFailoverGroupList contains a list of MSSQLFailoverGroups
type MSSQLFailoverGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSSQLFailoverGroup `json:"items"`
}

// Repository type metadata.
var (
	MSSQLFailoverGroup_Kind             = "MSSQLFailoverGroup"
	MSSQLFailoverGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MSSQLFailoverGroup_Kind}.String()
	MSSQLFailoverGroup_KindAPIVersion   = MSSQLFailoverGroup_Kind + "." + CRDGroupVersion.String()
	MSSQLFailoverGroup_GroupVersionKind = CRDGroupVersion.WithKind(MSSQLFailoverGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&MSSQLFailoverGroup{}, &MSSQLFailoverGroupList{})
}
