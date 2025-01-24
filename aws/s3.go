package aws

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	iLog "github.com/bartlomiej-jedrol/go-toolkit/log"
)

func GetS3Object(ctx context.Context, s3Client *s3.Client, bucket, objectKey string) (*io.ReadCloser, error) {
	function := "getS3Objectinfo"

	s3Input := s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
	}
	s3Output, err := s3Client.GetObject(ctx, &s3Input)
	if err != nil {
		iLog.Error("failed to get s3 object", nil, err, Service, function)
		return nil, err
	}
	iLog.Info("successfully get the s3 object", s3Output.ContentType, nil, Service, function)

	return &s3Output.Body, nil
}
