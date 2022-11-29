/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type AMITypes string

const (
	AMITypes_AL2_x86_64          AMITypes = "AL2_x86_64"
	AMITypes_AL2_x86_64_GPU      AMITypes = "AL2_x86_64_GPU"
	AMITypes_AL2_ARM_64          AMITypes = "AL2_ARM_64"
	AMITypes_CUSTOM              AMITypes = "CUSTOM"
	AMITypes_BOTTLEROCKET_ARM_64 AMITypes = "BOTTLEROCKET_ARM_64"
	AMITypes_BOTTLEROCKET_x86_64 AMITypes = "BOTTLEROCKET_x86_64"
)

type AddonIssueCode string

const (
	AddonIssueCode_AccessDenied                 AddonIssueCode = "AccessDenied"
	AddonIssueCode_InternalFailure              AddonIssueCode = "InternalFailure"
	AddonIssueCode_ClusterUnreachable           AddonIssueCode = "ClusterUnreachable"
	AddonIssueCode_InsufficientNumberOfReplicas AddonIssueCode = "InsufficientNumberOfReplicas"
	AddonIssueCode_ConfigurationConflict        AddonIssueCode = "ConfigurationConflict"
	AddonIssueCode_AdmissionRequestDenied       AddonIssueCode = "AdmissionRequestDenied"
	AddonIssueCode_UnsupportedAddonModification AddonIssueCode = "UnsupportedAddonModification"
	AddonIssueCode_K8sResourceNotFound          AddonIssueCode = "K8sResourceNotFound"
)

type AddonStatus_SDK string

const (
	AddonStatus_SDK_CREATING      AddonStatus_SDK = "CREATING"
	AddonStatus_SDK_ACTIVE        AddonStatus_SDK = "ACTIVE"
	AddonStatus_SDK_CREATE_FAILED AddonStatus_SDK = "CREATE_FAILED"
	AddonStatus_SDK_UPDATING      AddonStatus_SDK = "UPDATING"
	AddonStatus_SDK_DELETING      AddonStatus_SDK = "DELETING"
	AddonStatus_SDK_DELETE_FAILED AddonStatus_SDK = "DELETE_FAILED"
	AddonStatus_SDK_DEGRADED      AddonStatus_SDK = "DEGRADED"
)

type CapacityTypes string

const (
	CapacityTypes_ON_DEMAND CapacityTypes = "ON_DEMAND"
	CapacityTypes_SPOT      CapacityTypes = "SPOT"
)

type ClusterStatus string

const (
	ClusterStatus_CREATING ClusterStatus = "CREATING"
	ClusterStatus_ACTIVE   ClusterStatus = "ACTIVE"
	ClusterStatus_DELETING ClusterStatus = "DELETING"
	ClusterStatus_FAILED   ClusterStatus = "FAILED"
	ClusterStatus_UPDATING ClusterStatus = "UPDATING"
	ClusterStatus_PENDING  ClusterStatus = "PENDING"
)

type ConfigStatus string

const (
	ConfigStatus_CREATING ConfigStatus = "CREATING"
	ConfigStatus_DELETING ConfigStatus = "DELETING"
	ConfigStatus_ACTIVE   ConfigStatus = "ACTIVE"
)

type ConnectorConfigProvider string

const (
	ConnectorConfigProvider_EKS_ANYWHERE ConnectorConfigProvider = "EKS_ANYWHERE"
	ConnectorConfigProvider_ANTHOS       ConnectorConfigProvider = "ANTHOS"
	ConnectorConfigProvider_GKE          ConnectorConfigProvider = "GKE"
	ConnectorConfigProvider_AKS          ConnectorConfigProvider = "AKS"
	ConnectorConfigProvider_OPENSHIFT    ConnectorConfigProvider = "OPENSHIFT"
	ConnectorConfigProvider_TANZU        ConnectorConfigProvider = "TANZU"
	ConnectorConfigProvider_RANCHER      ConnectorConfigProvider = "RANCHER"
	ConnectorConfigProvider_EC2          ConnectorConfigProvider = "EC2"
	ConnectorConfigProvider_OTHER        ConnectorConfigProvider = "OTHER"
)

type ErrorCode string

