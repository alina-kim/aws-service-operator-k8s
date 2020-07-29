// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package db_proxy

import (
	"context"

	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	ackerr "github.com/aws/aws-controllers-k8s/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/rds"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws/aws-controllers-k8s/services/rds/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = &svcsdk.RDS{}
	_ = &svcapitypes.DBProxy{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	// TODO(jaypipes): Map out the ReadMany codepath

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()
	return &resource{ko}, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a new resource with any fields in the Status field filled in
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newCreateRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.CreateDBProxyWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	ko.Status.CreatedDate = &metav1.Time{*resp.DBProxy.CreatedDate}
	ko.Status.Endpoint = resp.DBProxy.Endpoint
	ko.Status.Status = resp.DBProxy.Status
	ko.Status.UpdatedDate = &metav1.Time{*resp.DBProxy.UpdatedDate}

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	r *resource,
) (*svcsdk.CreateDBProxyInput, error) {
	res := &svcsdk.CreateDBProxyInput{}

	f0 := []*svcsdk.UserAuthConfig{}
	for _, f0iter := range r.ko.Spec.Auth {
		f0elem := &svcsdk.UserAuthConfig{}
		f0elem.SetAuthScheme(*f0iter.AuthScheme)
		f0elem.SetDescription(*f0iter.Description)
		f0elem.SetIAMAuth(*f0iter.IAMAuth)
		f0elem.SetSecretArn(*f0iter.SecretARN)
		f0elem.SetUserName(*f0iter.UserName)
		f0 = append(f0, f0elem)
	}
	res.SetAuth(f0)
	res.SetDBProxyName(*r.ko.Spec.DBProxyName)
	res.SetDebugLogging(*r.ko.Spec.DebugLogging)
	res.SetEngineFamily(*r.ko.Spec.EngineFamily)
	res.SetIdleClientTimeout(*r.ko.Spec.IdleClientTimeout)
	res.SetRequireTLS(*r.ko.Spec.RequireTLS)
	res.SetRoleArn(*r.ko.Spec.RoleARN)
	f7 := []*svcsdk.Tag{}
	for _, f7iter := range r.ko.Spec.Tags {
		f7elem := &svcsdk.Tag{}
		f7elem.SetKey(*f7iter.Key)
		f7elem.SetValue(*f7iter.Value)
		f7 = append(f7, f7elem)
	}
	res.SetTags(f7)
	f8 := []*string{}
	for _, f8iter := range r.ko.Spec.VPCSecurityGroupIDs {
		var f8elem string
		f8elem = *f8iter
		f8 = append(f8, &f8elem)
	}
	res.SetVpcSecurityGroupIds(f8)
	f9 := []*string{}
	for _, f9iter := range r.ko.Spec.VPCSubnetIDs {
		var f9elem string
		f9elem = *f9iter
		f9 = append(f9, &f9elem)
	}
	res.SetVpcSubnetIds(f9)

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newUpdateRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.ModifyDBProxyWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	ko.Status.CreatedDate = &metav1.Time{*resp.DBProxy.CreatedDate}
	ko.Status.Endpoint = resp.DBProxy.Endpoint
	ko.Status.Status = resp.DBProxy.Status
	ko.Status.UpdatedDate = &metav1.Time{*resp.DBProxy.UpdatedDate}

	return &resource{ko}, nil
}
// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	r *resource,
) (*svcsdk.ModifyDBProxyInput, error) {
	res := &svcsdk.ModifyDBProxyInput{}

	f0 := []*svcsdk.UserAuthConfig{}
	for _, f0iter := range r.ko.Spec.Auth {
		f0elem := &svcsdk.UserAuthConfig{}
		f0elem.SetAuthScheme(*f0iter.AuthScheme)
		f0elem.SetDescription(*f0iter.Description)
		f0elem.SetIAMAuth(*f0iter.IAMAuth)
		f0elem.SetSecretArn(*f0iter.SecretARN)
		f0elem.SetUserName(*f0iter.UserName)
		f0 = append(f0, f0elem)
	}
	res.SetAuth(f0)
	res.SetDBProxyName(*r.ko.Spec.DBProxyName)
	res.SetDebugLogging(*r.ko.Spec.DebugLogging)
	res.SetIdleClientTimeout(*r.ko.Spec.IdleClientTimeout)
	res.SetRequireTLS(*r.ko.Spec.RequireTLS)
	res.SetRoleArn(*r.ko.Spec.RoleARN)

	return res, nil
}


// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) error {
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	_, respErr := rm.sdkapi.DeleteDBProxyWithContext(ctx, input)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteDBProxyInput, error) {
	res := &svcsdk.DeleteDBProxyInput{}

	res.SetDBProxyName(*r.ko.Spec.DBProxyName)

	return res, nil
}