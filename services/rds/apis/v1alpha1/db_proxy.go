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

// DBProxySpec defines the desired state of DBProxy
type DBProxySpec struct {
	Auth []*UserAuthConfig `json:"auth,omitempty"`
	DBProxyName *string `json:"dbProxyName,omitempty"`
	DebugLogging *bool `json:"debugLogging,omitempty"`
	EngineFamily *string `json:"engineFamily,omitempty"`
	IdleClientTimeout *int64 `json:"idleClientTimeout,omitempty"`
	RequireTLS *bool `json:"requireTLS,omitempty"`
	RoleARN *string `json:"roleARN,omitempty"`
	Tags []*Tag `json:"tags,omitempty"`
	VPCSecurityGroupIDs []*string `json:"vpcSecurityGroupIDs,omitempty"`
	VPCSubnetIDs []*string `json:"vpcSubnetIDs,omitempty"`
}

// DBProxyStatus defines the observed state of DBProxy
type DBProxyStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	CreatedDate *metav1.Time `json:"createdDate,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	Status *string `json:"status,omitempty"`
	UpdatedDate *metav1.Time `json:"updatedDate,omitempty"`
}

// DBProxy is the Schema for the DBProxies API
// +kubebuilder:object:root=true
type DBProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   DBProxySpec   `json:"spec,omitempty"`
	Status DBProxyStatus `json:"status,omitempty"`
}

// DBProxyList contains a list of DBProxy
// +kubebuilder:object:root=true
type DBProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items []DBProxy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DBProxy{}, &DBProxyList{})
}