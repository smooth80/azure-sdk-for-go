//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// VirtualMachineScaleSetRollingUpgradesClient contains the methods for the VirtualMachineScaleSetRollingUpgrades group.
// Don't use this type directly, use NewVirtualMachineScaleSetRollingUpgradesClient() instead.
type VirtualMachineScaleSetRollingUpgradesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualMachineScaleSetRollingUpgradesClient creates a new instance of VirtualMachineScaleSetRollingUpgradesClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualMachineScaleSetRollingUpgradesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VirtualMachineScaleSetRollingUpgradesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &VirtualMachineScaleSetRollingUpgradesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCancel - Cancels the current virtual machine scale set rolling upgrade.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// vmScaleSetName - The name of the VM scale set.
// options - VirtualMachineScaleSetRollingUpgradesClientBeginCancelOptions contains the optional parameters for the VirtualMachineScaleSetRollingUpgradesClient.BeginCancel
// method.
func (client *VirtualMachineScaleSetRollingUpgradesClient) BeginCancel(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesClientBeginCancelOptions) (VirtualMachineScaleSetRollingUpgradesClientCancelPollerResponse, error) {
	resp, err := client.cancel(ctx, resourceGroupName, vmScaleSetName, options)
	if err != nil {
		return VirtualMachineScaleSetRollingUpgradesClientCancelPollerResponse{}, err
	}
	result := VirtualMachineScaleSetRollingUpgradesClientCancelPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualMachineScaleSetRollingUpgradesClient.Cancel", "", resp, client.pl)
	if err != nil {
		return VirtualMachineScaleSetRollingUpgradesClientCancelPollerResponse{}, err
	}
	result.Poller = &VirtualMachineScaleSetRollingUpgradesClientCancelPoller{
		pt: pt,
	}
	return result, nil
}

// Cancel - Cancels the current virtual machine scale set rolling upgrade.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualMachineScaleSetRollingUpgradesClient) cancel(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesClientBeginCancelOptions) (*http.Response, error) {
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, vmScaleSetName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *VirtualMachineScaleSetRollingUpgradesClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesClientBeginCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/rollingUpgrades/cancel"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// GetLatest - Gets the status of the latest virtual machine scale set rolling upgrade.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// vmScaleSetName - The name of the VM scale set.
// options - VirtualMachineScaleSetRollingUpgradesClientGetLatestOptions contains the optional parameters for the VirtualMachineScaleSetRollingUpgradesClient.GetLatest
// method.
func (client *VirtualMachineScaleSetRollingUpgradesClient) GetLatest(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesClientGetLatestOptions) (VirtualMachineScaleSetRollingUpgradesClientGetLatestResponse, error) {
	req, err := client.getLatestCreateRequest(ctx, resourceGroupName, vmScaleSetName, options)
	if err != nil {
		return VirtualMachineScaleSetRollingUpgradesClientGetLatestResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineScaleSetRollingUpgradesClientGetLatestResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineScaleSetRollingUpgradesClientGetLatestResponse{}, runtime.NewResponseError(resp)
	}
	return client.getLatestHandleResponse(resp)
}

// getLatestCreateRequest creates the GetLatest request.
func (client *VirtualMachineScaleSetRollingUpgradesClient) getLatestCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesClientGetLatestOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/rollingUpgrades/latest"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getLatestHandleResponse handles the GetLatest response.
func (client *VirtualMachineScaleSetRollingUpgradesClient) getLatestHandleResponse(resp *http.Response) (VirtualMachineScaleSetRollingUpgradesClientGetLatestResponse, error) {
	result := VirtualMachineScaleSetRollingUpgradesClientGetLatestResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RollingUpgradeStatusInfo); err != nil {
		return VirtualMachineScaleSetRollingUpgradesClientGetLatestResponse{}, err
	}
	return result, nil
}

