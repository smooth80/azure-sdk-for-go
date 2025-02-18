//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetapp_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp"
)

// x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-08-01/examples/SnapshotPolicies_List.json
func ExampleSnapshotPoliciesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetapp.NewSnapshotPoliciesClient("<subscription-id>", cred, nil)
	res, err := client.List(ctx,
		"<resource-group-name>",
		"<account-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SnapshotPoliciesClientListResult)
}

// x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-08-01/examples/SnapshotPolicies_Get.json
func ExampleSnapshotPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetapp.NewSnapshotPoliciesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<snapshot-policy-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SnapshotPoliciesClientGetResult)
}

// x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-08-01/examples/SnapshotPolicies_Create.json
func ExampleSnapshotPoliciesClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetapp.NewSnapshotPoliciesClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<snapshot-policy-name>",
		armnetapp.SnapshotPolicy{
			Location: to.StringPtr("<location>"),
			Properties: &armnetapp.SnapshotPolicyProperties{
				DailySchedule: &armnetapp.DailySchedule{
					Hour:            to.Int32Ptr(14),
					Minute:          to.Int32Ptr(30),
					SnapshotsToKeep: to.Int32Ptr(4),
				},
				Enabled: to.BoolPtr(true),
				HourlySchedule: &armnetapp.HourlySchedule{
					Minute:          to.Int32Ptr(50),
					SnapshotsToKeep: to.Int32Ptr(2),
				},
				MonthlySchedule: &armnetapp.MonthlySchedule{
					DaysOfMonth:     to.StringPtr("<days-of-month>"),
					Hour:            to.Int32Ptr(14),
					Minute:          to.Int32Ptr(15),
					SnapshotsToKeep: to.Int32Ptr(5),
				},
				WeeklySchedule: &armnetapp.WeeklySchedule{
					Day:             to.StringPtr("<day>"),
					Hour:            to.Int32Ptr(14),
					Minute:          to.Int32Ptr(45),
					SnapshotsToKeep: to.Int32Ptr(3),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SnapshotPoliciesClientCreateResult)
}

// x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-08-01/examples/SnapshotPolicies_Update.json
func ExampleSnapshotPoliciesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetapp.NewSnapshotPoliciesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<snapshot-policy-name>",
		armnetapp.SnapshotPolicyPatch{
			Location: to.StringPtr("<location>"),
			Properties: &armnetapp.SnapshotPolicyProperties{
				DailySchedule: &armnetapp.DailySchedule{
					Hour:            to.Int32Ptr(14),
					Minute:          to.Int32Ptr(30),
					SnapshotsToKeep: to.Int32Ptr(4),
				},
				Enabled: to.BoolPtr(true),
				HourlySchedule: &armnetapp.HourlySchedule{
					Minute:          to.Int32Ptr(50),
					SnapshotsToKeep: to.Int32Ptr(2),
				},
				MonthlySchedule: &armnetapp.MonthlySchedule{
					DaysOfMonth:     to.StringPtr("<days-of-month>"),
					Hour:            to.Int32Ptr(14),
					Minute:          to.Int32Ptr(15),
					SnapshotsToKeep: to.Int32Ptr(5),
				},
				WeeklySchedule: &armnetapp.WeeklySchedule{
					Day:             to.StringPtr("<day>"),
					Hour:            to.Int32Ptr(14),
					Minute:          to.Int32Ptr(45),
					SnapshotsToKeep: to.Int32Ptr(3),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SnapshotPoliciesClientUpdateResult)
}

// x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-08-01/examples/SnapshotPolicies_Delete.json
func ExampleSnapshotPoliciesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetapp.NewSnapshotPoliciesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<snapshot-policy-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-08-01/examples/SnapshotPolicies_ListVolumes.json
func ExampleSnapshotPoliciesClient_ListVolumes() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetapp.NewSnapshotPoliciesClient("<subscription-id>", cred, nil)
	res, err := client.ListVolumes(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<snapshot-policy-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SnapshotPoliciesClientListVolumesResult)
}
