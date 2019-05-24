// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package machinelearning

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateMLModelInput struct {
	_ struct{} `type:"structure"`

	// A user-supplied ID that uniquely identifies the MLModel.
	//
	// MLModelId is a required field
	MLModelId *string `min:"1" type:"string" required:"true"`

	// A user-supplied name or description of the MLModel.
	MLModelName *string `type:"string"`

	// The category of supervised learning that this MLModel will address. Choose
	// from the following types:
	//
	//    * Choose REGRESSION if the MLModel will be used to predict a numeric value.
	//
	//    * Choose BINARY if the MLModel result has two possible values.
	//
	//    * Choose MULTICLASS if the MLModel result has a limited number of values.
	//
	// For more information, see the Amazon Machine Learning Developer Guide (http://docs.aws.amazon.com/machine-learning/latest/dg).
	//
	// MLModelType is a required field
	MLModelType MLModelType `type:"string" required:"true" enum:"true"`

	// A list of the training parameters in the MLModel. The list is implemented
	// as a map of key-value pairs.
	//
	// The following is the current set of training parameters:
	//
	//    * sgd.maxMLModelSizeInBytes - The maximum allowed size of the model. Depending
	//    on the input data, the size of the model might affect its performance.
	//    The value is an integer that ranges from 100000 to 2147483648. The default
	//    value is 33554432.
	//
	//    * sgd.maxPasses - The number of times that the training process traverses
	//    the observations to build the MLModel. The value is an integer that ranges
	//    from 1 to 10000. The default value is 10.
	//
	//    * sgd.shuffleType - Whether Amazon ML shuffles the training data. Shuffling
	//    the data improves a model's ability to find the optimal solution for a
	//    variety of data types. The valid values are auto and none. The default
	//    value is none. We strongly recommend that you shuffle your data.
	//
	//    * sgd.l1RegularizationAmount - The coefficient regularization L1 norm.
	//    It controls overfitting the data by penalizing large coefficients. This
	//    tends to drive coefficients to zero, resulting in a sparse feature set.
	//    If you use this parameter, start by specifying a small value, such as
	//    1.0E-08. The value is a double that ranges from 0 to MAX_DOUBLE. The default
	//    is to not use L1 normalization. This parameter can't be used when L2 is
	//    specified. Use this parameter sparingly.
	//
	//    * sgd.l2RegularizationAmount - The coefficient regularization L2 norm.
	//    It controls overfitting the data by penalizing large coefficients. This
	//    tends to drive coefficients to small, nonzero values. If you use this
	//    parameter, start by specifying a small value, such as 1.0E-08. The value
	//    is a double that ranges from 0 to MAX_DOUBLE. The default is to not use
	//    L2 normalization. This parameter can't be used when L1 is specified. Use
	//    this parameter sparingly.
	Parameters map[string]string `type:"map"`

	// The data recipe for creating the MLModel. You must specify either the recipe
	// or its URI. If you don't specify a recipe or its URI, Amazon ML creates a
	// default.
	Recipe *string `type:"string"`

	// The Amazon Simple Storage Service (Amazon S3) location and file name that
	// contains the MLModel recipe. You must specify either the recipe or its URI.
	// If you don't specify a recipe or its URI, Amazon ML creates a default.
	RecipeUri *string `type:"string"`

	// The DataSource that points to the training data.
	//
	// TrainingDataSourceId is a required field
	TrainingDataSourceId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateMLModelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateMLModelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateMLModelInput"}

	if s.MLModelId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MLModelId"))
	}
	if s.MLModelId != nil && len(*s.MLModelId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MLModelId", 1))
	}
	if len(s.MLModelType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("MLModelType"))
	}

	if s.TrainingDataSourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrainingDataSourceId"))
	}
	if s.TrainingDataSourceId != nil && len(*s.TrainingDataSourceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TrainingDataSourceId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a CreateMLModel operation, and is an acknowledgement
// that Amazon ML received the request.
//
// The CreateMLModel operation is asynchronous. You can poll for status updates
// by using the GetMLModel operation and checking the Status parameter.
type CreateMLModelOutput struct {
	_ struct{} `type:"structure"`

	// A user-supplied ID that uniquely identifies the MLModel. This value should
	// be identical to the value of the MLModelId in the request.
	MLModelId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateMLModelOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateMLModel = "CreateMLModel"

// CreateMLModelRequest returns a request value for making API operation for
// Amazon Machine Learning.
//
// Creates a new MLModel using the DataSource and the recipe as information
// sources.
//
// An MLModel is nearly immutable. Users can update only the MLModelName and
// the ScoreThreshold in an MLModel without creating a new MLModel.
//
// CreateMLModel is an asynchronous operation. In response to CreateMLModel,
// Amazon Machine Learning (Amazon ML) immediately returns and sets the MLModel
// status to PENDING. After the MLModel has been created and ready is for use,
// Amazon ML sets the status to COMPLETED.
//
// You can use the GetMLModel operation to check the progress of the MLModel
// during the creation operation.
//
// CreateMLModel requires a DataSource with computed statistics, which can be
// created by setting ComputeStatistics to true in CreateDataSourceFromRDS,
// CreateDataSourceFromS3, or CreateDataSourceFromRedshift operations.
//
//    // Example sending a request using CreateMLModelRequest.
//    req := client.CreateMLModelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateMLModelRequest(input *CreateMLModelInput) CreateMLModelRequest {
	op := &aws.Operation{
		Name:       opCreateMLModel,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateMLModelInput{}
	}

	req := c.newRequest(op, input, &CreateMLModelOutput{})
	return CreateMLModelRequest{Request: req, Input: input, Copy: c.CreateMLModelRequest}
}

// CreateMLModelRequest is the request type for the
// CreateMLModel API operation.
type CreateMLModelRequest struct {
	*aws.Request
	Input *CreateMLModelInput
	Copy  func(*CreateMLModelInput) CreateMLModelRequest
}

// Send marshals and sends the CreateMLModel API request.
func (r CreateMLModelRequest) Send(ctx context.Context) (*CreateMLModelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateMLModelResponse{
		CreateMLModelOutput: r.Request.Data.(*CreateMLModelOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateMLModelResponse is the response type for the
// CreateMLModel API operation.
type CreateMLModelResponse struct {
	*CreateMLModelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateMLModel request.
func (r *CreateMLModelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
