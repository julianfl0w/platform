// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: policy/kasregistry/key_access_server_registry.proto

package kasregistryconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	kasregistry "github.com/opentdf/platform/protocol/go/policy/kasregistry"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// KeyAccessServerRegistryServiceName is the fully-qualified name of the
	// KeyAccessServerRegistryService service.
	KeyAccessServerRegistryServiceName = "policy.kasregistry.KeyAccessServerRegistryService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// KeyAccessServerRegistryServiceListKeyAccessServersProcedure is the fully-qualified name of the
	// KeyAccessServerRegistryService's ListKeyAccessServers RPC.
	KeyAccessServerRegistryServiceListKeyAccessServersProcedure = "/policy.kasregistry.KeyAccessServerRegistryService/ListKeyAccessServers"
	// KeyAccessServerRegistryServiceGetKeyAccessServerProcedure is the fully-qualified name of the
	// KeyAccessServerRegistryService's GetKeyAccessServer RPC.
	KeyAccessServerRegistryServiceGetKeyAccessServerProcedure = "/policy.kasregistry.KeyAccessServerRegistryService/GetKeyAccessServer"
	// KeyAccessServerRegistryServiceCreateKeyAccessServerProcedure is the fully-qualified name of the
	// KeyAccessServerRegistryService's CreateKeyAccessServer RPC.
	KeyAccessServerRegistryServiceCreateKeyAccessServerProcedure = "/policy.kasregistry.KeyAccessServerRegistryService/CreateKeyAccessServer"
	// KeyAccessServerRegistryServiceUpdateKeyAccessServerProcedure is the fully-qualified name of the
	// KeyAccessServerRegistryService's UpdateKeyAccessServer RPC.
	KeyAccessServerRegistryServiceUpdateKeyAccessServerProcedure = "/policy.kasregistry.KeyAccessServerRegistryService/UpdateKeyAccessServer"
	// KeyAccessServerRegistryServiceDeleteKeyAccessServerProcedure is the fully-qualified name of the
	// KeyAccessServerRegistryService's DeleteKeyAccessServer RPC.
	KeyAccessServerRegistryServiceDeleteKeyAccessServerProcedure = "/policy.kasregistry.KeyAccessServerRegistryService/DeleteKeyAccessServer"
	// KeyAccessServerRegistryServiceListKeyAccessServerGrantsProcedure is the fully-qualified name of
	// the KeyAccessServerRegistryService's ListKeyAccessServerGrants RPC.
	KeyAccessServerRegistryServiceListKeyAccessServerGrantsProcedure = "/policy.kasregistry.KeyAccessServerRegistryService/ListKeyAccessServerGrants"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	keyAccessServerRegistryServiceServiceDescriptor                         = kasregistry.File_policy_kasregistry_key_access_server_registry_proto.Services().ByName("KeyAccessServerRegistryService")
	keyAccessServerRegistryServiceListKeyAccessServersMethodDescriptor      = keyAccessServerRegistryServiceServiceDescriptor.Methods().ByName("ListKeyAccessServers")
	keyAccessServerRegistryServiceGetKeyAccessServerMethodDescriptor        = keyAccessServerRegistryServiceServiceDescriptor.Methods().ByName("GetKeyAccessServer")
	keyAccessServerRegistryServiceCreateKeyAccessServerMethodDescriptor     = keyAccessServerRegistryServiceServiceDescriptor.Methods().ByName("CreateKeyAccessServer")
	keyAccessServerRegistryServiceUpdateKeyAccessServerMethodDescriptor     = keyAccessServerRegistryServiceServiceDescriptor.Methods().ByName("UpdateKeyAccessServer")
	keyAccessServerRegistryServiceDeleteKeyAccessServerMethodDescriptor     = keyAccessServerRegistryServiceServiceDescriptor.Methods().ByName("DeleteKeyAccessServer")
	keyAccessServerRegistryServiceListKeyAccessServerGrantsMethodDescriptor = keyAccessServerRegistryServiceServiceDescriptor.Methods().ByName("ListKeyAccessServerGrants")
)

