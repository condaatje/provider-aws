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

type CustomStepStatus string

const (
	CustomStepStatus_SUCCESS CustomStepStatus = "SUCCESS"
	CustomStepStatus_FAILURE CustomStepStatus = "FAILURE"
)

type Domain string

const (
	Domain_S3  Domain = "S3"
	Domain_EFS Domain = "EFS"
)

type EndpointType string

const (
	EndpointType_PUBLIC       EndpointType = "PUBLIC"
	EndpointType_VPC          EndpointType = "VPC"
	EndpointType_VPC_ENDPOINT EndpointType = "VPC_ENDPOINT"
)

type ExecutionErrorType string

const (
	ExecutionErrorType_PERMISSION_DENIED     ExecutionErrorType = "PERMISSION_DENIED"
	ExecutionErrorType_CUSTOM_STEP_FAILED    ExecutionErrorType = "CUSTOM_STEP_FAILED"
	ExecutionErrorType_THROTTLED             ExecutionErrorType = "THROTTLED"
	ExecutionErrorType_ALREADY_EXISTS        ExecutionErrorType = "ALREADY_EXISTS"
	ExecutionErrorType_NOT_FOUND             ExecutionErrorType = "NOT_FOUND"
	ExecutionErrorType_BAD_REQUEST           ExecutionErrorType = "BAD_REQUEST"
	ExecutionErrorType_TIMEOUT               ExecutionErrorType = "TIMEOUT"
	ExecutionErrorType_INTERNAL_SERVER_ERROR ExecutionErrorType = "INTERNAL_SERVER_ERROR"
)

type ExecutionStatus string

const (
	ExecutionStatus_IN_PROGRESS        ExecutionStatus = "IN_PROGRESS"
	ExecutionStatus_COMPLETED          ExecutionStatus = "COMPLETED"
	ExecutionStatus_EXCEPTION          ExecutionStatus = "EXCEPTION"
	ExecutionStatus_HANDLING_EXCEPTION ExecutionStatus = "HANDLING_EXCEPTION"
)

type HomeDirectoryType string

const (
	HomeDirectoryType_PATH    HomeDirectoryType = "PATH"
	HomeDirectoryType_LOGICAL HomeDirectoryType = "LOGICAL"
)

type IdentityProviderType string

const (
	IdentityProviderType_SERVICE_MANAGED       IdentityProviderType = "SERVICE_MANAGED"
	IdentityProviderType_API_GATEWAY           IdentityProviderType = "API_GATEWAY"
	IdentityProviderType_AWS_DIRECTORY_SERVICE IdentityProviderType = "AWS_DIRECTORY_SERVICE"
	IdentityProviderType_AWS_LAMBDA            IdentityProviderType = "AWS_LAMBDA"
)

type OverwriteExisting string

const (
	OverwriteExisting_TRUE  OverwriteExisting = "TRUE"
	OverwriteExisting_FALSE OverwriteExisting = "FALSE"
)

type Protocol string

const (
	Protocol_SFTP Protocol = "SFTP"
	Protocol_FTP  Protocol = "FTP"
	Protocol_FTPS Protocol = "FTPS"
)

type State string

const (
	State_OFFLINE      State = "OFFLINE"
	State_ONLINE       State = "ONLINE"
	State_STARTING     State = "STARTING"
	State_STOPPING     State = "STOPPING"
	State_START_FAILED State = "START_FAILED"
	State_STOP_FAILED  State = "STOP_FAILED"
)

type TLSSessionResumptionMode string

const (
	TLSSessionResumptionMode_DISABLED TLSSessionResumptionMode = "DISABLED"
	TLSSessionResumptionMode_ENABLED  TLSSessionResumptionMode = "ENABLED"
	TLSSessionResumptionMode_ENFORCED TLSSessionResumptionMode = "ENFORCED"
)

type WorkflowStepType string

const (
	WorkflowStepType_COPY   WorkflowStepType = "COPY"
	WorkflowStepType_CUSTOM WorkflowStepType = "CUSTOM"
	WorkflowStepType_TAG    WorkflowStepType = "TAG"
	WorkflowStepType_DELETE WorkflowStepType = "DELETE"
)
