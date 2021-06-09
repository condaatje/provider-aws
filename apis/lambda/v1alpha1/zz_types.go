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

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
)

type AccountLimit struct {
	CodeSizeUnzipped *int64 `json:"codeSizeUnzipped,omitempty"`

	CodeSizeZipped *int64 `json:"codeSizeZipped,omitempty"`

	TotalCodeSize *int64 `json:"totalCodeSize,omitempty"`
}

type AccountUsage struct {
	FunctionCount *int64 `json:"functionCount,omitempty"`

	TotalCodeSize *int64 `json:"totalCodeSize,omitempty"`
}

type CodeSigningConfig struct {
	CodeSigningConfigARN *string `json:"codeSigningConfigARN,omitempty"`

	Description *string `json:"description,omitempty"`

	LastModified *string `json:"lastModified,omitempty"`
}

type DeadLetterConfig struct {
	TargetARN *string `json:"targetARN,omitempty"`
}

type Environment struct {
	Variables map[string]*string `json:"variables,omitempty"`
}

type EnvironmentError struct {
	ErrorCode *string `json:"errorCode,omitempty"`

	Message *string `json:"message,omitempty"`
}

type EnvironmentResponse struct {
	// Error messages for environment variables that couldn't be applied.
	Error *EnvironmentError `json:"error,omitempty"`

	Variables map[string]*string `json:"variables,omitempty"`
}

type FileSystemConfig struct {
	ARN *string `json:"arn,omitempty"`

	LocalMountPath *string `json:"localMountPath,omitempty"`
}

type FunctionCode struct {
	ImageURI *string `json:"imageURI,omitempty"`
}

type FunctionCodeLocation struct {
	ImageURI *string `json:"imageURI,omitempty"`

	Location *string `json:"location,omitempty"`

	RepositoryType *string `json:"repositoryType,omitempty"`

	ResolvedImageURI *string `json:"resolvedImageURI,omitempty"`
}

type FunctionEventInvokeConfig struct {
	FunctionARN *string `json:"functionARN,omitempty"`
}

type ImageConfig struct {
	Command []*string `json:"command,omitempty"`

	EntryPoint []*string `json:"entryPoint,omitempty"`

	WorkingDirectory *string `json:"workingDirectory,omitempty"`
}

type ImageConfigError struct {
	ErrorCode *string `json:"errorCode,omitempty"`

	Message *string `json:"message,omitempty"`
}

type ImageConfigResponse struct {
	// Error response to GetFunctionConfiguration.
	Error *ImageConfigError `json:"error,omitempty"`
	// Configuration values that override the container image Dockerfile settings.
	// See Container settings (https://docs.aws.amazon.com/lambda/latest/dg/images-create.html#images-parms).
	ImageConfig *ImageConfig `json:"imageConfig,omitempty"`
}

type Layer struct {
	ARN *string `json:"arn,omitempty"`

	CodeSize *int64 `json:"codeSize,omitempty"`

	SigningJobARN *string `json:"signingJobARN,omitempty"`

	SigningProfileVersionARN *string `json:"signingProfileVersionARN,omitempty"`
}

type LayerVersionContentOutput struct {
	CodeSHA256 *string `json:"codeSHA256,omitempty"`

	CodeSize *int64 `json:"codeSize,omitempty"`

	Location *string `json:"location,omitempty"`

	SigningJobARN *string `json:"signingJobARN,omitempty"`

	SigningProfileVersionARN *string `json:"signingProfileVersionARN,omitempty"`
}

type LayerVersionsListItem struct {
	CreatedDate *string `json:"createdDate,omitempty"`

	Description *string `json:"description,omitempty"`

	LayerVersionARN *string `json:"layerVersionARN,omitempty"`
}

type ProvisionedConcurrencyConfigListItem struct {
	FunctionARN *string `json:"functionARN,omitempty"`

	LastModified *string `json:"lastModified,omitempty"`

	StatusReason *string `json:"statusReason,omitempty"`
}

type TracingConfig struct {
	Mode *string `json:"mode,omitempty"`
}

type TracingConfigResponse struct {
	Mode *string `json:"mode,omitempty"`
}

type VPCConfig struct {
	SecurityGroupIDs []*string `json:"securityGroupIDs,omitempty"`

	SubnetIDs []*string `json:"subnetIDs,omitempty"`
}

type VPCConfigResponse struct {
	SecurityGroupIDs []*string `json:"securityGroupIDs,omitempty"`

	SubnetIDs []*string `json:"subnetIDs,omitempty"`

	VPCID *string `json:"vpcID,omitempty"`
}
