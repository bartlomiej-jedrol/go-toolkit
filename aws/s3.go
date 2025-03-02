package aws

import (
	"context"
	"io"
	"os"

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

func SaveS3ObjectToLambdaTemp(ctx context.Context, s3Client *s3.Client, bucket, objectKey, lambdaTempPath, fileName, fileExtension string) error {
	function := "SaveS3ObjectToLambdaTemp"

	s3ObjectBody, err := GetS3Object(ctx, s3Client, bucket, objectKey)
	if err != nil {
		return err
	}

	tmpFile, err := os.CreateTemp(lambdaTempPath, fileName+"-*."+fileExtension)
	if err != nil {
		iLog.Error("failed to create pdf file", nil, err, Service, function)
		return err
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	_, err = io.Copy(tmpFile, *s3ObjectBody)
	if err != nil {
		iLog.Error("failed to save S3 object to lambda tmp path", nil, err, Service, function)
		return err
	}
	iLog.Info("pdf saved to", tmpFile.Name(), nil, Service, function)

	return nil
}
