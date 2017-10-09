package provisioner

import (
	"btrz/aws"
	"errors"
	"fmt"
	"github.com/hashicorp/packer/packer"
	"github.com/mitchellh/mapstructure"
	"os"
	"reflect"
)

// S3Loader - s3 loader provisioner
type S3Loader struct {
	runnerData RunnerData
	awsInfo    aws.BtrzAwsAuthenticationInfo
}

// Prepare - prepering
func (loader *S3Loader) Prepare(params ...interface{}) error {
	if len(params) > 1 {
		mapstructure.Decode(params[0], loader.runnerData)
		packerMainMap, ok := params[1].(map[interface{}]interface{})
		if !ok {
			return fmt.Errorf("S3Loader: Can't get packer parameters,type %s,\n%v", reflect.TypeOf(params[1]), params[1])
		}
		userInfoMap, ok := packerMainMap["packer_user_variables"].(map[interface{}]interface{})
		if !ok {
			return fmt.Errorf("Can't get aws parameters list %s", reflect.TypeOf(packerMainMap["packer_user_variables"]))
		}
		loader.awsInfo.AwsKey, ok = userInfoMap["aws_access_key"].(string)
		if !ok {
			return fmt.Errorf("Can't get aws parameters %s", reflect.TypeOf(packerMainMap["packer_user_variables"]))
		}
		loader.awsInfo.AwsSecret = userInfoMap["aws_secret_key"].(string)
	}

	return nil
}

// Provision - do work
func (loader *S3Loader) Provision(ui packer.Ui, communicator packer.Communicator) error {
	sess, err := aws.GetAWSSessionWithKey(loader.awsInfo.AwsKey, loader.awsInfo.AwsSecret)
	if err != nil {
		return err
	}
	ui.Message(fmt.Sprintf("loader data: %v", loader.runnerData))
	ok, fullPathName, fileName, err := aws.LoadLastObjectFromBucket("btrz-scaling-repo", "/home/tal/temp", "connex", sess)
	if !ok {
		ui.Error("couldn't connect to s3")
		return errors.New("not ok returning s3 object")
	}
	if err != nil {
		return err
	}
	ui.Message("loading " + fileName + " to remote host")
	fileHandle, err := os.Open(fullPathName)
	if err != nil {
		return err
	}
	defer fileHandle.Close()
	communicator.Upload("/home/ubuntu/"+fileName, fileHandle, nil)
	return nil
}

// Cancel - cancel provision
func (loader *S3Loader) Cancel() {

}