// BeginStartExtensionUpgrade - Starts a rolling upgrade to move all extensions for all virtual machine scale set instances
// to the latest available extension version. Instances which are already running the latest extension versions
// are not affected.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// vmScaleSetName - The name of the VM scale set.
// options - VirtualMachineScaleSetRollingUpgradesClientBeginStartExtensionUpgradeOptions contains the optional parameters
// for the VirtualMachineScaleSetRollingUpgradesClient.BeginStartExtensionUpgrade method.
func (client *VirtualMachineScaleSetRollingUpgradesClient) BeginStartExtensionUpgrade(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesClientBeginStartExtensionUpgradeOptions) (VirtualMachineScaleSetRollingUpgradesClientStartExtensionUpgradePollerResponse, error) {
	resp, err := client.startExtensionUpgrade(ctx, resourceGroupName, vmScaleSetName, options)
	if err != nil {
		return VirtualMachineScaleSetRollingUpgradesClientStartExtensionUpgradePollerResponse{}, err
	}
	result := VirtualMachineScaleSetRollingUpgradesClientStartExtensionUpgradePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualMachineScaleSetRollingUpgradesClient.StartExtensionUpgrade", "", resp, client.pl)
	if err != nil {
		return VirtualMachineScaleSetRollingUpgradesClientStartExtensionUpgradePollerResponse{}, err
	}
	result.Poller = &VirtualMachineScaleSetRollingUpgradesClientStartExtensionUpgradePoller{
		pt: pt,
	}
	return result, nil
}

// StartExtensionUpgrade - Starts a rolling upgrade to move all extensions for all virtual machine scale set instances to
// the latest available extension version. Instances which are already running the latest extension versions
// are not affected.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualMachineScaleSetRollingUpgradesClient) startExtensionUpgrade(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesClientBeginStartExtensionUpgradeOptions) (*http.Response, error) {
	req, err := client.startExtensionUpgradeCreateRequest(ctx, resourceGroupName, vmScaleSetName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// startExtensionUpgradeCreateRequest creates the StartExtensionUpgrade request.
func (client *VirtualMachineScaleSetRollingUpgradesClient) startExtensionUpgradeCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesClientBeginStartExtensionUpgradeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensionRollingUpgrade"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// BeginStartOSUpgrade - Starts a rolling upgrade to move all virtual machine scale set instances to the latest available
// Platform Image OS version. Instances which are already running the latest available OS version are not
// affected.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// vmScaleSetName - The name of the VM scale set.
// options - VirtualMachineScaleSetRollingUpgradesClientBeginStartOSUpgradeOptions contains the optional parameters for the
// VirtualMachineScaleSetRollingUpgradesClient.BeginStartOSUpgrade method.
func (client *VirtualMachineScaleSetRollingUpgradesClient) BeginStartOSUpgrade(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesClientBeginStartOSUpgradeOptions) (VirtualMachineScaleSetRollingUpgradesClientStartOSUpgradePollerResponse, error) {
	resp, err := client.startOSUpgrade(ctx, resourceGroupName, vmScaleSetName, options)
	if err != nil {
		return VirtualMachineScaleSetRollingUpgradesClientStartOSUpgradePollerResponse{}, err
	}
	result := VirtualMachineScaleSetRollingUpgradesClientStartOSUpgradePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualMachineScaleSetRollingUpgradesClient.StartOSUpgrade", "", resp, client.pl)
	if err != nil {
		return VirtualMachineScaleSetRollingUpgradesClientStartOSUpgradePollerResponse{}, err
	}
	result.Poller = &VirtualMachineScaleSetRollingUpgradesClientStartOSUpgradePoller{
		pt: pt,
	}
	return result, nil
}

// StartOSUpgrade - Starts a rolling upgrade to move all virtual machine scale set instances to the latest available Platform
// Image OS version. Instances which are already running the latest available OS version are not
// affected.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualMachineScaleSetRollingUpgradesClient) startOSUpgrade(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesClientBeginStartOSUpgradeOptions) (*http.Response, error) {
	req, err := client.startOSUpgradeCreateRequest(ctx, resourceGroupName, vmScaleSetName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// startOSUpgradeCreateRequest creates the StartOSUpgrade request.
func (client *VirtualMachineScaleSetRollingUpgradesClient) startOSUpgradeCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesClientBeginStartOSUpgradeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/osRollingUpgrade"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}
