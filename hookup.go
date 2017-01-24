package hookup

type DirectConnect interface {
	DescribeDirectConnects(*DescribeDirectConnectsInput) (*DescribeDirectConnectsOutput, error)
}

type DescribeDirectConnectsInput struct {

}

type DescribeDirectConnectsOutput struct {

}


type AwsDirectConnect struct {

}