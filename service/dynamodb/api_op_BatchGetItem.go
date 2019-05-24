// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a BatchGetItem operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/BatchGetItemInput
type BatchGetItemInput struct {
	_ struct{} `type:"structure"`

	// A map of one or more table names and, for each table, a map that describes
	// one or more items to retrieve from that table. Each table name can be used
	// only once per BatchGetItem request.
	//
	// Each element in the map of items to retrieve consists of the following:
	//
	//    * ConsistentRead - If true, a strongly consistent read is used; if false
	//    (the default), an eventually consistent read is used.
	//
	//    * ExpressionAttributeNames - One or more substitution tokens for attribute
	//    names in the ProjectionExpression parameter. The following are some use
	//    cases for using ExpressionAttributeNames: To access an attribute whose
	//    name conflicts with a DynamoDB reserved word. To create a placeholder
	//    for repeating occurrences of an attribute name in an expression. To prevent
	//    special characters in an attribute name from being misinterpreted in an
	//    expression. Use the # character in an expression to dereference an attribute
	//    name. For example, consider the following attribute name: Percentile The
	//    name of this attribute conflicts with a reserved word, so it cannot be
	//    used directly in an expression. (For the complete list of reserved words,
	//    see Reserved Words (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html)
	//    in the Amazon DynamoDB Developer Guide). To work around this, you could
	//    specify the following for ExpressionAttributeNames: {"#P":"Percentile"}
	//    You could then use this substitution in an expression, as in this example:
	//    #P = :val Tokens that begin with the : character are expression attribute
	//    values, which are placeholders for the actual value at runtime. For more
	//    information on expression attribute names, see Accessing Item Attributes
	//    (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html)
	//    in the Amazon DynamoDB Developer Guide.
	//
	//    * Keys - An array of primary key attribute values that define specific
	//    items in the table. For each primary key, you must provide all of the
	//    key attributes. For example, with a simple primary key, you only need
	//    to provide the partition key value. For a composite key, you must provide
	//    both the partition key value and the sort key value.
	//
	//    * ProjectionExpression - A string that identifies one or more attributes
	//    to retrieve from the table. These attributes can include scalars, sets,
	//    or elements of a JSON document. The attributes in the expression must
	//    be separated by commas. If no attribute names are specified, then all
	//    attributes will be returned. If any of the requested attributes are not
	//    found, they will not appear in the result. For more information, see Accessing
	//    Item Attributes (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html)
	//    in the Amazon DynamoDB Developer Guide.
	//
	//    * AttributesToGet - This is a legacy parameter. Use ProjectionExpression
	//    instead. For more information, see AttributesToGet (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.AttributesToGet.html)
	//    in the Amazon DynamoDB Developer Guide.
	//
	// RequestItems is a required field
	RequestItems map[string]KeysAndAttributes `min:"1" type:"map" required:"true"`

	// Determines the level of detail about provisioned throughput consumption that
	// is returned in the response:
	//
	//    * INDEXES - The response includes the aggregate ConsumedCapacity for the
	//    operation, together with ConsumedCapacity for each table and secondary
	//    index that was accessed. Note that some operations, such as GetItem and
	//    BatchGetItem, do not access any indexes at all. In these cases, specifying
	//    INDEXES will only return ConsumedCapacity information for table(s).
	//
	//    * TOTAL - The response includes only the aggregate ConsumedCapacity for
	//    the operation.
	//
	//    * NONE - No ConsumedCapacity details are included in the response.
	ReturnConsumedCapacity ReturnConsumedCapacity `type:"string" enum:"true"`
}

