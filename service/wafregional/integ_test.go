// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// +build go1.10,integration

package wafregional_test

import (
	"context"
	"testing"
	"time"

	"github.com/journeymidnight/aws-sdk-go/aws"
	"github.com/journeymidnight/aws-sdk-go/aws/awserr"
	"github.com/journeymidnight/aws-sdk-go/aws/request"
	"github.com/journeymidnight/aws-sdk-go/awstesting/integration"
	"github.com/journeymidnight/aws-sdk-go/service/waf"
	"github.com/journeymidnight/aws-sdk-go/service/wafregional"
)

var _ aws.Config
var _ awserr.Error
var _ request.Request

func TestInteg_00_ListRules(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-east-1")
	svc := wafregional.New(sess)
	params := &waf.ListRulesInput{
		Limit: aws.Int64(20),
	}
	_, err := svc.ListRulesWithContext(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
func TestInteg_01_CreateSqlInjectionMatchSet(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-east-1")
	svc := wafregional.New(sess)
	params := &waf.CreateSqlInjectionMatchSetInput{
		ChangeToken: aws.String("fake_token"),
		Name:        aws.String("fake_name"),
	}
	_, err := svc.CreateSqlInjectionMatchSetWithContext(ctx, params)
	if err == nil {
		t.Fatalf("expect request to fail")
	}
	aerr, ok := err.(awserr.RequestFailure)
	if !ok {
		t.Fatalf("expect awserr, was %T", err)
	}
	if len(aerr.Code()) == 0 {
		t.Errorf("expect non-empty error code")
	}
	if v := aerr.Code(); v == request.ErrCodeSerialization {
		t.Errorf("expect API error code got serialization failure")
	}
}
