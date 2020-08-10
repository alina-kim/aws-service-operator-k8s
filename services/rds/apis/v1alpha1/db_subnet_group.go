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
	corev1 "k8s.io/api/core/v1"
)

// Throwaway variable to avoid unused import error
var (
	_ = &corev1.SecretReference{}
)

// DBSubnetGroupSpec defines the desired state of DBSubnetGroup
type DBSubnetGroupSpec struct {
	DBSubnetGroupDescription *string `json:"dbSubnetGroupDescription,omitempty"`
	DBSubnetGroupName *string `json:"dbSubnetGroupName,omitempty"`
	SubnetIDs []*string `json:"subnetIDs,omitempty"`
	Tags []*Tag `json:"tags,omitempty"`
}

// DBSubnetGroupStatus defines the observed state of DBSubnetGroup
type DBSubnetGroupStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	SubnetGroupStatus *string `json:"subnetGroupStatus,omitempty"`
	Subnets []*Subnet `json:"subnets,omitempty"`
	VPCID *string `json:"vpcID,omitempty"`
}

// DBSubnetGroup is the Schema for the DBSubnetGroups API
// +kubebuilder:object:root=true
type DBSubnetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   DBSubnetGroupSpec   `json:"spec,omitempty"`
	Status DBSubnetGroupStatus `json:"status,omitempty"`
}

// DBSubnetGroupList contains a list of DBSubnetGroup
// +kubebuilder:object:root=true
type DBSubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items []DBSubnetGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DBSubnetGroup{}, &DBSubnetGroupList{})
}