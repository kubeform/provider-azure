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




type MSSQLOutboundFirewallRuleInitParameters struct {

}


type MSSQLOutboundFirewallRuleObservation struct {


// The SQL Outbound Firewall Rule ID.
ID *string `json:"id,omitempty" tf:"id,omitempty"`

// The resource ID of the SQL Server on which to create the Outbound Firewall Rule. Changing this forces a new resource to be created.
ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`
}


type MSSQLOutboundFirewallRuleParameters struct {


// The resource ID of the SQL Server on which to create the Outbound Firewall Rule. Changing this forces a new resource to be created.
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
}

// MSSQLOutboundFirewallRuleSpec defines the desired state of MSSQLOutboundFirewallRule
type MSSQLOutboundFirewallRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider       MSSQLOutboundFirewallRuleParameters `json:"forProvider"`
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
	InitProvider       MSSQLOutboundFirewallRuleInitParameters `json:"initProvider,omitempty"`
}

// MSSQLOutboundFirewallRuleStatus defines the observed state of MSSQLOutboundFirewallRule.
type MSSQLOutboundFirewallRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider          MSSQLOutboundFirewallRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLOutboundFirewallRule is the Schema for the MSSQLOutboundFirewallRules API. Manages an Azure SQL Outbound Firewall Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MSSQLOutboundFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MSSQLOutboundFirewallRuleSpec   `json:"spec"`
	Status            MSSQLOutboundFirewallRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLOutboundFirewallRuleList contains a list of MSSQLOutboundFirewallRules
type MSSQLOutboundFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSSQLOutboundFirewallRule `json:"items"`
}

// Repository type metadata.
var (
	MSSQLOutboundFirewallRule_Kind             = "MSSQLOutboundFirewallRule"
	MSSQLOutboundFirewallRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MSSQLOutboundFirewallRule_Kind}.String()
	MSSQLOutboundFirewallRule_KindAPIVersion   = MSSQLOutboundFirewallRule_Kind + "." + CRDGroupVersion.String()
	MSSQLOutboundFirewallRule_GroupVersionKind = CRDGroupVersion.WithKind(MSSQLOutboundFirewallRule_Kind)
)

func init() {
	SchemeBuilder.Register(&MSSQLOutboundFirewallRule{}, &MSSQLOutboundFirewallRuleList{})
}
