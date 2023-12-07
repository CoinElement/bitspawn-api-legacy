package s3

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"testing"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/awstesting/unit"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/stretchr/testify/assert"
)

func TestHandleS3AvatarUpload(t *testing.T) {
	client := NewS3UploadClientWithUploader(getUploader(), nil, "us-west-1", "test")
	file1, err := os.Open("s3_upload.go")
	assert.NoError(t, err, "No Error is expected")
	location, err := client.HandleS3AvatarUpload("test", "file1", file1)
	assert.Contains(t, location, "https://test.s3.mock-region.amazonaws.com/userAvatar/file1/")
	assert.NoError(t, err, "No Error is expected")
}
func TestHandleS3Download(t *testing.T) {
	client := NewS3UploadClientWithUploader(nil, getDownloader(), "us-west-1", "test")
	file, err := client.HandleS3Download("test", "file1", "urlkey")
	assert.NoError(t, err, "No Error is expected")
	assert.NotNil(t, file)
	file1, err := os.Open("file1")
	assert.NoError(t, err, "No Error is expected")
	fileInfo, err := file1.Stat()
	assert.NoError(t, err, "No Error is expected")
	assert.True(t, fileInfo.Size() > 0, "size should be greater than zero")
	assert.NoError(t, err, "No Error is expected")
	err = os.Remove(fileInfo.Name())
	assert.NoError(t, err, "No Error is expected")
}

func TestHandleS3TournamentBannerUpload(t *testing.T) {
	client := NewS3UploadClientWithUploader(getUploader(), nil, "us-west-1", "test")
	file1, err := os.Open("s3_upload.go")
	assert.NoError(t, err, "No Error is expected")
	location, err := client.HandleS3TournamentBannerUpload("test", "file1", file1)
	assert.Contains(t, location, "https://test.s3.mock-region.amazonaws.com/tournamentBanner/file1/")
	assert.NoError(t, err, "No Error is expected")
}

func getDownloader() *s3manager.Downloader {

	var locker sync.Mutex
	payload := []byte(`some content`)

	svc := s3.New(unit.Session)
	svc.Handlers.Send.Clear()
	svc.Handlers.Send.PushBack(func(r *request.Request) {
		locker.Lock()
		defer locker.Unlock()
		r.HTTPResponse = &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewReader(payload)),
			Header:     http.Header{},
		}
		r.HTTPResponse.Header.Set("Content-Length", "1")
	})

	return s3manager.NewDownloaderWithClient(svc, func(d *s3manager.Downloader) {
		d.Concurrency = 1
		d.PartSize = 1
	})
}

func getUploader() *s3manager.Uploader {
	svc := s3.New(unit.Session)
	svc.Handlers.Unmarshal.Clear()
	svc.Handlers.UnmarshalMeta.Clear()
	svc.Handlers.UnmarshalError.Clear()
	svc.Handlers.Send.Clear()
	payload := []byte(`some text file`)

	// contentLen := ""
	svc.Handlers.Send.PushBack(func(r *request.Request) {
		r.HTTPResponse = &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(payload)),
		}
	})

	mgr := s3manager.NewUploaderWithClient(svc, func(u *s3manager.Uploader) {
		u.Concurrency = 1
		u.PartSize = 5242880
	})
	return mgr
}
