package awsurl

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"time"
)

func PreSign(obj string) string{
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	svc := s3.New(sess, (&aws.Config{Region: aws.String("us-west-2")}))
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("testBucket"),
		Key:    aws.String(obj),
	})
	urlStr, err := req.Presign(15 * time.Minute)

	if err != nil {
		log.Println("Failed to sign request", err)
	}

	log.Println("Signed URL: ", urlStr)
	return urlStr
}
