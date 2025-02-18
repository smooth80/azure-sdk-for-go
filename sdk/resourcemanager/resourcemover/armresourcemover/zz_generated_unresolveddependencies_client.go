//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcemover

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

// UnresolvedDependenciesClient contains the methods for the UnresolvedDependencies group.
// Don't use this type directly, use NewUnresolvedDependenciesClient() instead.
type UnresolvedDependenciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewUnresolvedDependenciesClient creates a new instance of UnresolvedDependenciesClient with the specified values.
// subscriptionID - The Subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewUnresolvedDependenciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *UnresolvedDependenciesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &UnresolvedDependenciesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Gets a list of unresolved dependencies.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The Resource Group Name.
// moveCollectionName - The Move Collection Name.
// options - UnresolvedDependenciesClientGetOptions contains the optional parameters for the UnresolvedDependenciesClient.Get
// method.
func (client *UnresolvedDependenciesClient) Get(resourceGroupName string, moveCollectionName string, options *UnresolvedDependenciesClientGetOptions) *UnresolvedDependenciesClientGetPager {
	return &UnresolvedDependenciesClientGetPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getCreateRequest(ctx, resourceGroupName, moveCollectionName, options)
		},
		advancer: func(ctx context.Context, resp UnresolvedDependenciesClientGetResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.UnresolvedDependencyCollection.NextLink)
		},
	}
}

// getCreateRequest creates the Get request.
func (client *UnresolvedDependenciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, moveCollectionName string, options *UnresolvedDependenciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/moveCollections/{moveCollectionName}/unresolvedDependencies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if moveCollectionName == "" {
		return nil, errors.New("parameter moveCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{moveCollectionName}", url.PathEscape(moveCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.DependencyLevel != nil {
		reqQP.Set("dependencyLevel", string(*options.DependencyLevel))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2021-08-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *UnresolvedDependenciesClient) getHandleResponse(resp *http.Response) (UnresolvedDependenciesClientGetResponse, error) {
	result := UnresolvedDependenciesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.UnresolvedDependencyCollection); err != nil {
		return UnresolvedDependenciesClientGetResponse{}, err
	}
	return result, nil
}
