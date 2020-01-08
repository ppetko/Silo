package main

import (
	"log"
	"os"
	"time"

	"github.com/ppetko/silo/aws"
	"github.com/urfave/cli"
)

var (
	region string
)

func main() {
	app := cli.NewApp()
	app.Name = "Silo"
	app.Usage = "Cloud Native Backup Solution"
	app.Version = "0.1-beta"
	app.Compiled = time.Now()

	app.Commands = []*cli.Command{
		{
			Name:    "configure",
			Aliases: []string{"c"},
			Usage:   "setup aws credentials",
			Action: func(c *cli.Context) error {
				aws.SetupAWSAuth()
				return nil
			},
		},
		{
			Name:  "glacier",
			Usage: "glacier operations",
			Subcommands: []*cli.Command{
				{
					Name:  "create-vault",
					Usage: "create new vault",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Usage: "vault name",
						},
						&cli.StringFlag{
							Name:        "region",
							Usage:       "aws region",
							EnvVars:     []string{"AWS_DEFAULT_REGION"},
							Destination: &region,
						},
					},
					Action: func(c *cli.Context) error {
						if c.String("name") == "" || region == "" {
							return cli.NewExitError("specify vault name and region using --name and --region", 2)
						}
						aws.CreateVault(region, c.String("name"))
						return nil
					},
				},
				{
					Name:  "list-vaults",
					Usage: "list valuts in region",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:        "region",
							Usage:       "aws region",
							EnvVars:     []string{"AWS_DEFAULT_REGION"},
							Destination: &region,
						},
					},
					Action: func(c *cli.Context) error {
						aws.ListVault(region)
						return nil
					},
				},
				{
					Name:  "list-jobs",
					Usage: "list jobs per vault",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Usage: "vault name",
						},
						&cli.StringFlag{
							Name:        "region",
							Usage:       "aws region",
							EnvVars:     []string{"AWS_DEFAULT_REGION"},
							Destination: &region,
						},
					},
					Action: func(c *cli.Context) error {
						if c.String("name") == "" || region == "" {
							return cli.NewExitError("specify vault name and region using --name and --region", 2)
						}
						aws.ListJobs(region, c.String("name"))
						return nil
					},
				},
				{
					Name:  "describe-vault",
					Usage: "get information about a previously initiated job",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Usage: "vault name",
						},
						&cli.StringFlag{
							Name:        "region",
							Usage:       "aws region",
							EnvVars:     []string{"AWS_DEFAULT_REGION"},
							Destination: &region,
						},
					},
					Action: func(c *cli.Context) error {
						if c.String("name") == "" || region == "" {
							return cli.NewExitError("specify vault name and region using --name and --region", 2)
						}
						aws.DescriveVault(region, c.String("name"))
						return nil
					},
				},
				{
					Name:  "describe-job",
					Usage: "get information about a vault",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Usage: "vault name",
						},
						&cli.StringFlag{
							Name:  "jobID",
							Usage: "specify a job ID",
						},
						&cli.StringFlag{
							Name:        "region",
							Usage:       "aws region",
							EnvVars:     []string{"AWS_DEFAULT_REGION"},
							Destination: &region,
						},
					},
					Action: func(c *cli.Context) error {
						if c.String("name") == "" || c.String("region") == "" || c.String("jobID") == "" {
							return cli.NewExitError("specify vault name, region and job ID using --name, --region and --jobID", 2)
						}
						aws.DescribeJob(region, c.String("name"), c.String("jobID"))
						return nil
					},
				},
				{
					Name:  "upload-archive",
					Usage: "upload data to a vault",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Usage: "vault name",
						},
						&cli.StringFlag{
							Name:        "region",
							Usage:       "aws region",
							EnvVars:     []string{"AWS_DEFAULT_REGION"},
							Destination: &region,
						},
						&cli.StringFlag{
							Name:  "file",
							Usage: "upload file name",
						},
					},
					Action: func(c *cli.Context) error {
						if c.String("name") == "" || c.String("file") == "" || region == "" {
							return cli.NewExitError("specify vault name, region and upload file using --name, --region and --file", 2)
						}
						aws.UploadArchive(region, c.String("name"), c.String("file"))
						return nil
					},
				},
				{
					Name:  "init-inventory-retrieval",
					Usage: "initiate an inventory-retrieval job for vault",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Usage: "vault name",
						},
						&cli.StringFlag{
							Name:  "desc",
							Value: "inventory-retrieval-job",
							Usage: "description for a job",
						},
						&cli.StringFlag{
							Name:        "region",
							Usage:       "aws region",
							EnvVars:     []string{"AWS_DEFAULT_REGION"},
							Destination: &region,
						},
					},
					Action: func(c *cli.Context) error {
						if c.String("name") == "" || region == "" {
							return cli.NewExitError("specify vault name and region using --name and --region", 2)
						}
						aws.InitInventoryRetrieval(region, c.String("name"), c.String("desc"))
						return nil
					},
				},
				{
					Name:  "init-archive-retrieval",
					Usage: "initiate an archive-retrieval job for vault",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Usage: "vault name",
						},
						&cli.StringFlag{
							Name:  "desc",
							Value: "archive-retrieval-job",
							Usage: "description for a job",
						},
						&cli.StringFlag{
							Name:  "jobID",
							Usage: "specify a job ID",
						},
						&cli.StringFlag{
							Name:        "region",
							Usage:       "aws region",
							EnvVars:     []string{"AWS_DEFAULT_REGION"},
							Destination: &region,
						},
					},
					Action: func(c *cli.Context) error {
						if c.String("name") == "" || region == "" || c.String("jobID") == "" {
							return cli.NewExitError("specify vault name, region and jobID using --name, --region and --jobID", 2)
						}
						aws.InitArchiveRetrieval(region, c.String("name"), c.String("desc"), c.String("jobID"))
						return nil
					},
				},
				{
					Name:  "get-inventory",
					Usage: "get output of inventory retrieval job",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Usage: "vault name",
						},
						&cli.StringFlag{
							Name:  "jobID",
							Usage: "specify a job ID",
						},
						&cli.StringFlag{
							Name:        "region",
							Usage:       "aws region",
							EnvVars:     []string{"AWS_DEFAULT_REGION"},
							Destination: &region,
						},
					},
					Action: func(c *cli.Context) error {
						if c.String("name") == "" || c.String("jobID") == "" || region == "" {
							return cli.NewExitError("specify vault name, region and jobID using --name, --region and --jobID", 2)
						}
						aws.GetVautlInventory(region, c.String("name"), c.String("jobID"))
						return nil
					},
				},
				{
					Name:  "get-archive",
					Usage: "get output of archive retrieval job",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Usage: "vault name",
						},
						&cli.StringFlag{
							Name:  "jobID",
							Usage: "specify a job ID",
						},
						&cli.StringFlag{
							Name:  "file",
							Usage: "specify file including path",
						},
						&cli.StringFlag{
							Name:        "region",
							Usage:       "aws region",
							EnvVars:     []string{"AWS_DEFAULT_REGION"},
							Destination: &region,
						},
					},
					Action: func(c *cli.Context) error {
						if c.String("name") == "" || c.String("jobID") == "" || region == "" || c.String("file") == "" {
							return cli.NewExitError("specify vault name, region, jobID and file using --name, --region, --jobID and --file", 2)
						}
						aws.GetVaultArchive(region, c.String("name"), c.String("jobID"), c.String("file"))
						return nil
					},
				},
				{
					Name:  "get-vaultlock",
					Usage: "get information about vault's policy",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Usage: "vault name",
						},
						&cli.StringFlag{
							Name:        "region",
							Usage:       "aws region",
							EnvVars:     []string{"AWS_DEFAULT_REGION"},
							Destination: &region,
						},
					},
					Action: func(c *cli.Context) error {
						if c.String("name") == "" || region == "" {
							return cli.NewExitError("specify vault name and region using --name and --region", 2)
						}
						aws.GetVaultLock(region, c.String("name"))
						return nil
					},
				},
				{
					Name:  "init-vaultlock",
					Usage: "init vault lock policy on the specified vault",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Usage: "vault name",
						},
						&cli.StringFlag{
							Name:  "policy",
							Usage: "vault policy",
						},
						&cli.StringFlag{
							Name:        "region",
							Usage:       "aws region",
							EnvVars:     []string{"AWS_DEFAULT_REGION"},
							Destination: &region,
						},
					},
					Action: func(c *cli.Context) error {
						if c.String("name") == "" || region == "" {
							return cli.NewExitError("specify vault name, region and policy using --name, --region and --policy", 2)
						}
						aws.InitiateVaultLock(region, c.String("name"), c.String("policy"))
						return nil
					},
				},
				{
					Name:  "abort-vaultlock",
					Usage: "aborts the vault locking process if not already in locked state",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Usage: "vault name",
						},
						&cli.StringFlag{
							Name:        "region",
							Usage:       "aws region",
							EnvVars:     []string{"AWS_DEFAULT_REGION"},
							Destination: &region,
						},
					},
					Action: func(c *cli.Context) error {
						if c.String("name") == "" || region == "" {
							return cli.NewExitError("specify vault name and region using --name and --region", 2)
						}
						aws.AbortVaultLock(region, c.String("name"))
						return nil
					},
				},
				{
					Name:  "complete-vaultlock",
					Usage: "complete vault lock in process.",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Usage: "vault name",
						},
						&cli.StringFlag{
							Name:        "region",
							Usage:       "aws region",
							EnvVars:     []string{"AWS_DEFAULT_REGION"},
							Destination: &region,
						},
						&cli.StringFlag{
							Name:  "lockID",
							Usage: "vault lockID",
						},
					},
					Action: func(c *cli.Context) error {
						if c.String("name") == "" || region == "" || c.String("lockID") == "" {
							return cli.NewExitError("specify vault name, region and lockID using --name, --region and --lockID", 2)
						}
						aws.CompleteVaultLock(region, c.String("name"), c.String("lockID"))
						return nil
					},
				},
				{
					Name:  "get-retrieval-policy",
					Usage: "get the current data retrieval policy",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:        "region",
							Usage:       "aws region",
							EnvVars:     []string{"AWS_DEFAULT_REGION"},
							Destination: &region,
						},
					},
					Action: func(c *cli.Context) error {
						if region == "" {
							return cli.NewExitError("specify region using --region", 2)
						}
						aws.GetRetrievalPolicy(region)
						return nil
					},
				},
				{
					Name:  "delete-archive",
					Usage: "delete archive from vault",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Usage: "vault name",
						},
						&cli.StringFlag{
							Name:  "archiveID",
							Usage: "specify a archive ID",
						},
						&cli.StringFlag{
							Name:        "region",
							Usage:       "aws region",
							EnvVars:     []string{"AWS_DEFAULT_REGION"},
							Destination: &region,
						},
					},
					Action: func(c *cli.Context) error {
						if c.String("name") == "" || c.String("archiveID") == "" || region == "" {
							return cli.NewExitError("specify vault name, region and archiveID using --name, --region and --archiveID", 2)
						}
						aws.DeleteArchive(region, c.String("name"), c.String("archiveID"))
						return nil
					},
				},
				{
					Name:  "delete-vault",
					Usage: "delete empty vault",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Usage: "vault name",
						},
						&cli.StringFlag{
							Name:        "region",
							Usage:       "aws region",
							EnvVars:     []string{"AWS_DEFAULT_REGION"},
							Destination: &region,
						},
					},
					Action: func(c *cli.Context) error {
						if c.String("name") == "" || region == "" {
							return cli.NewExitError("specify vault name and region using --name and --region", 2)
						}
						aws.DeleteVault(region, c.String("name"))
						return nil
					},
				},
			},
		}, // end glacier commands
		{
			Name:  "s3",
			Usage: "s3 operations",
			Subcommands: []*cli.Command{
				{
					Name:  "list-buckets",
					Usage: "list buckets in region",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:        "region",
							Usage:       "aws region",
							EnvVars:     []string{"AWS_DEFAULT_REGION"},
							Destination: &region,
						},
					},
					Action: func(c *cli.Context) error {
						if region == "" {
							return cli.NewExitError("specify region using --region", 2)
						}
						aws.ListBuckets(region)
						return nil
					},
				},
				{
					Name:  "create-bucket",
					Usage: "create a bucket",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Usage: "bucket name",
						},
						&cli.StringFlag{
							Name:        "region",
							Usage:       "aws region",
							EnvVars:     []string{"AWS_DEFAULT_REGION"},
							Destination: &region,
						},
					},
					Action: func(c *cli.Context) error {
						if c.String("name") == "" || region == "" {
							return cli.NewExitError("specify bucket name and region using --name and --region", 2)
						}
						aws.CreateBucket(region, c.String("name"))
						return nil
					},
				},
				{
					Name:  "upload-archive",
					Usage: "upload data to a bucket",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Usage: "bucket name",
						},
						&cli.StringFlag{
							Name:        "region",
							Usage:       "aws region",
							EnvVars:     []string{"AWS_DEFAULT_REGION"},
							Destination: &region,
						},
						&cli.StringFlag{
							Name:  "file",
							Usage: "upload file name",
						},
					},
					Action: func(c *cli.Context) error {
						if c.String("name") == "" || c.String("file") == "" || region == "" {
							return cli.NewExitError("specify bucket name, region and upload file using --name, --region and --file", 2)
						}
						aws.UploadBucket(region, c.String("name"), c.String("file"))
						return nil
					},
				},
				{
					Name:  "list-objects",
					Usage: "list objects in a bucket",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Usage: "bucket name",
						},
						&cli.StringFlag{
							Name:        "region",
							Usage:       "aws region",
							EnvVars:     []string{"AWS_DEFAULT_REGION"},
							Destination: &region,
						},
					},
					Action: func(c *cli.Context) error {
						if c.String("name") == "" || region == "" {
							return cli.NewExitError("specify bucket name and region using --name and --region", 2)
						}
						aws.ListObjects(region, c.String("name"))
						return nil
					},
				},
				{
					Name:  "delete-object",
					Usage: "delete object from bucket",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Usage: "bucket name",
						},
						&cli.StringFlag{
							Name:        "region",
							Usage:       "aws region",
							EnvVars:     []string{"AWS_DEFAULT_REGION"},
							Destination: &region,
						},
						&cli.StringFlag{
							Name:  "objectKey",
							Usage: "object key",
						},
					},
					Action: func(c *cli.Context) error {
						if c.String("name") == "" || c.String("objectKey") == "" || region == "" {
							return cli.NewExitError("specify bucket name, region and object key using --name, --region and --objectKey", 2)
						}
						aws.DeleteObjects(region, c.String("name"), c.String("objectKey"))
						return nil
					},
				},
				{
					Name:  "delete-bucket",
					Usage: "delete a bucket",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Usage: "bucket name",
						},
						&cli.StringFlag{
							Name:        "region",
							Usage:       "aws region",
							EnvVars:     []string{"AWS_DEFAULT_REGION"},
							Destination: &region,
						},
					},
					Action: func(c *cli.Context) error {
						if c.String("name") == "" || region == "" {
							return cli.NewExitError("specify bucket name and region using --name and --region", 2)
						}
						aws.DeleteBucket(region, c.String("name"))
						return nil
					},
				},
			}, // end of s3 operations
		}, // cli.Command
	} // app.Commands

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

} //end of main
