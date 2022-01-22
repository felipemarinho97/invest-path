package clients

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

/*
* CODE GENERATED AUTOMATICALLY WITH github.com/felipemarinho97/go-mock-gen
 */

// IS3UploaderClient generic client
type IS3UploaderClient interface {
	Upload(arg1 context.Context, arg2 *s3.PutObjectInput, arg3 ...func(*manager.Uploader)) (*manager.UploadOutput, error)
}

// S3UploaderClientMock generic client mock
type S3UploaderClientMock struct {
	UploadMock func(arg1 context.Context, arg2 *s3.PutObjectInput, arg3 ...func(*manager.Uploader)) (*manager.UploadOutput, error)
}

// Upload mocked function
func (m S3UploaderClientMock) Upload(arg1 context.Context, arg2 *s3.PutObjectInput, arg3 ...func(*manager.Uploader)) (*manager.UploadOutput, error) {
	return m.UploadMock(arg1, arg2, arg3...)
}
