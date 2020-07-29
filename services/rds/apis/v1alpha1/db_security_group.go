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

package v1alpha1

import (
	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DBSecurityGroupSpec defines the desired state of DBSecurityGroup
type DBSecurityGroupSpec struct {
	DBSecurityGroupDescription *string `json:"dbSecurityGroupDescription,omitempty"`
	DBSecurityGroupName *string `json:"dbSecurityGroupName,omitempty"`
	Tags []*Tag `json:"tags,omitempty"`
}

// DBSecurityGroupStatus defines the observed state of DBSecurityGroup
type DBSecurityGroupStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	EC2SecurityGroups []*EC2SecurityGroup `json:"ec2SecurityGroups,omitempty"`
	IPRanges []*IPRange `json:"ipRanges,omitempty"`
	OwnerID *string `json:"ownerID,omitempty"`
	VPCID *string `json:"vpcID,omitempty"`
}

// DBSecurityGroup is the Schema for the DBSecurityGroups API
// +kubebuilder:object:root=true
type DBSecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   DBSecurityGroupSpec   `json:"spec,omitempty"`
	Status DBSecurityGroupStatus `json:"status,omitempty"`
}

// DBSecurityGroupList contains a list of DBSecurityGroup
// +kubebuilder:object:root=true
type DBSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items []DBSecurityGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DBSecurityGroup{}, &DBSecurityGroupList{})
}