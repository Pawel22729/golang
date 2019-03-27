package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var regoin = "eu-west-2"

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func createSess() *ssm.SSM {
	sess, err := session.NewSession()
	checkErr(err)
	svc := ssm.New(sess, aws.NewConfig().WithRegion(regoin))
	return svc
}

func getParam(path string) *ssm.GetParameterOutput {
	svc := createSess()
	desc := true
	param, err := svc.GetParameter(&ssm.GetParameterInput{
		Name:           &path,
		WithDecryption: &desc,
	})
	checkErr(err)
	return param
}

func getParams(path string) *ssm.GetParametersByPathOutput {
	svc := createSess()
	recur := true
	desc := true
	params, err := svc.GetParametersByPath(&ssm.GetParametersByPathInput{
		Path:           &path,
		Recursive:      &recur,
		WithDecryption: &desc,
	})
	checkErr(err)
	return params
}

func putParam(path, value string) *ssm.PutParameterOutput {
	svc := createSess()
	owrite := true
	ptype := "String"
	param, err := svc.PutParameter(&ssm.PutParameterInput{
		Name:      &path,
		Type:      &ptype,
		Overwrite: &owrite,
		Value:     &value,
	})
	checkErr(err)
	return param
}

func delParam(path string) *ssm.DeleteParameterOutput {
	svc := createSess()
	param, err := svc.DeleteParameter(&ssm.DeleteParameterInput{
		Name: &path,
	})
	checkErr(err)
	return param
}
