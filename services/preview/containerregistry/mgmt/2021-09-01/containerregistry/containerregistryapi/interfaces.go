package containerregistryapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/containerregistry/mgmt/2021-09-01/containerregistry"
)

// RegistriesClientAPI contains the set of methods on the RegistriesClient type.
type RegistriesClientAPI interface {
	CheckNameAvailability(ctx context.Context, registryNameCheckRequest containerregistry.RegistryNameCheckRequest) (result containerregistry.RegistryNameStatus, err error)
	Create(ctx context.Context, resourceGroupName string, registryName string, registry containerregistry.Registry) (result containerregistry.RegistriesCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.RegistriesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.Registry, err error)
	GetBuildSourceUploadURL(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.SourceUploadDefinition, err error)
	GetPrivateLinkResource(ctx context.Context, resourceGroupName string, registryName string, groupName string) (result containerregistry.PrivateLinkResource, err error)
	ImportImage(ctx context.Context, resourceGroupName string, registryName string, parameters containerregistry.ImportImageParameters) (result containerregistry.RegistriesImportImageFuture, err error)
	List(ctx context.Context) (result containerregistry.RegistryListResultPage, err error)
	ListComplete(ctx context.Context) (result containerregistry.RegistryListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result containerregistry.RegistryListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result containerregistry.RegistryListResultIterator, err error)
	ListCredentials(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.RegistryListCredentialsResult, err error)
	ListPrivateLinkResources(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.PrivateLinkResourceListResultPage, err error)
	ListPrivateLinkResourcesComplete(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.PrivateLinkResourceListResultIterator, err error)
	ListUsages(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.RegistryUsageListResult, err error)
	RegenerateCredential(ctx context.Context, resourceGroupName string, registryName string, regenerateCredentialParameters containerregistry.RegenerateCredentialParameters) (result containerregistry.RegistryListCredentialsResult, err error)
	ScheduleRun(ctx context.Context, resourceGroupName string, registryName string, runRequest containerregistry.BasicRunRequest) (result containerregistry.RegistriesScheduleRunFuture, err error)
	Update(ctx context.Context, resourceGroupName string, registryName string, registryUpdateParameters containerregistry.RegistryUpdateParameters) (result containerregistry.RegistriesUpdateFuture, err error)
}

var _ RegistriesClientAPI = (*containerregistry.RegistriesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result containerregistry.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result containerregistry.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*containerregistry.OperationsClient)(nil)

// PrivateEndpointConnectionsClientAPI contains the set of methods on the PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, registryName string, privateEndpointConnectionName string, privateEndpointConnection containerregistry.PrivateEndpointConnection) (result containerregistry.PrivateEndpointConnectionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, registryName string, privateEndpointConnectionName string) (result containerregistry.PrivateEndpointConnectionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, registryName string, privateEndpointConnectionName string) (result containerregistry.PrivateEndpointConnection, err error)
	List(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.PrivateEndpointConnectionListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.PrivateEndpointConnectionListResultIterator, err error)
}

var _ PrivateEndpointConnectionsClientAPI = (*containerregistry.PrivateEndpointConnectionsClient)(nil)

// ReplicationsClientAPI contains the set of methods on the ReplicationsClient type.
type ReplicationsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replication containerregistry.Replication) (result containerregistry.ReplicationsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, registryName string, replicationName string) (result containerregistry.ReplicationsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, registryName string, replicationName string) (result containerregistry.Replication, err error)
	List(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.ReplicationListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.ReplicationListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replicationUpdateParameters containerregistry.ReplicationUpdateParameters) (result containerregistry.ReplicationsUpdateFuture, err error)
}

var _ ReplicationsClientAPI = (*containerregistry.ReplicationsClient)(nil)

