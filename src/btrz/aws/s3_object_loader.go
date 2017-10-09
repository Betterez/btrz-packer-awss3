package aws

import (
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// GetLastObject getting the last object in a bucket
func GetLastObject(bucketName, bucketPathName string, sess *session.Session) (*s3.Object, error) {
	var selectedObject *s3.Object
	var batchSize int64 = 500
	svc := s3.New(sess)
	listingParams := &s3.ListObjectsInput{
		Bucket:    aws.String(bucketName), // Required
		Delimiter: aws.String("Delimiter"),
		MaxKeys:   &batchSize,
	}
	for {
		resp, err := svc.ListObjects(listingParams)
		if err != nil {
			return nil, err
		}
		for _, currentObject := range resp.Contents {
			if !strings.Contains(*currentObject.Key, bucketPathName) {
				continue
			}
			if selectedObject == nil {
				selectedObject = currentObject
			} else {
				if currentObject.LastModified.After(*selectedObject.LastModified) {
					selectedObject = currentObject
				}
			}
		}
		if len(resp.Contents) < int(batchSize) {
			break
		} else {
			listingParams.Marker = selectedObject.Key
		}
	}
	return selectedObject, nil
}

func loadBackup(selectedObject *s3.Object, bucketName, destinationFolderName string, sess *session.Session) (outputFileName, fileName string, err error) {
	svc := s3.New(sess)
	params := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    selectedObject.Key,
	}
	resp, err := svc.GetObject(params)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()
	fullFileName, fileName := getObjectNamePath(selectedObject.Key, destinationFolderName)
	destinationFile, err := os.Create(fullFileName)
	if err != nil {
		return "", "", err
	}
	defer destinationFile.Close()
	dataBytes := make([]byte, 4096)
	for {
		bytesRead, _ := resp.Body.Read(dataBytes)
		if err != nil {
			return "", "", err
		}
		if bytesRead > 0 {
			bytesToWrite := dataBytes[:bytesRead]
			destinationFile.Write(bytesToWrite)
		} else {
			break
		}
	}
	return fullFileName, fileName, nil
}

//LoadLastObjectFromBucket loading last object from a bucket in to destination file
func LoadLastObjectFromBucket(bucketName, destinationFolderName, bucketPathName string, sess *session.Session) (bool, string, string, error) {
	lastObject, err := GetLastObject(bucketName, bucketPathName, sess)
	if err != nil {
		return false, "", "", err
	}
	if lastObject == nil {
		return false, "", "", nil
	}
	fullPathName, outputFileName, err := loadBackup(lastObject, bucketName, destinationFolderName, sess)
	if err != nil {
		return false, "", "", err
	}
	return true, fullPathName, outputFileName, nil
}

func getObjectNamePath(objectKeyValue *string, destinationFolderName string) (string, string) {
	var fileName string
	fullFileName := destinationFolderName
	if fullFileName[len(fullFileName)-1] != os.PathSeparator {
		fullFileName += string(os.PathSeparator)
	}
	if objectKeyValue != nil {
		pathValues := strings.Split(*objectKeyValue, "/")
		fileName = pathValues[len(pathValues)-1]
		fullFileName += fileName
	}
	return fullFileName, fileName
}
