package awsurl

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3PreSign presigns the url for s3 GET allowing for signed downloads of an s3 asset.
// obj to be signed and returned as a presigned string
// awsRegion is the aws region the s3 bucket resides
// s3Bucket is the s3Bucket the object is located
// presignTime is the length of time the presignedURL is good for in minutes
// Returns the signed url as a string.
func S3PreSign(obj string, awsRegion string, s3Bucket string, presignTime int) string {
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	svc := s3.New(sess, (&aws.Config{Region: aws.String(awsRegion)}))
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(obj),
	})
	urlStr, err := req.Presign(time.Duration(presignTime) * time.Minute)

	if err != nil {
		log.Println("Failed to sign request", err)
	}

	log.Println("Signed URL: ", urlStr)
	return urlStr
}
