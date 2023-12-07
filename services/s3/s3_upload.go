package s3

import (
	"bytes"
	"io/ioutil"
	"log"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	// "github.com/aws/aws-sdk-go/service/s3"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	uuid "github.com/satori/go.uuid"
)

type S3UploadClient struct {
	S3Region     string
	S3BucketName string
	uploader     *s3manager.Uploader
	downloader   *s3manager.Downloader
}

// NewS3Upload ...
func NewS3UploadClient(s3Region string, s3BucketName string) *S3UploadClient {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(s3Region)},
	)
	if err != nil {
		exitErrorf("couldn't create new session based on the region%s :%v", s3Region, err)
	}
	return &S3UploadClient{
		S3Region:     s3Region,
		S3BucketName: s3BucketName,
		uploader:     s3manager.NewUploader(sess),
		downloader:   s3manager.NewDownloader(sess),
	}
}

func NewS3UploadClientWithUploader(uploader *s3manager.Uploader, downloader *s3manager.Downloader, s3Region string, s3BucketName string) *S3UploadClient {
	return &S3UploadClient{
		S3Region:     s3Region,
		S3BucketName: s3BucketName,
		uploader:     uploader,
		downloader:   downloader,
	}
}
func (s3upload *S3UploadClient) Upload(bucket, filename string, file multipart.File) (string, error) {
	out, err := s3upload.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   file,
		ACL:    aws.String("public-read"), // could be private if you want it to be access by only authorized users
	})
	if err != nil {
		// Print the error and exit.
		exitErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
		return "", err
	}

	fmt.Printf("Successfully uploaded %q to %q\n", filename, bucket)

	return out.Location, nil
}

func (s3upload *S3UploadClient) HandleS3AvatarUpload(bucketname string, filenameforupload string, file multipart.File) (string, error) {
	return s3upload.Upload(bucketname, fmt.Sprintf("userAvatar/%s/%s", filenameforupload, uuid.NewV4().String()), file)
}
func (s3upload *S3UploadClient) HandleS3TeamAvatarUpload(bucketname string, filenameforupload string, file multipart.File) (string, error) {
	return s3upload.Upload(bucketname, fmt.Sprintf("teamAvatar/%s/%s", filenameforupload, uuid.NewV4().String()), file)
}
func (s3upload *S3UploadClient) HandleS3ProfileBannerUpload(bucketname string, filenameforupload string, file multipart.File) (string, error) {

	return s3upload.Upload(bucketname, fmt.Sprintf("userProfileBanner/%s/%s", filenameforupload, uuid.NewV4().String()), file)
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func (s3upload *S3UploadClient) HandleS3TournamentBannerUpload(bucketname string, subfoldername string, file multipart.File) (string, error) {

	return s3upload.Upload(bucketname, fmt.Sprintf("tournamentBanner/%s/%s", subfoldername, uuid.NewV4().String()), file)
}

func (s3upload *S3UploadClient) HandleS3TournamentLogoUpload(bucketname string, subfoldername string, file multipart.File) (string, error) {

	return s3upload.Upload(bucketname, fmt.Sprintf("tournamentLogo/%s/%s", subfoldername, uuid.NewV4().String()), file)
}

func (s3upload *S3UploadClient) HandleS3TournamentThumbnailUpload(bucketname string, subfoldername string, file multipart.File) (string, error) {

	return s3upload.Upload(bucketname, fmt.Sprintf("tournamentThumbnail/%s/%s", subfoldername, uuid.NewV4().String()), file)
}

func (s3upload *S3UploadClient) HandleS3BadgeUpload(bucketname string, file multipart.File) (string, error) {
	return s3upload.Upload(bucketname, fmt.Sprintf("badge/%s", uuid.NewV4().String()), file)
}

func (s3upload *S3UploadClient) HandleS3ScreenshotUpload(bucketname string, filenameforupload string, file multipart.File) (string, error) {
	return s3upload.Upload(bucketname, fmt.Sprintf("screenshot/%s", uuid.NewV4().String()), file)
}

func (s3upload *S3UploadClient) HandleS3Download(bucketname string, filename string, urlkey string) (multipart.File /*bytes.Reader*/, error) {

	// Create a file to write the S3 Object contents to.
	downloadFile, err := os.Create(filename)
	if err != nil {
		fmt.Printf("failed to create file %q, %v", filename, err)
		return nil, err
	}

	// Write the contents of S3 Object to the file
	n, err := s3upload.downloader.Download(downloadFile, &s3.GetObjectInput{
		Bucket: aws.String(bucketname),
		Key:    aws.String(urlkey),
	})
	if err != nil {
		fmt.Printf("failed to download file, %v", err)
		return nil, err
	}
	fmt.Printf("file downloaded, %d bytes\n", n)
	defer downloadFile.Close()
	downloadBuffer, err := ioutil.ReadAll(downloadFile)
	if err != nil {
		fmt.Printf("failed to read file to buffer, %v", err)
		return nil, err
	}
	reader := bytes.NewReader(downloadBuffer)

	// io.Copy(downloadBuffer, downloadFile)
	log.Println("downloadBuffer", downloadBuffer)
	log.Println("downloadFile", reader)

	return downloadFile, nil
}

func (s3upload *S3UploadClient) HandleClubAvatarUpload(bucketname string, filenameforupload string, file multipart.File) (string, error) {
	return s3upload.Upload(bucketname, fmt.Sprintf("clubAvatar/%s/%s", filenameforupload, uuid.NewV4().String()), file)
}

func (s3upload *S3UploadClient) HandleClubBannerUpload(bucketname string, filenameforupload string, file multipart.File) (string, error) {
	return s3upload.Upload(bucketname, fmt.Sprintf("clubBanner/%s/%s", filenameforupload, uuid.NewV4().String()), file)
}
