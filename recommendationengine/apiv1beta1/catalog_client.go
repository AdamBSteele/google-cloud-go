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

package recommendationengine

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	recommendationenginepb "cloud.google.com/go/recommendationengine/apiv1beta1/recommendationenginepb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newCatalogClientHook clientHook

// CatalogCallOptions contains the retry settings for each method of CatalogClient.
type CatalogCallOptions struct {
	CreateCatalogItem  []gax.CallOption
	GetCatalogItem     []gax.CallOption
	ListCatalogItems   []gax.CallOption
	UpdateCatalogItem  []gax.CallOption
	DeleteCatalogItem  []gax.CallOption
	ImportCatalogItems []gax.CallOption
}

func defaultCatalogGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("recommendationengine.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("recommendationengine.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("recommendationengine.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://recommendationengine.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCatalogCallOptions() *CatalogCallOptions {
	return &CatalogCallOptions{
		CreateCatalogItem: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetCatalogItem: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListCatalogItems: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateCatalogItem: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteCatalogItem: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ImportCatalogItems: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

func defaultCatalogRESTCallOptions() *CatalogCallOptions {
	return &CatalogCallOptions{
		CreateCatalogItem: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable,
					http.StatusGatewayTimeout)
			}),
		},
		GetCatalogItem: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable,
					http.StatusGatewayTimeout)
			}),
		},
		ListCatalogItems: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable,
					http.StatusGatewayTimeout)
			}),
		},
		UpdateCatalogItem: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable,
					http.StatusGatewayTimeout)
			}),
		},
		DeleteCatalogItem: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable,
					http.StatusGatewayTimeout)
			}),
		},
		ImportCatalogItems: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable,
					http.StatusGatewayTimeout)
			}),
		},
	}
}

// internalCatalogClient is an interface that defines the methods available from Recommendations AI.
type internalCatalogClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateCatalogItem(context.Context, *recommendationenginepb.CreateCatalogItemRequest, ...gax.CallOption) (*recommendationenginepb.CatalogItem, error)
	GetCatalogItem(context.Context, *recommendationenginepb.GetCatalogItemRequest, ...gax.CallOption) (*recommendationenginepb.CatalogItem, error)
	ListCatalogItems(context.Context, *recommendationenginepb.ListCatalogItemsRequest, ...gax.CallOption) *CatalogItemIterator
	UpdateCatalogItem(context.Context, *recommendationenginepb.UpdateCatalogItemRequest, ...gax.CallOption) (*recommendationenginepb.CatalogItem, error)
	DeleteCatalogItem(context.Context, *recommendationenginepb.DeleteCatalogItemRequest, ...gax.CallOption) error
	ImportCatalogItems(context.Context, *recommendationenginepb.ImportCatalogItemsRequest, ...gax.CallOption) (*ImportCatalogItemsOperation, error)
	ImportCatalogItemsOperation(name string) *ImportCatalogItemsOperation
}

