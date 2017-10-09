package aws

import (
	"testing"
)

const (
	jsonTestFilename = "../../../secrets/aws_config.json"
)

func LoadJSONInfo(t *testing.T) {
	info := BtrzAwsAuthenticationInfo{}
	err := info.LoadFromJSONFile(jsonTestFilename)
	if err != nil {
		t.SkipNow()
	}
	if info.AwsKey == "" {
		t.Fatal("Can't get aws key.")
	}
}

func GetSessionFromParameters(t *testing.T) {
	info := BtrzAwsAuthenticationInfo{}
	err := info.LoadFromJSONFile(jsonTestFilename)
	if err != nil {
		t.SkipNow()
	}
	sess, err := GetAWSSessionWithKey(info.AwsKey, info.AwsSecret)
	if err != nil {
		t.Fatal("err creating aws session", err)
	}
	if sess == nil {
		t.Fatal("aws session is nil")
	}
}
