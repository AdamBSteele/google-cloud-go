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

package apigeeconnect

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	apigeeconnectpb "cloud.google.com/go/apigeeconnect/apiv1/apigeeconnectpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

var newConnectionClientHook clientHook

// ConnectionCallOptions contains the retry settings for each method of ConnectionClient.
type ConnectionCallOptions struct {
	ListConnections []gax.CallOption
}

func defaultConnectionGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("apigeeconnect.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("apigeeconnect.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("apigeeconnect.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://apigeeconnect.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultConnectionCallOptions() *ConnectionCallOptions {
	return &ConnectionCallOptions{
		ListConnections: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.Unknown,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalConnectionClient is an interface that defines the methods available from Apigee Connect API.
type internalConnectionClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListConnections(context.Context, *apigeeconnectpb.ListConnectionsRequest, ...gax.CallOption) *ConnectionIterator
}

// ConnectionClient is a client for interacting with Apigee Connect API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service Interface for the Apigee Connect connection management APIs.
type ConnectionClient struct {
	// The internal transport-dependent client.
	internalClient internalConnectionClient

	// The call options for this service.
	CallOptions *ConnectionCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ConnectionClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ConnectionClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ConnectionClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListConnections lists connections that are currently active for the given Apigee Connect
// endpoint.
func (c *ConnectionClient) ListConnections(ctx context.Context, req *apigeeconnectpb.ListConnectionsRequest, opts ...gax.CallOption) *ConnectionIterator {
	return c.internalClient.ListConnections(ctx, req, opts...)
}

// connectionGRPCClient is a client for interacting with Apigee Connect API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type connectionGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing ConnectionClient
	CallOptions **ConnectionCallOptions

	// The gRPC API client.
	connectionClient apigeeconnectpb.ConnectionServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewConnectionClient creates a new connection service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service Interface for the Apigee Connect connection management APIs.
func NewConnectionClient(ctx context.Context, opts ...option.ClientOption) (*ConnectionClient, error) {
	clientOpts := defaultConnectionGRPCClientOptions()
	if newConnectionClientHook != nil {
		hookOpts, err := newConnectionClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ConnectionClient{CallOptions: defaultConnectionCallOptions()}

	c := &connectionGRPCClient{
		connPool:         connPool,
		connectionClient: apigeeconnectpb.NewConnectionServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *connectionGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *connectionGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *connectionGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *connectionGRPCClient) ListConnections(ctx context.Context, req *apigeeconnectpb.ListConnectionsRequest, opts ...gax.CallOption) *ConnectionIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListConnections[0:len((*c.CallOptions).ListConnections):len((*c.CallOptions).ListConnections)], opts...)
	it := &ConnectionIterator{}
	req = proto.Clone(req).(*apigeeconnectpb.ListConnectionsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*apigeeconnectpb.Connection, string, error) {
		resp := &apigeeconnectpb.ListConnectionsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.connectionClient.ListConnections(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetConnections(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}