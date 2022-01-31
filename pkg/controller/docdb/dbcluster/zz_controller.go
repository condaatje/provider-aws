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

package dbcluster

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/docdb"
	svcsdk "github.com/aws/aws-sdk-go/service/docdb"
	svcsdkapi "github.com/aws/aws-sdk-go/service/docdb/docdbiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane/provider-aws/apis/docdb/v1alpha1"
	awsclient "github.com/crossplane/provider-aws/pkg/clients"
)

const (
	errUnexpectedObject = "managed resource is not an DBCluster resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create DBCluster in AWS"
	errUpdate        = "cannot update DBCluster in AWS"
	errDescribe      = "failed to describe DBCluster"
	errDelete        = "failed to delete DBCluster"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.DBCluster)
	if !ok {
		return nil, errors.New(errUnexpectedObject)
	}
	sess, err := awsclient.GetConfigV1(ctx, c.kube, mg, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, mg cpresource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*svcapitypes.DBCluster)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateDescribeDBClustersInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.DescribeDBClustersWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	resp = e.filterList(cr, resp)
	if len(resp.DBClusters) == 0 {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateDBCluster(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)

	upToDate, err := e.isUpToDate(cr, resp)
	if err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "isUpToDate check failed")
	}
	return e.postObserve(ctx, cr, resp, managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        upToDate,
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
	}, nil)
}

