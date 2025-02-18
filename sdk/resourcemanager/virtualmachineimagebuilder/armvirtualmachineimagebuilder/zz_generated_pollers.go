//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvirtualmachineimagebuilder

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// VirtualMachineImageTemplatesClientCancelPoller provides polling facilities until the operation reaches a terminal state.
type VirtualMachineImageTemplatesClientCancelPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VirtualMachineImageTemplatesClientCancelPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VirtualMachineImageTemplatesClientCancelPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VirtualMachineImageTemplatesClientCancelResponse will be returned.
func (p *VirtualMachineImageTemplatesClientCancelPoller) FinalResponse(ctx context.Context) (VirtualMachineImageTemplatesClientCancelResponse, error) {
	respType := VirtualMachineImageTemplatesClientCancelResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return VirtualMachineImageTemplatesClientCancelResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VirtualMachineImageTemplatesClientCancelPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VirtualMachineImageTemplatesClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type VirtualMachineImageTemplatesClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VirtualMachineImageTemplatesClientCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VirtualMachineImageTemplatesClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VirtualMachineImageTemplatesClientCreateOrUpdateResponse will be returned.
func (p *VirtualMachineImageTemplatesClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (VirtualMachineImageTemplatesClientCreateOrUpdateResponse, error) {
	respType := VirtualMachineImageTemplatesClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ImageTemplate)
	if err != nil {
		return VirtualMachineImageTemplatesClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VirtualMachineImageTemplatesClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VirtualMachineImageTemplatesClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type VirtualMachineImageTemplatesClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VirtualMachineImageTemplatesClientDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VirtualMachineImageTemplatesClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VirtualMachineImageTemplatesClientDeleteResponse will be returned.
func (p *VirtualMachineImageTemplatesClientDeletePoller) FinalResponse(ctx context.Context) (VirtualMachineImageTemplatesClientDeleteResponse, error) {
	respType := VirtualMachineImageTemplatesClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return VirtualMachineImageTemplatesClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VirtualMachineImageTemplatesClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VirtualMachineImageTemplatesClientRunPoller provides polling facilities until the operation reaches a terminal state.
type VirtualMachineImageTemplatesClientRunPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VirtualMachineImageTemplatesClientRunPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VirtualMachineImageTemplatesClientRunPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VirtualMachineImageTemplatesClientRunResponse will be returned.
func (p *VirtualMachineImageTemplatesClientRunPoller) FinalResponse(ctx context.Context) (VirtualMachineImageTemplatesClientRunResponse, error) {
	respType := VirtualMachineImageTemplatesClientRunResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return VirtualMachineImageTemplatesClientRunResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VirtualMachineImageTemplatesClientRunPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VirtualMachineImageTemplatesClientUpdatePoller provides polling facilities until the operation reaches a terminal state.
type VirtualMachineImageTemplatesClientUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VirtualMachineImageTemplatesClientUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VirtualMachineImageTemplatesClientUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VirtualMachineImageTemplatesClientUpdateResponse will be returned.
func (p *VirtualMachineImageTemplatesClientUpdatePoller) FinalResponse(ctx context.Context) (VirtualMachineImageTemplatesClientUpdateResponse, error) {
	respType := VirtualMachineImageTemplatesClientUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ImageTemplate)
	if err != nil {
		return VirtualMachineImageTemplatesClientUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VirtualMachineImageTemplatesClientUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}
