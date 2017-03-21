package awsurl

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// S3Session builds an interface to s3
type S3Session struct {
	Client s3iface.S3API
}

// s3signer signs urls for aws S3
func (s *S3Session) s3signer(obj string, s3Bucket string, presignTime int) (string, error) {
	req, _ := s.Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(obj),
	})
	urlStr, err := req.Presign(time.Duration(presignTime) * time.Minute)

	if err != nil {
		log.Println("Failed to sign request", err)
	}

	log.Println("Signed URL: ", urlStr)
	return urlStr, nil
}

// S3PreSign presigns the url for s3 GET allowing for signed downloads of an s3 asset.
// obj to be signed and returned as a presigned string
// awsRegion is the aws region the s3 bucket resides
// s3Bucket is the s3Bucket the object is located
// presignTime is the length of time the presignedURL is good for in minutes
// Returns the signed url as a string.
func S3PreSign(obj string, s3Bucket string, presignTime int) string {
	sess, err := session.NewSession()
	if err != nil {
		log.Panic(err)
	}

	svc := S3Session{Client: s3.New(sess)}
	urlStr, _ := svc.s3signer(obj, s3Bucket, presignTime)
	return urlStr
}
