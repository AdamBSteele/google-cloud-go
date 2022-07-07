// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package iot_test

import (
	"context"

	iot "cloud.google.com/go/iot/apiv1"
	"google.golang.org/api/iterator"
	iotpb "google.golang.org/genproto/googleapis/cloud/iot/v1"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
)

func ExampleNewDeviceManagerClient() {
	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleDeviceManagerClient_CreateDeviceRegistry() {
	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iotpb.CreateDeviceRegistryRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/iot/v1#CreateDeviceRegistryRequest.
	}
	resp, err := c.CreateDeviceRegistry(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_GetDeviceRegistry() {
	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iotpb.GetDeviceRegistryRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/iot/v1#GetDeviceRegistryRequest.
	}
	resp, err := c.GetDeviceRegistry(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_UpdateDeviceRegistry() {
	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iotpb.UpdateDeviceRegistryRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/iot/v1#UpdateDeviceRegistryRequest.
	}
	resp, err := c.UpdateDeviceRegistry(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_DeleteDeviceRegistry() {
	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iotpb.DeleteDeviceRegistryRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/iot/v1#DeleteDeviceRegistryRequest.
	}
	err = c.DeleteDeviceRegistry(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleDeviceManagerClient_ListDeviceRegistries() {
	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iotpb.ListDeviceRegistriesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/iot/v1#ListDeviceRegistriesRequest.
	}
	it := c.ListDeviceRegistries(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleDeviceManagerClient_CreateDevice() {
	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iotpb.CreateDeviceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/iot/v1#CreateDeviceRequest.
	}
	resp, err := c.CreateDevice(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_GetDevice() {
	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iotpb.GetDeviceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/iot/v1#GetDeviceRequest.
	}
	resp, err := c.GetDevice(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_UpdateDevice() {
	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iotpb.UpdateDeviceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/iot/v1#UpdateDeviceRequest.
	}
	resp, err := c.UpdateDevice(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_DeleteDevice() {
	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iotpb.DeleteDeviceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/iot/v1#DeleteDeviceRequest.
	}
	err = c.DeleteDevice(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleDeviceManagerClient_ListDevices() {
	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iotpb.ListDevicesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/iot/v1#ListDevicesRequest.
	}
	it := c.ListDevices(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleDeviceManagerClient_ModifyCloudToDeviceConfig() {
	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iotpb.ModifyCloudToDeviceConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/iot/v1#ModifyCloudToDeviceConfigRequest.
	}
	resp, err := c.ModifyCloudToDeviceConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_ListDeviceConfigVersions() {
	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iotpb.ListDeviceConfigVersionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/iot/v1#ListDeviceConfigVersionsRequest.
	}
	resp, err := c.ListDeviceConfigVersions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_ListDeviceStates() {
	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iotpb.ListDeviceStatesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/iot/v1#ListDeviceStatesRequest.
	}
	resp, err := c.ListDeviceStates(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_SetIamPolicy() {
	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.SetIamPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/iam/v1#SetIamPolicyRequest.
	}
	resp, err := c.SetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_GetIamPolicy() {
	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.GetIamPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/iam/v1#GetIamPolicyRequest.
	}
	resp, err := c.GetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_TestIamPermissions() {
	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.TestIamPermissionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/iam/v1#TestIamPermissionsRequest.
	}
	resp, err := c.TestIamPermissions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_SendCommandToDevice() {
	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iotpb.SendCommandToDeviceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/iot/v1#SendCommandToDeviceRequest.
	}
	resp, err := c.SendCommandToDevice(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_BindDeviceToGateway() {
	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iotpb.BindDeviceToGatewayRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/iot/v1#BindDeviceToGatewayRequest.
	}
	resp, err := c.BindDeviceToGateway(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDeviceManagerClient_UnbindDeviceFromGateway() {
	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iotpb.UnbindDeviceFromGatewayRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/iot/v1#UnbindDeviceFromGatewayRequest.
	}
	resp, err := c.UnbindDeviceFromGateway(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
