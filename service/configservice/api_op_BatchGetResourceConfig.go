// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/BatchGetResourceConfigRequest
type BatchGetResourceConfigInput struct {
	_ struct{} `type:"structure"`

	// A list of resource keys to be processed with the current request. Each element
	// in the list consists of the resource type and resource ID.
	//
	// ResourceKeys is a required field
	ResourceKeys []ResourceKey `locationName:"resourceKeys" min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchGetResourceConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchGetResourceConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchGetResourceConfigInput"}

	if s.ResourceKeys == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceKeys"))
	}
	if s.ResourceKeys != nil && len(s.ResourceKeys) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceKeys", 1))
	}
	if s.ResourceKeys != nil {
		for i, v := range s.ResourceKeys {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ResourceKeys", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/BatchGetResourceConfigResponse
type BatchGetResourceConfigOutput struct {
	_ struct{} `type:"structure"`

	// A list that contains the current configuration of one or more resources.
	BaseConfigurationItems []BaseConfigurationItem `locationName:"baseConfigurationItems" type:"list"`

	// A list of resource keys that were not processed with the current response.
	// The unprocessesResourceKeys value is in the same form as ResourceKeys, so
	// the value can be directly provided to a subsequent BatchGetResourceConfig
	// operation. If there are no unprocessed resource keys, the response contains
	// an empty unprocessedResourceKeys list.
	UnprocessedResourceKeys []ResourceKey `locationName:"unprocessedResourceKeys" min:"1" type:"list"`
}

// String returns the string representation
func (s BatchGetResourceConfigOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchGetResourceConfig = "BatchGetResourceConfig"

// BatchGetResourceConfigRequest returns a request value for making API operation for
// AWS Config.
//
// Returns the current configuration for one or more requested resources. The
// operation also returns a list of resources that are not processed in the
// current request. If there are no unprocessed resources, the operation returns
// an empty unprocessedResourceKeys list.
//
//    * The API does not return results for deleted resources.
//
//    * The API does not return any tags for the requested resources. This information
//    is filtered out of the supplementaryConfiguration section of the API response.
//
//    // Example sending a request using BatchGetResourceConfigRequest.
//    req := client.BatchGetResourceConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/BatchGetResourceConfig
func (c *Client) BatchGetResourceConfigRequest(input *BatchGetResourceConfigInput) BatchGetResourceConfigRequest {
	op := &aws.Operation{
		Name:       opBatchGetResourceConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchGetResourceConfigInput{}
	}

	req := c.newRequest(op, input, &BatchGetResourceConfigOutput{})
	return BatchGetResourceConfigRequest{Request: req, Input: input, Copy: c.BatchGetResourceConfigRequest}
}

// BatchGetResourceConfigRequest is the request type for the
// BatchGetResourceConfig API operation.
type BatchGetResourceConfigRequest struct {
	*aws.Request
	Input *BatchGetResourceConfigInput
	Copy  func(*BatchGetResourceConfigInput) BatchGetResourceConfigRequest
}

// Send marshals and sends the BatchGetResourceConfig API request.
func (r BatchGetResourceConfigRequest) Send(ctx context.Context) (*BatchGetResourceConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchGetResourceConfigResponse{
		BatchGetResourceConfigOutput: r.Request.Data.(*BatchGetResourceConfigOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchGetResourceConfigResponse is the response type for the
// BatchGetResourceConfig API operation.
type BatchGetResourceConfigResponse struct {
	*BatchGetResourceConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchGetResourceConfig request.
func (r *BatchGetResourceConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
