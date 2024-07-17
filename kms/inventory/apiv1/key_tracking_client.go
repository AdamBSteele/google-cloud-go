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

package inventory

import (
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	inventorypb "cloud.google.com/go/kms/inventory/apiv1/inventorypb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newKeyTrackingClientHook clientHook

// KeyTrackingCallOptions contains the retry settings for each method of KeyTrackingClient.
type KeyTrackingCallOptions struct {
	GetProtectedResourcesSummary []gax.CallOption
	SearchProtectedResources     []gax.CallOption
}

func defaultKeyTrackingGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("kmsinventory.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("kmsinventory.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("kmsinventory.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://kmsinventory.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultKeyTrackingCallOptions() *KeyTrackingCallOptions {
	return &KeyTrackingCallOptions{
		GetProtectedResourcesSummary: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		SearchProtectedResources: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

func defaultKeyTrackingRESTCallOptions() *KeyTrackingCallOptions {
	return &KeyTrackingCallOptions{
		GetProtectedResourcesSummary: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		SearchProtectedResources: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

// internalKeyTrackingClient is an interface that defines the methods available from KMS Inventory API.
type internalKeyTrackingClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetProtectedResourcesSummary(context.Context, *inventorypb.GetProtectedResourcesSummaryRequest, ...gax.CallOption) (*inventorypb.ProtectedResourcesSummary, error)
	SearchProtectedResources(context.Context, *inventorypb.SearchProtectedResourcesRequest, ...gax.CallOption) *ProtectedResourceIterator
}

// KeyTrackingClient is a client for interacting with KMS Inventory API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Returns information about the resources in an org that are protected by a
// given Cloud KMS key via CMEK.
type KeyTrackingClient struct {
	// The internal transport-dependent client.
	internalClient internalKeyTrackingClient

	// The call options for this service.
	CallOptions *KeyTrackingCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *KeyTrackingClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *KeyTrackingClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *KeyTrackingClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetProtectedResourcesSummary returns aggregate information about the resources protected by the given
// Cloud KMS CryptoKey. Only resources within
// the same Cloud organization as the key will be returned. The project that
// holds the key must be part of an organization in order for this call to
// succeed.
func (c *KeyTrackingClient) GetProtectedResourcesSummary(ctx context.Context, req *inventorypb.GetProtectedResourcesSummaryRequest, opts ...gax.CallOption) (*inventorypb.ProtectedResourcesSummary, error) {
	return c.internalClient.GetProtectedResourcesSummary(ctx, req, opts...)
}

// SearchProtectedResources returns metadata about the resources protected by the given Cloud KMS
// CryptoKey in the given Cloud organization.
func (c *KeyTrackingClient) SearchProtectedResources(ctx context.Context, req *inventorypb.SearchProtectedResourcesRequest, opts ...gax.CallOption) *ProtectedResourceIterator {
	return c.internalClient.SearchProtectedResources(ctx, req, opts...)
}

// keyTrackingGRPCClient is a client for interacting with KMS Inventory API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type keyTrackingGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing KeyTrackingClient
	CallOptions **KeyTrackingCallOptions

	// The gRPC API client.
	keyTrackingClient inventorypb.KeyTrackingServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewKeyTrackingClient creates a new key tracking service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Returns information about the resources in an org that are protected by a
// given Cloud KMS key via CMEK.
func NewKeyTrackingClient(ctx context.Context, opts ...option.ClientOption) (*KeyTrackingClient, error) {
	clientOpts := defaultKeyTrackingGRPCClientOptions()
	if newKeyTrackingClientHook != nil {
		hookOpts, err := newKeyTrackingClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := KeyTrackingClient{CallOptions: defaultKeyTrackingCallOptions()}

	c := &keyTrackingGRPCClient{
		connPool:          connPool,
		keyTrackingClient: inventorypb.NewKeyTrackingServiceClient(connPool),
		CallOptions:       &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *keyTrackingGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *keyTrackingGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *keyTrackingGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type keyTrackingRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing KeyTrackingClient
	CallOptions **KeyTrackingCallOptions
}

// NewKeyTrackingRESTClient creates a new key tracking service rest client.
//
// Returns information about the resources in an org that are protected by a
// given Cloud KMS key via CMEK.
func NewKeyTrackingRESTClient(ctx context.Context, opts ...option.ClientOption) (*KeyTrackingClient, error) {
	clientOpts := append(defaultKeyTrackingRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultKeyTrackingRESTCallOptions()
	c := &keyTrackingRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &KeyTrackingClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultKeyTrackingRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://kmsinventory.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://kmsinventory.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://kmsinventory.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://kmsinventory.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *keyTrackingRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *keyTrackingRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *keyTrackingRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *keyTrackingGRPCClient) GetProtectedResourcesSummary(ctx context.Context, req *inventorypb.GetProtectedResourcesSummaryRequest, opts ...gax.CallOption) (*inventorypb.ProtectedResourcesSummary, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetProtectedResourcesSummary[0:len((*c.CallOptions).GetProtectedResourcesSummary):len((*c.CallOptions).GetProtectedResourcesSummary)], opts...)
	var resp *inventorypb.ProtectedResourcesSummary
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.keyTrackingClient.GetProtectedResourcesSummary(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *keyTrackingGRPCClient) SearchProtectedResources(ctx context.Context, req *inventorypb.SearchProtectedResourcesRequest, opts ...gax.CallOption) *ProtectedResourceIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "scope", url.QueryEscape(req.GetScope()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SearchProtectedResources[0:len((*c.CallOptions).SearchProtectedResources):len((*c.CallOptions).SearchProtectedResources)], opts...)
	it := &ProtectedResourceIterator{}
	req = proto.Clone(req).(*inventorypb.SearchProtectedResourcesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*inventorypb.ProtectedResource, string, error) {
		resp := &inventorypb.SearchProtectedResourcesResponse{}
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
			resp, err = c.keyTrackingClient.SearchProtectedResources(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetProtectedResources(), resp.GetNextPageToken(), nil
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

// GetProtectedResourcesSummary returns aggregate information about the resources protected by the given
// Cloud KMS CryptoKey. Only resources within
// the same Cloud organization as the key will be returned. The project that
// holds the key must be part of an organization in order for this call to
// succeed.
func (c *keyTrackingRESTClient) GetProtectedResourcesSummary(ctx context.Context, req *inventorypb.GetProtectedResourcesSummaryRequest, opts ...gax.CallOption) (*inventorypb.ProtectedResourcesSummary, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v/protectedResourcesSummary", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetProtectedResourcesSummary[0:len((*c.CallOptions).GetProtectedResourcesSummary):len((*c.CallOptions).GetProtectedResourcesSummary)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &inventorypb.ProtectedResourcesSummary{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// SearchProtectedResources returns metadata about the resources protected by the given Cloud KMS
// CryptoKey in the given Cloud organization.
func (c *keyTrackingRESTClient) SearchProtectedResources(ctx context.Context, req *inventorypb.SearchProtectedResourcesRequest, opts ...gax.CallOption) *ProtectedResourceIterator {
	it := &ProtectedResourceIterator{}
	req = proto.Clone(req).(*inventorypb.SearchProtectedResourcesRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*inventorypb.ProtectedResource, string, error) {
		resp := &inventorypb.SearchProtectedResourcesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v1/%v/protectedResources:search", req.GetScope())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		params.Add("cryptoKey", fmt.Sprintf("%v", req.GetCryptoKey()))
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if items := req.GetResourceTypes(); len(items) > 0 {
			for _, item := range items {
				params.Add("resourceTypes", fmt.Sprintf("%v", item))
			}
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			httpRsp, err := c.httpClient.Do(httpReq)
			if err != nil {
				return err
			}
			defer httpRsp.Body.Close()

			if err = googleapi.CheckResponse(httpRsp); err != nil {
				return err
			}

			buf, err := io.ReadAll(httpRsp.Body)
			if err != nil {
				return err
			}

			if err := unm.Unmarshal(buf, resp); err != nil {
				return err
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetProtectedResources(), resp.GetNextPageToken(), nil
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
