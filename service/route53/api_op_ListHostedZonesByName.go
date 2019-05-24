// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Retrieves a list of the public and private hosted zones that are associated
// with the current AWS account in ASCII order by domain name.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/ListHostedZonesByNameRequest
type ListHostedZonesByNameInput struct {
	_ struct{} `type:"structure"`

	// (Optional) For your first request to ListHostedZonesByName, include the dnsname
	// parameter only if you want to specify the name of the first hosted zone in
	// the response. If you don't include the dnsname parameter, Amazon Route 53
	// returns all of the hosted zones that were created by the current AWS account,
	// in ASCII order. For subsequent requests, include both dnsname and hostedzoneid
	// parameters. For dnsname, specify the value of NextDNSName from the previous
	// response.
	DNSName *string `location:"querystring" locationName:"dnsname" type:"string"`

	// (Optional) For your first request to ListHostedZonesByName, do not include
	// the hostedzoneid parameter.
	//
	// If you have more hosted zones than the value of maxitems, ListHostedZonesByName
	// returns only the first maxitems hosted zones. To get the next group of maxitems
	// hosted zones, submit another request to ListHostedZonesByName and include
	// both dnsname and hostedzoneid parameters. For the value of hostedzoneid,
	// specify the value of the NextHostedZoneId element from the previous response.
	HostedZoneId *string `location:"querystring" locationName:"hostedzoneid" type:"string"`

	// The maximum number of hosted zones to be included in the response body for
	// this request. If you have more than maxitems hosted zones, then the value
	// of the IsTruncated element in the response is true, and the values of NextDNSName
	// and NextHostedZoneId specify the first hosted zone in the next group of maxitems
	// hosted zones.
	MaxItems *string `location:"querystring" locationName:"maxitems" type:"string"`
}

// String returns the string representation
func (s ListHostedZonesByNameInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListHostedZonesByNameInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.DNSName != nil {
		v := *s.DNSName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "dnsname", protocol.StringValue(v), metadata)
	}
	if s.HostedZoneId != nil {
		v := *s.HostedZoneId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "hostedzoneid", protocol.StringValue(v), metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxitems", protocol.StringValue(v), metadata)
	}
	return nil
}

// A complex type that contains the response information for the request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/ListHostedZonesByNameResponse
type ListHostedZonesByNameOutput struct {
	_ struct{} `type:"structure"`

	// For the second and subsequent calls to ListHostedZonesByName, DNSName is
	// the value that you specified for the dnsname parameter in the request that
	// produced the current response.
	DNSName *string `type:"string"`

	// The ID that Amazon Route 53 assigned to the hosted zone when you created
	// it.
	HostedZoneId *string `type:"string"`

	// A complex type that contains general information about the hosted zone.
	//
	// HostedZones is a required field
	HostedZones []HostedZone `locationNameList:"HostedZone" type:"list" required:"true"`

	// A flag that indicates whether there are more hosted zones to be listed. If
	// the response was truncated, you can get the next group of maxitems hosted
	// zones by calling ListHostedZonesByName again and specifying the values of
	// NextDNSName and NextHostedZoneId elements in the dnsname and hostedzoneid
	// parameters.
	//
	// IsTruncated is a required field
	IsTruncated *bool `type:"boolean" required:"true"`

	// The value that you specified for the maxitems parameter in the call to ListHostedZonesByName
	// that produced the current response.
	//
	// MaxItems is a required field
	MaxItems *string `type:"string" required:"true"`

	// If IsTruncated is true, the value of NextDNSName is the name of the first
	// hosted zone in the next group of maxitems hosted zones. Call ListHostedZonesByName
	// again and specify the value of NextDNSName and NextHostedZoneId in the dnsname
	// and hostedzoneid parameters, respectively.
	//
	// This element is present only if IsTruncated is true.
	NextDNSName *string `type:"string"`

	// If IsTruncated is true, the value of NextHostedZoneId identifies the first
	// hosted zone in the next group of maxitems hosted zones. Call ListHostedZonesByName
	// again and specify the value of NextDNSName and NextHostedZoneId in the dnsname
	// and hostedzoneid parameters, respectively.
	//
	// This element is present only if IsTruncated is true.
	NextHostedZoneId *string `type:"string"`
}

