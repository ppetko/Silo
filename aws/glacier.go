package aws

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glacier"
)

// VaultInventory - Vault inventory struct used for unmarshaling data
type VaultInventory struct {
	VaultARN      string    `json:"VaultARN"`
	InventoryDate time.Time `json:"InventoryDate"`
	ArchiveList   []struct {
		ArchiveID          string    `json:"ArchiveId"`
		ArchiveDescription string    `json:"ArchiveDescription"`
		CreationDate       time.Time `json:"CreationDate"`
		Size               int       `json:"Size"`
		SHA256TreeHash     string    `json:"SHA256TreeHash"`
	} `json:"ArchiveList"`
}

// GetVaultLock - Retrieve vault lock-policy related attributes that are set on a vault
func GetVaultLock(awsRegion, vaultName string) {
	svc := glacier.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &glacier.GetVaultLockInput{
		AccountId: aws.String("-"),
		VaultName: aws.String(vaultName),
	}
	result, err := svc.GetVaultLock(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glacier.ErrCodeResourceNotFoundException:
				fmt.Println(glacier.ErrCodeResourceNotFoundException, aerr.Error())
			case glacier.ErrCodeInvalidParameterValueException:
				fmt.Println(glacier.ErrCodeInvalidParameterValueException, aerr.Error())
			case glacier.ErrCodeMissingParameterValueException:
				fmt.Println(glacier.ErrCodeMissingParameterValueException, aerr.Error())
			case glacier.ErrCodeServiceUnavailableException:
				fmt.Println(glacier.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}

// GetVaultAccessPolicy - Get the access-policy set on the vault
func GetVaultAccessPolicy(awsRegion, vaultName string) {
	svc := glacier.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &glacier.GetVaultAccessPolicyInput{
		AccountId: aws.String("-"),
		VaultName: aws.String(vaultName),
	}
	result, err := svc.GetVaultAccessPolicy(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glacier.ErrCodeResourceNotFoundException:
				fmt.Println(glacier.ErrCodeResourceNotFoundException, aerr.Error())
			case glacier.ErrCodeInvalidParameterValueException:
				fmt.Println(glacier.ErrCodeInvalidParameterValueException, aerr.Error())
			case glacier.ErrCodeMissingParameterValueException:
				fmt.Println(glacier.ErrCodeMissingParameterValueException, aerr.Error())
			case glacier.ErrCodeServiceUnavailableException:
				fmt.Println(glacier.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}

// DeleteArchive - Delete archive
func DeleteArchive(awsRegion, vaultName, ArchiveID string) {
	svc := glacier.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &glacier.DeleteArchiveInput{
		AccountId: aws.String("-"),
		ArchiveId: aws.String(ArchiveID),
		VaultName: aws.String(vaultName),
	}
	result, err := svc.DeleteArchive(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glacier.ErrCodeResourceNotFoundException:
				fmt.Println(glacier.ErrCodeResourceNotFoundException, aerr.Error())
			case glacier.ErrCodeInvalidParameterValueException:
				fmt.Println(glacier.ErrCodeInvalidParameterValueException, aerr.Error())
			case glacier.ErrCodeMissingParameterValueException:
				fmt.Println(glacier.ErrCodeMissingParameterValueException, aerr.Error())
			case glacier.ErrCodeServiceUnavailableException:
				fmt.Println(glacier.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}

// InitInventoryRetrieval - Initiate an inventory-retrieval job based on vault name
func InitInventoryRetrieval(awsRegion, vaultName, jobDescription string) {
	svc := glacier.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &glacier.InitiateJobInput{
		AccountId: aws.String("-"),
		JobParameters: &glacier.JobParameters{
			Description: aws.String(jobDescription),
			Type:        aws.String("inventory-retrieval"),
		},
		VaultName: aws.String(vaultName),
	}
	result, err := svc.InitiateJob(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glacier.ErrCodeResourceNotFoundException:
				fmt.Println(glacier.ErrCodeResourceNotFoundException, aerr.Error())
			case glacier.ErrCodePolicyEnforcedException:
				fmt.Println(glacier.ErrCodePolicyEnforcedException, aerr.Error())
			case glacier.ErrCodeInvalidParameterValueException:
				fmt.Println(glacier.ErrCodeInvalidParameterValueException, aerr.Error())
			case glacier.ErrCodeMissingParameterValueException:
				fmt.Println(glacier.ErrCodeMissingParameterValueException, aerr.Error())
			case glacier.ErrCodeInsufficientCapacityException:
				fmt.Println(glacier.ErrCodeInsufficientCapacityException, aerr.Error())
			case glacier.ErrCodeServiceUnavailableException:
				fmt.Println(glacier.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}

// InitArchiveRetrieval - Initiate an archive-retrieval job based on vault name
func InitArchiveRetrieval(awsRegion, vaultName, jobDescription, archiveID string) {
	svc := glacier.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &glacier.InitiateJobInput{
		AccountId: aws.String("-"),
		JobParameters: &glacier.JobParameters{
			ArchiveId:   aws.String(archiveID),
			Description: aws.String(jobDescription),
			Type:        aws.String("archive-retrieval"),
		},
		VaultName: aws.String(vaultName),
	}
	result, err := svc.InitiateJob(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glacier.ErrCodeResourceNotFoundException:
				fmt.Println(glacier.ErrCodeResourceNotFoundException, aerr.Error())
			case glacier.ErrCodePolicyEnforcedException:
				fmt.Println(glacier.ErrCodePolicyEnforcedException, aerr.Error())
			case glacier.ErrCodeInvalidParameterValueException:
				fmt.Println(glacier.ErrCodeInvalidParameterValueException, aerr.Error())
			case glacier.ErrCodeMissingParameterValueException:
				fmt.Println(glacier.ErrCodeMissingParameterValueException, aerr.Error())
			case glacier.ErrCodeInsufficientCapacityException:
				fmt.Println(glacier.ErrCodeInsufficientCapacityException, aerr.Error())
			case glacier.ErrCodeServiceUnavailableException:
				fmt.Println(glacier.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}

// InitJobInput - Initiate an inventory-retrieval job based on vault name
// Resource - https://docs.aws.amazon.com/sdk-for-go/api/service/glacier/#example_Glacier_InitiateJob_shared00
// This operation initiates a job of the specified type, which can be a select, an archival retrieval, or a vault retrieval.
func InitJobInput(awsRegion, vaultName string) {
	svc := glacier.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &glacier.InitiateJobInput{
		AccountId: aws.String("-"),
		JobParameters: &glacier.JobParameters{
			Description: aws.String("My inventory job"),
			//Format:      aws.String("CSV"),
			//SNSTopic:    aws.String("arn:aws:sns:us-west-2:111111111111:Glacier-InventoryRetrieval-topic-Example"),
			//  archive-retrieval - Retrieve an archive
			// inventory-retrieval - Inventory a vault
			Type: aws.String("archive-retrieval"),
		},
		VaultName: aws.String(vaultName),
	}
	result, err := svc.InitiateJob(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glacier.ErrCodeResourceNotFoundException:
				fmt.Println(glacier.ErrCodeResourceNotFoundException, aerr.Error())
			case glacier.ErrCodePolicyEnforcedException:
				fmt.Println(glacier.ErrCodePolicyEnforcedException, aerr.Error())
			case glacier.ErrCodeInvalidParameterValueException:
				fmt.Println(glacier.ErrCodeInvalidParameterValueException, aerr.Error())
			case glacier.ErrCodeMissingParameterValueException:
				fmt.Println(glacier.ErrCodeMissingParameterValueException, aerr.Error())
			case glacier.ErrCodeInsufficientCapacityException:
				fmt.Println(glacier.ErrCodeInsufficientCapacityException, aerr.Error())
			case glacier.ErrCodeServiceUnavailableException:
				fmt.Println(glacier.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}

// GetVautlInventory - Get the output of a previously initiated job for inventory retrieval that is identified by the job ID
func GetVautlInventory(awsRegion, vaultName, jobID string) {
	svc := glacier.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &glacier.GetJobOutputInput{
		AccountId: aws.String("-"),
		JobId:     aws.String(jobID),
		Range:     aws.String(""),
		VaultName: aws.String(vaultName),
	}
	result, err := svc.GetJobOutput(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glacier.ErrCodeResourceNotFoundException:
				fmt.Println(glacier.ErrCodeResourceNotFoundException, aerr.Error())
			case glacier.ErrCodeInvalidParameterValueException:
				fmt.Println(glacier.ErrCodeInvalidParameterValueException, aerr.Error())
			case glacier.ErrCodeMissingParameterValueException:
				fmt.Println(glacier.ErrCodeMissingParameterValueException, aerr.Error())
			case glacier.ErrCodeServiceUnavailableException:
				fmt.Println(glacier.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	//fmt.Println(result)

	body, err := ioutil.ReadAll(result.Body)

	var inventory VaultInventory
	err1 := json.Unmarshal(body, &inventory)
	if err1 != nil {
		fmt.Println("error:", err)
	}
	t := inventory.ArchiveList

	fmt.Printf("VaultARN: %s\n", inventory.VaultARN)
	fmt.Printf("InventoryDate: %s\n", inventory.InventoryDate)
	for _, j := range t {
		fmt.Printf("ArchiveID: %s\n", j.ArchiveID)
		fmt.Printf("ArchiveDescription %s\n", j.ArchiveDescription)
		fmt.Printf("CreationDate %s\n", j.CreationDate)
		fmt.Printf("Size: %v\n", j.Size)
		fmt.Printf("SHA256TreeHash: %v\n", j.SHA256TreeHash)

	}
}

/* TODO: In the case of an archive retrieval job, depending on the byte range you specify,
Amazon Glacier returns the checksum for the portion of the data.
To ensure the portion you downloaded is the correct data, compute the checksum on the client,
verify that the values match, and verify that the size is what you expected.
*/

// GetVaultArchive - Get the output of a previously initiated job for archive retrieval
// GetJobOutput -  Get the output of a previously initiated job, for instance inventory retrieval job that is identified by the job ID
// https://docs.aws.amazon.com/amazonglacier/latest/dev/api-job-output-get.html
func GetVaultArchive(awsRegion, vaultName, jobID, fileName string) {
	svc := glacier.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &glacier.GetJobOutputInput{
		AccountId: aws.String("-"),
		JobId:     aws.String(jobID),
		Range:     aws.String(""),
		VaultName: aws.String(vaultName),
	}
	result, err := svc.GetJobOutput(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glacier.ErrCodeResourceNotFoundException:
				fmt.Println(glacier.ErrCodeResourceNotFoundException, aerr.Error())
			case glacier.ErrCodeInvalidParameterValueException:
				fmt.Println(glacier.ErrCodeInvalidParameterValueException, aerr.Error())
			case glacier.ErrCodeMissingParameterValueException:
				fmt.Println(glacier.ErrCodeMissingParameterValueException, aerr.Error())
			case glacier.ErrCodeServiceUnavailableException:
				fmt.Println(glacier.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)

	outFile, err1 := os.Create(fileName)
	if err1 != nil {
		panic(err)
	}
	defer outFile.Close()
	_, err1 = io.Copy(outFile, result.Body)

}

// ListJobs - List all pending jobs per vault
func ListJobs(awsRegion, vaultName string) {
	svc := glacier.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &glacier.ListJobsInput{
		AccountId: aws.String("-"),
		VaultName: aws.String(vaultName),
	}
	result, err := svc.ListJobs(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glacier.ErrCodeResourceNotFoundException:
				fmt.Println(glacier.ErrCodeResourceNotFoundException, aerr.Error())
			case glacier.ErrCodeInvalidParameterValueException:
				fmt.Println(glacier.ErrCodeInvalidParameterValueException, aerr.Error())
			case glacier.ErrCodeMissingParameterValueException:
				fmt.Println(glacier.ErrCodeMissingParameterValueException, aerr.Error())
			case glacier.ErrCodeServiceUnavailableException:
				fmt.Println(glacier.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}

// DescribeJob - Get information about a previously initiated job, specified by the job ID.
func DescribeJob(awsRegion, vaultName, jobID string) {
	svc := glacier.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &glacier.DescribeJobInput{
		AccountId: aws.String("-"),
		JobId:     aws.String(jobID),
		VaultName: aws.String(vaultName),
	}
	result, err := svc.DescribeJob(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glacier.ErrCodeResourceNotFoundException:
				fmt.Println(glacier.ErrCodeResourceNotFoundException, aerr.Error())
			case glacier.ErrCodeInvalidParameterValueException:
				fmt.Println(glacier.ErrCodeInvalidParameterValueException, aerr.Error())
			case glacier.ErrCodeMissingParameterValueException:
				fmt.Println(glacier.ErrCodeMissingParameterValueException, aerr.Error())
			case glacier.ErrCodeServiceUnavailableException:
				fmt.Println(glacier.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}

// DescriveVault - Retrieve information about a vault
func DescriveVault(awsRegion, vaultName string) {
	svc := glacier.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &glacier.DescribeVaultInput{
		AccountId: aws.String("-"),
		VaultName: aws.String(vaultName),
	}
	result, err := svc.DescribeVault(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glacier.ErrCodeResourceNotFoundException:
				fmt.Println(glacier.ErrCodeResourceNotFoundException, aerr.Error())
			case glacier.ErrCodeInvalidParameterValueException:
				fmt.Println(glacier.ErrCodeInvalidParameterValueException, aerr.Error())
			case glacier.ErrCodeMissingParameterValueException:
				fmt.Println(glacier.ErrCodeMissingParameterValueException, aerr.Error())
			case glacier.ErrCodeServiceUnavailableException:
				fmt.Println(glacier.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}

// UploadArchive - Upload archive to vault
// Reference - https://docs.aws.amazon.com/amazonglacier/latest/dev/api-archive-post.html
// More - https://docs.aws.amazon.com/amazonglacier/latest/dev/uploading-an-archive.html
func UploadArchive(awsRegion, vaultName, fileUpload string) {
	svc := glacier.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &glacier.UploadArchiveInput{
		AccountId:          aws.String("-"),
		ArchiveDescription: aws.String(getFilename(fileUpload)),
		Body:               aws.ReadSeekCloser(strings.NewReader(isFile(fileUpload))),
		Checksum:           aws.String(""),
		VaultName:          aws.String(vaultName),
	}
	result, err := svc.UploadArchive(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glacier.ErrCodeResourceNotFoundException:
				fmt.Println(glacier.ErrCodeResourceNotFoundException, aerr.Error())
			case glacier.ErrCodeInvalidParameterValueException:
				fmt.Println(glacier.ErrCodeInvalidParameterValueException, aerr.Error())
			case glacier.ErrCodeMissingParameterValueException:
				fmt.Println(glacier.ErrCodeMissingParameterValueException, aerr.Error())
			case glacier.ErrCodeRequestTimeoutException:
				fmt.Println(glacier.ErrCodeRequestTimeoutException, aerr.Error())
			case glacier.ErrCodeServiceUnavailableException:
				fmt.Println(glacier.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}

// DeleteVault - Delete vault based on name and region
func DeleteVault(awsRegion, vaultName string) {
	svc := glacier.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &glacier.DeleteVaultInput{
		AccountId: aws.String("-"),
		VaultName: aws.String(vaultName),
	}
	result, err := svc.DeleteVault(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glacier.ErrCodeResourceNotFoundException:
				fmt.Println(glacier.ErrCodeResourceNotFoundException, aerr.Error())
			case glacier.ErrCodeInvalidParameterValueException:
				fmt.Println(glacier.ErrCodeInvalidParameterValueException, aerr.Error())
			case glacier.ErrCodeMissingParameterValueException:
				fmt.Println(glacier.ErrCodeMissingParameterValueException, aerr.Error())
			case glacier.ErrCodeServiceUnavailableException:
				fmt.Println(glacier.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}

// CreateVault - Create new vault based on name and region
func CreateVault(awsRegion, vaultName string) {
	svc := glacier.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &glacier.CreateVaultInput{
		AccountId: aws.String("-"),
		VaultName: aws.String(vaultName),
		//Limit:     aws.String(""),
		//Marker:    aws.String(""),
	}
	result, err := svc.CreateVault(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glacier.ErrCodeInvalidParameterValueException:
				fmt.Println(glacier.ErrCodeInvalidParameterValueException, aerr.Error())
			case glacier.ErrCodeMissingParameterValueException:
				fmt.Println(glacier.ErrCodeMissingParameterValueException, aerr.Error())
			case glacier.ErrCodeServiceUnavailableException:
				fmt.Println(glacier.ErrCodeServiceUnavailableException, aerr.Error())
			case glacier.ErrCodeLimitExceededException:
				fmt.Println(glacier.ErrCodeLimitExceededException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}

// ListVault - List all vaults based on region
func ListVault(awsRegion string) {
	svc := glacier.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &glacier.ListVaultsInput{
		AccountId: aws.String("-"),
		//Limit:     aws.String(""),
		//Marker:    aws.String(""),
	}
	result, err := svc.ListVaults(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glacier.ErrCodeResourceNotFoundException:
				fmt.Println(glacier.ErrCodeResourceNotFoundException, aerr.Error())
			case glacier.ErrCodeInvalidParameterValueException:
				fmt.Println(glacier.ErrCodeInvalidParameterValueException, aerr.Error())
			case glacier.ErrCodeMissingParameterValueException:
				fmt.Println(glacier.ErrCodeMissingParameterValueException, aerr.Error())
			case glacier.ErrCodeServiceUnavailableException:
				fmt.Println(glacier.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}

// GetRetrievalPolicy - Get the current data retrieval policy for an account
func GetRetrievalPolicy(awsRegion string) {
	svc := glacier.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &glacier.GetDataRetrievalPolicyInput{
		AccountId: aws.String("-"),
	}
	result, err := svc.GetDataRetrievalPolicy(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glacier.ErrCodeInvalidParameterValueException:
				fmt.Println(glacier.ErrCodeInvalidParameterValueException, aerr.Error())
			case glacier.ErrCodeMissingParameterValueException:
				fmt.Println(glacier.ErrCodeMissingParameterValueException, aerr.Error())
			case glacier.ErrCodeServiceUnavailableException:
				fmt.Println(glacier.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}

// TODO SetDataRetrievalPolicyFreeTier and SetDataRetrievalPolicy needs testing

// SetDataRetrievalPolicyFreeTier - Set FreeTier retrieval policy
func SetDataRetrievalPolicyFreeTier(awsRegion string) {
	svc := glacier.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &glacier.SetDataRetrievalPolicyInput{
		Policy: &glacier.DataRetrievalPolicy{
			Rules: []*glacier.DataRetrievalRule{
				{
					Strategy: aws.String("FreeTier"),
				},
			},
		},
		AccountId: aws.String("-"),
	}
	result, err := svc.SetDataRetrievalPolicy(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glacier.ErrCodeInvalidParameterValueException:
				fmt.Println(glacier.ErrCodeInvalidParameterValueException, aerr.Error())
			case glacier.ErrCodeMissingParameterValueException:
				fmt.Println(glacier.ErrCodeMissingParameterValueException, aerr.Error())
			case glacier.ErrCodeServiceUnavailableException:
				fmt.Println(glacier.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}

// SetDataRetrievalPolicy - Set and then enact a data retrieval policy
func SetDataRetrievalPolicy(awsRegion, strategyPolicy string, bytesPerHour int64) {
	svc := glacier.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &glacier.SetDataRetrievalPolicyInput{
		Policy: &glacier.DataRetrievalPolicy{
			Rules: []*glacier.DataRetrievalRule{
				{
					BytesPerHour: aws.Int64(bytesPerHour),
					Strategy:     aws.String(strategyPolicy),
				},
			},
		},
		AccountId: aws.String("-"),
	}
	result, err := svc.SetDataRetrievalPolicy(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glacier.ErrCodeInvalidParameterValueException:
				fmt.Println(glacier.ErrCodeInvalidParameterValueException, aerr.Error())
			case glacier.ErrCodeMissingParameterValueException:
				fmt.Println(glacier.ErrCodeMissingParameterValueException, aerr.Error())
			case glacier.ErrCodeServiceUnavailableException:
				fmt.Println(glacier.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// InitiateVaultLock - Installing a vault lock policy on the specified vault.
// Setting the lock state of vault lock to InProgress.
// Returning a lock ID, which is used to complete the vault locking process.
func InitiateVaultLock(awsRegion, vaultName, vaultPolicy string) {
	svc := glacier.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &glacier.InitiateVaultLockInput{
		AccountId: aws.String("-"),
		Policy: &glacier.VaultLockPolicy{
			Policy: aws.String(vaultPolicy),
			//Policy: aws.String("{\"Version\":\"2012-10-17\",\"Statement\":[{\"Sid\":\"Define-vault-lock\",\"Effect\":\"Deny\",\"Principal\":{\"AWS\":\"*\"},\"Action\":\"glacier:DeleteArchive\",\"Resource\":\"arn:aws:glacier:us-east-2:757758175257:vaults/my-vault\",\"Condition\":{\"NumericLessThanEquals\":{\"glacier:ArchiveAgeinDays\":\"365\"}}}]}"),
		},
		VaultName: aws.String(vaultName),
	}
	result, err := svc.InitiateVaultLock(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glacier.ErrCodeResourceNotFoundException:
				fmt.Println(glacier.ErrCodeResourceNotFoundException, aerr.Error())
			case glacier.ErrCodeInvalidParameterValueException:
				fmt.Println(glacier.ErrCodeInvalidParameterValueException, aerr.Error())
			case glacier.ErrCodeMissingParameterValueException:
				fmt.Println(glacier.ErrCodeMissingParameterValueException, aerr.Error())
			case glacier.ErrCodeServiceUnavailableException:
				fmt.Println(glacier.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}

// AbortVaultLock - This operation aborts the vault locking process if the vault lock is not in the Locked state.
// If the vault lock is in the Locked state when this operation is requested, the operation returns an AccessDeniedException error.
// Aborting the vault locking process removes the vault lock policy from the specified vault.
func AbortVaultLock(awsRegion, vaultName string) {
	svc := glacier.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &glacier.AbortVaultLockInput{
		AccountId: aws.String("-"),
		VaultName: aws.String(vaultName),
	}
	_, err := svc.AbortVaultLock(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glacier.ErrCodeResourceNotFoundException:
				fmt.Println(glacier.ErrCodeResourceNotFoundException, aerr.Error())
			case glacier.ErrCodeInvalidParameterValueException:
				fmt.Println(glacier.ErrCodeInvalidParameterValueException, aerr.Error())
			case glacier.ErrCodeMissingParameterValueException:
				fmt.Println(glacier.ErrCodeMissingParameterValueException, aerr.Error())
			case glacier.ErrCodeServiceUnavailableException:
				fmt.Println(glacier.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Printf("Vault lock aborted on %s \n", vaultName)
}

// CompleteVaultLock - This operation completes the vault locking process by transitioning the vault lock
// from the InProgress state to the Locked state, which causes the vault lock policy to become unchangeable.
func CompleteVaultLock(awsRegion, vaultName, lockID string) {
	svc := glacier.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &glacier.CompleteVaultLockInput{
		AccountId: aws.String("-"),
		LockId:    aws.String(lockID),
		VaultName: aws.String(vaultName),
	}
	result, err := svc.CompleteVaultLock(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glacier.ErrCodeResourceNotFoundException:
				fmt.Println(glacier.ErrCodeResourceNotFoundException, aerr.Error())
			case glacier.ErrCodeInvalidParameterValueException:
				fmt.Println(glacier.ErrCodeInvalidParameterValueException, aerr.Error())
			case glacier.ErrCodeMissingParameterValueException:
				fmt.Println(glacier.ErrCodeMissingParameterValueException, aerr.Error())
			case glacier.ErrCodeServiceUnavailableException:
				fmt.Println(glacier.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}
