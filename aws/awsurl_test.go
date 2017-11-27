package awsurl

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type mockS3Session struct {
	s3iface.S3API
}

func TestS3PreSign(t *testing.T) {
	obj := "/this-is-a-test-file"
	bucket := "test-bucket"
	time := 5

	// mockSession := &mockS3Session{}

	S3PreSign(obj, bucket, time)
	// mockSession.s3signer(obj, bucket, time)
}