func (e *external) Create(ctx context.Context, mg cpresource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*svcapitypes.DBCluster)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateDBClusterInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateDBClusterWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, awsclient.Wrap(err, errCreate)
	}

	if resp.DBCluster.AssociatedRoles != nil {
		f0 := []*svcapitypes.DBClusterRole{}
		for _, f0iter := range resp.DBCluster.AssociatedRoles {
			f0elem := &svcapitypes.DBClusterRole{}
			if f0iter.RoleArn != nil {
				f0elem.RoleARN = f0iter.RoleArn
			}
			if f0iter.Status != nil {
				f0elem.Status = f0iter.Status
			}
			f0 = append(f0, f0elem)
		}
		cr.Status.AtProvider.AssociatedRoles = f0
	} else {
		cr.Status.AtProvider.AssociatedRoles = nil
	}
	if resp.DBCluster.AvailabilityZones != nil {
		f1 := []*string{}
		for _, f1iter := range resp.DBCluster.AvailabilityZones {
			var f1elem string
			f1elem = *f1iter
			f1 = append(f1, &f1elem)
		}
		cr.Spec.ForProvider.AvailabilityZones = f1
	} else {
		cr.Spec.ForProvider.AvailabilityZones = nil
	}
	if resp.DBCluster.BackupRetentionPeriod != nil {
		cr.Spec.ForProvider.BackupRetentionPeriod = resp.DBCluster.BackupRetentionPeriod
	} else {
		cr.Spec.ForProvider.BackupRetentionPeriod = nil
	}
	if resp.DBCluster.ClusterCreateTime != nil {
		cr.Status.AtProvider.ClusterCreateTime = &metav1.Time{*resp.DBCluster.ClusterCreateTime}
	} else {
		cr.Status.AtProvider.ClusterCreateTime = nil
	}
	if resp.DBCluster.DBClusterArn != nil {
		cr.Status.AtProvider.DBClusterARN = resp.DBCluster.DBClusterArn
	} else {
		cr.Status.AtProvider.DBClusterARN = nil
	}
	if resp.DBCluster.DBClusterIdentifier != nil {
		cr.Status.AtProvider.DBClusterIdentifier = resp.DBCluster.DBClusterIdentifier
	} else {
		cr.Status.AtProvider.DBClusterIdentifier = nil
	}
	if resp.DBCluster.DBClusterMembers != nil {
		f6 := []*svcapitypes.DBClusterMember{}
		for _, f6iter := range resp.DBCluster.DBClusterMembers {
			f6elem := &svcapitypes.DBClusterMember{}
			if f6iter.DBClusterParameterGroupStatus != nil {
				f6elem.DBClusterParameterGroupStatus = f6iter.DBClusterParameterGroupStatus
			}
			if f6iter.DBInstanceIdentifier != nil {
				f6elem.DBInstanceIdentifier = f6iter.DBInstanceIdentifier
			}
			if f6iter.IsClusterWriter != nil {
				f6elem.IsClusterWriter = f6iter.IsClusterWriter
			}
			if f6iter.PromotionTier != nil {
				f6elem.PromotionTier = f6iter.PromotionTier
			}
			f6 = append(f6, f6elem)
		}
		cr.Status.AtProvider.DBClusterMembers = f6
	} else {
		cr.Status.AtProvider.DBClusterMembers = nil
	}
	if resp.DBCluster.DBClusterParameterGroup != nil {
		cr.Status.AtProvider.DBClusterParameterGroup = resp.DBCluster.DBClusterParameterGroup
	} else {
		cr.Status.AtProvider.DBClusterParameterGroup = nil
	}
	if resp.DBCluster.DBSubnetGroup != nil {
		cr.Status.AtProvider.DBSubnetGroup = resp.DBCluster.DBSubnetGroup
	} else {
		cr.Status.AtProvider.DBSubnetGroup = nil
	}
	if resp.DBCluster.DbClusterResourceId != nil {
		cr.Status.AtProvider.DBClusterResourceID = resp.DBCluster.DbClusterResourceId
	} else {
		cr.Status.AtProvider.DBClusterResourceID = nil
	}
	if resp.DBCluster.DeletionProtection != nil {
		cr.Spec.ForProvider.DeletionProtection = resp.DBCluster.DeletionProtection
	} else {
		cr.Spec.ForProvider.DeletionProtection = nil
	}
	if resp.DBCluster.EarliestRestorableTime != nil {
		cr.Status.AtProvider.EarliestRestorableTime = &metav1.Time{*resp.DBCluster.EarliestRestorableTime}
	} else {
		cr.Status.AtProvider.EarliestRestorableTime = nil
	}
	if resp.DBCluster.EnabledCloudwatchLogsExports != nil {
		f12 := []*string{}
		for _, f12iter := range resp.DBCluster.EnabledCloudwatchLogsExports {
			var f12elem string
			f12elem = *f12iter
			f12 = append(f12, &f12elem)
		}
		cr.Status.AtProvider.EnabledCloudwatchLogsExports = f12
	} else {
		cr.Status.AtProvider.EnabledCloudwatchLogsExports = nil
	}
	if resp.DBCluster.Endpoint != nil {
		cr.Status.AtProvider.Endpoint = resp.DBCluster.Endpoint
	} else {
		cr.Status.AtProvider.Endpoint = nil
	}
	if resp.DBCluster.Engine != nil {
		cr.Spec.ForProvider.Engine = resp.DBCluster.Engine
	} else {
		cr.Spec.ForProvider.Engine = nil
	}
	if resp.DBCluster.EngineVersion != nil {
		cr.Spec.ForProvider.EngineVersion = resp.DBCluster.EngineVersion
	} else {
		cr.Spec.ForProvider.EngineVersion = nil
	}
	if resp.DBCluster.HostedZoneId != nil {
		cr.Status.AtProvider.HostedZoneID = resp.DBCluster.HostedZoneId
	} else {
		cr.Status.AtProvider.HostedZoneID = nil
	}
	if resp.DBCluster.KmsKeyId != nil {
		cr.Spec.ForProvider.KMSKeyID = resp.DBCluster.KmsKeyId
	} else {
		cr.Spec.ForProvider.KMSKeyID = nil
	}
	if resp.DBCluster.LatestRestorableTime != nil {
		cr.Status.AtProvider.LatestRestorableTime = &metav1.Time{*resp.DBCluster.LatestRestorableTime}
	} else {
		cr.Status.AtProvider.LatestRestorableTime = nil
	}
	if resp.DBCluster.MasterUsername != nil {
		cr.Spec.ForProvider.MasterUsername = resp.DBCluster.MasterUsername
	} else {
		cr.Spec.ForProvider.MasterUsername = nil
	}
	if resp.DBCluster.MultiAZ != nil {
		cr.Status.AtProvider.MultiAZ = resp.DBCluster.MultiAZ
	} else {
		cr.Status.AtProvider.MultiAZ = nil
	}
	if resp.DBCluster.PercentProgress != nil {
		cr.Status.AtProvider.PercentProgress = resp.DBCluster.PercentProgress
	} else {
		cr.Status.AtProvider.PercentProgress = nil
	}
	if resp.DBCluster.Port != nil {
		cr.Spec.ForProvider.Port = resp.DBCluster.Port
	} else {
		cr.Spec.ForProvider.Port = nil
	}
	if resp.DBCluster.PreferredBackupWindow != nil {
		cr.Spec.ForProvider.PreferredBackupWindow = resp.DBCluster.PreferredBackupWindow
	} else {
		cr.Spec.ForProvider.PreferredBackupWindow = nil
	}
	if resp.DBCluster.PreferredMaintenanceWindow != nil {
		cr.Spec.ForProvider.PreferredMaintenanceWindow = resp.DBCluster.PreferredMaintenanceWindow
	} else {
		cr.Spec.ForProvider.PreferredMaintenanceWindow = nil
	}
	if resp.DBCluster.ReaderEndpoint != nil {
		cr.Status.AtProvider.ReaderEndpoint = resp.DBCluster.ReaderEndpoint
	} else {
		cr.Status.AtProvider.ReaderEndpoint = nil
	}
	if resp.DBCluster.Status != nil {
		cr.Status.AtProvider.Status = resp.DBCluster.Status
	} else {
		cr.Status.AtProvider.Status = nil
	}
	if resp.DBCluster.StorageEncrypted != nil {
		cr.Spec.ForProvider.StorageEncrypted = resp.DBCluster.StorageEncrypted
	} else {
		cr.Spec.ForProvider.StorageEncrypted = nil
	}
	if resp.DBCluster.VpcSecurityGroups != nil {
		f28 := []*svcapitypes.VPCSecurityGroupMembership{}
		for _, f28iter := range resp.DBCluster.VpcSecurityGroups {
			f28elem := &svcapitypes.VPCSecurityGroupMembership{}
			if f28iter.Status != nil {
				f28elem.Status = f28iter.Status
			}
			if f28iter.VpcSecurityGroupId != nil {
				f28elem.VPCSecurityGroupID = f28iter.VpcSecurityGroupId
			}
			f28 = append(f28, f28elem)
		}
		cr.Status.AtProvider.VPCSecurityGroups = f28
	} else {
		cr.Status.AtProvider.VPCSecurityGroups = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.DBCluster)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateModifyDBClusterInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.ModifyDBClusterWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, awsclient.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.DBCluster)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteDBClusterInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteDBClusterWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.DocDBAPI, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		filterList:     nopFilterList,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		preDelete:      nopPreDelete,
		postDelete:     nopPostDelete,
		preUpdate:      nopPreUpdate,
		postUpdate:     nopPostUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.DocDBAPI
	preObserve     func(context.Context, *svcapitypes.DBCluster, *svcsdk.DescribeDBClustersInput) error
	postObserve    func(context.Context, *svcapitypes.DBCluster, *svcsdk.DescribeDBClustersOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	filterList     func(*svcapitypes.DBCluster, *svcsdk.DescribeDBClustersOutput) *svcsdk.DescribeDBClustersOutput
	lateInitialize func(*svcapitypes.DBClusterParameters, *svcsdk.DescribeDBClustersOutput) error
	isUpToDate     func(*svcapitypes.DBCluster, *svcsdk.DescribeDBClustersOutput) (bool, error)
	preCreate      func(context.Context, *svcapitypes.DBCluster, *svcsdk.CreateDBClusterInput) error
	postCreate     func(context.Context, *svcapitypes.DBCluster, *svcsdk.CreateDBClusterOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.DBCluster, *svcsdk.DeleteDBClusterInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.DBCluster, *svcsdk.DeleteDBClusterOutput, error) error
	preUpdate      func(context.Context, *svcapitypes.DBCluster, *svcsdk.ModifyDBClusterInput) error
	postUpdate     func(context.Context, *svcapitypes.DBCluster, *svcsdk.ModifyDBClusterOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.DBCluster, *svcsdk.DescribeDBClustersInput) error {
	return nil
}
func nopPostObserve(_ context.Context, _ *svcapitypes.DBCluster, _ *svcsdk.DescribeDBClustersOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopFilterList(_ *svcapitypes.DBCluster, list *svcsdk.DescribeDBClustersOutput) *svcsdk.DescribeDBClustersOutput {
	return list
}

func nopLateInitialize(*svcapitypes.DBClusterParameters, *svcsdk.DescribeDBClustersOutput) error {
	return nil
}
func alwaysUpToDate(*svcapitypes.DBCluster, *svcsdk.DescribeDBClustersOutput) (bool, error) {
	return true, nil
}

func nopPreCreate(context.Context, *svcapitypes.DBCluster, *svcsdk.CreateDBClusterInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.DBCluster, _ *svcsdk.CreateDBClusterOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.DBCluster, *svcsdk.DeleteDBClusterInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.DBCluster, _ *svcsdk.DeleteDBClusterOutput, err error) error {
	return err
}
func nopPreUpdate(context.Context, *svcapitypes.DBCluster, *svcsdk.ModifyDBClusterInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.DBCluster, _ *svcsdk.ModifyDBClusterOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
