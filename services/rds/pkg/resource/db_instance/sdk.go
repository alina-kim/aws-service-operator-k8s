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

package db_instance

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
	_ = &svcapitypes.DBInstance{}
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

	resp, respErr := rm.sdkapi.CreateDBInstanceWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	f1 := []*svcapitypes.DBInstanceRole{}
	for _, f1iter := range resp.DBInstance.AssociatedRoles {
		f1elem := &svcapitypes.DBInstanceRole{}
		f1elem.FeatureName = f1iter.FeatureName
		f1elem.RoleARN = f1iter.RoleArn
		f1elem.Status = f1iter.Status
		f1 = append(f1, f1elem)
	}
	ko.Status.AssociatedRoles = f1
	ko.Status.CACertificateIdentifier = resp.DBInstance.CACertificateIdentifier
	ko.Status.DBInstanceStatus = resp.DBInstance.DBInstanceStatus
	f14 := []*svcapitypes.DBParameterGroupStatus_SDK{}
	for _, f14iter := range resp.DBInstance.DBParameterGroups {
		f14elem := &svcapitypes.DBParameterGroupStatus_SDK{}
		f14elem.DBParameterGroupName = f14iter.DBParameterGroupName
		f14elem.ParameterApplyStatus = f14iter.ParameterApplyStatus
		f14 = append(f14, f14elem)
	}
	ko.Status.DBParameterGroups = f14
	f16 := &svcapitypes.DBSubnetGroup_SDK{}
	f16.DBSubnetGroupARN = resp.DBInstance.DBSubnetGroup.DBSubnetGroupArn
	f16.DBSubnetGroupDescription = resp.DBInstance.DBSubnetGroup.DBSubnetGroupDescription
	f16.DBSubnetGroupName = resp.DBInstance.DBSubnetGroup.DBSubnetGroupName
	f16.SubnetGroupStatus = resp.DBInstance.DBSubnetGroup.SubnetGroupStatus
	f16f4 := []*svcapitypes.Subnet{}
	for _, f16f4iter := range resp.DBInstance.DBSubnetGroup.Subnets {
		f16f4elem := &svcapitypes.Subnet{}
		f16f4elemf0 := &svcapitypes.AvailabilityZone{}
		f16f4elemf0.Name = f16f4iter.SubnetAvailabilityZone.Name
		f16f4elem.SubnetAvailabilityZone = f16f4elemf0
		f16f4elem.SubnetIdentifier = f16f4iter.SubnetIdentifier
		f16f4elemf2 := &svcapitypes.Outpost{}
		f16f4elemf2.ARN = f16f4iter.SubnetOutpost.Arn
		f16f4elem.SubnetOutpost = f16f4elemf2
		f16f4elem.SubnetStatus = f16f4iter.SubnetStatus
		f16f4 = append(f16f4, f16f4elem)
	}
	f16.Subnets = f16f4
	f16.VPCID = resp.DBInstance.DBSubnetGroup.VpcId
	ko.Status.DBSubnetGroup = f16
	ko.Status.DBInstancePort = resp.DBInstance.DbInstancePort
	ko.Status.DBIResourceID = resp.DBInstance.DbiResourceId
	f20 := []*svcapitypes.DomainMembership{}
	for _, f20iter := range resp.DBInstance.DomainMemberships {
		f20elem := &svcapitypes.DomainMembership{}
		f20elem.Domain = f20iter.Domain
		f20elem.FQDN = f20iter.FQDN
		f20elem.IAMRoleName = f20iter.IAMRoleName
		f20elem.Status = f20iter.Status
		f20 = append(f20, f20elem)
	}
	ko.Status.DomainMemberships = f20
	f21 := []*string{}
	for _, f21iter := range resp.DBInstance.EnabledCloudwatchLogsExports {
		var f21elem string
		f21elem = *f21iter
		f21 = append(f21, &f21elem)
	}
	ko.Status.EnabledCloudwatchLogsExports = f21
	f22 := &svcapitypes.Endpoint{}
	f22.Address = resp.DBInstance.Endpoint.Address
	f22.HostedZoneID = resp.DBInstance.Endpoint.HostedZoneId
	f22.Port = resp.DBInstance.Endpoint.Port
	ko.Status.Endpoint = f22
	ko.Status.EnhancedMonitoringResourceARN = resp.DBInstance.EnhancedMonitoringResourceArn
	ko.Status.IAMDatabaseAuthenticationEnabled = resp.DBInstance.IAMDatabaseAuthenticationEnabled
	ko.Status.InstanceCreateTime = &metav1.Time{*resp.DBInstance.InstanceCreateTime}
	ko.Status.LatestRestorableTime = &metav1.Time{*resp.DBInstance.LatestRestorableTime}
	f32 := &svcapitypes.Endpoint{}
	f32.Address = resp.DBInstance.ListenerEndpoint.Address
	f32.HostedZoneID = resp.DBInstance.ListenerEndpoint.HostedZoneId
	f32.Port = resp.DBInstance.ListenerEndpoint.Port
	ko.Status.ListenerEndpoint = f32
	f38 := []*svcapitypes.OptionGroupMembership{}
	for _, f38iter := range resp.DBInstance.OptionGroupMemberships {
		f38elem := &svcapitypes.OptionGroupMembership{}
		f38elem.OptionGroupName = f38iter.OptionGroupName
		f38elem.Status = f38iter.Status
		f38 = append(f38, f38elem)
	}
	ko.Status.OptionGroupMemberships = f38
	f39 := &svcapitypes.PendingModifiedValues{}
	f39.AllocatedStorage = resp.DBInstance.PendingModifiedValues.AllocatedStorage
	f39.BackupRetentionPeriod = resp.DBInstance.PendingModifiedValues.BackupRetentionPeriod
	f39.CACertificateIdentifier = resp.DBInstance.PendingModifiedValues.CACertificateIdentifier
	f39.DBInstanceClass = resp.DBInstance.PendingModifiedValues.DBInstanceClass
	f39.DBInstanceIdentifier = resp.DBInstance.PendingModifiedValues.DBInstanceIdentifier
	f39.DBSubnetGroupName = resp.DBInstance.PendingModifiedValues.DBSubnetGroupName
	f39.EngineVersion = resp.DBInstance.PendingModifiedValues.EngineVersion
	f39.IOPS = resp.DBInstance.PendingModifiedValues.Iops
	f39.LicenseModel = resp.DBInstance.PendingModifiedValues.LicenseModel
	f39.MasterUserPassword = resp.DBInstance.PendingModifiedValues.MasterUserPassword
	f39.MultiAZ = resp.DBInstance.PendingModifiedValues.MultiAZ
	f39f11 := &svcapitypes.PendingCloudwatchLogsExports{}
	f39f11f0 := []*string{}
	for _, f39f11f0iter := range resp.DBInstance.PendingModifiedValues.PendingCloudwatchLogsExports.LogTypesToDisable {
		var f39f11f0elem string
		f39f11f0elem = *f39f11f0iter
		f39f11f0 = append(f39f11f0, &f39f11f0elem)
	}
	f39f11.LogTypesToDisable = f39f11f0
	f39f11f1 := []*string{}
	for _, f39f11f1iter := range resp.DBInstance.PendingModifiedValues.PendingCloudwatchLogsExports.LogTypesToEnable {
		var f39f11f1elem string
		f39f11f1elem = *f39f11f1iter
		f39f11f1 = append(f39f11f1, &f39f11f1elem)
	}
	f39f11.LogTypesToEnable = f39f11f1
	f39.PendingCloudwatchLogsExports = f39f11
	f39.Port = resp.DBInstance.PendingModifiedValues.Port
	f39f13 := []*svcapitypes.ProcessorFeature{}
	for _, f39f13iter := range resp.DBInstance.PendingModifiedValues.ProcessorFeatures {
		f39f13elem := &svcapitypes.ProcessorFeature{}
		f39f13elem.Name = f39f13iter.Name
		f39f13elem.Value = f39f13iter.Value
		f39f13 = append(f39f13, f39f13elem)
	}
	f39.ProcessorFeatures = f39f13
	f39.StorageType = resp.DBInstance.PendingModifiedValues.StorageType
	ko.Status.PendingModifiedValues = f39
	ko.Status.PerformanceInsightsEnabled = resp.DBInstance.PerformanceInsightsEnabled
	f48 := []*string{}
	for _, f48iter := range resp.DBInstance.ReadReplicaDBClusterIdentifiers {
		var f48elem string
		f48elem = *f48iter
		f48 = append(f48, &f48elem)
	}
	ko.Status.ReadReplicaDBClusterIdentifiers = f48
	f49 := []*string{}
	for _, f49iter := range resp.DBInstance.ReadReplicaDBInstanceIdentifiers {
		var f49elem string
		f49elem = *f49iter
		f49 = append(f49, &f49elem)
	}
	ko.Status.ReadReplicaDBInstanceIdentifiers = f49
	ko.Status.ReadReplicaSourceDBInstanceIdentifier = resp.DBInstance.ReadReplicaSourceDBInstanceIdentifier
	ko.Status.SecondaryAvailabilityZone = resp.DBInstance.SecondaryAvailabilityZone
	f52 := []*svcapitypes.DBInstanceStatusInfo{}
	for _, f52iter := range resp.DBInstance.StatusInfos {
		f52elem := &svcapitypes.DBInstanceStatusInfo{}
		f52elem.Message = f52iter.Message
		f52elem.Normal = f52iter.Normal
		f52elem.Status = f52iter.Status
		f52elem.StatusType = f52iter.StatusType
		f52 = append(f52, f52elem)
	}
	ko.Status.StatusInfos = f52
	f57 := []*svcapitypes.VpcSecurityGroupMembership{}
	for _, f57iter := range resp.DBInstance.VpcSecurityGroups {
		f57elem := &svcapitypes.VpcSecurityGroupMembership{}
		f57elem.Status = f57iter.Status
		f57elem.VPCSecurityGroupID = f57iter.VpcSecurityGroupId
		f57 = append(f57, f57elem)
	}
	ko.Status.VPCSecurityGroups = f57

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	r *resource,
) (*svcsdk.CreateDBInstanceInput, error) {
	res := &svcsdk.CreateDBInstanceInput{}

	res.SetAllocatedStorage(*r.ko.Spec.AllocatedStorage)
	res.SetAutoMinorVersionUpgrade(*r.ko.Spec.AutoMinorVersionUpgrade)
	res.SetAvailabilityZone(*r.ko.Spec.AvailabilityZone)
	res.SetBackupRetentionPeriod(*r.ko.Spec.BackupRetentionPeriod)
	res.SetCharacterSetName(*r.ko.Spec.CharacterSetName)
	res.SetCopyTagsToSnapshot(*r.ko.Spec.CopyTagsToSnapshot)
	res.SetDBClusterIdentifier(*r.ko.Spec.DBClusterIdentifier)
	res.SetDBInstanceClass(*r.ko.Spec.DBInstanceClass)
	res.SetDBInstanceIdentifier(*r.ko.Spec.DBInstanceIdentifier)
	res.SetDBName(*r.ko.Spec.DBName)
	res.SetDBParameterGroupName(*r.ko.Spec.DBParameterGroupName)
	f11 := []*string{}
	for _, f11iter := range r.ko.Spec.DBSecurityGroups {
		var f11elem string
		f11elem = *f11iter
		f11 = append(f11, &f11elem)
	}
	res.SetDBSecurityGroups(f11)
	res.SetDBSubnetGroupName(*r.ko.Spec.DBSubnetGroupName)
	res.SetDeletionProtection(*r.ko.Spec.DeletionProtection)
	res.SetDomain(*r.ko.Spec.Domain)
	res.SetDomainIAMRoleName(*r.ko.Spec.DomainIAMRoleName)
	f16 := []*string{}
	for _, f16iter := range r.ko.Spec.EnableCloudwatchLogsExports {
		var f16elem string
		f16elem = *f16iter
		f16 = append(f16, &f16elem)
	}
	res.SetEnableCloudwatchLogsExports(f16)
	res.SetEnableIAMDatabaseAuthentication(*r.ko.Spec.EnableIAMDatabaseAuthentication)
	res.SetEnablePerformanceInsights(*r.ko.Spec.EnablePerformanceInsights)
	res.SetEngine(*r.ko.Spec.Engine)
	res.SetEngineVersion(*r.ko.Spec.EngineVersion)
	res.SetIops(*r.ko.Spec.IOPS)
	res.SetKmsKeyId(*r.ko.Spec.KMSKeyID)
	res.SetLicenseModel(*r.ko.Spec.LicenseModel)
	f24 := string
	f24elem = *r.ko.Spec.MasterUserPassword
	log.Print(f24elem.Name)
	secretName := f24elem.Name
	secret, err := rm.client.secrets.Get(secretName)
	if err != nil {
 		return err
	}
	f24, err = b64.StdEncoding.DecodeString(secret)
 	if err != nil {
 		return err
	}
	res.SetMasterUserPassword(f24)
	res.SetMasterUsername(*r.ko.Spec.MasterUsername)
	res.SetMaxAllocatedStorage(*r.ko.Spec.MaxAllocatedStorage)
	res.SetMonitoringInterval(*r.ko.Spec.MonitoringInterval)
	res.SetMonitoringRoleArn(*r.ko.Spec.MonitoringRoleARN)
	res.SetMultiAZ(*r.ko.Spec.MultiAZ)
	res.SetOptionGroupName(*r.ko.Spec.OptionGroupName)
	res.SetPerformanceInsightsKMSKeyId(*r.ko.Spec.PerformanceInsightsKMSKeyID)
	res.SetPerformanceInsightsRetentionPeriod(*r.ko.Spec.PerformanceInsightsRetentionPeriod)
	res.SetPort(*r.ko.Spec.Port)
	res.SetPreferredBackupWindow(*r.ko.Spec.PreferredBackupWindow)
	res.SetPreferredMaintenanceWindow(*r.ko.Spec.PreferredMaintenanceWindow)
	f36 := []*svcsdk.ProcessorFeature{}
	for _, f36iter := range r.ko.Spec.ProcessorFeatures {
		f36elem := &svcsdk.ProcessorFeature{}
		f36elem.SetName(*f36iter.Name)
		f36elem.SetValue(*f36iter.Value)
		f36 = append(f36, f36elem)
	}
	res.SetProcessorFeatures(f36)
	res.SetPromotionTier(*r.ko.Spec.PromotionTier)
	res.SetPubliclyAccessible(*r.ko.Spec.PubliclyAccessible)
	res.SetStorageEncrypted(*r.ko.Spec.StorageEncrypted)
	res.SetStorageType(*r.ko.Spec.StorageType)
	f41 := []*svcsdk.Tag{}
	for _, f41iter := range r.ko.Spec.Tags {
		f41elem := &svcsdk.Tag{}
		f41elem.SetKey(*f41iter.Key)
		f41elem.SetValue(*f41iter.Value)
		f41 = append(f41, f41elem)
	}
	res.SetTags(f41)
	res.SetTdeCredentialArn(*r.ko.Spec.TDECredentialARN)
	res.SetTdeCredentialPassword(*r.ko.Spec.TDECredentialPassword)
	res.SetTimezone(*r.ko.Spec.Timezone)
	f45 := []*string{}
	for _, f45iter := range r.ko.Spec.VPCSecurityGroupIDs {
		var f45elem string
		f45elem = *f45iter
		f45 = append(f45, &f45elem)
	}
	res.SetVpcSecurityGroupIds(f45)

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

	resp, respErr := rm.sdkapi.ModifyDBInstanceWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	f1 := []*svcapitypes.DBInstanceRole{}
	for _, f1iter := range resp.DBInstance.AssociatedRoles {
		f1elem := &svcapitypes.DBInstanceRole{}
		f1elem.FeatureName = f1iter.FeatureName
		f1elem.RoleARN = f1iter.RoleArn
		f1elem.Status = f1iter.Status
		f1 = append(f1, f1elem)
	}
	ko.Status.AssociatedRoles = f1
	ko.Status.CACertificateIdentifier = resp.DBInstance.CACertificateIdentifier
	ko.Status.DBInstanceStatus = resp.DBInstance.DBInstanceStatus
	f14 := []*svcapitypes.DBParameterGroupStatus_SDK{}
	for _, f14iter := range resp.DBInstance.DBParameterGroups {
		f14elem := &svcapitypes.DBParameterGroupStatus_SDK{}
		f14elem.DBParameterGroupName = f14iter.DBParameterGroupName
		f14elem.ParameterApplyStatus = f14iter.ParameterApplyStatus
		f14 = append(f14, f14elem)
	}
	ko.Status.DBParameterGroups = f14
	f16 := &svcapitypes.DBSubnetGroup_SDK{}
	f16.DBSubnetGroupARN = resp.DBInstance.DBSubnetGroup.DBSubnetGroupArn
	f16.DBSubnetGroupDescription = resp.DBInstance.DBSubnetGroup.DBSubnetGroupDescription
	f16.DBSubnetGroupName = resp.DBInstance.DBSubnetGroup.DBSubnetGroupName
	f16.SubnetGroupStatus = resp.DBInstance.DBSubnetGroup.SubnetGroupStatus
	f16f4 := []*svcapitypes.Subnet{}
	for _, f16f4iter := range resp.DBInstance.DBSubnetGroup.Subnets {
		f16f4elem := &svcapitypes.Subnet{}
		f16f4elemf0 := &svcapitypes.AvailabilityZone{}
		f16f4elemf0.Name = f16f4iter.SubnetAvailabilityZone.Name
		f16f4elem.SubnetAvailabilityZone = f16f4elemf0
		f16f4elem.SubnetIdentifier = f16f4iter.SubnetIdentifier
		f16f4elemf2 := &svcapitypes.Outpost{}
		f16f4elemf2.ARN = f16f4iter.SubnetOutpost.Arn
		f16f4elem.SubnetOutpost = f16f4elemf2
		f16f4elem.SubnetStatus = f16f4iter.SubnetStatus
		f16f4 = append(f16f4, f16f4elem)
	}
	f16.Subnets = f16f4
	f16.VPCID = resp.DBInstance.DBSubnetGroup.VpcId
	ko.Status.DBSubnetGroup = f16
	ko.Status.DBInstancePort = resp.DBInstance.DbInstancePort
	ko.Status.DBIResourceID = resp.DBInstance.DbiResourceId
	f20 := []*svcapitypes.DomainMembership{}
	for _, f20iter := range resp.DBInstance.DomainMemberships {
		f20elem := &svcapitypes.DomainMembership{}
		f20elem.Domain = f20iter.Domain
		f20elem.FQDN = f20iter.FQDN
		f20elem.IAMRoleName = f20iter.IAMRoleName
		f20elem.Status = f20iter.Status
		f20 = append(f20, f20elem)
	}
	ko.Status.DomainMemberships = f20
	f21 := []*string{}
	for _, f21iter := range resp.DBInstance.EnabledCloudwatchLogsExports {
		var f21elem string
		f21elem = *f21iter
		f21 = append(f21, &f21elem)
	}
	ko.Status.EnabledCloudwatchLogsExports = f21
	f22 := &svcapitypes.Endpoint{}
	f22.Address = resp.DBInstance.Endpoint.Address
	f22.HostedZoneID = resp.DBInstance.Endpoint.HostedZoneId
	f22.Port = resp.DBInstance.Endpoint.Port
	ko.Status.Endpoint = f22
	ko.Status.EnhancedMonitoringResourceARN = resp.DBInstance.EnhancedMonitoringResourceArn
	ko.Status.IAMDatabaseAuthenticationEnabled = resp.DBInstance.IAMDatabaseAuthenticationEnabled
	ko.Status.InstanceCreateTime = &metav1.Time{*resp.DBInstance.InstanceCreateTime}
	ko.Status.LatestRestorableTime = &metav1.Time{*resp.DBInstance.LatestRestorableTime}
	f32 := &svcapitypes.Endpoint{}
	f32.Address = resp.DBInstance.ListenerEndpoint.Address
	f32.HostedZoneID = resp.DBInstance.ListenerEndpoint.HostedZoneId
	f32.Port = resp.DBInstance.ListenerEndpoint.Port
	ko.Status.ListenerEndpoint = f32
	f38 := []*svcapitypes.OptionGroupMembership{}
	for _, f38iter := range resp.DBInstance.OptionGroupMemberships {
		f38elem := &svcapitypes.OptionGroupMembership{}
		f38elem.OptionGroupName = f38iter.OptionGroupName
		f38elem.Status = f38iter.Status
		f38 = append(f38, f38elem)
	}
	ko.Status.OptionGroupMemberships = f38
	f39 := &svcapitypes.PendingModifiedValues{}
	f39.AllocatedStorage = resp.DBInstance.PendingModifiedValues.AllocatedStorage
	f39.BackupRetentionPeriod = resp.DBInstance.PendingModifiedValues.BackupRetentionPeriod
	f39.CACertificateIdentifier = resp.DBInstance.PendingModifiedValues.CACertificateIdentifier
	f39.DBInstanceClass = resp.DBInstance.PendingModifiedValues.DBInstanceClass
	f39.DBInstanceIdentifier = resp.DBInstance.PendingModifiedValues.DBInstanceIdentifier
	f39.DBSubnetGroupName = resp.DBInstance.PendingModifiedValues.DBSubnetGroupName
	f39.EngineVersion = resp.DBInstance.PendingModifiedValues.EngineVersion
	f39.IOPS = resp.DBInstance.PendingModifiedValues.Iops
	f39.LicenseModel = resp.DBInstance.PendingModifiedValues.LicenseModel
	f39.MasterUserPassword = resp.DBInstance.PendingModifiedValues.MasterUserPassword
	f39.MultiAZ = resp.DBInstance.PendingModifiedValues.MultiAZ
	f39f11 := &svcapitypes.PendingCloudwatchLogsExports{}
	f39f11f0 := []*string{}
	for _, f39f11f0iter := range resp.DBInstance.PendingModifiedValues.PendingCloudwatchLogsExports.LogTypesToDisable {
		var f39f11f0elem string
		f39f11f0elem = *f39f11f0iter
		f39f11f0 = append(f39f11f0, &f39f11f0elem)
	}
	f39f11.LogTypesToDisable = f39f11f0
	f39f11f1 := []*string{}
	for _, f39f11f1iter := range resp.DBInstance.PendingModifiedValues.PendingCloudwatchLogsExports.LogTypesToEnable {
		var f39f11f1elem string
		f39f11f1elem = *f39f11f1iter
		f39f11f1 = append(f39f11f1, &f39f11f1elem)
	}
	f39f11.LogTypesToEnable = f39f11f1
	f39.PendingCloudwatchLogsExports = f39f11
	f39.Port = resp.DBInstance.PendingModifiedValues.Port
	f39f13 := []*svcapitypes.ProcessorFeature{}
	for _, f39f13iter := range resp.DBInstance.PendingModifiedValues.ProcessorFeatures {
		f39f13elem := &svcapitypes.ProcessorFeature{}
		f39f13elem.Name = f39f13iter.Name
		f39f13elem.Value = f39f13iter.Value
		f39f13 = append(f39f13, f39f13elem)
	}
	f39.ProcessorFeatures = f39f13
	f39.StorageType = resp.DBInstance.PendingModifiedValues.StorageType
	ko.Status.PendingModifiedValues = f39
	ko.Status.PerformanceInsightsEnabled = resp.DBInstance.PerformanceInsightsEnabled
	f48 := []*string{}
	for _, f48iter := range resp.DBInstance.ReadReplicaDBClusterIdentifiers {
		var f48elem string
		f48elem = *f48iter
		f48 = append(f48, &f48elem)
	}
	ko.Status.ReadReplicaDBClusterIdentifiers = f48
	f49 := []*string{}
	for _, f49iter := range resp.DBInstance.ReadReplicaDBInstanceIdentifiers {
		var f49elem string
		f49elem = *f49iter
		f49 = append(f49, &f49elem)
	}
	ko.Status.ReadReplicaDBInstanceIdentifiers = f49
	ko.Status.ReadReplicaSourceDBInstanceIdentifier = resp.DBInstance.ReadReplicaSourceDBInstanceIdentifier
	ko.Status.SecondaryAvailabilityZone = resp.DBInstance.SecondaryAvailabilityZone
	f52 := []*svcapitypes.DBInstanceStatusInfo{}
	for _, f52iter := range resp.DBInstance.StatusInfos {
		f52elem := &svcapitypes.DBInstanceStatusInfo{}
		f52elem.Message = f52iter.Message
		f52elem.Normal = f52iter.Normal
		f52elem.Status = f52iter.Status
		f52elem.StatusType = f52iter.StatusType
		f52 = append(f52, f52elem)
	}
	ko.Status.StatusInfos = f52
	f57 := []*svcapitypes.VpcSecurityGroupMembership{}
	for _, f57iter := range resp.DBInstance.VpcSecurityGroups {
		f57elem := &svcapitypes.VpcSecurityGroupMembership{}
		f57elem.Status = f57iter.Status
		f57elem.VPCSecurityGroupID = f57iter.VpcSecurityGroupId
		f57 = append(f57, f57elem)
	}
	ko.Status.VPCSecurityGroups = f57

	return &resource{ko}, nil
}
// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	r *resource,
) (*svcsdk.ModifyDBInstanceInput, error) {
	res := &svcsdk.ModifyDBInstanceInput{}

	res.SetAllocatedStorage(*r.ko.Spec.AllocatedStorage)
	res.SetAutoMinorVersionUpgrade(*r.ko.Spec.AutoMinorVersionUpgrade)
	res.SetBackupRetentionPeriod(*r.ko.Spec.BackupRetentionPeriod)
	res.SetCACertificateIdentifier(*r.ko.Status.CACertificateIdentifier)
	res.SetCopyTagsToSnapshot(*r.ko.Spec.CopyTagsToSnapshot)
	res.SetDBInstanceClass(*r.ko.Spec.DBInstanceClass)
	res.SetDBInstanceIdentifier(*r.ko.Spec.DBInstanceIdentifier)
	res.SetDBParameterGroupName(*r.ko.Spec.DBParameterGroupName)
	f13 := []*string{}
	for _, f13iter := range r.ko.Spec.DBSecurityGroups {
		var f13elem string
		f13elem = *f13iter
		f13 = append(f13, &f13elem)
	}
	res.SetDBSecurityGroups(f13)
	res.SetDBSubnetGroupName(*r.ko.Spec.DBSubnetGroupName)
	res.SetDeletionProtection(*r.ko.Spec.DeletionProtection)
	res.SetDomain(*r.ko.Spec.Domain)
	res.SetDomainIAMRoleName(*r.ko.Spec.DomainIAMRoleName)
	res.SetEnableIAMDatabaseAuthentication(*r.ko.Spec.EnableIAMDatabaseAuthentication)
	res.SetEnablePerformanceInsights(*r.ko.Spec.EnablePerformanceInsights)
	res.SetEngineVersion(*r.ko.Spec.EngineVersion)
	res.SetIops(*r.ko.Spec.IOPS)
	res.SetLicenseModel(*r.ko.Spec.LicenseModel)
	f23 := string
	f23elem = *r.ko.Spec.MasterUserPassword
	log.Print(f23elem.Name)
	secretName := f23elem.Name
	secret, err := rm.client.secrets.Get(secretName)
	if err != nil {
 		return err
	}
	f23, err = b64.StdEncoding.DecodeString(secret)
 	if err != nil {
 		return err
	}
	res.SetMasterUserPassword(f23)
	res.SetMaxAllocatedStorage(*r.ko.Spec.MaxAllocatedStorage)
	res.SetMonitoringInterval(*r.ko.Spec.MonitoringInterval)
	res.SetMonitoringRoleArn(*r.ko.Spec.MonitoringRoleARN)
	res.SetMultiAZ(*r.ko.Spec.MultiAZ)
	res.SetOptionGroupName(*r.ko.Spec.OptionGroupName)
	res.SetPerformanceInsightsKMSKeyId(*r.ko.Spec.PerformanceInsightsKMSKeyID)
	res.SetPerformanceInsightsRetentionPeriod(*r.ko.Spec.PerformanceInsightsRetentionPeriod)
	res.SetPreferredBackupWindow(*r.ko.Spec.PreferredBackupWindow)
	res.SetPreferredMaintenanceWindow(*r.ko.Spec.PreferredMaintenanceWindow)
	f34 := []*svcsdk.ProcessorFeature{}
	for _, f34iter := range r.ko.Spec.ProcessorFeatures {
		f34elem := &svcsdk.ProcessorFeature{}
		f34elem.SetName(*f34iter.Name)
		f34elem.SetValue(*f34iter.Value)
		f34 = append(f34, f34elem)
	}
	res.SetProcessorFeatures(f34)
	res.SetPromotionTier(*r.ko.Spec.PromotionTier)
	res.SetPubliclyAccessible(*r.ko.Spec.PubliclyAccessible)
	res.SetStorageType(*r.ko.Spec.StorageType)
	res.SetTdeCredentialArn(*r.ko.Spec.TDECredentialARN)
	res.SetTdeCredentialPassword(*r.ko.Spec.TDECredentialPassword)
	f41 := []*string{}
	for _, f41iter := range r.ko.Spec.VPCSecurityGroupIDs {
		var f41elem string
		f41elem = *f41iter
		f41 = append(f41, &f41elem)
	}
	res.SetVpcSecurityGroupIds(f41)

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
	_, respErr := rm.sdkapi.DeleteDBInstanceWithContext(ctx, input)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteDBInstanceInput, error) {
	res := &svcsdk.DeleteDBInstanceInput{}

	res.SetDBInstanceIdentifier(*r.ko.Spec.DBInstanceIdentifier)

	return res, nil
}