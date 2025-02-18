//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armserialconsole_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/serialconsole/armserialconsole"
)

// x-ms-original-file: specification/serialconsole/resource-manager/Microsoft.SerialConsole/stable/2018-05-01/examples/GetOperationsExample.json
func ExampleMicrosoftSerialConsoleClient_ListOperations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armserialconsole.NewMicrosoftSerialConsoleClient("<subscription-id>", cred, nil)
	res, err := client.ListOperations(ctx,
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MicrosoftSerialConsoleClientListOperationsResult)
}

// x-ms-original-file: specification/serialconsole/resource-manager/Microsoft.SerialConsole/stable/2018-05-01/examples/SerialConsoleStatus.json
func ExampleMicrosoftSerialConsoleClient_GetConsoleStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armserialconsole.NewMicrosoftSerialConsoleClient("<subscription-id>", cred, nil)
	res, err := client.GetConsoleStatus(ctx,
		"<default-param>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.Value)
}

// x-ms-original-file: specification/serialconsole/resource-manager/Microsoft.SerialConsole/stable/2018-05-01/examples/DisableConsoleExamples.json
func ExampleMicrosoftSerialConsoleClient_DisableConsole() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armserialconsole.NewMicrosoftSerialConsoleClient("<subscription-id>", cred, nil)
	res, err := client.DisableConsole(ctx,
		"<default-param>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.Value)
}

// x-ms-original-file: specification/serialconsole/resource-manager/Microsoft.SerialConsole/stable/2018-05-01/examples/EnableConsoleExamples.json
func ExampleMicrosoftSerialConsoleClient_EnableConsole() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armserialconsole.NewMicrosoftSerialConsoleClient("<subscription-id>", cred, nil)
	res, err := client.EnableConsole(ctx,
		"<default-param>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.Value)
}