// String returns the string representation
func (s BatchGetItemInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchGetItemInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchGetItemInput"}

	if s.RequestItems == nil {
		invalidParams.Add(aws.NewErrParamRequired("RequestItems"))
	}
	if s.RequestItems != nil && len(s.RequestItems) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RequestItems", 1))
	}
	if s.RequestItems != nil {
		for i, v := range s.RequestItems {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "RequestItems", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a BatchGetItem operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/BatchGetItemOutput
type BatchGetItemOutput struct {
	_ struct{} `type:"structure"`

	// The read capacity units consumed by the entire BatchGetItem operation.
	//
	// Each element consists of:
	//
	//    * TableName - The table that consumed the provisioned throughput.
	//
	//    * CapacityUnits - The total number of capacity units consumed.
	ConsumedCapacity []ConsumedCapacity `type:"list"`

	// A map of table name to a list of items. Each object in Responses consists
	// of a table name, along with a map of attribute data consisting of the data
	// type and attribute value.
	Responses map[string][]map[string]AttributeValue `type:"map"`

	// A map of tables and their respective keys that were not processed with the
	// current response. The UnprocessedKeys value is in the same form as RequestItems,
	// so the value can be provided directly to a subsequent BatchGetItem operation.
	// For more information, see RequestItems in the Request Parameters section.
	//
	// Each element consists of:
	//
	//    * Keys - An array of primary key attribute values that define specific
	//    items in the table.
	//
	//    * ProjectionExpression - One or more attributes to be retrieved from the
	//    table or index. By default, all attributes are returned. If a requested
	//    attribute is not found, it does not appear in the result.
	//
	//    * ConsistentRead - The consistency of a read operation. If set to true,
	//    then a strongly consistent read is used; otherwise, an eventually consistent
	//    read is used.
	//
	// If there are no unprocessed keys remaining, the response contains an empty
	// UnprocessedKeys map.
	UnprocessedKeys map[string]KeysAndAttributes `min:"1" type:"map"`
}

// String returns the string representation
func (s BatchGetItemOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchGetItem = "BatchGetItem"

// BatchGetItemRequest returns a request value for making API operation for
// Amazon DynamoDB.
//
// The BatchGetItem operation returns the attributes of one or more items from
// one or more tables. You identify requested items by primary key.
//
// A single operation can retrieve up to 16 MB of data, which can contain as
// many as 100 items. BatchGetItem will return a partial result if the response
// size limit is exceeded, the table's provisioned throughput is exceeded, or
// an internal processing failure occurs. If a partial result is returned, the
// operation returns a value for UnprocessedKeys. You can use this value to
// retry the operation starting with the next item to get.
//
// If you request more than 100 items BatchGetItem will return a ValidationException
// with the message "Too many items requested for the BatchGetItem call".
//
// For example, if you ask to retrieve 100 items, but each individual item is
// 300 KB in size, the system returns 52 items (so as not to exceed the 16 MB
// limit). It also returns an appropriate UnprocessedKeys value so you can get
// the next page of results. If desired, your application can include its own
// logic to assemble the pages of results into one data set.
//
// If none of the items can be processed due to insufficient provisioned throughput
// on all of the tables in the request, then BatchGetItem will return a ProvisionedThroughputExceededException.
// If at least one of the items is successfully processed, then BatchGetItem
// completes successfully, while returning the keys of the unread items in UnprocessedKeys.
//
// If DynamoDB returns any unprocessed items, you should retry the batch operation
// on those items. However, we strongly recommend that you use an exponential
// backoff algorithm. If you retry the batch operation immediately, the underlying
// read or write requests can still fail due to throttling on the individual
// tables. If you delay the batch operation using exponential backoff, the individual
// requests in the batch are much more likely to succeed.
//
// For more information, see Batch Operations and Error Handling (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ErrorHandling.html#BatchOperations)
// in the Amazon DynamoDB Developer Guide.
//
// By default, BatchGetItem performs eventually consistent reads on every table
// in the request. If you want strongly consistent reads instead, you can set
// ConsistentRead to true for any or all tables.
//
// In order to minimize response latency, BatchGetItem retrieves items in parallel.
//
// When designing your application, keep in mind that DynamoDB does not return
// items in any particular order. To help parse the response by item, include
// the primary key values for the items in your request in the ProjectionExpression
// parameter.
//
// If a requested item does not exist, it is not returned in the result. Requests
// for nonexistent items consume the minimum read capacity units according to
// the type of read. For more information, see Capacity Units Calculations (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#CapacityUnitCalculations)
// in the Amazon DynamoDB Developer Guide.
//
//    // Example sending a request using BatchGetItemRequest.
//    req := client.BatchGetItemRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/BatchGetItem
func (c *Client) BatchGetItemRequest(input *BatchGetItemInput) BatchGetItemRequest {
	op := &aws.Operation{
		Name:       opBatchGetItem,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"RequestItems"},
			OutputTokens:    []string{"UnprocessedKeys"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &BatchGetItemInput{}
	}

	req := c.newRequest(op, input, &BatchGetItemOutput{})
	return BatchGetItemRequest{Request: req, Input: input, Copy: c.BatchGetItemRequest}
}

// BatchGetItemRequest is the request type for the
// BatchGetItem API operation.
type BatchGetItemRequest struct {
	*aws.Request
	Input *BatchGetItemInput
	Copy  func(*BatchGetItemInput) BatchGetItemRequest
}

// Send marshals and sends the BatchGetItem API request.
func (r BatchGetItemRequest) Send(ctx context.Context) (*BatchGetItemResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchGetItemResponse{
		BatchGetItemOutput: r.Request.Data.(*BatchGetItemOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewBatchGetItemRequestPaginator returns a paginator for BatchGetItem.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.BatchGetItemRequest(input)
//   p := dynamodb.NewBatchGetItemRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewBatchGetItemPaginator(req BatchGetItemRequest) BatchGetItemPaginator {
	return BatchGetItemPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *BatchGetItemInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// BatchGetItemPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type BatchGetItemPaginator struct {
	aws.Pager
}

func (p *BatchGetItemPaginator) CurrentPage() *BatchGetItemOutput {
	return p.Pager.CurrentPage().(*BatchGetItemOutput)
}

// BatchGetItemResponse is the response type for the
// BatchGetItem API operation.
type BatchGetItemResponse struct {
	*BatchGetItemOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchGetItem request.
func (r *BatchGetItemResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
