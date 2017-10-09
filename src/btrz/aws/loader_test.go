package aws

import (
	"os"
	"testing"
)

func TestLoader(t *testing.T) {
	info := &BtrzAwsAuthenticationInfo{}
	const storagePath = "/home/tal/temp"
	err := info.LoadFromJSONFile("../../../secrets/aws_config.json")
	if err != nil {
		t.SkipNow()
	}
	sess, err := GetAWSSessionWithKey(info.AwsKey, info.AwsSecret)
	if err != nil {
		t.Fatal("error getting session", err)
	}
	if _, err = os.Stat(storagePath); os.IsNotExist(err) {
		os.Mkdir(storagePath, 0777)
	}
	ok, fullPathName, fileName, err := LoadLastObjectFromBucket("btrz-scaling-repo", storagePath, "connex", sess)
	if err != nil {
		t.Fatal("error getting object", err)
	}
	if !ok {
		t.Fatal("LoadLastObjectFromBucket returned an false")
	}
	if fullPathName == "" {
		t.Fatal("fullPathName is empty")
	}
	if fileName == "" {
		t.Fatal("fileName is empty")
	}
}
