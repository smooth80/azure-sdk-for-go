//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// WorkflowRunActionScopeRepetitionsClient contains the methods for the WorkflowRunActionScopeRepetitions group.
// Don't use this type directly, use NewWorkflowRunActionScopeRepetitionsClient() instead.
type WorkflowRunActionScopeRepetitionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkflowRunActionScopeRepetitionsClient creates a new instance of WorkflowRunActionScopeRepetitionsClient with the specified values.
// subscriptionID - The subscription id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkflowRunActionScopeRepetitionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *WorkflowRunActionScopeRepetitionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &WorkflowRunActionScopeRepetitionsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Get a workflow run action scoped repetition.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// workflowName - The workflow name.
// runName - The workflow run name.
// actionName - The workflow action name.
// repetitionName - The workflow repetition.
// options - WorkflowRunActionScopeRepetitionsClientGetOptions contains the optional parameters for the WorkflowRunActionScopeRepetitionsClient.Get
// method.
func (client *WorkflowRunActionScopeRepetitionsClient) Get(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, repetitionName string, options *WorkflowRunActionScopeRepetitionsClientGetOptions) (WorkflowRunActionScopeRepetitionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workflowName, runName, actionName, repetitionName, options)
	if err != nil {
		return WorkflowRunActionScopeRepetitionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkflowRunActionScopeRepetitionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkflowRunActionScopeRepetitionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkflowRunActionScopeRepetitionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, repetitionName string, options *WorkflowRunActionScopeRepetitionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/scopeRepetitions/{repetitionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if runName == "" {
		return nil, errors.New("parameter runName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runName}", url.PathEscape(runName))
	if actionName == "" {
		return nil, errors.New("parameter actionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{actionName}", url.PathEscape(actionName))
	if repetitionName == "" {
		return nil, errors.New("parameter repetitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{repetitionName}", url.PathEscape(repetitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkflowRunActionScopeRepetitionsClient) getHandleResponse(resp *http.Response) (WorkflowRunActionScopeRepetitionsClientGetResponse, error) {
	result := WorkflowRunActionScopeRepetitionsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowRunActionRepetitionDefinition); err != nil {
		return WorkflowRunActionScopeRepetitionsClientGetResponse{}, err
	}
	return result, nil
}

// List - List the workflow run action scoped repetitions.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// workflowName - The workflow name.
// runName - The workflow run name.
// actionName - The workflow action name.
// options - WorkflowRunActionScopeRepetitionsClientListOptions contains the optional parameters for the WorkflowRunActionScopeRepetitionsClient.List
// method.
func (client *WorkflowRunActionScopeRepetitionsClient) List(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, options *WorkflowRunActionScopeRepetitionsClientListOptions) (WorkflowRunActionScopeRepetitionsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, workflowName, runName, actionName, options)
	if err != nil {
		return WorkflowRunActionScopeRepetitionsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkflowRunActionScopeRepetitionsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkflowRunActionScopeRepetitionsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *WorkflowRunActionScopeRepetitionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, options *WorkflowRunActionScopeRepetitionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/scopeRepetitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if runName == "" {
		return nil, errors.New("parameter runName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runName}", url.PathEscape(runName))
	if actionName == "" {
		return nil, errors.New("parameter actionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{actionName}", url.PathEscape(actionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkflowRunActionScopeRepetitionsClient) listHandleResponse(resp *http.Response) (WorkflowRunActionScopeRepetitionsClientListResponse, error) {
	result := WorkflowRunActionScopeRepetitionsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowRunActionRepetitionDefinitionCollection); err != nil {
		return WorkflowRunActionScopeRepetitionsClientListResponse{}, err
	}
	return result, nil
}
