#  Silo - Cloud Native Backup Solution

### Usage

```
./silo 
NAME:
   Silo - Cloud Native Backup Solution

USAGE:
   silo [global options] command [command options] [arguments...]

VERSION:
   0.1-beta

COMMANDS:
   configure, c  setup aws credentials
   glacier       glacier operations
   s3            s3 operations
   help, h       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

### Create AWS IAM
* Create AWS user with policy permissions: AmazonGlacierFullAccess and AmazonS3FullAccess. 
* DO NOT use admin account.

### Get Started 
```
$ ./silo configure 
AWS Access Key ID: ******
AWS Secret Access: ******
Default region name: us-east-2
Default output format:[json] json
2020/01/08 17:14:04 config file generated ~/.aws/config
2020/01/08 17:14:04 credential file generated ~/.aws/credentials
```

### Export AWS Region

```
export AWS_DEFAULT_REGION=us-east-2
```

### List all vaults 

```
$ ./silo glacier list-vaults 
{
  VaultList: []
}
```
## Pull requests welcome!