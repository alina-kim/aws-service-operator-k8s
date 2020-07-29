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

package db_security_group

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
	_ = &svcapitypes.DBSecurityGroup{}
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

	resp, respErr := rm.sdkapi.CreateDBSecurityGroupWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	f3 := []*svcapitypes.EC2SecurityGroup{}
	for _, f3iter := range resp.DBSecurityGroup.EC2SecurityGroups {
		f3elem := &svcapitypes.EC2SecurityGroup{}
		f3elem.EC2SecurityGroupID = f3iter.EC2SecurityGroupId
		f3elem.EC2SecurityGroupName = f3iter.EC2SecurityGroupName
		f3elem.EC2SecurityGroupOwnerID = f3iter.EC2SecurityGroupOwnerId
		f3elem.Status = f3iter.Status
		f3 = append(f3, f3elem)
	}
	ko.Status.EC2SecurityGroups = f3
	f4 := []*svcapitypes.IPRange{}
	for _, f4iter := range resp.DBSecurityGroup.IPRanges {
		f4elem := &svcapitypes.IPRange{}
		f4elem.CIDRIP = f4iter.CIDRIP
		f4elem.Status = f4iter.Status
		f4 = append(f4, f4elem)
	}
	ko.Status.IPRanges = f4
	ko.Status.OwnerID = resp.DBSecurityGroup.OwnerId
	ko.Status.VPCID = resp.DBSecurityGroup.VpcId

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	r *resource,
) (*svcsdk.CreateDBSecurityGroupInput, error) {
	res := &svcsdk.CreateDBSecurityGroupInput{}

	res.SetDBSecurityGroupDescription(*r.ko.Spec.DBSecurityGroupDescription)
	res.SetDBSecurityGroupName(*r.ko.Spec.DBSecurityGroupName)
	f2 := []*svcsdk.Tag{}
	for _, f2iter := range r.ko.Spec.Tags {
		f2elem := &svcsdk.Tag{}
		f2elem.SetKey(*f2iter.Key)
		f2elem.SetValue(*f2iter.Value)
		f2 = append(f2, f2elem)
	}
	res.SetTags(f2)

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	// TODO(jaypipes): Figure this out...
	return nil, nil
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
	_, respErr := rm.sdkapi.DeleteDBSecurityGroupWithContext(ctx, input)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteDBSecurityGroupInput, error) {
	res := &svcsdk.DeleteDBSecurityGroupInput{}

	res.SetDBSecurityGroupName(*r.ko.Spec.DBSecurityGroupName)

	return res, nil
}