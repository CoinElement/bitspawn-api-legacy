package signURL

import (
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sirupsen/logrus"
)

type SignURLManager struct {
	svc        *s3.S3
	log        *logrus.Entry
	BucketName string
}

// NewSignURLManager is used to create a signURLManager
func NewSignURLManager(s3Region string, log *logrus.Logger, s3BucketName string) *SignURLManager {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(s3Region)},
	)
	if err != nil {
		return nil
	}
	svc := s3.New(sess)
	return &SignURLManager{
		svc:        svc,
		log:        log.WithField("service", "signURL"),
		BucketName: s3BucketName,
	}
}

// GetSignedURL is used to create a pre-signed url based on the provided bucket and key
func (s *SignURLManager) GetSignedURL(filename string) (string, http.Header, error) {
	req, _ := s.svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(filename),
		ACL:    aws.String("public-read"),
	})
	urlStr, headers, err := req.PresignRequest(3 * time.Minute)
	if err != nil {
		s.log.Error("error in generating presignURL: " + err.Error())
		return "", nil, err
	}

	return urlStr, headers, nil
}

// GetSignedURL is used to create a pre-signed url based on the provided bucket and key
func (s *SignURLManager) GetReadOnlySignedURL(filename string) (string, error) {
	req, _ := s.svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(filename),
	})
	urlStr, err := req.Presign(3 * time.Minute)
	if err != nil {
		s.log.Error("error in generating presignURL: " + err.Error())
		return "", err
	}

	return urlStr, nil
}