// String returns the string representation
func (s ListHostedZonesByNameOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListHostedZonesByNameOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DNSName != nil {
		v := *s.DNSName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DNSName", protocol.StringValue(v), metadata)
	}
	if s.HostedZoneId != nil {
		v := *s.HostedZoneId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "HostedZoneId", protocol.StringValue(v), metadata)
	}
	if len(s.HostedZones) > 0 {
		v := s.HostedZones

		metadata := protocol.Metadata{ListLocationName: "HostedZone"}
		ls0 := e.List(protocol.BodyTarget, "HostedZones", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.IsTruncated != nil {
		v := *s.IsTruncated

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IsTruncated", protocol.BoolValue(v), metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxItems", protocol.StringValue(v), metadata)
	}
	if s.NextDNSName != nil {
		v := *s.NextDNSName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextDNSName", protocol.StringValue(v), metadata)
	}
	if s.NextHostedZoneId != nil {
		v := *s.NextHostedZoneId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextHostedZoneId", protocol.StringValue(v), metadata)
	}
	return nil
}

const opListHostedZonesByName = "ListHostedZonesByName"

// ListHostedZonesByNameRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Retrieves a list of your hosted zones in lexicographic order. The response
// includes a HostedZones child element for each hosted zone created by the
// current AWS account.
//
// ListHostedZonesByName sorts hosted zones by name with the labels reversed.
// For example:
//
// com.example.www.
//
// Note the trailing dot, which can change the sort order in some circumstances.
//
// If the domain name includes escape characters or Punycode, ListHostedZonesByName
// alphabetizes the domain name using the escaped or Punycoded value, which
// is the format that Amazon Route 53 saves in its database. For example, to
// create a hosted zone for exämple.com, you specify ex\344mple.com for the
// domain name. ListHostedZonesByName alphabetizes it as:
//
// com.ex\344mple.
//
// The labels are reversed and alphabetized using the escaped value. For more
// information about valid domain name formats, including internationalized
// domain names, see DNS Domain Name Format (http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DomainNameFormat.html)
// in the Amazon Route 53 Developer Guide.
//
// Route 53 returns up to 100 items in each response. If you have a lot of hosted
// zones, use the MaxItems parameter to list them in groups of up to 100. The
// response includes values that help navigate from one group of MaxItems hosted
// zones to the next:
//
//    * The DNSName and HostedZoneId elements in the response contain the values,
//    if any, specified for the dnsname and hostedzoneid parameters in the request
//    that produced the current response.
//
//    * The MaxItems element in the response contains the value, if any, that
//    you specified for the maxitems parameter in the request that produced
//    the current response.
//
//    * If the value of IsTruncated in the response is true, there are more
//    hosted zones associated with the current AWS account. If IsTruncated is
//    false, this response includes the last hosted zone that is associated
//    with the current account. The NextDNSName element and NextHostedZoneId
//    elements are omitted from the response.
//
//    * The NextDNSName and NextHostedZoneId elements in the response contain
//    the domain name and the hosted zone ID of the next hosted zone that is
//    associated with the current AWS account. If you want to list more hosted
//    zones, make another call to ListHostedZonesByName, and specify the value
//    of NextDNSName and NextHostedZoneId in the dnsname and hostedzoneid parameters,
//    respectively.
//
//    // Example sending a request using ListHostedZonesByNameRequest.
//    req := client.ListHostedZonesByNameRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/ListHostedZonesByName
func (c *Client) ListHostedZonesByNameRequest(input *ListHostedZonesByNameInput) ListHostedZonesByNameRequest {
	op := &aws.Operation{
		Name:       opListHostedZonesByName,
		HTTPMethod: "GET",
		HTTPPath:   "/2013-04-01/hostedzonesbyname",
	}

	if input == nil {
		input = &ListHostedZonesByNameInput{}
	}

	req := c.newRequest(op, input, &ListHostedZonesByNameOutput{})
	return ListHostedZonesByNameRequest{Request: req, Input: input, Copy: c.ListHostedZonesByNameRequest}
}

// ListHostedZonesByNameRequest is the request type for the
// ListHostedZonesByName API operation.
type ListHostedZonesByNameRequest struct {
	*aws.Request
	Input *ListHostedZonesByNameInput
	Copy  func(*ListHostedZonesByNameInput) ListHostedZonesByNameRequest
}

// Send marshals and sends the ListHostedZonesByName API request.
func (r ListHostedZonesByNameRequest) Send(ctx context.Context) (*ListHostedZonesByNameResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListHostedZonesByNameResponse{
		ListHostedZonesByNameOutput: r.Request.Data.(*ListHostedZonesByNameOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListHostedZonesByNameResponse is the response type for the
// ListHostedZonesByName API operation.
type ListHostedZonesByNameResponse struct {
	*ListHostedZonesByNameOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListHostedZonesByName request.
func (r *ListHostedZonesByNameResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
