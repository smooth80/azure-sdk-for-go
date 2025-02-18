//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconfidentialledger_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confidentialledger/armconfidentialledger"
)

// x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2021-05-13-preview/examples/ConfidentialLedger_Get.json
func ExampleLedgerClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armconfidentialledger.NewLedgerClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<ledger-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.LedgerClientGetResult)
}

// x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2021-05-13-preview/examples/ConfidentialLedger_Delete.json
func ExampleLedgerClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armconfidentialledger.NewLedgerClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<ledger-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2021-05-13-preview/examples/ConfidentialLedger_Create.json
func ExampleLedgerClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armconfidentialledger.NewLedgerClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<ledger-name>",
		armconfidentialledger.ConfidentialLedger{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"additionalProps1": to.StringPtr("additional properties"),
			},
			Properties: &armconfidentialledger.LedgerProperties{
				AADBasedSecurityPrincipals: []*armconfidentialledger.AADBasedSecurityPrincipal{
					{
						LedgerRoleName: armconfidentialledger.LedgerRoleName("Administrator").ToPtr(),
						PrincipalID:    to.StringPtr("<principal-id>"),
						TenantID:       to.StringPtr("<tenant-id>"),
					}},
				CertBasedSecurityPrincipals: []*armconfidentialledger.CertBasedSecurityPrincipal{
					{
						Cert:           to.StringPtr("<cert>"),
						LedgerRoleName: armconfidentialledger.LedgerRoleName("Reader").ToPtr(),
					}},
				LedgerType: armconfidentialledger.LedgerType("Public").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.LedgerClientCreateResult)
}

// x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2021-05-13-preview/examples/ConfidentialLedger_Update.json
func ExampleLedgerClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armconfidentialledger.NewLedgerClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<ledger-name>",
		armconfidentialledger.ConfidentialLedger{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"additionProps2":   to.StringPtr("additional property value"),
				"additionalProps1": to.StringPtr("additional properties"),
			},
			Properties: &armconfidentialledger.LedgerProperties{},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.LedgerClientUpdateResult)
}

// x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2021-05-13-preview/examples/ConfidentialLedger_List.json
func ExampleLedgerClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armconfidentialledger.NewLedgerClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		&armconfidentialledger.LedgerClientListByResourceGroupOptions{Filter: nil})
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

// x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2021-05-13-preview/examples/ConfidentialLedger_ListBySub.json
func ExampleLedgerClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armconfidentialledger.NewLedgerClient("<subscription-id>", cred, nil)
	pager := client.ListBySubscription(&armconfidentialledger.LedgerClientListBySubscriptionOptions{Filter: nil})
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
