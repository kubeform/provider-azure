//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPolicyInitParameters) DeepCopyInto(out *AccessPolicyInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPolicyInitParameters.
func (in *AccessPolicyInitParameters) DeepCopy() *AccessPolicyInitParameters {
	if in == nil {
		return nil
	}
	out := new(AccessPolicyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPolicyObservation) DeepCopyInto(out *AccessPolicyObservation) {
	*out = *in
	if in.ApplicationID != nil {
		in, out := &in.ApplicationID, &out.ApplicationID
		*out = new(string)
		**out = **in
	}
	if in.CertificatePermissions != nil {
		in, out := &in.CertificatePermissions, &out.CertificatePermissions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.KeyPermissions != nil {
		in, out := &in.KeyPermissions, &out.KeyPermissions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ObjectID != nil {
		in, out := &in.ObjectID, &out.ObjectID
		*out = new(string)
		**out = **in
	}
	if in.SecretPermissions != nil {
		in, out := &in.SecretPermissions, &out.SecretPermissions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.StoragePermissions != nil {
		in, out := &in.StoragePermissions, &out.StoragePermissions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPolicyObservation.
func (in *AccessPolicyObservation) DeepCopy() *AccessPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(AccessPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPolicyParameters) DeepCopyInto(out *AccessPolicyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPolicyParameters.
func (in *AccessPolicyParameters) DeepCopy() *AccessPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(AccessPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutomaticInitParameters) DeepCopyInto(out *AutomaticInitParameters) {
	*out = *in
	if in.TimeAfterCreation != nil {
		in, out := &in.TimeAfterCreation, &out.TimeAfterCreation
		*out = new(string)
		**out = **in
	}
	if in.TimeBeforeExpiry != nil {
		in, out := &in.TimeBeforeExpiry, &out.TimeBeforeExpiry
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutomaticInitParameters.
func (in *AutomaticInitParameters) DeepCopy() *AutomaticInitParameters {
	if in == nil {
		return nil
	}
	out := new(AutomaticInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutomaticObservation) DeepCopyInto(out *AutomaticObservation) {
	*out = *in
	if in.TimeAfterCreation != nil {
		in, out := &in.TimeAfterCreation, &out.TimeAfterCreation
		*out = new(string)
		**out = **in
	}
	if in.TimeBeforeExpiry != nil {
		in, out := &in.TimeBeforeExpiry, &out.TimeBeforeExpiry
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutomaticObservation.
func (in *AutomaticObservation) DeepCopy() *AutomaticObservation {
	if in == nil {
		return nil
	}
	out := new(AutomaticObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutomaticParameters) DeepCopyInto(out *AutomaticParameters) {
	*out = *in
	if in.TimeAfterCreation != nil {
		in, out := &in.TimeAfterCreation, &out.TimeAfterCreation
		*out = new(string)
		**out = **in
	}
	if in.TimeBeforeExpiry != nil {
		in, out := &in.TimeBeforeExpiry, &out.TimeBeforeExpiry
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutomaticParameters.
func (in *AutomaticParameters) DeepCopy() *AutomaticParameters {
	if in == nil {
		return nil
	}
	out := new(AutomaticParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactInitParameters) DeepCopyInto(out *ContactInitParameters) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Phone != nil {
		in, out := &in.Phone, &out.Phone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactInitParameters.
func (in *ContactInitParameters) DeepCopy() *ContactInitParameters {
	if in == nil {
		return nil
	}
	out := new(ContactInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactObservation) DeepCopyInto(out *ContactObservation) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Phone != nil {
		in, out := &in.Phone, &out.Phone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactObservation.
func (in *ContactObservation) DeepCopy() *ContactObservation {
	if in == nil {
		return nil
	}
	out := new(ContactObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactParameters) DeepCopyInto(out *ContactParameters) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Phone != nil {
		in, out := &in.Phone, &out.Phone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactParameters.
func (in *ContactParameters) DeepCopy() *ContactParameters {
	if in == nil {
		return nil
	}
	out := new(ContactParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Key) DeepCopyInto(out *Key) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Key.
func (in *Key) DeepCopy() *Key {
	if in == nil {
		return nil
	}
	out := new(Key)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Key) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyInitParameters) DeepCopyInto(out *KeyInitParameters) {
	*out = *in
	if in.Curve != nil {
		in, out := &in.Curve, &out.Curve
		*out = new(string)
		**out = **in
	}
	if in.ExpirationDate != nil {
		in, out := &in.ExpirationDate, &out.ExpirationDate
		*out = new(string)
		**out = **in
	}
	if in.KeyOpts != nil {
		in, out := &in.KeyOpts, &out.KeyOpts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.KeySize != nil {
		in, out := &in.KeySize, &out.KeySize
		*out = new(float64)
		**out = **in
	}
	if in.KeyType != nil {
		in, out := &in.KeyType, &out.KeyType
		*out = new(string)
		**out = **in
	}
	if in.NotBeforeDate != nil {
		in, out := &in.NotBeforeDate, &out.NotBeforeDate
		*out = new(string)
		**out = **in
	}
	if in.RotationPolicy != nil {
		in, out := &in.RotationPolicy, &out.RotationPolicy
		*out = make([]RotationPolicyInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyInitParameters.
func (in *KeyInitParameters) DeepCopy() *KeyInitParameters {
	if in == nil {
		return nil
	}
	out := new(KeyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyList) DeepCopyInto(out *KeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Key, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyList.
func (in *KeyList) DeepCopy() *KeyList {
	if in == nil {
		return nil
	}
	out := new(KeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyObservation) DeepCopyInto(out *KeyObservation) {
	*out = *in
	if in.Curve != nil {
		in, out := &in.Curve, &out.Curve
		*out = new(string)
		**out = **in
	}
	if in.E != nil {
		in, out := &in.E, &out.E
		*out = new(string)
		**out = **in
	}
	if in.ExpirationDate != nil {
		in, out := &in.ExpirationDate, &out.ExpirationDate
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.KeyOpts != nil {
		in, out := &in.KeyOpts, &out.KeyOpts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.KeySize != nil {
		in, out := &in.KeySize, &out.KeySize
		*out = new(float64)
		**out = **in
	}
	if in.KeyType != nil {
		in, out := &in.KeyType, &out.KeyType
		*out = new(string)
		**out = **in
	}
	if in.KeyVaultID != nil {
		in, out := &in.KeyVaultID, &out.KeyVaultID
		*out = new(string)
		**out = **in
	}
	if in.N != nil {
		in, out := &in.N, &out.N
		*out = new(string)
		**out = **in
	}
	if in.NotBeforeDate != nil {
		in, out := &in.NotBeforeDate, &out.NotBeforeDate
		*out = new(string)
		**out = **in
	}
	if in.PublicKeyOpenssh != nil {
		in, out := &in.PublicKeyOpenssh, &out.PublicKeyOpenssh
		*out = new(string)
		**out = **in
	}
	if in.PublicKeyPem != nil {
		in, out := &in.PublicKeyPem, &out.PublicKeyPem
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.ResourceVersionlessID != nil {
		in, out := &in.ResourceVersionlessID, &out.ResourceVersionlessID
		*out = new(string)
		**out = **in
	}
	if in.RotationPolicy != nil {
		in, out := &in.RotationPolicy, &out.RotationPolicy
		*out = make([]RotationPolicyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.VersionlessID != nil {
		in, out := &in.VersionlessID, &out.VersionlessID
		*out = new(string)
		**out = **in
	}
	if in.X != nil {
		in, out := &in.X, &out.X
		*out = new(string)
		**out = **in
	}
	if in.Y != nil {
		in, out := &in.Y, &out.Y
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyObservation.
func (in *KeyObservation) DeepCopy() *KeyObservation {
	if in == nil {
		return nil
	}
	out := new(KeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyParameters) DeepCopyInto(out *KeyParameters) {
	*out = *in
	if in.Curve != nil {
		in, out := &in.Curve, &out.Curve
		*out = new(string)
		**out = **in
	}
	if in.ExpirationDate != nil {
		in, out := &in.ExpirationDate, &out.ExpirationDate
		*out = new(string)
		**out = **in
	}
	if in.KeyOpts != nil {
		in, out := &in.KeyOpts, &out.KeyOpts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.KeySize != nil {
		in, out := &in.KeySize, &out.KeySize
		*out = new(float64)
		**out = **in
	}
	if in.KeyType != nil {
		in, out := &in.KeyType, &out.KeyType
		*out = new(string)
		**out = **in
	}
	if in.KeyVaultID != nil {
		in, out := &in.KeyVaultID, &out.KeyVaultID
		*out = new(string)
		**out = **in
	}
	if in.KeyVaultIDRef != nil {
		in, out := &in.KeyVaultIDRef, &out.KeyVaultIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyVaultIDSelector != nil {
		in, out := &in.KeyVaultIDSelector, &out.KeyVaultIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.NotBeforeDate != nil {
		in, out := &in.NotBeforeDate, &out.NotBeforeDate
		*out = new(string)
		**out = **in
	}
	if in.RotationPolicy != nil {
		in, out := &in.RotationPolicy, &out.RotationPolicy
		*out = make([]RotationPolicyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyParameters.
func (in *KeyParameters) DeepCopy() *KeyParameters {
	if in == nil {
		return nil
	}
	out := new(KeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeySpec) DeepCopyInto(out *KeySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeySpec.
func (in *KeySpec) DeepCopy() *KeySpec {
	if in == nil {
		return nil
	}
	out := new(KeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyStatus) DeepCopyInto(out *KeyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyStatus.
func (in *KeyStatus) DeepCopy() *KeyStatus {
	if in == nil {
		return nil
	}
	out := new(KeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAclsInitParameters) DeepCopyInto(out *NetworkAclsInitParameters) {
	*out = *in
	if in.Bypass != nil {
		in, out := &in.Bypass, &out.Bypass
		*out = new(string)
		**out = **in
	}
	if in.DefaultAction != nil {
		in, out := &in.DefaultAction, &out.DefaultAction
		*out = new(string)
		**out = **in
	}
	if in.IPRules != nil {
		in, out := &in.IPRules, &out.IPRules
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VirtualNetworkSubnetIds != nil {
		in, out := &in.VirtualNetworkSubnetIds, &out.VirtualNetworkSubnetIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAclsInitParameters.
func (in *NetworkAclsInitParameters) DeepCopy() *NetworkAclsInitParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkAclsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAclsObservation) DeepCopyInto(out *NetworkAclsObservation) {
	*out = *in
	if in.Bypass != nil {
		in, out := &in.Bypass, &out.Bypass
		*out = new(string)
		**out = **in
	}
	if in.DefaultAction != nil {
		in, out := &in.DefaultAction, &out.DefaultAction
		*out = new(string)
		**out = **in
	}
	if in.IPRules != nil {
		in, out := &in.IPRules, &out.IPRules
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VirtualNetworkSubnetIds != nil {
		in, out := &in.VirtualNetworkSubnetIds, &out.VirtualNetworkSubnetIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAclsObservation.
func (in *NetworkAclsObservation) DeepCopy() *NetworkAclsObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkAclsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAclsParameters) DeepCopyInto(out *NetworkAclsParameters) {
	*out = *in
	if in.Bypass != nil {
		in, out := &in.Bypass, &out.Bypass
		*out = new(string)
		**out = **in
	}
	if in.DefaultAction != nil {
		in, out := &in.DefaultAction, &out.DefaultAction
		*out = new(string)
		**out = **in
	}
	if in.IPRules != nil {
		in, out := &in.IPRules, &out.IPRules
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VirtualNetworkSubnetIds != nil {
		in, out := &in.VirtualNetworkSubnetIds, &out.VirtualNetworkSubnetIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAclsParameters.
func (in *NetworkAclsParameters) DeepCopy() *NetworkAclsParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkAclsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RotationPolicyInitParameters) DeepCopyInto(out *RotationPolicyInitParameters) {
	*out = *in
	if in.Automatic != nil {
		in, out := &in.Automatic, &out.Automatic
		*out = make([]AutomaticInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExpireAfter != nil {
		in, out := &in.ExpireAfter, &out.ExpireAfter
		*out = new(string)
		**out = **in
	}
	if in.NotifyBeforeExpiry != nil {
		in, out := &in.NotifyBeforeExpiry, &out.NotifyBeforeExpiry
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RotationPolicyInitParameters.
func (in *RotationPolicyInitParameters) DeepCopy() *RotationPolicyInitParameters {
	if in == nil {
		return nil
	}
	out := new(RotationPolicyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RotationPolicyObservation) DeepCopyInto(out *RotationPolicyObservation) {
	*out = *in
	if in.Automatic != nil {
		in, out := &in.Automatic, &out.Automatic
		*out = make([]AutomaticObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExpireAfter != nil {
		in, out := &in.ExpireAfter, &out.ExpireAfter
		*out = new(string)
		**out = **in
	}
	if in.NotifyBeforeExpiry != nil {
		in, out := &in.NotifyBeforeExpiry, &out.NotifyBeforeExpiry
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RotationPolicyObservation.
func (in *RotationPolicyObservation) DeepCopy() *RotationPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(RotationPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RotationPolicyParameters) DeepCopyInto(out *RotationPolicyParameters) {
	*out = *in
	if in.Automatic != nil {
		in, out := &in.Automatic, &out.Automatic
		*out = make([]AutomaticParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExpireAfter != nil {
		in, out := &in.ExpireAfter, &out.ExpireAfter
		*out = new(string)
		**out = **in
	}
	if in.NotifyBeforeExpiry != nil {
		in, out := &in.NotifyBeforeExpiry, &out.NotifyBeforeExpiry
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RotationPolicyParameters.
func (in *RotationPolicyParameters) DeepCopy() *RotationPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(RotationPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vault) DeepCopyInto(out *Vault) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vault.
func (in *Vault) DeepCopy() *Vault {
	if in == nil {
		return nil
	}
	out := new(Vault)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Vault) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultInitParameters) DeepCopyInto(out *VaultInitParameters) {
	*out = *in
	if in.Contact != nil {
		in, out := &in.Contact, &out.Contact
		*out = make([]ContactInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnableRbacAuthorization != nil {
		in, out := &in.EnableRbacAuthorization, &out.EnableRbacAuthorization
		*out = new(bool)
		**out = **in
	}
	if in.EnabledForDeployment != nil {
		in, out := &in.EnabledForDeployment, &out.EnabledForDeployment
		*out = new(bool)
		**out = **in
	}
	if in.EnabledForDiskEncryption != nil {
		in, out := &in.EnabledForDiskEncryption, &out.EnabledForDiskEncryption
		*out = new(bool)
		**out = **in
	}
	if in.EnabledForTemplateDeployment != nil {
		in, out := &in.EnabledForTemplateDeployment, &out.EnabledForTemplateDeployment
		*out = new(bool)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.NetworkAcls != nil {
		in, out := &in.NetworkAcls, &out.NetworkAcls
		*out = make([]NetworkAclsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PublicNetworkAccessEnabled != nil {
		in, out := &in.PublicNetworkAccessEnabled, &out.PublicNetworkAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.PurgeProtectionEnabled != nil {
		in, out := &in.PurgeProtectionEnabled, &out.PurgeProtectionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.SoftDeleteRetentionDays != nil {
		in, out := &in.SoftDeleteRetentionDays, &out.SoftDeleteRetentionDays
		*out = new(float64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultInitParameters.
func (in *VaultInitParameters) DeepCopy() *VaultInitParameters {
	if in == nil {
		return nil
	}
	out := new(VaultInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultList) DeepCopyInto(out *VaultList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Vault, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultList.
func (in *VaultList) DeepCopy() *VaultList {
	if in == nil {
		return nil
	}
	out := new(VaultList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultObservation) DeepCopyInto(out *VaultObservation) {
	*out = *in
	if in.AccessPolicy != nil {
		in, out := &in.AccessPolicy, &out.AccessPolicy
		*out = make([]AccessPolicyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Contact != nil {
		in, out := &in.Contact, &out.Contact
		*out = make([]ContactObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnableRbacAuthorization != nil {
		in, out := &in.EnableRbacAuthorization, &out.EnableRbacAuthorization
		*out = new(bool)
		**out = **in
	}
	if in.EnabledForDeployment != nil {
		in, out := &in.EnabledForDeployment, &out.EnabledForDeployment
		*out = new(bool)
		**out = **in
	}
	if in.EnabledForDiskEncryption != nil {
		in, out := &in.EnabledForDiskEncryption, &out.EnabledForDiskEncryption
		*out = new(bool)
		**out = **in
	}
	if in.EnabledForTemplateDeployment != nil {
		in, out := &in.EnabledForTemplateDeployment, &out.EnabledForTemplateDeployment
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.NetworkAcls != nil {
		in, out := &in.NetworkAcls, &out.NetworkAcls
		*out = make([]NetworkAclsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PublicNetworkAccessEnabled != nil {
		in, out := &in.PublicNetworkAccessEnabled, &out.PublicNetworkAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.PurgeProtectionEnabled != nil {
		in, out := &in.PurgeProtectionEnabled, &out.PurgeProtectionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.SoftDeleteRetentionDays != nil {
		in, out := &in.SoftDeleteRetentionDays, &out.SoftDeleteRetentionDays
		*out = new(float64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
	if in.VaultURI != nil {
		in, out := &in.VaultURI, &out.VaultURI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultObservation.
func (in *VaultObservation) DeepCopy() *VaultObservation {
	if in == nil {
		return nil
	}
	out := new(VaultObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultParameters) DeepCopyInto(out *VaultParameters) {
	*out = *in
	if in.Contact != nil {
		in, out := &in.Contact, &out.Contact
		*out = make([]ContactParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnableRbacAuthorization != nil {
		in, out := &in.EnableRbacAuthorization, &out.EnableRbacAuthorization
		*out = new(bool)
		**out = **in
	}
	if in.EnabledForDeployment != nil {
		in, out := &in.EnabledForDeployment, &out.EnabledForDeployment
		*out = new(bool)
		**out = **in
	}
	if in.EnabledForDiskEncryption != nil {
		in, out := &in.EnabledForDiskEncryption, &out.EnabledForDiskEncryption
		*out = new(bool)
		**out = **in
	}
	if in.EnabledForTemplateDeployment != nil {
		in, out := &in.EnabledForTemplateDeployment, &out.EnabledForTemplateDeployment
		*out = new(bool)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.NetworkAcls != nil {
		in, out := &in.NetworkAcls, &out.NetworkAcls
		*out = make([]NetworkAclsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PublicNetworkAccessEnabled != nil {
		in, out := &in.PublicNetworkAccessEnabled, &out.PublicNetworkAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.PurgeProtectionEnabled != nil {
		in, out := &in.PurgeProtectionEnabled, &out.PurgeProtectionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.SoftDeleteRetentionDays != nil {
		in, out := &in.SoftDeleteRetentionDays, &out.SoftDeleteRetentionDays
		*out = new(float64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultParameters.
func (in *VaultParameters) DeepCopy() *VaultParameters {
	if in == nil {
		return nil
	}
	out := new(VaultParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultSpec) DeepCopyInto(out *VaultSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultSpec.
func (in *VaultSpec) DeepCopy() *VaultSpec {
	if in == nil {
		return nil
	}
	out := new(VaultSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultStatus) DeepCopyInto(out *VaultStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultStatus.
func (in *VaultStatus) DeepCopy() *VaultStatus {
	if in == nil {
		return nil
	}
	out := new(VaultStatus)
	in.DeepCopyInto(out)
	return out
}
