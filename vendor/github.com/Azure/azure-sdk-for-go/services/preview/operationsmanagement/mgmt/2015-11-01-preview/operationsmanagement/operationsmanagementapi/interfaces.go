package operationsmanagementapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement"
	"github.com/Azure/go-autorest/autorest"
)

// SolutionsClientAPI contains the set of methods on the SolutionsClient type.
type SolutionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, solutionName string, parameters operationsmanagement.Solution) (result operationsmanagement.SolutionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, solutionName string) (result operationsmanagement.SolutionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, solutionName string) (result operationsmanagement.Solution, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result operationsmanagement.SolutionPropertiesList, err error)
	ListBySubscription(ctx context.Context) (result operationsmanagement.SolutionPropertiesList, err error)
	Update(ctx context.Context, resourceGroupName string, solutionName string, parameters operationsmanagement.SolutionPatch) (result operationsmanagement.SolutionsUpdateFuture, err error)
}

var _ SolutionsClientAPI = (*operationsmanagement.SolutionsClient)(nil)

// ManagementAssociationsClientAPI contains the set of methods on the ManagementAssociationsClient type.
type ManagementAssociationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, managementAssociationName string, parameters operationsmanagement.ManagementAssociation) (result operationsmanagement.ManagementAssociation, err error)
	Delete(ctx context.Context, resourceGroupName string, managementAssociationName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, managementAssociationName string) (result operationsmanagement.ManagementAssociation, err error)
	ListBySubscription(ctx context.Context) (result operationsmanagement.ManagementAssociationPropertiesList, err error)
}

var _ ManagementAssociationsClientAPI = (*operationsmanagement.ManagementAssociationsClient)(nil)

// ManagementConfigurationsClientAPI contains the set of methods on the ManagementConfigurationsClient type.
type ManagementConfigurationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, managementConfigurationName string, parameters operationsmanagement.ManagementConfiguration) (result operationsmanagement.ManagementConfiguration, err error)
	Delete(ctx context.Context, resourceGroupName string, managementConfigurationName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, managementConfigurationName string) (result operationsmanagement.ManagementConfiguration, err error)
	ListBySubscription(ctx context.Context) (result operationsmanagement.ManagementConfigurationPropertiesList, err error)
}

var _ ManagementConfigurationsClientAPI = (*operationsmanagement.ManagementConfigurationsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result operationsmanagement.OperationListResult, err error)
}

var _ OperationsClientAPI = (*operationsmanagement.OperationsClient)(nil)
