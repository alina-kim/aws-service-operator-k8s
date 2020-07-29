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

package db_cluster_endpoint

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
	_ = &svcapitypes.DBClusterEndpoint{}
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

	resp, respErr := rm.sdkapi.CreateDBClusterEndpointWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	ko.Status.CustomEndpointType = resp.CustomEndpointType
	ko.Status.DBClusterEndpointResourceIdentifier = resp.DBClusterEndpointResourceIdentifier
	ko.Status.Endpoint = resp.Endpoint
	ko.Status.Status = resp.Status

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	r *resource,
) (*svcsdk.CreateDBClusterEndpointInput, error) {
	res := &svcsdk.CreateDBClusterEndpointInput{}

	res.SetDBClusterEndpointIdentifier(*r.ko.Spec.DBClusterEndpointIdentifier)
	res.SetDBClusterIdentifier(*r.ko.Spec.DBClusterIdentifier)
	res.SetEndpointType(*r.ko.Spec.EndpointType)
	f3 := []*string{}
	for _, f3iter := range r.ko.Spec.ExcludedMembers {
		var f3elem string
		f3elem = *f3iter
		f3 = append(f3, &f3elem)
	}
	res.SetExcludedMembers(f3)
	f4 := []*string{}
	for _, f4iter := range r.ko.Spec.StaticMembers {
		var f4elem string
		f4elem = *f4iter
		f4 = append(f4, &f4elem)
	}
	res.SetStaticMembers(f4)
	f5 := []*svcsdk.Tag{}
	for _, f5iter := range r.ko.Spec.Tags {
		f5elem := &svcsdk.Tag{}
		f5elem.SetKey(*f5iter.Key)
		f5elem.SetValue(*f5iter.Value)
		f5 = append(f5, f5elem)
	}
	res.SetTags(f5)

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

	resp, respErr := rm.sdkapi.ModifyDBClusterEndpointWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	ko.Status.CustomEndpointType = resp.CustomEndpointType
	ko.Status.DBClusterEndpointResourceIdentifier = resp.DBClusterEndpointResourceIdentifier
	ko.Status.Endpoint = resp.Endpoint
	ko.Status.Status = resp.Status

	return &resource{ko}, nil
}
// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	r *resource,
) (*svcsdk.ModifyDBClusterEndpointInput, error) {
	res := &svcsdk.ModifyDBClusterEndpointInput{}

	res.SetDBClusterEndpointIdentifier(*r.ko.Spec.DBClusterEndpointIdentifier)
	res.SetEndpointType(*r.ko.Spec.EndpointType)
	f2 := []*string{}
	for _, f2iter := range r.ko.Spec.ExcludedMembers {
		var f2elem string
		f2elem = *f2iter
		f2 = append(f2, &f2elem)
	}
	res.SetExcludedMembers(f2)
	f3 := []*string{}
	for _, f3iter := range r.ko.Spec.StaticMembers {
		var f3elem string
		f3elem = *f3iter
		f3 = append(f3, &f3elem)
	}
	res.SetStaticMembers(f3)

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
	_, respErr := rm.sdkapi.DeleteDBClusterEndpointWithContext(ctx, input)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteDBClusterEndpointInput, error) {
	res := &svcsdk.DeleteDBClusterEndpointInput{}

	res.SetDBClusterEndpointIdentifier(*r.ko.Spec.DBClusterEndpointIdentifier)

	return res, nil
}