// CatalogClient is a client for interacting with Recommendations AI.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for ingesting catalog information of the customer’s website.
type CatalogClient struct {
	// The internal transport-dependent client.
	internalClient internalCatalogClient

	// The call options for this service.
	CallOptions *CatalogCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CatalogClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CatalogClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *CatalogClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateCatalogItem creates a catalog item.
func (c *CatalogClient) CreateCatalogItem(ctx context.Context, req *recommendationenginepb.CreateCatalogItemRequest, opts ...gax.CallOption) (*recommendationenginepb.CatalogItem, error) {
	return c.internalClient.CreateCatalogItem(ctx, req, opts...)
}

// GetCatalogItem gets a specific catalog item.
func (c *CatalogClient) GetCatalogItem(ctx context.Context, req *recommendationenginepb.GetCatalogItemRequest, opts ...gax.CallOption) (*recommendationenginepb.CatalogItem, error) {
	return c.internalClient.GetCatalogItem(ctx, req, opts...)
}

// ListCatalogItems gets a list of catalog items.
func (c *CatalogClient) ListCatalogItems(ctx context.Context, req *recommendationenginepb.ListCatalogItemsRequest, opts ...gax.CallOption) *CatalogItemIterator {
	return c.internalClient.ListCatalogItems(ctx, req, opts...)
}

// UpdateCatalogItem updates a catalog item. Partial updating is supported. Non-existing
// items will be created.
func (c *CatalogClient) UpdateCatalogItem(ctx context.Context, req *recommendationenginepb.UpdateCatalogItemRequest, opts ...gax.CallOption) (*recommendationenginepb.CatalogItem, error) {
	return c.internalClient.UpdateCatalogItem(ctx, req, opts...)
}

// DeleteCatalogItem deletes a catalog item.
func (c *CatalogClient) DeleteCatalogItem(ctx context.Context, req *recommendationenginepb.DeleteCatalogItemRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteCatalogItem(ctx, req, opts...)
}

// ImportCatalogItems bulk import of multiple catalog items. Request processing may be
// synchronous. No partial updating supported. Non-existing items will be
// created.
//
// Operation.response is of type ImportResponse. Note that it is
// possible for a subset of the items to be successfully updated.
func (c *CatalogClient) ImportCatalogItems(ctx context.Context, req *recommendationenginepb.ImportCatalogItemsRequest, opts ...gax.CallOption) (*ImportCatalogItemsOperation, error) {
	return c.internalClient.ImportCatalogItems(ctx, req, opts...)
}

// ImportCatalogItemsOperation returns a new ImportCatalogItemsOperation from a given name.
// The name must be that of a previously created ImportCatalogItemsOperation, possibly from a different process.
func (c *CatalogClient) ImportCatalogItemsOperation(name string) *ImportCatalogItemsOperation {
	return c.internalClient.ImportCatalogItemsOperation(name)
}

// catalogGRPCClient is a client for interacting with Recommendations AI over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type catalogGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing CatalogClient
	CallOptions **CatalogCallOptions

	// The gRPC API client.
	catalogClient recommendationenginepb.CatalogServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewCatalogClient creates a new catalog service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for ingesting catalog information of the customer’s website.
func NewCatalogClient(ctx context.Context, opts ...option.ClientOption) (*CatalogClient, error) {
	clientOpts := defaultCatalogGRPCClientOptions()
	if newCatalogClientHook != nil {
		hookOpts, err := newCatalogClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := CatalogClient{CallOptions: defaultCatalogCallOptions()}

	c := &catalogGRPCClient{
		connPool:      connPool,
		catalogClient: recommendationenginepb.NewCatalogServiceClient(connPool),
		CallOptions:   &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *catalogGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *catalogGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *catalogGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type catalogRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing CatalogClient
	CallOptions **CatalogCallOptions
}

// NewCatalogRESTClient creates a new catalog service rest client.
//
// Service for ingesting catalog information of the customer’s website.
func NewCatalogRESTClient(ctx context.Context, opts ...option.ClientOption) (*CatalogClient, error) {
	clientOpts := append(defaultCatalogRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultCatalogRESTCallOptions()
	c := &catalogRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	lroOpts := []option.ClientOption{
		option.WithHTTPClient(httpClient),
		option.WithEndpoint(endpoint),
	}
	opClient, err := lroauto.NewOperationsRESTClient(ctx, lroOpts...)
	if err != nil {
		return nil, err
	}
	c.LROClient = &opClient

	return &CatalogClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultCatalogRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://recommendationengine.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://recommendationengine.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://recommendationengine.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://recommendationengine.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *catalogRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *catalogRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *catalogRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *catalogGRPCClient) CreateCatalogItem(ctx context.Context, req *recommendationenginepb.CreateCatalogItemRequest, opts ...gax.CallOption) (*recommendationenginepb.CatalogItem, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateCatalogItem[0:len((*c.CallOptions).CreateCatalogItem):len((*c.CallOptions).CreateCatalogItem)], opts...)
	var resp *recommendationenginepb.CatalogItem
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.catalogClient.CreateCatalogItem(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *catalogGRPCClient) GetCatalogItem(ctx context.Context, req *recommendationenginepb.GetCatalogItemRequest, opts ...gax.CallOption) (*recommendationenginepb.CatalogItem, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetCatalogItem[0:len((*c.CallOptions).GetCatalogItem):len((*c.CallOptions).GetCatalogItem)], opts...)
	var resp *recommendationenginepb.CatalogItem
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.catalogClient.GetCatalogItem(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *catalogGRPCClient) ListCatalogItems(ctx context.Context, req *recommendationenginepb.ListCatalogItemsRequest, opts ...gax.CallOption) *CatalogItemIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListCatalogItems[0:len((*c.CallOptions).ListCatalogItems):len((*c.CallOptions).ListCatalogItems)], opts...)
	it := &CatalogItemIterator{}
	req = proto.Clone(req).(*recommendationenginepb.ListCatalogItemsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*recommendationenginepb.CatalogItem, string, error) {
		resp := &recommendationenginepb.ListCatalogItemsResponse{}
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
			resp, err = c.catalogClient.ListCatalogItems(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetCatalogItems(), resp.GetNextPageToken(), nil
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

func (c *catalogGRPCClient) UpdateCatalogItem(ctx context.Context, req *recommendationenginepb.UpdateCatalogItemRequest, opts ...gax.CallOption) (*recommendationenginepb.CatalogItem, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateCatalogItem[0:len((*c.CallOptions).UpdateCatalogItem):len((*c.CallOptions).UpdateCatalogItem)], opts...)
	var resp *recommendationenginepb.CatalogItem
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.catalogClient.UpdateCatalogItem(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *catalogGRPCClient) DeleteCatalogItem(ctx context.Context, req *recommendationenginepb.DeleteCatalogItemRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteCatalogItem[0:len((*c.CallOptions).DeleteCatalogItem):len((*c.CallOptions).DeleteCatalogItem)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.catalogClient.DeleteCatalogItem(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *catalogGRPCClient) ImportCatalogItems(ctx context.Context, req *recommendationenginepb.ImportCatalogItemsRequest, opts ...gax.CallOption) (*ImportCatalogItemsOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ImportCatalogItems[0:len((*c.CallOptions).ImportCatalogItems):len((*c.CallOptions).ImportCatalogItems)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.catalogClient.ImportCatalogItems(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ImportCatalogItemsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// CreateCatalogItem creates a catalog item.
func (c *catalogRESTClient) CreateCatalogItem(ctx context.Context, req *recommendationenginepb.CreateCatalogItemRequest, opts ...gax.CallOption) (*recommendationenginepb.CatalogItem, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetCatalogItem()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v/catalogItems", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).CreateCatalogItem[0:len((*c.CallOptions).CreateCatalogItem):len((*c.CallOptions).CreateCatalogItem)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &recommendationenginepb.CatalogItem{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
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

// GetCatalogItem gets a specific catalog item.
func (c *catalogRESTClient) GetCatalogItem(ctx context.Context, req *recommendationenginepb.GetCatalogItemRequest, opts ...gax.CallOption) (*recommendationenginepb.CatalogItem, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetCatalogItem[0:len((*c.CallOptions).GetCatalogItem):len((*c.CallOptions).GetCatalogItem)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &recommendationenginepb.CatalogItem{}
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

// ListCatalogItems gets a list of catalog items.
func (c *catalogRESTClient) ListCatalogItems(ctx context.Context, req *recommendationenginepb.ListCatalogItemsRequest, opts ...gax.CallOption) *CatalogItemIterator {
	it := &CatalogItemIterator{}
	req = proto.Clone(req).(*recommendationenginepb.ListCatalogItemsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*recommendationenginepb.CatalogItem, string, error) {
		resp := &recommendationenginepb.ListCatalogItemsResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v1beta1/%v/catalogItems", req.GetParent())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetFilter() != "" {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
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
		return resp.GetCatalogItems(), resp.GetNextPageToken(), nil
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

// UpdateCatalogItem updates a catalog item. Partial updating is supported. Non-existing
// items will be created.
func (c *catalogRESTClient) UpdateCatalogItem(ctx context.Context, req *recommendationenginepb.UpdateCatalogItemRequest, opts ...gax.CallOption) (*recommendationenginepb.CatalogItem, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetCatalogItem()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetUpdateMask() != nil {
		updateMask, err := protojson.Marshal(req.GetUpdateMask())
		if err != nil {
			return nil, err
		}
		params.Add("updateMask", string(updateMask[1:len(updateMask)-1]))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).UpdateCatalogItem[0:len((*c.CallOptions).UpdateCatalogItem):len((*c.CallOptions).UpdateCatalogItem)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &recommendationenginepb.CatalogItem{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
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

// DeleteCatalogItem deletes a catalog item.
func (c *catalogRESTClient) DeleteCatalogItem(ctx context.Context, req *recommendationenginepb.DeleteCatalogItemRequest, opts ...gax.CallOption) error {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	return gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
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

		// Returns nil if there is no error, otherwise wraps
		// the response code and body into a non-nil error
		return googleapi.CheckResponse(httpRsp)
	}, opts...)
}

// ImportCatalogItems bulk import of multiple catalog items. Request processing may be
// synchronous. No partial updating supported. Non-existing items will be
// created.
//
// Operation.response is of type ImportResponse. Note that it is
// possible for a subset of the items to be successfully updated.
func (c *catalogRESTClient) ImportCatalogItems(ctx context.Context, req *recommendationenginepb.ImportCatalogItemsRequest, opts ...gax.CallOption) (*ImportCatalogItemsOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v/catalogItems:import", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
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

	override := fmt.Sprintf("/v1beta1/%s", resp.GetName())
	return &ImportCatalogItemsOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// ImportCatalogItemsOperation returns a new ImportCatalogItemsOperation from a given name.
// The name must be that of a previously created ImportCatalogItemsOperation, possibly from a different process.
func (c *catalogGRPCClient) ImportCatalogItemsOperation(name string) *ImportCatalogItemsOperation {
	return &ImportCatalogItemsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// ImportCatalogItemsOperation returns a new ImportCatalogItemsOperation from a given name.
// The name must be that of a previously created ImportCatalogItemsOperation, possibly from a different process.
func (c *catalogRESTClient) ImportCatalogItemsOperation(name string) *ImportCatalogItemsOperation {
	override := fmt.Sprintf("/v1beta1/%s", name)
	return &ImportCatalogItemsOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}