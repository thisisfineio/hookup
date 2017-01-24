package hookup

import (
	"github.com/aws/aws-sdk-go/service/directconnect"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/thisisfineio/common"
	"github.com/aws/aws-sdk-go/aws/session"
	"errors"
)



type DirectConnect interface {
	Provider() int
	DescribeDirectConnects(*DescribeDirectConnectsInput) (*DescribeDirectConnectsOutput, error)
}

type DescribeDirectConnectsInput struct {
	AwsInput *directconnect.DescribeConnectionsInput
}

type DescribeDirectConnectsOutput struct {

}

func NewDirectConnect(provider int, region string) (DirectConnect, error) {
	switch provider {
	case common.AwsProvider:
		return DirectConnect(NewAwsDirectConnect(region)), nil
	}
	return nil, common.InvalidProviderErr("NewDirectConnect", "hookup")
}

func NewAwsDirectConnectWithConfig(c *aws.Config) *AwsDirectConnect {
	return &AwsDirectConnect{directconnect.New(session.New(c))}
}

type AwsDirectConnect struct {
	service *directconnect.DirectConnect
}

func NewAwsDirectConnect(region string) *AwsDirectConnect {
	return NewAwsDirectConnectWithConfig(&aws.Config{Region: aws.String(region)})
}

func (a *AwsDirectConnect) DescribeDirectConnects(input *DescribeDirectConnectsInput) (*DescribeDirectConnectsOutput, error) {
	return nil, nil
}

func (a *AwsDirectConnect) Provider() int {
	return common.AwsProvider
}