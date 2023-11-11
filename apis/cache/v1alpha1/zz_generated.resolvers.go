/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	rconfig "kubedb.dev/provider-azure/apis/rconfig"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this RedisLinkedServer.
func (mg *RedisLinkedServer) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LinkedRedisCacheID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.LinkedRedisCacheIDRef,
		Selector:     mg.Spec.ForProvider.LinkedRedisCacheIDSelector,
		To: reference.To{
			List:    &RedisCacheList{},
			Managed: &RedisCache{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LinkedRedisCacheID")
	}
	mg.Spec.ForProvider.LinkedRedisCacheID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LinkedRedisCacheIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TargetRedisCacheName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TargetRedisCacheNameRef,
		Selector:     mg.Spec.ForProvider.TargetRedisCacheNameSelector,
		To: reference.To{
			List:    &RedisCacheList{},
			Managed: &RedisCache{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TargetRedisCacheName")
	}
	mg.Spec.ForProvider.TargetRedisCacheName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TargetRedisCacheNameRef = rsp.ResolvedReference

	return nil
}
