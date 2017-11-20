package awsurl

import (
	"testing"
)

func TestS3PreSign(t *testing.T) {
	obj := "/this-is-a-test-file"
	bucket := "test-bucket"
	time := 5

	S3PreSign(obj, bucket, time)
}
