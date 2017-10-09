package aws

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
)

// BtrzAwsAuthenticationInfo - authentication info
type BtrzAwsAuthenticationInfo struct {
	AwsKey    string `json:"aws_access_key_id"`
	AwsSecret string `json:"aws_secret_access_key"`
}

// LoadFromJSONFile - loads aws authentication info from a json file
func (info *BtrzAwsAuthenticationInfo) LoadFromJSONFile(filename string) error {
	fileHandle, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileHandle.Close()
	decoder := json.NewDecoder(fileHandle)
	decoder.Decode(info)
	return nil
}

// GetAWSSession -  creates an aws session
func GetAWSSession() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
	if err != nil {
		return nil, err
	}
	if sess == nil {
		return nil, errors.New("can't create aws session")
	}
	return sess, nil
}

// GetAWSSessionWithKey - get session from a key
func GetAWSSessionWithKey(AWSKey, AWSSecret string) (sess *session.Session, err error) {
	creds := credentials.NewStaticCredentials(AWSKey, AWSSecret, "")
	sess, err = session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String("us-east-1")})
	return sess, err
}
