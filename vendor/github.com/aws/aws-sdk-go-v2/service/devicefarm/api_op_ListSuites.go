// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to the list suites operation.
type ListSuitesInput struct {
	_ struct{} `type:"structure"`

	// The job's Amazon Resource Name (ARN).
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`
}

// String returns the string representation
func (s ListSuitesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSuitesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListSuitesInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}
	if s.NextToken != nil && len(*s.NextToken) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 4))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the result of a list suites request.
type ListSuitesOutput struct {
	_ struct{} `type:"structure"`

	// If the number of items that are returned is significantly large, this is
	// an identifier that is also returned. It can be used in a subsequent call
	// to this operation to return the next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`

	// Information about the suites.
	Suites []Suite `locationName:"suites" type:"list"`
}

// String returns the string representation
func (s ListSuitesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListSuites = "ListSuites"

// ListSuitesRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Gets information about test suites for a given job.
//
//    // Example sending a request using ListSuitesRequest.
//    req := client.ListSuitesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/ListSuites
func (c *Client) ListSuitesRequest(input *ListSuitesInput) ListSuitesRequest {
	op := &aws.Operation{
		Name:       opListSuites,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListSuitesInput{}
	}

	req := c.newRequest(op, input, &ListSuitesOutput{})
	return ListSuitesRequest{Request: req, Input: input, Copy: c.ListSuitesRequest}
}

// ListSuitesRequest is the request type for the
// ListSuites API operation.
type ListSuitesRequest struct {
	*aws.Request
	Input *ListSuitesInput
	Copy  func(*ListSuitesInput) ListSuitesRequest
}

// Send marshals and sends the ListSuites API request.
func (r ListSuitesRequest) Send(ctx context.Context) (*ListSuitesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSuitesResponse{
		ListSuitesOutput: r.Request.Data.(*ListSuitesOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListSuitesRequestPaginator returns a paginator for ListSuites.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListSuitesRequest(input)
//   p := devicefarm.NewListSuitesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListSuitesPaginator(req ListSuitesRequest) ListSuitesPaginator {
	return ListSuitesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListSuitesInput
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

// ListSuitesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListSuitesPaginator struct {
	aws.Pager
}

func (p *ListSuitesPaginator) CurrentPage() *ListSuitesOutput {
	return p.Pager.CurrentPage().(*ListSuitesOutput)
}

// ListSuitesResponse is the response type for the
// ListSuites API operation.
type ListSuitesResponse struct {
	*ListSuitesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSuites request.
func (r *ListSuitesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
