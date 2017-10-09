Packer s3 loader provisioner plugin
========================================
This plug in is used to have packer loading files from s3 buckets in to an ami image.

Naming
----------------------
The end file will be renamed to `packer-provisioner-s3loader`, since this is the convention packer expects.

Using parameters
----------------------


Running packer
----------------------
Sample packer command `packer build -var 'aws_access_key=' -var 'aws_secret_key=' -var 'source_ami=' ./packer-files/sample-provisioner.json`
