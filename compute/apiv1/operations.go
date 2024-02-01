// Copyright 2024 Google LLC
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

package compute

import (
	"context"
	"fmt"
	"time"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	gax "github.com/googleapis/gax-go/v2"
	"github.com/googleapis/gax-go/v2/apierror"
	"google.golang.org/api/googleapi"
)

// Operation represents a long-running operation for this API.
type Operation struct {
	operationHandle
}

// Done reports whether the long-running operation has completed.
func (o *Operation) Done() bool {
	return o.Proto().GetStatus() == computepb.Operation_DONE
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (o *Operation) Name() string {
	return o.Proto().GetName()
}

// Wait blocks until the operation is complete, polling regularly
// after an intial period of backing off between attempts.
func (o *Operation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	bo := gax.Backoff{
		Initial: time.Second,
		Max:     time.Minute,
	}
	for {
		if err := o.Poll(ctx, opts...); err != nil {
			return err
		}
		if o.Done() {
			return nil
		}
		if err := gax.Sleep(ctx, bo.Pause()); err != nil {
			return err
		}
	}
}

type operationHandle interface {
	// Poll retrieves the operation.
	Poll(ctx context.Context, opts ...gax.CallOption) error

	// Proto returns the long-running operation message.
	Proto() *computepb.Operation
}

// Implements the operationHandle interface for RegionOperations.
type regionOperationsHandle struct {
	c       *RegionOperationsClient
	proto   *computepb.Operation
	project string
	region  string
}

// Poll retrieves the latest data for the long-running operation.
func (h *regionOperationsHandle) Poll(ctx context.Context, opts ...gax.CallOption) error {
	resp, err := h.c.Get(ctx, &computepb.GetRegionOperationRequest{
		Operation: h.proto.GetName(),
		Project:   h.project,
		Region:    h.region,
	}, opts...)
	if err != nil {
		return err
	}
	h.proto = resp
	if resp.HttpErrorStatusCode != nil && (resp.GetHttpErrorStatusCode() < 200 || resp.GetHttpErrorStatusCode() > 299) {
		aErr := &googleapi.Error{
			Code:    int(resp.GetHttpErrorStatusCode()),
			Message: fmt.Sprintf("%s: %v", resp.GetHttpErrorMessage(), resp.GetError()),
		}
		err, _ := apierror.FromError(aErr)
		return err
	}
	return nil
}

// Proto returns the raw type this wraps.
func (h *regionOperationsHandle) Proto() *computepb.Operation {
	return h.proto
}

// Implements the operationHandle interface for ZoneOperations.
type zoneOperationsHandle struct {
	c       *ZoneOperationsClient
	proto   *computepb.Operation
	project string
	zone    string
}

// Poll retrieves the latest data for the long-running operation.
func (h *zoneOperationsHandle) Poll(ctx context.Context, opts ...gax.CallOption) error {
	resp, err := h.c.Get(ctx, &computepb.GetZoneOperationRequest{
		Operation: h.proto.GetName(),
		Project:   h.project,
		Zone:      h.zone,
	}, opts...)
	if err != nil {
		return err
	}
	h.proto = resp
	if resp.HttpErrorStatusCode != nil && (resp.GetHttpErrorStatusCode() < 200 || resp.GetHttpErrorStatusCode() > 299) {
		aErr := &googleapi.Error{
			Code:    int(resp.GetHttpErrorStatusCode()),
			Message: fmt.Sprintf("%s: %v", resp.GetHttpErrorMessage(), resp.GetError()),
		}
		err, _ := apierror.FromError(aErr)
		return err
	}
	return nil
}

// Proto returns the raw type this wraps.
func (h *zoneOperationsHandle) Proto() *computepb.Operation {
	return h.proto
}

// Implements the operationHandle interface for GlobalOperations.
type globalOperationsHandle struct {
	c       *GlobalOperationsClient
	proto   *computepb.Operation
	project string
}

// Poll retrieves the latest data for the long-running operation.
func (h *globalOperationsHandle) Poll(ctx context.Context, opts ...gax.CallOption) error {
	resp, err := h.c.Get(ctx, &computepb.GetGlobalOperationRequest{
		Operation: h.proto.GetName(),
		Project:   h.project,
	}, opts...)
	if err != nil {
		return err
	}
	h.proto = resp
	if resp.HttpErrorStatusCode != nil && (resp.GetHttpErrorStatusCode() < 200 || resp.GetHttpErrorStatusCode() > 299) {
		aErr := &googleapi.Error{
			Code:    int(resp.GetHttpErrorStatusCode()),
			Message: fmt.Sprintf("%s: %v", resp.GetHttpErrorMessage(), resp.GetError()),
		}
		err, _ := apierror.FromError(aErr)
		return err
	}
	return nil
}

// Proto returns the raw type this wraps.
func (h *globalOperationsHandle) Proto() *computepb.Operation {
	return h.proto
}

// Implements the operationHandle interface for GlobalOrganizationOperations.
type globalOrganizationOperationsHandle struct {
	c     *GlobalOrganizationOperationsClient
	proto *computepb.Operation
}

// Poll retrieves the latest data for the long-running operation.
func (h *globalOrganizationOperationsHandle) Poll(ctx context.Context, opts ...gax.CallOption) error {
	resp, err := h.c.Get(ctx, &computepb.GetGlobalOrganizationOperationRequest{
		Operation: h.proto.GetName(),
	}, opts...)
	if err != nil {
		return err
	}
	h.proto = resp
	if resp.HttpErrorStatusCode != nil && (resp.GetHttpErrorStatusCode() < 200 || resp.GetHttpErrorStatusCode() > 299) {
		aErr := &googleapi.Error{
			Code:    int(resp.GetHttpErrorStatusCode()),
			Message: fmt.Sprintf("%s: %v", resp.GetHttpErrorMessage(), resp.GetError()),
		}
		err, _ := apierror.FromError(aErr)
		return err
	}
	return nil
}

// Proto returns the raw type this wraps.
func (h *globalOrganizationOperationsHandle) Proto() *computepb.Operation {
	return h.proto
}
