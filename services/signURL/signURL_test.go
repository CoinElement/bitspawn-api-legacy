package signURL

import (
	"net/http"
	"testing"

	"github.com/aws/aws-sdk-go/awstesting/unit"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
)

func TestGetURL(t *testing.T) {
	svc := s3.New(unit.Session)

	s := SignURLManager{svc, nil, "mybucket"}
	url, headers, err := s.GetSignedURL("mykey")
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, http.Header{"x-amz-acl": []string{"public-read"}}, headers)
	assert.Contains(t, url, "https://mybucket.s3.mock-region.amazonaws.com/mykey?")
}
