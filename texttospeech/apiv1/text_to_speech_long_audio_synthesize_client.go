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

package texttospeech

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
	texttospeechpb "cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
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

var newTextToSpeechLongAudioSynthesizeClientHook clientHook

// TextToSpeechLongAudioSynthesizeCallOptions contains the retry settings for each method of TextToSpeechLongAudioSynthesizeClient.
type TextToSpeechLongAudioSynthesizeCallOptions struct {
	SynthesizeLongAudio []gax.CallOption
	GetOperation        []gax.CallOption
	ListOperations      []gax.CallOption
}

func defaultTextToSpeechLongAudioSynthesizeGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("texttospeech.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("texttospeech.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("texttospeech.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://texttospeech.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultTextToSpeechLongAudioSynthesizeCallOptions() *TextToSpeechLongAudioSynthesizeCallOptions {
	return &TextToSpeechLongAudioSynthesizeCallOptions{
		SynthesizeLongAudio: []gax.CallOption{
			gax.WithTimeout(5000000 * time.Millisecond),
		},
		GetOperation:   []gax.CallOption{},
		ListOperations: []gax.CallOption{},
	}
}

func defaultTextToSpeechLongAudioSynthesizeRESTCallOptions() *TextToSpeechLongAudioSynthesizeCallOptions {
	return &TextToSpeechLongAudioSynthesizeCallOptions{
		SynthesizeLongAudio: []gax.CallOption{
			gax.WithTimeout(5000000 * time.Millisecond),
		},
		GetOperation:   []gax.CallOption{},
		ListOperations: []gax.CallOption{},
	}
}

// internalTextToSpeechLongAudioSynthesizeClient is an interface that defines the methods available from Cloud Text-to-Speech API.
type internalTextToSpeechLongAudioSynthesizeClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	SynthesizeLongAudio(context.Context, *texttospeechpb.SynthesizeLongAudioRequest, ...gax.CallOption) (*SynthesizeLongAudioOperation, error)
	SynthesizeLongAudioOperation(name string) *SynthesizeLongAudioOperation
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
}

