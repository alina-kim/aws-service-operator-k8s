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

// DBClusterSpec defines the desired state of DBCluster
type DBClusterSpec struct {
	AvailabilityZones []*string `json:"availabilityZones,omitempty"`
	BacktrackWindow *int64 `json:"backtrackWindow,omitempty"`
	BackupRetentionPeriod *int64 `json:"backupRetentionPeriod,omitempty"`
	CharacterSetName *string `json:"characterSetName,omitempty"`
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot,omitempty"`
	DBClusterIdentifier *string `json:"dbClusterIdentifier,omitempty"`
	DBClusterParameterGroupName *string `json:"dbClusterParameterGroupName,omitempty"`
	DBSubnetGroupName *string `json:"dbSubnetGroupName,omitempty"`
	DatabaseName *string `json:"databaseName,omitempty"`
	DeletionProtection *bool `json:"deletionProtection,omitempty"`
	DestinationRegion *string `json:"destinationRegion,omitempty"`
	Domain *string `json:"domain,omitempty"`
	DomainIAMRoleName *string `json:"domainIAMRoleName,omitempty"`
	EnableCloudwatchLogsExports []*string `json:"enableCloudwatchLogsExports,omitempty"`
	EnableGlobalWriteForwarding *bool `json:"enableGlobalWriteForwarding,omitempty"`
	EnableHTTPEndpoint *bool `json:"enableHTTPEndpoint,omitempty"`
	EnableIAMDatabaseAuthentication *bool `json:"enableIAMDatabaseAuthentication,omitempty"`
	Engine *string `json:"engine,omitempty"`
	EngineMode *string `json:"engineMode,omitempty"`
	EngineVersion *string `json:"engineVersion,omitempty"`
	GlobalClusterIdentifier *string `json:"globalClusterIdentifier,omitempty"`
	KMSKeyID *string `json:"kmsKeyID,omitempty"`
	MasterUserPassword *SecretReference `json:"masterUserPassword,omitempty"`
	MasterUsername *string `json:"masterUsername,omitempty"`
	OptionGroupName *string `json:"optionGroupName,omitempty"`
	Port *int64 `json:"port,omitempty"`
	PreSignedURL *string `json:"preSignedURL,omitempty"`
	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty"`
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty"`
	ReplicationSourceIdentifier *string `json:"replicationSourceIdentifier,omitempty"`
	ScalingConfiguration *ScalingConfiguration `json:"scalingConfiguration,omitempty"`
	SourceRegion *string `json:"sourceRegion,omitempty"`
	StorageEncrypted *bool `json:"storageEncrypted,omitempty"`
	Tags []*Tag `json:"tags,omitempty"`
	VPCSecurityGroupIDs []*string `json:"vpcSecurityGroupIDs,omitempty"`
}

// DBClusterStatus defines the observed state of DBCluster
type DBClusterStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	ActivityStreamKinesisStreamName *string `json:"activityStreamKinesisStreamName,omitempty"`
	ActivityStreamKMSKeyID *string `json:"activityStreamKMSKeyID,omitempty"`
	ActivityStreamMode *string `json:"activityStreamMode,omitempty"`
	ActivityStreamStatus *string `json:"activityStreamStatus,omitempty"`
	AllocatedStorage *int64 `json:"allocatedStorage,omitempty"`
	AssociatedRoles []*DBClusterRole `json:"associatedRoles,omitempty"`
	BacktrackConsumedChangeRecords *int64 `json:"backtrackConsumedChangeRecords,omitempty"`
	Capacity *int64 `json:"capacity,omitempty"`
	CloneGroupID *string `json:"cloneGroupID,omitempty"`
	ClusterCreateTime *metav1.Time `json:"clusterCreateTime,omitempty"`
	CrossAccountClone *bool `json:"crossAccountClone,omitempty"`
	CustomEndpoints []*string `json:"customEndpoints,omitempty"`
	DBClusterMembers []*DBClusterMember `json:"dbClusterMembers,omitempty"`
	DBClusterOptionGroupMemberships []*DBClusterOptionGroupStatus `json:"dbClusterOptionGroupMemberships,omitempty"`
	DBClusterParameterGroup *string `json:"dbClusterParameterGroup,omitempty"`
	DBSubnetGroup *string `json:"dbSubnetGroup,omitempty"`
	DBClusterResourceID *string `json:"dbClusterResourceID,omitempty"`
	DomainMemberships []*DomainMembership `json:"domainMemberships,omitempty"`
	EarliestBacktrackTime *metav1.Time `json:"earliestBacktrackTime,omitempty"`
	EarliestRestorableTime *metav1.Time `json:"earliestRestorableTime,omitempty"`
	EnabledCloudwatchLogsExports []*string `json:"enabledCloudwatchLogsExports,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	GlobalWriteForwardingRequested *bool `json:"globalWriteForwardingRequested,omitempty"`
	GlobalWriteForwardingStatus *string `json:"globalWriteForwardingStatus,omitempty"`
	HostedZoneID *string `json:"hostedZoneID,omitempty"`
	HTTPEndpointEnabled *bool `json:"httpEndpointEnabled,omitempty"`
	IAMDatabaseAuthenticationEnabled *bool `json:"iamDatabaseAuthenticationEnabled,omitempty"`
	LatestRestorableTime *metav1.Time `json:"latestRestorableTime,omitempty"`
	MultiAZ *bool `json:"multiAZ,omitempty"`
	PercentProgress *string `json:"percentProgress,omitempty"`
	ReadReplicaIdentifiers []*string `json:"readReplicaIdentifiers,omitempty"`
	ReaderEndpoint *string `json:"readerEndpoint,omitempty"`
	ScalingConfigurationInfo *ScalingConfigurationInfo `json:"scalingConfigurationInfo,omitempty"`
	Status *string `json:"status,omitempty"`
	VPCSecurityGroups []*VPCSecurityGroupMembership `json:"vpcSecurityGroups,omitempty"`
}

// DBCluster is the Schema for the DBClusters API
// +kubebuilder:object:root=true
type DBCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   DBClusterSpec   `json:"spec,omitempty"`
	Status DBClusterStatus `json:"status,omitempty"`
}

// DBClusterList contains a list of DBCluster
// +kubebuilder:object:root=true
type DBClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items []DBCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DBCluster{}, &DBClusterList{})
}