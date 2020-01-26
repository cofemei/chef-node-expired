package main

import (
	"io/ioutil"
	"encoding/base64"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

// getpem
func getpem(profile string, region string ,fname string) (string, error) {
	ef, err := ioutil.ReadFile(fname)
	if err != nil {
		return "", err
	}
	var sess *session.Session
	if profile != "" {
		sess = session.Must(session.NewSessionWithOptions(session.Options{
			Config: aws.Config{Region: aws.String(region)},
			Profile: "prd",
	   }))
	} else {
		sess = session.Must(session.NewSessionWithOptions(session.Options{
			Config: aws.Config{Region: aws.String(region)},
	   }))
	}
	kmssvc := kms.New(sess)
	base64Blob, err := base64.StdEncoding.DecodeString(string(ef))
	if err != nil {
		return "", err
	}
	decryptOutput, err := kmssvc.Decrypt(&kms.DecryptInput{CiphertextBlob: []byte(base64Blob)})
	if err != nil {
		return "", err
	}
	return string(decryptOutput.Plaintext), nil
}

