package aws

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var (
	// AWS Vars
	awsConfPath       = "/.aws/"
	awsConfigFile     = "config"
	awsCredentialFile = "credentials"
)

// isError - Checks for error
func isError(err error) error {
	if err != nil {
		panic(err)
	}
	return err
}

// getFilename - Get file name from path
func getFilename(filePath string) string {
	return filepath.Base(filePath)
}

// isFile - Check if file exist
func isFile(filePath string) string {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("file %s does not exist", filePath)
		os.Exit(1)
	}
	return filePath
}

// UserHomeDir - Determine the user home path depends on OS Type
func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

// userInputConfigs -  Request the user input for aws credentionas and configurations
func userInputConfigs() (string, string, string, string) {
	fmt.Print("AWS Access Key ID: ")
	var accessKey string
	fmt.Scanln(&accessKey)

	fmt.Print("AWS Secret Access: ")
	var secretKey string
	fmt.Scanln(&secretKey)

	fmt.Print("Default region name: ")
	var region string
	fmt.Scanln(&region)

	fmt.Print("Default output format:[json] ")
	var output string
	fmt.Scanln(&output)

	return accessKey, secretKey, region, output
}

// SetupAWSAuth - Create ~/.aws folder containing the conf and credentials
func SetupAWSAuth() {
	// TODO: The generations of the files doesn't work under Windows

	// calling userInputConfigs for aws credenctioans and configurations
	accessKey, secretKey, region, output := userInputConfigs()

	// Set default the output json
	output = "json"

	userHomePath := UserHomeDir()
	err := os.RemoveAll(userHomePath + awsConfPath)
	isError(err)

	if _, err = os.Stat(userHomePath + awsConfPath); os.IsNotExist(err) {
		err := os.Mkdir(userHomePath+awsConfPath, 0775)
		isError(err)
	}

	// Generate ~/.aws/config 
	f, err := os.OpenFile(userHomePath+awsConfPath+awsConfigFile, os.O_RDWR|os.O_CREATE, 0600)
	isError(err)

	f.WriteString("[default]\noutput = " + output + "\n" + "region = " + region + "\n")
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	log.Println("config file generated ~/.aws/config")

	// Generate  ~/.aws/credentials 
	f1, err := os.OpenFile(userHomePath+awsConfPath+awsCredentialFile, os.O_RDWR|os.O_CREATE, 0600)
	isError(err)
	f1.WriteString("[default]\naws_access_key_id = " + accessKey + "\n" + "aws_secret_access_key = " + secretKey + "\n")
	if err := f1.Close(); err != nil {
		log.Fatal(err)
	}
	log.Println("credential file generated ~/.aws/credentials")
}
