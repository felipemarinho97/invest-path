package clients

import (
	"context"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

/*
* CODE GENERATED AUTOMATICALLY WITH github.com/felipemarinho97/go-mock-gen
 */

// IS3PresignClient generic client
type IS3PresignClient interface {
	// PresignGetObject is used to generate a presigned HTTP Request which contains
	// presigned URL, signed headers and HTTP method used.
	//
	PresignGetObject(arg1 context.Context, arg2 *s3.GetObjectInput, arg3 ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
	// PresignPutObject is used to generate a presigned HTTP Request which contains
	// presigned URL, signed headers and HTTP method used.
	//
	PresignPutObject(arg1 context.Context, arg2 *s3.PutObjectInput, arg3 ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
	// PresignUploadPart is used to generate a presigned HTTP Request which contains
	// presigned URL, signed headers and HTTP method used.
	//
	PresignUploadPart(arg1 context.Context, arg2 *s3.UploadPartInput, arg3 ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
}

// S3PresignClientMock generic client mock
type S3PresignClientMock struct {
	PresignGetObjectMock  func(arg1 context.Context, arg2 *s3.GetObjectInput, arg3 ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
	PresignPutObjectMock  func(arg1 context.Context, arg2 *s3.PutObjectInput, arg3 ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
	PresignUploadPartMock func(arg1 context.Context, arg2 *s3.UploadPartInput, arg3 ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
}

// PresignGetObject mocked funcition
func (m S3PresignClientMock) PresignGetObject(arg1 context.Context, arg2 *s3.GetObjectInput, arg3 ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	return m.PresignGetObjectMock(arg1, arg2, arg3...)
}

// PresignPutObject mocked funcition
func (m S3PresignClientMock) PresignPutObject(arg1 context.Context, arg2 *s3.PutObjectInput, arg3 ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	return m.PresignPutObjectMock(arg1, arg2, arg3...)
}

// PresignUploadPart mocked funcition
func (m S3PresignClientMock) PresignUploadPart(arg1 context.Context, arg2 *s3.UploadPartInput, arg3 ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	return m.PresignUploadPartMock(arg1, arg2, arg3...)
}
