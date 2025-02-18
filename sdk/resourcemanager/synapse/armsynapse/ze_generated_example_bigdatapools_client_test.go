//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/GetBigDataPool.json
func ExampleBigDataPoolsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewBigDataPoolsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<big-data-pool-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BigDataPoolsClientGetResult)
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/UpdateBigDataPool.json
func ExampleBigDataPoolsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewBigDataPoolsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<big-data-pool-name>",
		armsynapse.BigDataPoolPatchInfo{
			Tags: map[string]*string{
				"key": to.StringPtr("value"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BigDataPoolsClientUpdateResult)
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/CreateOrUpdateBigDataPool.json
func ExampleBigDataPoolsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewBigDataPoolsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<big-data-pool-name>",
		armsynapse.BigDataPoolResourceInfo{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"key": to.StringPtr("value"),
			},
			Properties: &armsynapse.BigDataPoolResourceProperties{
				AutoPause: &armsynapse.AutoPauseProperties{
					DelayInMinutes: to.Int32Ptr(15),
					Enabled:        to.BoolPtr(true),
				},
				AutoScale: &armsynapse.AutoScaleProperties{
					Enabled:      to.BoolPtr(true),
					MaxNodeCount: to.Int32Ptr(50),
					MinNodeCount: to.Int32Ptr(3),
				},
				DefaultSparkLogFolder: to.StringPtr("<default-spark-log-folder>"),
				LibraryRequirements: &armsynapse.LibraryRequirements{
					Content:  to.StringPtr("<content>"),
					Filename: to.StringPtr("<filename>"),
				},
				NodeCount:         to.Int32Ptr(4),
				NodeSize:          armsynapse.NodeSize("Medium").ToPtr(),
				NodeSizeFamily:    armsynapse.NodeSizeFamily("MemoryOptimized").ToPtr(),
				SparkEventsFolder: to.StringPtr("<spark-events-folder>"),
				SparkVersion:      to.StringPtr("<spark-version>"),
			},
		},
		&armsynapse.BigDataPoolsClientBeginCreateOrUpdateOptions{Force: nil})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BigDataPoolsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/DeleteBigDataPool.json
func ExampleBigDataPoolsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewBigDataPoolsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<big-data-pool-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BigDataPoolsClientDeleteResult)
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/ListBigDataPoolsInWorkspace.json
func ExampleBigDataPoolsClient_ListByWorkspace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewBigDataPoolsClient("<subscription-id>", cred, nil)
	pager := client.ListByWorkspace("<resource-group-name>",
		"<workspace-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}
