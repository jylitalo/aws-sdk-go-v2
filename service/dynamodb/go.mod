module github.com/aws/aws-sdk-go-v2/service/dynamodb

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v0.28.1-0.20201027184009-8eb8fc303e7c
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v0.3.0
	github.com/awslabs/smithy-go v0.3.0
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding => ../../service/internal/accept-encoding/