// TextToSpeechLongAudioSynthesizeClient is a client for interacting with Cloud Text-to-Speech API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service that implements Google Cloud Text-to-Speech API.
type TextToSpeechLongAudioSynthesizeClient struct {
	// The internal transport-dependent client.
	internalClient internalTextToSpeechLongAudioSynthesizeClient

	// The call options for this service.
	CallOptions *TextToSpeechLongAudioSynthesizeCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TextToSpeechLongAudioSynthesizeClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TextToSpeechLongAudioSynthesizeClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *TextToSpeechLongAudioSynthesizeClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// SynthesizeLongAudio synthesizes long form text asynchronously.
func (c *TextToSpeechLongAudioSynthesizeClient) SynthesizeLongAudio(ctx context.Context, req *texttospeechpb.SynthesizeLongAudioRequest, opts ...gax.CallOption) (*SynthesizeLongAudioOperation, error) {
	return c.internalClient.SynthesizeLongAudio(ctx, req, opts...)
}

// SynthesizeLongAudioOperation returns a new SynthesizeLongAudioOperation from a given name.
// The name must be that of a previously created SynthesizeLongAudioOperation, possibly from a different process.
func (c *TextToSpeechLongAudioSynthesizeClient) SynthesizeLongAudioOperation(name string) *SynthesizeLongAudioOperation {
	return c.internalClient.SynthesizeLongAudioOperation(name)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *TextToSpeechLongAudioSynthesizeClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *TextToSpeechLongAudioSynthesizeClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.ListOperations(ctx, req, opts...)
}

// textToSpeechLongAudioSynthesizeGRPCClient is a client for interacting with Cloud Text-to-Speech API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type textToSpeechLongAudioSynthesizeGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing TextToSpeechLongAudioSynthesizeClient
	CallOptions **TextToSpeechLongAudioSynthesizeCallOptions

	// The gRPC API client.
	textToSpeechLongAudioSynthesizeClient texttospeechpb.TextToSpeechLongAudioSynthesizeClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	operationsClient longrunningpb.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewTextToSpeechLongAudioSynthesizeClient creates a new text to speech long audio synthesize client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service that implements Google Cloud Text-to-Speech API.
func NewTextToSpeechLongAudioSynthesizeClient(ctx context.Context, opts ...option.ClientOption) (*TextToSpeechLongAudioSynthesizeClient, error) {
	clientOpts := defaultTextToSpeechLongAudioSynthesizeGRPCClientOptions()
	if newTextToSpeechLongAudioSynthesizeClientHook != nil {
		hookOpts, err := newTextToSpeechLongAudioSynthesizeClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := TextToSpeechLongAudioSynthesizeClient{CallOptions: defaultTextToSpeechLongAudioSynthesizeCallOptions()}

	c := &textToSpeechLongAudioSynthesizeGRPCClient{
		connPool:                              connPool,
		textToSpeechLongAudioSynthesizeClient: texttospeechpb.NewTextToSpeechLongAudioSynthesizeClient(connPool),
		CallOptions:                           &client.CallOptions,
		operationsClient:                      longrunningpb.NewOperationsClient(connPool),
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
func (c *textToSpeechLongAudioSynthesizeGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *textToSpeechLongAudioSynthesizeGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *textToSpeechLongAudioSynthesizeGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type textToSpeechLongAudioSynthesizeRESTClient struct {
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

	// Points back to the CallOptions field of the containing TextToSpeechLongAudioSynthesizeClient
	CallOptions **TextToSpeechLongAudioSynthesizeCallOptions
}

// NewTextToSpeechLongAudioSynthesizeRESTClient creates a new text to speech long audio synthesize rest client.
//
// Service that implements Google Cloud Text-to-Speech API.
func NewTextToSpeechLongAudioSynthesizeRESTClient(ctx context.Context, opts ...option.ClientOption) (*TextToSpeechLongAudioSynthesizeClient, error) {
	clientOpts := append(defaultTextToSpeechLongAudioSynthesizeRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultTextToSpeechLongAudioSynthesizeRESTCallOptions()
	c := &textToSpeechLongAudioSynthesizeRESTClient{
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

	return &TextToSpeechLongAudioSynthesizeClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultTextToSpeechLongAudioSynthesizeRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://texttospeech.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://texttospeech.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://texttospeech.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://texttospeech.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *textToSpeechLongAudioSynthesizeRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *textToSpeechLongAudioSynthesizeRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *textToSpeechLongAudioSynthesizeRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *textToSpeechLongAudioSynthesizeGRPCClient) SynthesizeLongAudio(ctx context.Context, req *texttospeechpb.SynthesizeLongAudioRequest, opts ...gax.CallOption) (*SynthesizeLongAudioOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SynthesizeLongAudio[0:len((*c.CallOptions).SynthesizeLongAudio):len((*c.CallOptions).SynthesizeLongAudio)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.textToSpeechLongAudioSynthesizeClient.SynthesizeLongAudio(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &SynthesizeLongAudioOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *textToSpeechLongAudioSynthesizeGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.operationsClient.GetOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *textToSpeechLongAudioSynthesizeGRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListOperations[0:len((*c.CallOptions).ListOperations):len((*c.CallOptions).ListOperations)], opts...)
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
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
			resp, err = c.operationsClient.ListOperations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
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

// SynthesizeLongAudio synthesizes long form text asynchronously.
func (c *textToSpeechLongAudioSynthesizeRESTClient) SynthesizeLongAudio(ctx context.Context, req *texttospeechpb.SynthesizeLongAudioRequest, opts ...gax.CallOption) (*SynthesizeLongAudioOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v:synthesizeLongAudio", req.GetParent())

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

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &SynthesizeLongAudioOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *textToSpeechLongAudioSynthesizeRESTClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
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

// ListOperations is a utility method from google.longrunning.Operations.
func (c *textToSpeechLongAudioSynthesizeRESTClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v1/%v/operations", req.GetName())

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
		return resp.GetOperations(), resp.GetNextPageToken(), nil
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

// SynthesizeLongAudioOperation returns a new SynthesizeLongAudioOperation from a given name.
// The name must be that of a previously created SynthesizeLongAudioOperation, possibly from a different process.
func (c *textToSpeechLongAudioSynthesizeGRPCClient) SynthesizeLongAudioOperation(name string) *SynthesizeLongAudioOperation {
	return &SynthesizeLongAudioOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// SynthesizeLongAudioOperation returns a new SynthesizeLongAudioOperation from a given name.
// The name must be that of a previously created SynthesizeLongAudioOperation, possibly from a different process.
func (c *textToSpeechLongAudioSynthesizeRESTClient) SynthesizeLongAudioOperation(name string) *SynthesizeLongAudioOperation {
	override := fmt.Sprintf("/v1/%s", name)
	return &SynthesizeLongAudioOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}