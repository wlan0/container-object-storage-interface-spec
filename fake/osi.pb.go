package fake

import (
	"context"

	osi "sigs.k8s.io/osi-spec"

	"google.golang.org/grpc"
)

// this ensures that the mock implements the client interface
var _ osi.ProvisionerClient = (*MockProvisionerClient)(nil)

// MockProvisionerClient is a type that implements all the methods for RolePolicyAttachmentClient interface
type MockProvisionerClient struct {
	GetInfo            func(ctx context.Context, in *osi.ProvisionerGetInfoRequest, opts ...grpc.CallOption) (*osi.ProvisionerGetInfoResponse, error)
	CreateBucket       func(ctx context.Context, in *osi.ProvisionerCreateBucketRequest, opts ...grpc.CallOption) (*osi.ProvisionerCreateBucketResponse, error)
	DeleteBucket       func(ctx context.Context, in *osi.ProvisionerDeleteBucketRequest, opts ...grpc.CallOption) (*osi.ProvisionerDeleteBucketResponse, error)
	GrantBucketAccess  func(ctx context.Context, in *osi.ProvisionerGrantBucketAccessRequest, opts ...grpc.CallOption) (*osi.ProvisionerGrantBucketAccessResponse, error)
	RevokeBucketAccess func(ctx context.Context, in *osi.ProvisionerRevokeBucketAccessRequest, opts ...grpc.CallOption) (*osi.ProvisionerRevokeBucketAccessResponse, error)
}

// ProvisionerCreateBucket mocks GetBucketPolicyRequest method
func (m *MockProvisionerClient) ProvisionerCreateBucket(ctx context.Context, in *osi.ProvisionerCreateBucketRequest, opts ...grpc.CallOption) (*osi.ProvisionerCreateBucketResponse, error) {
	return m.CreateBucket(ctx, in, opts...)
}

// ProvisionerDeleteBucket mocks PutBucketPolicyRequest method
func (m *MockProvisionerClient) ProvisionerDeleteBucket(ctx context.Context, in *osi.ProvisionerDeleteBucketRequest, opts ...grpc.CallOption) (*osi.ProvisionerDeleteBucketResponse, error) {
	return m.DeleteBucket(ctx, in, opts...)
}

// ProvisionerGrantBucketAccess mocks DeleteBucketPolicyRequest method
func (m *MockProvisionerClient) ProvisionerGrantBucketAccess(ctx context.Context, in *osi.ProvisionerGrantBucketAccessRequest, opts ...grpc.CallOption) (*osi.ProvisionerGrantBucketAccessResponse, error) {
	return m.GrantBucketAccess(ctx, in, opts...)
}

// ProvisionerRevokeBucketAccess mocks DeleteBucketPolicyRequest method
func (m *MockProvisionerClient) ProvisionerRevokeBucketAccess(ctx context.Context, in *osi.ProvisionerRevokeBucketAccessRequest, opts ...grpc.CallOption) (*osi.ProvisionerRevokeBucketAccessResponse, error) {
	return m.RevokeBucketAccess(ctx, in, opts...)
}

// ProvisionerRevokeBucketAccess mocks DeleteBucketPolicyRequest method
func (m *MockProvisionerClient) ProvisionerGetInfo(ctx context.Context, in *osi.ProvisionerGetInfoRequest, opts ...grpc.CallOption) (*osi.ProvisionerGetInfoResponse, error) {
	return m.GetInfo(ctx, in, opts...)
}
