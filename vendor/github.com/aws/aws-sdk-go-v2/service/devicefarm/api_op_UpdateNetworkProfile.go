// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateNetworkProfileInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the project for which you want to update
	// network profile settings.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`

	// The description of the network profile about which you are returning information.
	Description *string `locationName:"description" type:"string"`

	// The data throughput rate in bits per second, as an integer from 0 to 104857600.
	DownlinkBandwidthBits *int64 `locationName:"downlinkBandwidthBits" type:"long"`

	// Delay time for all packets to destination in milliseconds as an integer from
	// 0 to 2000.
	DownlinkDelayMs *int64 `locationName:"downlinkDelayMs" type:"long"`

	// Time variation in the delay of received packets in milliseconds as an integer
	// from 0 to 2000.
	DownlinkJitterMs *int64 `locationName:"downlinkJitterMs" type:"long"`

	// Proportion of received packets that fail to arrive from 0 to 100 percent.
	DownlinkLossPercent *int64 `locationName:"downlinkLossPercent" type:"integer"`

	// The name of the network profile about which you are returning information.
	Name *string `locationName:"name" type:"string"`

	// The type of network profile to return information about. Valid values are
	// listed here.
	Type NetworkProfileType `locationName:"type" type:"string" enum:"true"`

	// The data throughput rate in bits per second, as an integer from 0 to 104857600.
	UplinkBandwidthBits *int64 `locationName:"uplinkBandwidthBits" type:"long"`

	// Delay time for all packets to destination in milliseconds as an integer from
	// 0 to 2000.
	UplinkDelayMs *int64 `locationName:"uplinkDelayMs" type:"long"`

	// Time variation in the delay of received packets in milliseconds as an integer
	// from 0 to 2000.
	UplinkJitterMs *int64 `locationName:"uplinkJitterMs" type:"long"`

	// Proportion of transmitted packets that fail to arrive from 0 to 100 percent.
	UplinkLossPercent *int64 `locationName:"uplinkLossPercent" type:"integer"`
}

// String returns the string representation
func (s UpdateNetworkProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateNetworkProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateNetworkProfileInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateNetworkProfileOutput struct {
	_ struct{} `type:"structure"`

	// A list of the available network profiles.
	NetworkProfile *NetworkProfile `locationName:"networkProfile" type:"structure"`
}

// String returns the string representation
func (s UpdateNetworkProfileOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateNetworkProfile = "UpdateNetworkProfile"

// UpdateNetworkProfileRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Updates the network profile.
//
//    // Example sending a request using UpdateNetworkProfileRequest.
//    req := client.UpdateNetworkProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/UpdateNetworkProfile
func (c *Client) UpdateNetworkProfileRequest(input *UpdateNetworkProfileInput) UpdateNetworkProfileRequest {
	op := &aws.Operation{
		Name:       opUpdateNetworkProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateNetworkProfileInput{}
	}

	req := c.newRequest(op, input, &UpdateNetworkProfileOutput{})
	return UpdateNetworkProfileRequest{Request: req, Input: input, Copy: c.UpdateNetworkProfileRequest}
}

// UpdateNetworkProfileRequest is the request type for the
// UpdateNetworkProfile API operation.
type UpdateNetworkProfileRequest struct {
	*aws.Request
	Input *UpdateNetworkProfileInput
	Copy  func(*UpdateNetworkProfileInput) UpdateNetworkProfileRequest
}

// Send marshals and sends the UpdateNetworkProfile API request.
func (r UpdateNetworkProfileRequest) Send(ctx context.Context) (*UpdateNetworkProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateNetworkProfileResponse{
		UpdateNetworkProfileOutput: r.Request.Data.(*UpdateNetworkProfileOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateNetworkProfileResponse is the response type for the
// UpdateNetworkProfile API operation.
type UpdateNetworkProfileResponse struct {
	*UpdateNetworkProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateNetworkProfile request.
func (r *UpdateNetworkProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