const (
	ErrorCode_SubnetNotFound               ErrorCode = "SubnetNotFound"
	ErrorCode_SecurityGroupNotFound        ErrorCode = "SecurityGroupNotFound"
	ErrorCode_EniLimitReached              ErrorCode = "EniLimitReached"
	ErrorCode_IpNotAvailable               ErrorCode = "IpNotAvailable"
	ErrorCode_AccessDenied                 ErrorCode = "AccessDenied"
	ErrorCode_OperationNotPermitted        ErrorCode = "OperationNotPermitted"
	ErrorCode_VpcIdNotFound                ErrorCode = "VpcIdNotFound"
	ErrorCode_Unknown                      ErrorCode = "Unknown"
	ErrorCode_NodeCreationFailure          ErrorCode = "NodeCreationFailure"
	ErrorCode_PodEvictionFailure           ErrorCode = "PodEvictionFailure"
	ErrorCode_InsufficientFreeAddresses    ErrorCode = "InsufficientFreeAddresses"
	ErrorCode_ClusterUnreachable           ErrorCode = "ClusterUnreachable"
	ErrorCode_InsufficientNumberOfReplicas ErrorCode = "InsufficientNumberOfReplicas"
	ErrorCode_ConfigurationConflict        ErrorCode = "ConfigurationConflict"
	ErrorCode_AdmissionRequestDenied       ErrorCode = "AdmissionRequestDenied"
	ErrorCode_UnsupportedAddonModification ErrorCode = "UnsupportedAddonModification"
	ErrorCode_K8sResourceNotFound          ErrorCode = "K8sResourceNotFound"
)

type FargateProfileStatus string

const (
	FargateProfileStatus_CREATING      FargateProfileStatus = "CREATING"
	FargateProfileStatus_ACTIVE        FargateProfileStatus = "ACTIVE"
	FargateProfileStatus_DELETING      FargateProfileStatus = "DELETING"
	FargateProfileStatus_CREATE_FAILED FargateProfileStatus = "CREATE_FAILED"
	FargateProfileStatus_DELETE_FAILED FargateProfileStatus = "DELETE_FAILED"
)

type IPFamily string

const (
	IPFamily_ipv4 IPFamily = "ipv4"
	IPFamily_ipv6 IPFamily = "ipv6"
)

type LogType string

const (
	LogType_api               LogType = "api"
	LogType_audit             LogType = "audit"
	LogType_authenticator     LogType = "authenticator"
	LogType_controllerManager LogType = "controllerManager"
	LogType_scheduler         LogType = "scheduler"
)

type NodegroupIssueCode string

const (
	NodegroupIssueCode_AutoScalingGroupNotFound             NodegroupIssueCode = "AutoScalingGroupNotFound"
	NodegroupIssueCode_AutoScalingGroupInvalidConfiguration NodegroupIssueCode = "AutoScalingGroupInvalidConfiguration"
	NodegroupIssueCode_Ec2SecurityGroupNotFound             NodegroupIssueCode = "Ec2SecurityGroupNotFound"
	NodegroupIssueCode_Ec2SecurityGroupDeletionFailure      NodegroupIssueCode = "Ec2SecurityGroupDeletionFailure"
	NodegroupIssueCode_Ec2LaunchTemplateNotFound            NodegroupIssueCode = "Ec2LaunchTemplateNotFound"
	NodegroupIssueCode_Ec2LaunchTemplateVersionMismatch     NodegroupIssueCode = "Ec2LaunchTemplateVersionMismatch"
	NodegroupIssueCode_Ec2SubnetNotFound                    NodegroupIssueCode = "Ec2SubnetNotFound"
	NodegroupIssueCode_Ec2SubnetInvalidConfiguration        NodegroupIssueCode = "Ec2SubnetInvalidConfiguration"
	NodegroupIssueCode_IamInstanceProfileNotFound           NodegroupIssueCode = "IamInstanceProfileNotFound"
	NodegroupIssueCode_IamLimitExceeded                     NodegroupIssueCode = "IamLimitExceeded"
	NodegroupIssueCode_IamNodeRoleNotFound                  NodegroupIssueCode = "IamNodeRoleNotFound"
	NodegroupIssueCode_NodeCreationFailure                  NodegroupIssueCode = "NodeCreationFailure"
	NodegroupIssueCode_AsgInstanceLaunchFailures            NodegroupIssueCode = "AsgInstanceLaunchFailures"
	NodegroupIssueCode_InstanceLimitExceeded                NodegroupIssueCode = "InstanceLimitExceeded"
	NodegroupIssueCode_InsufficientFreeAddresses            NodegroupIssueCode = "InsufficientFreeAddresses"
	NodegroupIssueCode_AccessDenied                         NodegroupIssueCode = "AccessDenied"
	NodegroupIssueCode_InternalFailure                      NodegroupIssueCode = "InternalFailure"
	NodegroupIssueCode_ClusterUnreachable                   NodegroupIssueCode = "ClusterUnreachable"
	NodegroupIssueCode_Ec2SubnetMissingIpv6Assignment       NodegroupIssueCode = "Ec2SubnetMissingIpv6Assignment"
)