// KeyAccessServerRegistryServiceClient is a client for the
// policy.kasregistry.KeyAccessServerRegistryService service.
type KeyAccessServerRegistryServiceClient interface {
	ListKeyAccessServers(context.Context, *connect.Request[kasregistry.ListKeyAccessServersRequest]) (*connect.Response[kasregistry.ListKeyAccessServersResponse], error)
	GetKeyAccessServer(context.Context, *connect.Request[kasregistry.GetKeyAccessServerRequest]) (*connect.Response[kasregistry.GetKeyAccessServerResponse], error)
	CreateKeyAccessServer(context.Context, *connect.Request[kasregistry.CreateKeyAccessServerRequest]) (*connect.Response[kasregistry.CreateKeyAccessServerResponse], error)
	UpdateKeyAccessServer(context.Context, *connect.Request[kasregistry.UpdateKeyAccessServerRequest]) (*connect.Response[kasregistry.UpdateKeyAccessServerResponse], error)
	DeleteKeyAccessServer(context.Context, *connect.Request[kasregistry.DeleteKeyAccessServerRequest]) (*connect.Response[kasregistry.DeleteKeyAccessServerResponse], error)
	ListKeyAccessServerGrants(context.Context, *connect.Request[kasregistry.ListKeyAccessServerGrantsRequest]) (*connect.Response[kasregistry.ListKeyAccessServerGrantsResponse], error)
}

