# Release History

## 0.2.0 (2022-01-13)
### Breaking Changes

- Function `*IscsiTargetsClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, string, IscsiTargetUpdate, *IscsiTargetsBeginUpdateOptions)` to `(context.Context, string, string, string, IscsiTargetUpdate, *IscsiTargetsClientBeginUpdateOptions)`
- Function `*IscsiTargetsClient.BeginUpdate` return value(s) have been changed from `(IscsiTargetsUpdatePollerResponse, error)` to `(IscsiTargetsClientUpdatePollerResponse, error)`
- Function `*IscsiTargetsClient.ListByDiskPool` parameter(s) have been changed from `(string, string, *IscsiTargetsListByDiskPoolOptions)` to `(string, string, *IscsiTargetsClientListByDiskPoolOptions)`
- Function `*IscsiTargetsClient.ListByDiskPool` return value(s) have been changed from `(*IscsiTargetsListByDiskPoolPager)` to `(*IscsiTargetsClientListByDiskPoolPager)`
- Function `*DiskPoolsClient.BeginStart` parameter(s) have been changed from `(context.Context, string, string, *DiskPoolsBeginStartOptions)` to `(context.Context, string, string, *DiskPoolsClientBeginStartOptions)`
- Function `*DiskPoolsClient.BeginStart` return value(s) have been changed from `(DiskPoolsStartPollerResponse, error)` to `(DiskPoolsClientStartPollerResponse, error)`
- Function `*DiskPoolsClient.Get` parameter(s) have been changed from `(context.Context, string, string, *DiskPoolsGetOptions)` to `(context.Context, string, string, *DiskPoolsClientGetOptions)`
- Function `*DiskPoolsClient.Get` return value(s) have been changed from `(DiskPoolsGetResponse, error)` to `(DiskPoolsClientGetResponse, error)`
- Function `*DiskPoolsClient.BeginDeallocate` parameter(s) have been changed from `(context.Context, string, string, *DiskPoolsBeginDeallocateOptions)` to `(context.Context, string, string, *DiskPoolsClientBeginDeallocateOptions)`
- Function `*DiskPoolsClient.BeginDeallocate` return value(s) have been changed from `(DiskPoolsDeallocatePollerResponse, error)` to `(DiskPoolsClientDeallocatePollerResponse, error)`
- Function `*DiskPoolsClient.ListByResourceGroup` parameter(s) have been changed from `(string, *DiskPoolsListByResourceGroupOptions)` to `(string, *DiskPoolsClientListByResourceGroupOptions)`
- Function `*DiskPoolsClient.ListByResourceGroup` return value(s) have been changed from `(*DiskPoolsListByResourceGroupPager)` to `(*DiskPoolsClientListByResourceGroupPager)`
- Function `*OperationsClient.List` parameter(s) have been changed from `(context.Context, *OperationsListOptions)` to `(context.Context, *OperationsClientListOptions)`
- Function `*OperationsClient.List` return value(s) have been changed from `(OperationsListResponse, error)` to `(OperationsClientListResponse, error)`
- Function `*DiskPoolsClient.BeginCreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, DiskPoolCreate, *DiskPoolsBeginCreateOrUpdateOptions)` to `(context.Context, string, string, DiskPoolCreate, *DiskPoolsClientBeginCreateOrUpdateOptions)`
- Function `*DiskPoolsClient.BeginCreateOrUpdate` return value(s) have been changed from `(DiskPoolsCreateOrUpdatePollerResponse, error)` to `(DiskPoolsClientCreateOrUpdatePollerResponse, error)`
- Function `*DiskPoolZonesClient.List` parameter(s) have been changed from `(string, *DiskPoolZonesListOptions)` to `(string, *DiskPoolZonesClientListOptions)`
- Function `*DiskPoolZonesClient.List` return value(s) have been changed from `(*DiskPoolZonesListPager)` to `(*DiskPoolZonesClientListPager)`
- Function `*IscsiTargetsClient.Get` parameter(s) have been changed from `(context.Context, string, string, string, *IscsiTargetsGetOptions)` to `(context.Context, string, string, string, *IscsiTargetsClientGetOptions)`
- Function `*IscsiTargetsClient.Get` return value(s) have been changed from `(IscsiTargetsGetResponse, error)` to `(IscsiTargetsClientGetResponse, error)`
- Function `*ResourceSKUsClient.List` parameter(s) have been changed from `(string, *ResourceSKUsListOptions)` to `(string, *ResourceSKUsClientListOptions)`
- Function `*ResourceSKUsClient.List` return value(s) have been changed from `(*ResourceSKUsListPager)` to `(*ResourceSKUsClientListPager)`
- Function `*DiskPoolsClient.ListBySubscription` parameter(s) have been changed from `(*DiskPoolsListBySubscriptionOptions)` to `(*DiskPoolsClientListBySubscriptionOptions)`
- Function `*DiskPoolsClient.ListBySubscription` return value(s) have been changed from `(*DiskPoolsListBySubscriptionPager)` to `(*DiskPoolsClientListBySubscriptionPager)`
- Function `*DiskPoolsClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, DiskPoolUpdate, *DiskPoolsBeginUpdateOptions)` to `(context.Context, string, string, DiskPoolUpdate, *DiskPoolsClientBeginUpdateOptions)`
- Function `*DiskPoolsClient.BeginUpdate` return value(s) have been changed from `(DiskPoolsUpdatePollerResponse, error)` to `(DiskPoolsClientUpdatePollerResponse, error)`
- Function `*DiskPoolsClient.ListOutboundNetworkDependenciesEndpoints` parameter(s) have been changed from `(string, string, *DiskPoolsListOutboundNetworkDependenciesEndpointsOptions)` to `(string, string, *DiskPoolsClientListOutboundNetworkDependenciesEndpointsOptions)`
- Function `*DiskPoolsClient.ListOutboundNetworkDependenciesEndpoints` return value(s) have been changed from `(*DiskPoolsListOutboundNetworkDependenciesEndpointsPager)` to `(*DiskPoolsClientListOutboundNetworkDependenciesEndpointsPager)`
- Function `*DiskPoolsClient.BeginUpgrade` parameter(s) have been changed from `(context.Context, string, string, *DiskPoolsBeginUpgradeOptions)` to `(context.Context, string, string, *DiskPoolsClientBeginUpgradeOptions)`
- Function `*DiskPoolsClient.BeginUpgrade` return value(s) have been changed from `(DiskPoolsUpgradePollerResponse, error)` to `(DiskPoolsClientUpgradePollerResponse, error)`
- Function `*IscsiTargetsClient.BeginCreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, IscsiTargetCreate, *IscsiTargetsBeginCreateOrUpdateOptions)` to `(context.Context, string, string, string, IscsiTargetCreate, *IscsiTargetsClientBeginCreateOrUpdateOptions)`
- Function `*IscsiTargetsClient.BeginCreateOrUpdate` return value(s) have been changed from `(IscsiTargetsCreateOrUpdatePollerResponse, error)` to `(IscsiTargetsClientCreateOrUpdatePollerResponse, error)`
- Function `*IscsiTargetsClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, string, string, *IscsiTargetsBeginDeleteOptions)` to `(context.Context, string, string, string, *IscsiTargetsClientBeginDeleteOptions)`
- Function `*IscsiTargetsClient.BeginDelete` return value(s) have been changed from `(IscsiTargetsDeletePollerResponse, error)` to `(IscsiTargetsClientDeletePollerResponse, error)`
- Function `*DiskPoolsClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, string, *DiskPoolsBeginDeleteOptions)` to `(context.Context, string, string, *DiskPoolsClientBeginDeleteOptions)`
- Function `*DiskPoolsClient.BeginDelete` return value(s) have been changed from `(DiskPoolsDeletePollerResponse, error)` to `(DiskPoolsClientDeletePollerResponse, error)`
- Function `*IscsiTargetsUpdatePoller.FinalResponse` has been removed
- Function `StoragePoolOperationListResult.MarshalJSON` has been removed
- Function `*DiskPoolsUpdatePoller.Done` has been removed
- Function `*DiskPoolsUpgradePoller.Done` has been removed
- Function `*DiskPoolsCreateOrUpdatePoller.Done` has been removed
- Function `*DiskPoolsUpdatePoller.ResumeToken` has been removed
- Function `DiskPoolsUpdatePollerResponse.PollUntilDone` has been removed
- Function `*IscsiTargetsUpdatePoller.ResumeToken` has been removed
- Function `*DiskPoolsDeletePoller.Poll` has been removed
- Function `*IscsiTargetsListByDiskPoolPager.PageResponse` has been removed
- Function `Error.Error` has been removed
- Function `*DiskPoolsUpgradePoller.ResumeToken` has been removed
- Function `*IscsiTargetsCreateOrUpdatePollerResponse.Resume` has been removed
- Function `*DiskPoolsCreateOrUpdatePoller.ResumeToken` has been removed
- Function `*DiskPoolsStartPoller.Done` has been removed
- Function `*DiskPoolsUpdatePollerResponse.Resume` has been removed
- Function `*DiskPoolsUpdatePoller.FinalResponse` has been removed
- Function `*ResourceSKUsListPager.Err` has been removed
- Function `*DiskPoolsListOutboundNetworkDependenciesEndpointsPager.PageResponse` has been removed
- Function `Resource.MarshalJSON` has been removed
- Function `*IscsiTargetsDeletePoller.Done` has been removed
- Function `*IscsiTargetsCreateOrUpdatePoller.ResumeToken` has been removed
- Function `*DiskPoolsCreateOrUpdatePoller.Poll` has been removed
- Function `*IscsiTargetsUpdatePoller.Done` has been removed
- Function `*DiskPoolsStartPollerResponse.Resume` has been removed
- Function `*DiskPoolsDeallocatePoller.FinalResponse` has been removed
- Function `*DiskPoolsDeallocatePoller.Done` has been removed
- Function `DiskPoolsDeletePollerResponse.PollUntilDone` has been removed
- Function `*DiskPoolZonesListPager.PageResponse` has been removed
- Function `*DiskPoolsCreateOrUpdatePollerResponse.Resume` has been removed
- Function `*DiskPoolsDeallocatePoller.Poll` has been removed
- Function `*DiskPoolsListBySubscriptionPager.Err` has been removed
- Function `*DiskPoolsStartPoller.ResumeToken` has been removed
- Function `*DiskPoolsListOutboundNetworkDependenciesEndpointsPager.Err` has been removed
- Function `*IscsiTargetsCreateOrUpdatePoller.Poll` has been removed
- Function `*IscsiTargetsUpdatePoller.Poll` has been removed
- Function `*DiskPoolsCreateOrUpdatePoller.FinalResponse` has been removed
- Function `*IscsiTargetsDeletePollerResponse.Resume` has been removed
- Function `DiskPoolsUpgradePollerResponse.PollUntilDone` has been removed
- Function `*ResourceSKUsListPager.PageResponse` has been removed
- Function `*DiskPoolsStartPoller.FinalResponse` has been removed
- Function `*IscsiTargetsUpdatePollerResponse.Resume` has been removed
- Function `DiskPoolsDeallocatePollerResponse.PollUntilDone` has been removed
- Function `*IscsiTargetsCreateOrUpdatePoller.FinalResponse` has been removed
- Function `*DiskPoolsDeletePollerResponse.Resume` has been removed
- Function `*IscsiTargetsListByDiskPoolPager.Err` has been removed
- Function `*DiskPoolsListByResourceGroupPager.NextPage` has been removed
- Function `*IscsiTargetsDeletePoller.ResumeToken` has been removed
- Function `*DiskPoolZonesListPager.Err` has been removed
- Function `*DiskPoolsUpgradePoller.Poll` has been removed
- Function `IscsiTargetsCreateOrUpdatePollerResponse.PollUntilDone` has been removed
- Function `*DiskPoolsListOutboundNetworkDependenciesEndpointsPager.NextPage` has been removed
- Function `*DiskPoolZonesListPager.NextPage` has been removed
- Function `*DiskPoolsStartPoller.Poll` has been removed
- Function `*DiskPoolsListByResourceGroupPager.Err` has been removed
- Function `*DiskPoolsListByResourceGroupPager.PageResponse` has been removed
- Function `*DiskPoolsDeallocatePoller.ResumeToken` has been removed
- Function `*DiskPoolsDeletePoller.FinalResponse` has been removed
- Function `IscsiTargetsDeletePollerResponse.PollUntilDone` has been removed
- Function `*DiskPoolsListBySubscriptionPager.NextPage` has been removed
- Function `*IscsiTargetsDeletePoller.Poll` has been removed
- Function `*DiskPoolsUpdatePoller.Poll` has been removed
- Function `*DiskPoolsDeallocatePollerResponse.Resume` has been removed
- Function `DiskPoolsCreateOrUpdatePollerResponse.PollUntilDone` has been removed
- Function `*ResourceSKUsListPager.NextPage` has been removed
- Function `*DiskPoolsUpgradePoller.FinalResponse` has been removed
- Function `IscsiTargetsUpdatePollerResponse.PollUntilDone` has been removed
- Function `*IscsiTargetsCreateOrUpdatePoller.Done` has been removed
- Function `*DiskPoolsDeletePoller.Done` has been removed
- Function `DiskPoolsStartPollerResponse.PollUntilDone` has been removed
- Function `*DiskPoolsListBySubscriptionPager.PageResponse` has been removed
- Function `*DiskPoolsUpgradePollerResponse.Resume` has been removed
- Function `*IscsiTargetsDeletePoller.FinalResponse` has been removed
- Function `*IscsiTargetsListByDiskPoolPager.NextPage` has been removed
- Function `*DiskPoolsDeletePoller.ResumeToken` has been removed
- Struct `DiskPoolZonesListOptions` has been removed
- Struct `DiskPoolZonesListPager` has been removed
- Struct `DiskPoolZonesListResponse` has been removed
- Struct `DiskPoolZonesListResult` has been removed
- Struct `DiskPoolsBeginCreateOrUpdateOptions` has been removed
- Struct `DiskPoolsBeginDeallocateOptions` has been removed
- Struct `DiskPoolsBeginDeleteOptions` has been removed
- Struct `DiskPoolsBeginStartOptions` has been removed
- Struct `DiskPoolsBeginUpdateOptions` has been removed
- Struct `DiskPoolsBeginUpgradeOptions` has been removed
- Struct `DiskPoolsCreateOrUpdatePoller` has been removed
- Struct `DiskPoolsCreateOrUpdatePollerResponse` has been removed
- Struct `DiskPoolsCreateOrUpdateResponse` has been removed
- Struct `DiskPoolsCreateOrUpdateResult` has been removed
- Struct `DiskPoolsDeallocatePoller` has been removed
- Struct `DiskPoolsDeallocatePollerResponse` has been removed
- Struct `DiskPoolsDeallocateResponse` has been removed
- Struct `DiskPoolsDeletePoller` has been removed
- Struct `DiskPoolsDeletePollerResponse` has been removed
- Struct `DiskPoolsDeleteResponse` has been removed
- Struct `DiskPoolsGetOptions` has been removed
- Struct `DiskPoolsGetResponse` has been removed
- Struct `DiskPoolsGetResult` has been removed
- Struct `DiskPoolsListByResourceGroupOptions` has been removed
- Struct `DiskPoolsListByResourceGroupPager` has been removed
- Struct `DiskPoolsListByResourceGroupResponse` has been removed
- Struct `DiskPoolsListByResourceGroupResult` has been removed
- Struct `DiskPoolsListBySubscriptionOptions` has been removed
- Struct `DiskPoolsListBySubscriptionPager` has been removed
- Struct `DiskPoolsListBySubscriptionResponse` has been removed
- Struct `DiskPoolsListBySubscriptionResult` has been removed
- Struct `DiskPoolsListOutboundNetworkDependenciesEndpointsOptions` has been removed
- Struct `DiskPoolsListOutboundNetworkDependenciesEndpointsPager` has been removed
- Struct `DiskPoolsListOutboundNetworkDependenciesEndpointsResponse` has been removed
- Struct `DiskPoolsListOutboundNetworkDependenciesEndpointsResult` has been removed
- Struct `DiskPoolsStartPoller` has been removed
- Struct `DiskPoolsStartPollerResponse` has been removed
- Struct `DiskPoolsStartResponse` has been removed
- Struct `DiskPoolsUpdatePoller` has been removed
- Struct `DiskPoolsUpdatePollerResponse` has been removed
- Struct `DiskPoolsUpdateResponse` has been removed
- Struct `DiskPoolsUpdateResult` has been removed
- Struct `DiskPoolsUpgradePoller` has been removed
- Struct `DiskPoolsUpgradePollerResponse` has been removed
- Struct `DiskPoolsUpgradeResponse` has been removed
- Struct `IscsiTargetsBeginCreateOrUpdateOptions` has been removed
- Struct `IscsiTargetsBeginDeleteOptions` has been removed
- Struct `IscsiTargetsBeginUpdateOptions` has been removed
- Struct `IscsiTargetsCreateOrUpdatePoller` has been removed
- Struct `IscsiTargetsCreateOrUpdatePollerResponse` has been removed
- Struct `IscsiTargetsCreateOrUpdateResponse` has been removed
- Struct `IscsiTargetsCreateOrUpdateResult` has been removed
- Struct `IscsiTargetsDeletePoller` has been removed
- Struct `IscsiTargetsDeletePollerResponse` has been removed
- Struct `IscsiTargetsDeleteResponse` has been removed
- Struct `IscsiTargetsGetOptions` has been removed
- Struct `IscsiTargetsGetResponse` has been removed
- Struct `IscsiTargetsGetResult` has been removed
- Struct `IscsiTargetsListByDiskPoolOptions` has been removed
- Struct `IscsiTargetsListByDiskPoolPager` has been removed
- Struct `IscsiTargetsListByDiskPoolResponse` has been removed
- Struct `IscsiTargetsListByDiskPoolResult` has been removed
- Struct `IscsiTargetsUpdatePoller` has been removed
- Struct `IscsiTargetsUpdatePollerResponse` has been removed
- Struct `IscsiTargetsUpdateResponse` has been removed
- Struct `IscsiTargetsUpdateResult` has been removed
- Struct `OperationsListOptions` has been removed
- Struct `OperationsListResponse` has been removed
- Struct `OperationsListResult` has been removed
- Struct `ResourceSKUsListOptions` has been removed
- Struct `ResourceSKUsListPager` has been removed
- Struct `ResourceSKUsListResponse` has been removed
- Struct `ResourceSKUsListResult` has been removed
- Struct `StoragePoolOperationDisplay` has been removed
- Struct `StoragePoolOperationListResult` has been removed
- Struct `StoragePoolRPOperation` has been removed
- Field `TrackedResource` of struct `DiskPool` has been removed
- Field `ProxyResource` of struct `IscsiTargetCreate` has been removed
- Field `Resource` of struct `TrackedResource` has been removed
- Field `InnerError` of struct `Error` has been removed
- Field `ProxyResource` of struct `IscsiTarget` has been removed
- Field `Resource` of struct `ProxyResource` has been removed
- Field `ProxyResource` of struct `IscsiTargetUpdate` has been removed

