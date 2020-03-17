// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to return the requested sending authorization policies
// for an identity. Sending authorization is an Amazon SES feature that enables
// you to authorize other senders to use your identities. For information, see
// the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
type GetIdentityPoliciesInput struct {
	_ struct{} `type:"structure"`

	// The identity for which the policies will be retrieved. You can specify an
	// identity by using its name or by using its Amazon Resource Name (ARN). Examples:
	// user@example.com, example.com, arn:aws:ses:us-east-1:123456789012:identity/example.com.
	//
	// To successfully call this API, you must own the identity.
	//
	// Identity is a required field
	Identity *string `type:"string" required:"true"`

	// A list of the names of policies to be retrieved. You can retrieve a maximum
	// of 20 policies at a time. If you do not know the names of the policies that
	// are attached to the identity, you can use ListIdentityPolicies.
	//
	// PolicyNames is a required field
	PolicyNames []string `type:"list" required:"true"`
}

// String returns the string representation
func (s GetIdentityPoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetIdentityPoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetIdentityPoliciesInput"}

	if s.Identity == nil {
		invalidParams.Add(aws.NewErrParamRequired("Identity"))
	}

	if s.PolicyNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyNames"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the requested sending authorization policies.
type GetIdentityPoliciesOutput struct {
	_ struct{} `type:"structure"`

	// A map of policy names to policies.
	//
	// Policies is a required field
	Policies map[string]string `type:"map" required:"true"`
}

// String returns the string representation
func (s GetIdentityPoliciesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetIdentityPolicies = "GetIdentityPolicies"

// GetIdentityPoliciesRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Returns the requested sending authorization policies for the given identity
// (an email address or a domain). The policies are returned as a map of policy
// names to policy contents. You can retrieve a maximum of 20 policies at a
// time.
//
// This API is for the identity owner only. If you have not verified the identity,
// this API will return an error.
//
// Sending authorization is a feature that enables an identity owner to authorize
// other senders to use its identities. For information about using sending
// authorization, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using GetIdentityPoliciesRequest.
//    req := client.GetIdentityPoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/GetIdentityPolicies
func (c *Client) GetIdentityPoliciesRequest(input *GetIdentityPoliciesInput) GetIdentityPoliciesRequest {
	op := &aws.Operation{
		Name:       opGetIdentityPolicies,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetIdentityPoliciesInput{}
	}

	req := c.newRequest(op, input, &GetIdentityPoliciesOutput{})
	return GetIdentityPoliciesRequest{Request: req, Input: input, Copy: c.GetIdentityPoliciesRequest}
}

// GetIdentityPoliciesRequest is the request type for the
// GetIdentityPolicies API operation.
type GetIdentityPoliciesRequest struct {
	*aws.Request
	Input *GetIdentityPoliciesInput
	Copy  func(*GetIdentityPoliciesInput) GetIdentityPoliciesRequest
}

// Send marshals and sends the GetIdentityPolicies API request.
func (r GetIdentityPoliciesRequest) Send(ctx context.Context) (*GetIdentityPoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetIdentityPoliciesResponse{
		GetIdentityPoliciesOutput: r.Request.Data.(*GetIdentityPoliciesOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetIdentityPoliciesResponse is the response type for the
// GetIdentityPolicies API operation.
type GetIdentityPoliciesResponse struct {
	*GetIdentityPoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetIdentityPolicies request.
func (r *GetIdentityPoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