// NewKeyAccessServerRegistryServiceClient constructs a client for the
// policy.kasregistry.KeyAccessServerRegistryService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewKeyAccessServerRegistryServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) KeyAccessServerRegistryServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &keyAccessServerRegistryServiceClient{
		listKeyAccessServers: connect.NewClient[kasregistry.ListKeyAccessServersRequest, kasregistry.ListKeyAccessServersResponse](
			httpClient,
			baseURL+KeyAccessServerRegistryServiceListKeyAccessServersProcedure,
			connect.WithSchema(keyAccessServerRegistryServiceListKeyAccessServersMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getKeyAccessServer: connect.NewClient[kasregistry.GetKeyAccessServerRequest, kasregistry.GetKeyAccessServerResponse](
			httpClient,
			baseURL+KeyAccessServerRegistryServiceGetKeyAccessServerProcedure,
			connect.WithSchema(keyAccessServerRegistryServiceGetKeyAccessServerMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		createKeyAccessServer: connect.NewClient[kasregistry.CreateKeyAccessServerRequest, kasregistry.CreateKeyAccessServerResponse](
			httpClient,
			baseURL+KeyAccessServerRegistryServiceCreateKeyAccessServerProcedure,
			connect.WithSchema(keyAccessServerRegistryServiceCreateKeyAccessServerMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateKeyAccessServer: connect.NewClient[kasregistry.UpdateKeyAccessServerRequest, kasregistry.UpdateKeyAccessServerResponse](
			httpClient,
			baseURL+KeyAccessServerRegistryServiceUpdateKeyAccessServerProcedure,
			connect.WithSchema(keyAccessServerRegistryServiceUpdateKeyAccessServerMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteKeyAccessServer: connect.NewClient[kasregistry.DeleteKeyAccessServerRequest, kasregistry.DeleteKeyAccessServerResponse](
			httpClient,
			baseURL+KeyAccessServerRegistryServiceDeleteKeyAccessServerProcedure,
			connect.WithSchema(keyAccessServerRegistryServiceDeleteKeyAccessServerMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listKeyAccessServerGrants: connect.NewClient[kasregistry.ListKeyAccessServerGrantsRequest, kasregistry.ListKeyAccessServerGrantsResponse](
			httpClient,
			baseURL+KeyAccessServerRegistryServiceListKeyAccessServerGrantsProcedure,
			connect.WithSchema(keyAccessServerRegistryServiceListKeyAccessServerGrantsMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
	}
}

// keyAccessServerRegistryServiceClient implements KeyAccessServerRegistryServiceClient.
type keyAccessServerRegistryServiceClient struct {
	listKeyAccessServers      *connect.Client[kasregistry.ListKeyAccessServersRequest, kasregistry.ListKeyAccessServersResponse]
	getKeyAccessServer        *connect.Client[kasregistry.GetKeyAccessServerRequest, kasregistry.GetKeyAccessServerResponse]
	createKeyAccessServer     *connect.Client[kasregistry.CreateKeyAccessServerRequest, kasregistry.CreateKeyAccessServerResponse]
	updateKeyAccessServer     *connect.Client[kasregistry.UpdateKeyAccessServerRequest, kasregistry.UpdateKeyAccessServerResponse]
	deleteKeyAccessServer     *connect.Client[kasregistry.DeleteKeyAccessServerRequest, kasregistry.DeleteKeyAccessServerResponse]
	listKeyAccessServerGrants *connect.Client[kasregistry.ListKeyAccessServerGrantsRequest, kasregistry.ListKeyAccessServerGrantsResponse]
}

// ListKeyAccessServers calls
// policy.kasregistry.KeyAccessServerRegistryService.ListKeyAccessServers.
func (c *keyAccessServerRegistryServiceClient) ListKeyAccessServers(ctx context.Context, req *connect.Request[kasregistry.ListKeyAccessServersRequest]) (*connect.Response[kasregistry.ListKeyAccessServersResponse], error) {
	return c.listKeyAccessServers.CallUnary(ctx, req)
}

// GetKeyAccessServer calls policy.kasregistry.KeyAccessServerRegistryService.GetKeyAccessServer.
func (c *keyAccessServerRegistryServiceClient) GetKeyAccessServer(ctx context.Context, req *connect.Request[kasregistry.GetKeyAccessServerRequest]) (*connect.Response[kasregistry.GetKeyAccessServerResponse], error) {
	return c.getKeyAccessServer.CallUnary(ctx, req)
}

// CreateKeyAccessServer calls
// policy.kasregistry.KeyAccessServerRegistryService.CreateKeyAccessServer.
func (c *keyAccessServerRegistryServiceClient) CreateKeyAccessServer(ctx context.Context, req *connect.Request[kasregistry.CreateKeyAccessServerRequest]) (*connect.Response[kasregistry.CreateKeyAccessServerResponse], error) {
	return c.createKeyAccessServer.CallUnary(ctx, req)
}

// UpdateKeyAccessServer calls
// policy.kasregistry.KeyAccessServerRegistryService.UpdateKeyAccessServer.
func (c *keyAccessServerRegistryServiceClient) UpdateKeyAccessServer(ctx context.Context, req *connect.Request[kasregistry.UpdateKeyAccessServerRequest]) (*connect.Response[kasregistry.UpdateKeyAccessServerResponse], error) {
	return c.updateKeyAccessServer.CallUnary(ctx, req)
}

// DeleteKeyAccessServer calls
// policy.kasregistry.KeyAccessServerRegistryService.DeleteKeyAccessServer.
func (c *keyAccessServerRegistryServiceClient) DeleteKeyAccessServer(ctx context.Context, req *connect.Request[kasregistry.DeleteKeyAccessServerRequest]) (*connect.Response[kasregistry.DeleteKeyAccessServerResponse], error) {
	return c.deleteKeyAccessServer.CallUnary(ctx, req)
}

// ListKeyAccessServerGrants calls
// policy.kasregistry.KeyAccessServerRegistryService.ListKeyAccessServerGrants.
func (c *keyAccessServerRegistryServiceClient) ListKeyAccessServerGrants(ctx context.Context, req *connect.Request[kasregistry.ListKeyAccessServerGrantsRequest]) (*connect.Response[kasregistry.ListKeyAccessServerGrantsResponse], error) {
	return c.listKeyAccessServerGrants.CallUnary(ctx, req)
}

// KeyAccessServerRegistryServiceHandler is an implementation of the
// policy.kasregistry.KeyAccessServerRegistryService service.
type KeyAccessServerRegistryServiceHandler interface {
	ListKeyAccessServers(context.Context, *connect.Request[kasregistry.ListKeyAccessServersRequest]) (*connect.Response[kasregistry.ListKeyAccessServersResponse], error)
	GetKeyAccessServer(context.Context, *connect.Request[kasregistry.GetKeyAccessServerRequest]) (*connect.Response[kasregistry.GetKeyAccessServerResponse], error)
	CreateKeyAccessServer(context.Context, *connect.Request[kasregistry.CreateKeyAccessServerRequest]) (*connect.Response[kasregistry.CreateKeyAccessServerResponse], error)
	UpdateKeyAccessServer(context.Context, *connect.Request[kasregistry.UpdateKeyAccessServerRequest]) (*connect.Response[kasregistry.UpdateKeyAccessServerResponse], error)
	DeleteKeyAccessServer(context.Context, *connect.Request[kasregistry.DeleteKeyAccessServerRequest]) (*connect.Response[kasregistry.DeleteKeyAccessServerResponse], error)
	ListKeyAccessServerGrants(context.Context, *connect.Request[kasregistry.ListKeyAccessServerGrantsRequest]) (*connect.Response[kasregistry.ListKeyAccessServerGrantsResponse], error)
}

// NewKeyAccessServerRegistryServiceHandler builds an HTTP handler from the service implementation.
// It returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewKeyAccessServerRegistryServiceHandler(svc KeyAccessServerRegistryServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	keyAccessServerRegistryServiceListKeyAccessServersHandler := connect.NewUnaryHandler(
		KeyAccessServerRegistryServiceListKeyAccessServersProcedure,
		svc.ListKeyAccessServers,
		connect.WithSchema(keyAccessServerRegistryServiceListKeyAccessServersMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	keyAccessServerRegistryServiceGetKeyAccessServerHandler := connect.NewUnaryHandler(
		KeyAccessServerRegistryServiceGetKeyAccessServerProcedure,
		svc.GetKeyAccessServer,
		connect.WithSchema(keyAccessServerRegistryServiceGetKeyAccessServerMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	keyAccessServerRegistryServiceCreateKeyAccessServerHandler := connect.NewUnaryHandler(
		KeyAccessServerRegistryServiceCreateKeyAccessServerProcedure,
		svc.CreateKeyAccessServer,
		connect.WithSchema(keyAccessServerRegistryServiceCreateKeyAccessServerMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	keyAccessServerRegistryServiceUpdateKeyAccessServerHandler := connect.NewUnaryHandler(
		KeyAccessServerRegistryServiceUpdateKeyAccessServerProcedure,
		svc.UpdateKeyAccessServer,
		connect.WithSchema(keyAccessServerRegistryServiceUpdateKeyAccessServerMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	keyAccessServerRegistryServiceDeleteKeyAccessServerHandler := connect.NewUnaryHandler(
		KeyAccessServerRegistryServiceDeleteKeyAccessServerProcedure,
		svc.DeleteKeyAccessServer,
		connect.WithSchema(keyAccessServerRegistryServiceDeleteKeyAccessServerMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	keyAccessServerRegistryServiceListKeyAccessServerGrantsHandler := connect.NewUnaryHandler(
		KeyAccessServerRegistryServiceListKeyAccessServerGrantsProcedure,
		svc.ListKeyAccessServerGrants,
		connect.WithSchema(keyAccessServerRegistryServiceListKeyAccessServerGrantsMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	return "/policy.kasregistry.KeyAccessServerRegistryService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case KeyAccessServerRegistryServiceListKeyAccessServersProcedure:
			keyAccessServerRegistryServiceListKeyAccessServersHandler.ServeHTTP(w, r)
		case KeyAccessServerRegistryServiceGetKeyAccessServerProcedure:
			keyAccessServerRegistryServiceGetKeyAccessServerHandler.ServeHTTP(w, r)
		case KeyAccessServerRegistryServiceCreateKeyAccessServerProcedure:
			keyAccessServerRegistryServiceCreateKeyAccessServerHandler.ServeHTTP(w, r)
		case KeyAccessServerRegistryServiceUpdateKeyAccessServerProcedure:
			keyAccessServerRegistryServiceUpdateKeyAccessServerHandler.ServeHTTP(w, r)
		case KeyAccessServerRegistryServiceDeleteKeyAccessServerProcedure:
			keyAccessServerRegistryServiceDeleteKeyAccessServerHandler.ServeHTTP(w, r)
		case KeyAccessServerRegistryServiceListKeyAccessServerGrantsProcedure:
			keyAccessServerRegistryServiceListKeyAccessServerGrantsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedKeyAccessServerRegistryServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedKeyAccessServerRegistryServiceHandler struct{}

func (UnimplementedKeyAccessServerRegistryServiceHandler) ListKeyAccessServers(context.Context, *connect.Request[kasregistry.ListKeyAccessServersRequest]) (*connect.Response[kasregistry.ListKeyAccessServersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("policy.kasregistry.KeyAccessServerRegistryService.ListKeyAccessServers is not implemented"))
}

func (UnimplementedKeyAccessServerRegistryServiceHandler) GetKeyAccessServer(context.Context, *connect.Request[kasregistry.GetKeyAccessServerRequest]) (*connect.Response[kasregistry.GetKeyAccessServerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("policy.kasregistry.KeyAccessServerRegistryService.GetKeyAccessServer is not implemented"))
}

func (UnimplementedKeyAccessServerRegistryServiceHandler) CreateKeyAccessServer(context.Context, *connect.Request[kasregistry.CreateKeyAccessServerRequest]) (*connect.Response[kasregistry.CreateKeyAccessServerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("policy.kasregistry.KeyAccessServerRegistryService.CreateKeyAccessServer is not implemented"))
}

func (UnimplementedKeyAccessServerRegistryServiceHandler) UpdateKeyAccessServer(context.Context, *connect.Request[kasregistry.UpdateKeyAccessServerRequest]) (*connect.Response[kasregistry.UpdateKeyAccessServerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("policy.kasregistry.KeyAccessServerRegistryService.UpdateKeyAccessServer is not implemented"))
}

func (UnimplementedKeyAccessServerRegistryServiceHandler) DeleteKeyAccessServer(context.Context, *connect.Request[kasregistry.DeleteKeyAccessServerRequest]) (*connect.Response[kasregistry.DeleteKeyAccessServerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("policy.kasregistry.KeyAccessServerRegistryService.DeleteKeyAccessServer is not implemented"))
}

func (UnimplementedKeyAccessServerRegistryServiceHandler) ListKeyAccessServerGrants(context.Context, *connect.Request[kasregistry.ListKeyAccessServerGrantsRequest]) (*connect.Response[kasregistry.ListKeyAccessServerGrantsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("policy.kasregistry.KeyAccessServerRegistryService.ListKeyAccessServerGrants is not implemented"))
}
