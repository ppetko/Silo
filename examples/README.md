#  Silo - Cloud Native Backup Solution

### Help menu

```
$ ./silo 
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

### Glacier menu

```
$ ./silo glacier 
NAME:
   Silo glacier - glacier operations

USAGE:
   Silo glacier command [command options] [arguments...]

COMMANDS:
   create-vault              create new vault
   list-vaults               list valuts in region
   list-jobs                 list jobs per vault
   describe-vault            get information about a previously initiated job
   describe-job              get information about a vault
   upload-archive            upload data to a vault
   init-inventory-retrieval  initiate an inventory-retrieval job for vault
   init-archive-retrieval    initiate an archive-retrieval job for vault
   get-inventory             get output of inventory retrieval job
   get-archive               get output of archive retrieval job
   get-vaultlock             get information about vault's policy
   init-vaultlock            init vault lock policy on the specified vault
   abort-vaultlock           aborts the vault locking process if not already in locked state
   complete-vaultlock        complete vault lock in process.
   get-retrieval-policy      get the current data retrieval policy
   delete-archive            delete archive from vault
   delete-vault              delete empty vault
   help, h                   Shows a list of commands or help for one command

OPTIONS:
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

### Apply policy on existing vault

```
$ ./silo glacier init-vaultlock --name my-vault --region us-east-2 --policy "{\"Version\":\"2012-10-17\",\"Statement\":[{\"Sid\":\"Define-vault-lock\",\"Effect\":\"Deny\",\"Principal\":{\"AWS\":\"*\"},\"Action\":\"glacier:DeleteArchive\",\"Resource\":\"arn:aws:glacier:us-east-2:757758175257:vaults/my-vault\",\"Condition\":{\"NumericLessThanEquals\":{\"glacier:ArchiveAgeinDays\":\"365\"}}}]}"
{
  LockId: "Juqai_nVz5z6ZSeZA7GRnHSL"
}

$ ./silo glacier abort-vaultlock --name my-vault --region us-east-2
Vault lock aborted on my-vault

```