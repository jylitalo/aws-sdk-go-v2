// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Requests API Gateway to create a Resource resource.
type CreateResourceInput struct {
	_ struct{} `type:"structure"`

	// [Required] The parent resource's identifier.
	//
	// ParentId is a required field
	ParentId *string `location:"uri" locationName:"parent_id" type:"string" required:"true"`

	// The last path segment for this resource.
	//
	// PathPart is a required field
	PathPart *string `locationName:"pathPart" type:"string" required:"true"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateResourceInput"}

	if s.ParentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParentId"))
	}

	if s.PathPart == nil {
		invalidParams.Add(aws.NewErrParamRequired("PathPart"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateResourceInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.PathPart != nil {
		v := *s.PathPart

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "pathPart", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ParentId != nil {
		v := *s.ParentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "parent_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Represents an API resource.
//
// Create an API (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-create-api.html)
type CreateResourceOutput struct {
	_ struct{} `type:"structure"`

	// The resource's identifier.
	Id *string `locationName:"id" type:"string"`

	// The parent resource's identifier.
	ParentId *string `locationName:"parentId" type:"string"`

	// The full path for this resource.
	Path *string `locationName:"path" type:"string"`

	// The last path segment for this resource.
	PathPart *string `locationName:"pathPart" type:"string"`

	// Gets an API resource's method of a given HTTP verb.
	//
	// The resource methods are a map of methods indexed by methods' HTTP verbs
	// enabled on the resource. This method map is included in the 200 OK response
	// of the GET /restapis/{restapi_id}/resources/{resource_id} or GET /restapis/{restapi_id}/resources/{resource_id}?embed=methods
	// request.
	//
	// Example: Get the GET method of an API resource
	//
	// Request
	//  GET /restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET HTTP/1.1 Content-Type:
	//  application/json Host: apigateway.us-east-1.amazonaws.com X-Amz-Date: 20170223T031827Z
	//  Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20170223/us-east-1/apigateway/aws4_request,
	//  SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
	// Response
	//  { "_links": { "curies": [ { "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-integration-{rel}.html",
	//  "name": "integration", "templated": true }, { "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-integration-response-{rel}.html",
	//  "name": "integrationresponse", "templated": true }, { "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-method-{rel}.html",
	//  "name": "method", "templated": true }, { "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-method-response-{rel}.html",
	//  "name": "methodresponse", "templated": true } ], "self": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET",
	//  "name": "GET", "title": "GET" }, "integration:put": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration"
	//  }, "method:delete": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET"
	//  }, "method:integration": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration"
	//  }, "method:responses": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200",
	//  "name": "200", "title": "200" }, "method:update": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET"
	//  }, "methodresponse:put": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/{status_code}",
	//  "templated": true } }, "apiKeyRequired": false, "authorizationType": "NONE",
	//  "httpMethod": "GET", "_embedded": { "method:integration": { "_links": {
	//  "self": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration"
	//  }, "integration:delete": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration"
	//  }, "integration:responses": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200",
	//  "name": "200", "title": "200" }, "integration:update": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration"
	//  }, "integrationresponse:put": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/{status_code}",
	//  "templated": true } }, "cacheKeyParameters": [], "cacheNamespace": "3kzxbg5sa2",
	//  "credentials": "arn:aws:iam::123456789012:role/apigAwsProxyRole", "httpMethod":
	//  "POST", "passthroughBehavior": "WHEN_NO_MATCH", "requestParameters": { "integration.request.header.Content-Type":
	//  "'application/x-amz-json-1.1'" }, "requestTemplates": { "application/json":
	//  "{\n}" }, "type": "AWS", "uri": "arn:aws:apigateway:us-east-1:kinesis:action/ListStreams",
	//  "_embedded": { "integration:responses": { "_links": { "self": { "href":
	//  "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200",
	//  "name": "200", "title": "200" }, "integrationresponse:delete": { "href":
	//  "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200"
	//  }, "integrationresponse:update": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200"
	//  } }, "responseParameters": { "method.response.header.Content-Type": "'application/xml'"
	//  }, "responseTemplates": { "application/json": "$util.urlDecode(\"%3CkinesisStreams%3E#foreach($stream
	//  in $input.path('$.StreamNames'))%3Cstream%3E%3Cname%3E$stream%3C/name%3E%3C/stream%3E#end%3C/kinesisStreams%3E\")\n"
	//  }, "statusCode": "200" } } }, "method:responses": { "_links": { "self":
	//  { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200",
	//  "name": "200", "title": "200" }, "methodresponse:delete": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200"
	//  }, "methodresponse:update": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200"
	//  } }, "responseModels": { "application/json": "Empty" }, "responseParameters":
	//  { "method.response.header.Content-Type": false }, "statusCode": "200" }
	//  } }
	// If the OPTIONS is enabled on the resource, you can follow the example here
	// to get that method. Just replace the GET of the last path segment in the
	// request URL with OPTIONS.
	ResourceMethods map[string]Method `locationName:"resourceMethods" type:"map"`
}

// String returns the string representation
func (s CreateResourceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateResourceOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ParentId != nil {
		v := *s.ParentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "parentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Path != nil {
		v := *s.Path

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "path", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PathPart != nil {
		v := *s.PathPart

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "pathPart", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ResourceMethods) > 0 {
		v := s.ResourceMethods

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "resourceMethods", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	return nil
}

const opCreateResource = "CreateResource"

// CreateResourceRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Creates a Resource resource.
//
//    // Example sending a request using CreateResourceRequest.
//    req := client.CreateResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateResourceRequest(input *CreateResourceInput) CreateResourceRequest {
	op := &aws.Operation{
		Name:       opCreateResource,
		HTTPMethod: "POST",
		HTTPPath:   "/restapis/{restapi_id}/resources/{parent_id}",
	}

	if input == nil {
		input = &CreateResourceInput{}
	}

	req := c.newRequest(op, input, &CreateResourceOutput{})
	return CreateResourceRequest{Request: req, Input: input, Copy: c.CreateResourceRequest}
}

// CreateResourceRequest is the request type for the
// CreateResource API operation.
type CreateResourceRequest struct {
	*aws.Request
	Input *CreateResourceInput
	Copy  func(*CreateResourceInput) CreateResourceRequest
}

// Send marshals and sends the CreateResource API request.
func (r CreateResourceRequest) Send(ctx context.Context) (*CreateResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateResourceResponse{
		CreateResourceOutput: r.Request.Data.(*CreateResourceOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateResourceResponse is the response type for the
// CreateResource API operation.
type CreateResourceResponse struct {
	*CreateResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateResource request.
func (r *CreateResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