// WebhooksClientAPI contains the set of methods on the WebhooksClient type.
type WebhooksClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, registryName string, webhookName string, webhookCreateParameters containerregistry.WebhookCreateParameters) (result containerregistry.WebhooksCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, registryName string, webhookName string) (result containerregistry.WebhooksDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, registryName string, webhookName string) (result containerregistry.Webhook, err error)
	GetCallbackConfig(ctx context.Context, resourceGroupName string, registryName string, webhookName string) (result containerregistry.CallbackConfig, err error)
	List(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.WebhookListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.WebhookListResultIterator, err error)
	ListEvents(ctx context.Context, resourceGroupName string, registryName string, webhookName string) (result containerregistry.EventListResultPage, err error)
	ListEventsComplete(ctx context.Context, resourceGroupName string, registryName string, webhookName string) (result containerregistry.EventListResultIterator, err error)
	Ping(ctx context.Context, resourceGroupName string, registryName string, webhookName string) (result containerregistry.EventInfo, err error)
	Update(ctx context.Context, resourceGroupName string, registryName string, webhookName string, webhookUpdateParameters containerregistry.WebhookUpdateParameters) (result containerregistry.WebhooksUpdateFuture, err error)
}

var _ WebhooksClientAPI = (*containerregistry.WebhooksClient)(nil)

// AgentPoolsClientAPI contains the set of methods on the AgentPoolsClient type.
type AgentPoolsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, agentPool containerregistry.AgentPool) (result containerregistry.AgentPoolsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string) (result containerregistry.AgentPoolsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string) (result containerregistry.AgentPool, err error)
	GetQueueStatus(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string) (result containerregistry.AgentPoolQueueStatus, err error)
	List(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.AgentPoolListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.AgentPoolListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, updateParameters containerregistry.AgentPoolUpdateParameters) (result containerregistry.AgentPoolsUpdateFuture, err error)
}

var _ AgentPoolsClientAPI = (*containerregistry.AgentPoolsClient)(nil)

// RunsClientAPI contains the set of methods on the RunsClient type.
type RunsClientAPI interface {
	Cancel(ctx context.Context, resourceGroupName string, registryName string, runID string) (result containerregistry.RunsCancelFuture, err error)
	Get(ctx context.Context, resourceGroupName string, registryName string, runID string) (result containerregistry.Run, err error)
	GetLogSasURL(ctx context.Context, resourceGroupName string, registryName string, runID string) (result containerregistry.RunGetLogResult, err error)
	List(ctx context.Context, resourceGroupName string, registryName string, filter string, top *int32) (result containerregistry.RunListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, registryName string, filter string, top *int32) (result containerregistry.RunListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, registryName string, runID string, runUpdateParameters containerregistry.RunUpdateParameters) (result containerregistry.RunsUpdateFuture, err error)
}

var _ RunsClientAPI = (*containerregistry.RunsClient)(nil)

// TaskRunsClientAPI contains the set of methods on the TaskRunsClient type.
type TaskRunsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, registryName string, taskRunName string, taskRun containerregistry.TaskRun) (result containerregistry.TaskRunsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, registryName string, taskRunName string) (result containerregistry.TaskRunsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, registryName string, taskRunName string) (result containerregistry.TaskRun, err error)
	GetDetails(ctx context.Context, resourceGroupName string, registryName string, taskRunName string) (result containerregistry.TaskRun, err error)
	List(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.TaskRunListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.TaskRunListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, registryName string, taskRunName string, updateParameters containerregistry.TaskRunUpdateParameters) (result containerregistry.TaskRunsUpdateFuture, err error)
}

var _ TaskRunsClientAPI = (*containerregistry.TaskRunsClient)(nil)

// TasksClientAPI contains the set of methods on the TasksClient type.
type TasksClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, registryName string, taskName string, taskCreateParameters containerregistry.Task) (result containerregistry.TasksCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, registryName string, taskName string) (result containerregistry.TasksDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, registryName string, taskName string) (result containerregistry.Task, err error)
	GetDetails(ctx context.Context, resourceGroupName string, registryName string, taskName string) (result containerregistry.Task, err error)
	List(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.TaskListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.TaskListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, registryName string, taskName string, taskUpdateParameters containerregistry.TaskUpdateParameters) (result containerregistry.TasksUpdateFuture, err error)
}

var _ TasksClientAPI = (*containerregistry.TasksClient)(nil)