### Features Added

- New function `*DiskPoolsClientUpgradePoller.Done() bool`
- New function `*DiskPoolsClientStartPoller.FinalResponse(context.Context) (DiskPoolsClientStartResponse, error)`
- New function `*DiskPoolsClientListBySubscriptionPager.Err() error`
- New function `*IscsiTargetsClientCreateOrUpdatePoller.Poll(context.Context) (*http.Response, error)`
- New function `DiskPoolsClientStartPollerResponse.PollUntilDone(context.Context, time.Duration) (DiskPoolsClientStartResponse, error)`
- New function `*DiskPoolsClientUpgradePoller.ResumeToken() (string, error)`
- New function `*IscsiTargetsClientUpdatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*IscsiTargetsClientUpdatePoller.FinalResponse(context.Context) (IscsiTargetsClientUpdateResponse, error)`
- New function `*DiskPoolsClientListBySubscriptionPager.PageResponse() DiskPoolsClientListBySubscriptionResponse`
- New function `*DiskPoolsClientStartPoller.Done() bool`
- New function `*DiskPoolsClientListByResourceGroupPager.PageResponse() DiskPoolsClientListByResourceGroupResponse`
- New function `*DiskPoolsClientStartPollerResponse.Resume(context.Context, *DiskPoolsClient, string) error`
- New function `*IscsiTargetsClientCreateOrUpdatePollerResponse.Resume(context.Context, *IscsiTargetsClient, string) error`
- New function `*ResourceSKUsClientListPager.NextPage(context.Context) bool`
- New function `IscsiTargetsClientCreateOrUpdatePollerResponse.PollUntilDone(context.Context, time.Duration) (IscsiTargetsClientCreateOrUpdateResponse, error)`
- New function `*DiskPoolsClientUpgradePoller.Poll(context.Context) (*http.Response, error)`
- New function `*DiskPoolsClientUpgradePollerResponse.Resume(context.Context, *DiskPoolsClient, string) error`
- New function `*DiskPoolZonesClientListPager.Err() error`
- New function `IscsiTargetsClientUpdatePollerResponse.PollUntilDone(context.Context, time.Duration) (IscsiTargetsClientUpdateResponse, error)`
- New function `*IscsiTargetsClientDeletePoller.Poll(context.Context) (*http.Response, error)`
- New function `*IscsiTargetsClientUpdatePoller.ResumeToken() (string, error)`
- New function `*IscsiTargetsClientCreateOrUpdatePoller.Done() bool`
- New function `*DiskPoolsClientDeletePollerResponse.Resume(context.Context, *DiskPoolsClient, string) error`
- New function `*ResourceSKUsClientListPager.PageResponse() ResourceSKUsClientListResponse`
- New function `*IscsiTargetsClientUpdatePoller.Done() bool`
- New function `*DiskPoolsClientListOutboundNetworkDependenciesEndpointsPager.Err() error`
- New function `DiskPoolsClientDeletePollerResponse.PollUntilDone(context.Context, time.Duration) (DiskPoolsClientDeleteResponse, error)`
- New function `*IscsiTargetsClientUpdatePollerResponse.Resume(context.Context, *IscsiTargetsClient, string) error`
- New function `*DiskPoolsClientCreateOrUpdatePollerResponse.Resume(context.Context, *DiskPoolsClient, string) error`
- New function `*DiskPoolsClientDeallocatePoller.FinalResponse(context.Context) (DiskPoolsClientDeallocateResponse, error)`
- New function `IscsiTargetsClientDeletePollerResponse.PollUntilDone(context.Context, time.Duration) (IscsiTargetsClientDeleteResponse, error)`
- New function `*DiskPoolsClientDeallocatePollerResponse.Resume(context.Context, *DiskPoolsClient, string) error`
- New function `*DiskPoolsClientUpdatePoller.Done() bool`
- New function `*IscsiTargetsClientDeletePoller.FinalResponse(context.Context) (IscsiTargetsClientDeleteResponse, error)`
- New function `*DiskPoolsClientDeletePoller.ResumeToken() (string, error)`
- New function `*DiskPoolsClientCreateOrUpdatePoller.Done() bool`
- New function `*IscsiTargetsClientListByDiskPoolPager.Err() error`
- New function `*DiskPoolsClientListByResourceGroupPager.NextPage(context.Context) bool`
- New function `*DiskPoolsClientUpdatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*IscsiTargetsClientListByDiskPoolPager.PageResponse() IscsiTargetsClientListByDiskPoolResponse`
- New function `*DiskPoolsClientUpdatePoller.FinalResponse(context.Context) (DiskPoolsClientUpdateResponse, error)`
- New function `DiskPoolsClientDeallocatePollerResponse.PollUntilDone(context.Context, time.Duration) (DiskPoolsClientDeallocateResponse, error)`
- New function `*DiskPoolsClientDeallocatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*DiskPoolsClientDeletePoller.FinalResponse(context.Context) (DiskPoolsClientDeleteResponse, error)`
- New function `*DiskPoolsClientStartPoller.Poll(context.Context) (*http.Response, error)`
- New function `*IscsiTargetsClientDeletePoller.ResumeToken() (string, error)`
- New function `*DiskPoolsClientDeletePoller.Done() bool`
- New function `*DiskPoolZonesClientListPager.NextPage(context.Context) bool`
- New function `*DiskPoolsClientDeallocatePoller.ResumeToken() (string, error)`
- New function `*DiskPoolsClientDeallocatePoller.Done() bool`
- New function `*DiskPoolZonesClientListPager.PageResponse() DiskPoolZonesClientListResponse`
- New function `*DiskPoolsClientListByResourceGroupPager.Err() error`
- New function `*DiskPoolsClientCreateOrUpdatePoller.Poll(context.Context) (*http.Response, error)`
- New function `DiskPoolsClientCreateOrUpdatePollerResponse.PollUntilDone(context.Context, time.Duration) (DiskPoolsClientCreateOrUpdateResponse, error)`
- New function `*DiskPoolsClientStartPoller.ResumeToken() (string, error)`
- New function `*DiskPoolsClientUpgradePoller.FinalResponse(context.Context) (DiskPoolsClientUpgradeResponse, error)`
- New function `*IscsiTargetsClientCreateOrUpdatePoller.FinalResponse(context.Context) (IscsiTargetsClientCreateOrUpdateResponse, error)`
- New function `*IscsiTargetsClientDeletePollerResponse.Resume(context.Context, *IscsiTargetsClient, string) error`
- New function `OperationListResult.MarshalJSON() ([]byte, error)`
- New function `*DiskPoolsClientListOutboundNetworkDependenciesEndpointsPager.NextPage(context.Context) bool`
- New function `*IscsiTargetsClientDeletePoller.Done() bool`
- New function `*DiskPoolsClientUpdatePollerResponse.Resume(context.Context, *DiskPoolsClient, string) error`
- New function `*DiskPoolsClientDeletePoller.Poll(context.Context) (*http.Response, error)`
- New function `*DiskPoolsClientCreateOrUpdatePoller.ResumeToken() (string, error)`
- New function `*ResourceSKUsClientListPager.Err() error`
- New function `*DiskPoolsClientUpdatePoller.ResumeToken() (string, error)`
- New function `*DiskPoolsClientListBySubscriptionPager.NextPage(context.Context) bool`
- New function `*DiskPoolsClientListOutboundNetworkDependenciesEndpointsPager.PageResponse() DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse`
- New function `*IscsiTargetsClientListByDiskPoolPager.NextPage(context.Context) bool`
- New function `*IscsiTargetsClientCreateOrUpdatePoller.ResumeToken() (string, error)`
- New function `*DiskPoolsClientCreateOrUpdatePoller.FinalResponse(context.Context) (DiskPoolsClientCreateOrUpdateResponse, error)`
- New function `DiskPoolsClientUpdatePollerResponse.PollUntilDone(context.Context, time.Duration) (DiskPoolsClientUpdateResponse, error)`
- New function `DiskPoolsClientUpgradePollerResponse.PollUntilDone(context.Context, time.Duration) (DiskPoolsClientUpgradeResponse, error)`
- New struct `DiskPoolZonesClientListOptions`
- New struct `DiskPoolZonesClientListPager`
- New struct `DiskPoolZonesClientListResponse`
- New struct `DiskPoolZonesClientListResult`
- New struct `DiskPoolsClientBeginCreateOrUpdateOptions`
- New struct `DiskPoolsClientBeginDeallocateOptions`
- New struct `DiskPoolsClientBeginDeleteOptions`
- New struct `DiskPoolsClientBeginStartOptions`
- New struct `DiskPoolsClientBeginUpdateOptions`
- New struct `DiskPoolsClientBeginUpgradeOptions`
- New struct `DiskPoolsClientCreateOrUpdatePoller`
- New struct `DiskPoolsClientCreateOrUpdatePollerResponse`
- New struct `DiskPoolsClientCreateOrUpdateResponse`
- New struct `DiskPoolsClientCreateOrUpdateResult`
- New struct `DiskPoolsClientDeallocatePoller`
- New struct `DiskPoolsClientDeallocatePollerResponse`
- New struct `DiskPoolsClientDeallocateResponse`
- New struct `DiskPoolsClientDeletePoller`
- New struct `DiskPoolsClientDeletePollerResponse`
- New struct `DiskPoolsClientDeleteResponse`
- New struct `DiskPoolsClientGetOptions`
- New struct `DiskPoolsClientGetResponse`
- New struct `DiskPoolsClientGetResult`
- New struct `DiskPoolsClientListByResourceGroupOptions`
- New struct `DiskPoolsClientListByResourceGroupPager`
- New struct `DiskPoolsClientListByResourceGroupResponse`
- New struct `DiskPoolsClientListByResourceGroupResult`
- New struct `DiskPoolsClientListBySubscriptionOptions`
- New struct `DiskPoolsClientListBySubscriptionPager`
- New struct `DiskPoolsClientListBySubscriptionResponse`
- New struct `DiskPoolsClientListBySubscriptionResult`
- New struct `DiskPoolsClientListOutboundNetworkDependenciesEndpointsOptions`
- New struct `DiskPoolsClientListOutboundNetworkDependenciesEndpointsPager`
- New struct `DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse`
- New struct `DiskPoolsClientListOutboundNetworkDependenciesEndpointsResult`
- New struct `DiskPoolsClientStartPoller`
- New struct `DiskPoolsClientStartPollerResponse`
- New struct `DiskPoolsClientStartResponse`
- New struct `DiskPoolsClientUpdatePoller`
- New struct `DiskPoolsClientUpdatePollerResponse`
- New struct `DiskPoolsClientUpdateResponse`
- New struct `DiskPoolsClientUpdateResult`
- New struct `DiskPoolsClientUpgradePoller`
- New struct `DiskPoolsClientUpgradePollerResponse`
- New struct `DiskPoolsClientUpgradeResponse`
- New struct `IscsiTargetsClientBeginCreateOrUpdateOptions`
- New struct `IscsiTargetsClientBeginDeleteOptions`
- New struct `IscsiTargetsClientBeginUpdateOptions`
- New struct `IscsiTargetsClientCreateOrUpdatePoller`
- New struct `IscsiTargetsClientCreateOrUpdatePollerResponse`
- New struct `IscsiTargetsClientCreateOrUpdateResponse`
- New struct `IscsiTargetsClientCreateOrUpdateResult`
- New struct `IscsiTargetsClientDeletePoller`
- New struct `IscsiTargetsClientDeletePollerResponse`
- New struct `IscsiTargetsClientDeleteResponse`
- New struct `IscsiTargetsClientGetOptions`
- New struct `IscsiTargetsClientGetResponse`
- New struct `IscsiTargetsClientGetResult`
- New struct `IscsiTargetsClientListByDiskPoolOptions`
- New struct `IscsiTargetsClientListByDiskPoolPager`
- New struct `IscsiTargetsClientListByDiskPoolResponse`
- New struct `IscsiTargetsClientListByDiskPoolResult`
- New struct `IscsiTargetsClientUpdatePoller`
- New struct `IscsiTargetsClientUpdatePollerResponse`
- New struct `IscsiTargetsClientUpdateResponse`
- New struct `IscsiTargetsClientUpdateResult`
- New struct `OperationDisplay`
- New struct `OperationListResult`
- New struct `OperationsClientListOptions`
- New struct `OperationsClientListResponse`
- New struct `OperationsClientListResult`
- New struct `RPOperation`
- New struct `ResourceSKUsClientListOptions`
- New struct `ResourceSKUsClientListPager`
- New struct `ResourceSKUsClientListResponse`
- New struct `ResourceSKUsClientListResult`
- New field `ID` in struct `IscsiTarget`
- New field `Name` in struct `IscsiTarget`
- New field `Type` in struct `IscsiTarget`
- New field `Error` in struct `Error`
- New field `ID` in struct `TrackedResource`
- New field `Name` in struct `TrackedResource`
- New field `Type` in struct `TrackedResource`
- New field `Name` in struct `IscsiTargetUpdate`
- New field `Type` in struct `IscsiTargetUpdate`
- New field `ID` in struct `IscsiTargetUpdate`
- New field `Type` in struct `IscsiTargetCreate`
- New field `ID` in struct `IscsiTargetCreate`
- New field `Name` in struct `IscsiTargetCreate`
- New field `Location` in struct `DiskPool`
- New field `Tags` in struct `DiskPool`
- New field `ID` in struct `DiskPool`
- New field `Name` in struct `DiskPool`
- New field `Type` in struct `DiskPool`
- New field `Name` in struct `ProxyResource`
- New field `Type` in struct `ProxyResource`
- New field `ID` in struct `ProxyResource`


## 0.1.0 (2021-12-16)

- Init release.