type NodegroupStatus string

const (
	NodegroupStatus_CREATING      NodegroupStatus = "CREATING"
	NodegroupStatus_ACTIVE        NodegroupStatus = "ACTIVE"
	NodegroupStatus_UPDATING      NodegroupStatus = "UPDATING"
	NodegroupStatus_DELETING      NodegroupStatus = "DELETING"
	NodegroupStatus_CREATE_FAILED NodegroupStatus = "CREATE_FAILED"
	NodegroupStatus_DELETE_FAILED NodegroupStatus = "DELETE_FAILED"
	NodegroupStatus_DEGRADED      NodegroupStatus = "DEGRADED"
)

type ResolveConflicts string

const (
	ResolveConflicts_OVERWRITE ResolveConflicts = "OVERWRITE"
	ResolveConflicts_NONE      ResolveConflicts = "NONE"
)

type TaintEffect string

const (
	TaintEffect_NO_SCHEDULE        TaintEffect = "NO_SCHEDULE"
	TaintEffect_NO_EXECUTE         TaintEffect = "NO_EXECUTE"
	TaintEffect_PREFER_NO_SCHEDULE TaintEffect = "PREFER_NO_SCHEDULE"
)

type UpdateParamType string

const (
	UpdateParamType_Version                  UpdateParamType = "Version"
	UpdateParamType_PlatformVersion          UpdateParamType = "PlatformVersion"
	UpdateParamType_EndpointPrivateAccess    UpdateParamType = "EndpointPrivateAccess"
	UpdateParamType_EndpointPublicAccess     UpdateParamType = "EndpointPublicAccess"
	UpdateParamType_ClusterLogging           UpdateParamType = "ClusterLogging"
	UpdateParamType_DesiredSize              UpdateParamType = "DesiredSize"
	UpdateParamType_LabelsToAdd              UpdateParamType = "LabelsToAdd"
	UpdateParamType_LabelsToRemove           UpdateParamType = "LabelsToRemove"
	UpdateParamType_TaintsToAdd              UpdateParamType = "TaintsToAdd"
	UpdateParamType_TaintsToRemove           UpdateParamType = "TaintsToRemove"
	UpdateParamType_MaxSize                  UpdateParamType = "MaxSize"
	UpdateParamType_MinSize                  UpdateParamType = "MinSize"
	UpdateParamType_ReleaseVersion           UpdateParamType = "ReleaseVersion"
	UpdateParamType_PublicAccessCidrs        UpdateParamType = "PublicAccessCidrs"
	UpdateParamType_LaunchTemplateName       UpdateParamType = "LaunchTemplateName"
	UpdateParamType_LaunchTemplateVersion    UpdateParamType = "LaunchTemplateVersion"
	UpdateParamType_IdentityProviderConfig   UpdateParamType = "IdentityProviderConfig"
	UpdateParamType_EncryptionConfig         UpdateParamType = "EncryptionConfig"
	UpdateParamType_AddonVersion             UpdateParamType = "AddonVersion"
	UpdateParamType_ServiceAccountRoleArn    UpdateParamType = "ServiceAccountRoleArn"
	UpdateParamType_ResolveConflicts         UpdateParamType = "ResolveConflicts"
	UpdateParamType_MaxUnavailable           UpdateParamType = "MaxUnavailable"
	UpdateParamType_MaxUnavailablePercentage UpdateParamType = "MaxUnavailablePercentage"
)

type UpdateStatus string

const (
	UpdateStatus_InProgress UpdateStatus = "InProgress"
	UpdateStatus_Failed     UpdateStatus = "Failed"
	UpdateStatus_Cancelled  UpdateStatus = "Cancelled"
	UpdateStatus_Successful UpdateStatus = "Successful"
)

type UpdateType string

const (
	UpdateType_VersionUpdate                      UpdateType = "VersionUpdate"
	UpdateType_EndpointAccessUpdate               UpdateType = "EndpointAccessUpdate"
	UpdateType_LoggingUpdate                      UpdateType = "LoggingUpdate"
	UpdateType_ConfigUpdate                       UpdateType = "ConfigUpdate"
	UpdateType_AssociateIdentityProviderConfig    UpdateType = "AssociateIdentityProviderConfig"
	UpdateType_DisassociateIdentityProviderConfig UpdateType = "DisassociateIdentityProviderConfig"
	UpdateType_AssociateEncryptionConfig          UpdateType = "AssociateEncryptionConfig"
	UpdateType_AddonUpdate                        UpdateType = "AddonUpdate"
)
