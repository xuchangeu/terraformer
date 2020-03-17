// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateRuleGroupInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// The RuleGroupId of the RuleGroup that you want to update. RuleGroupId is
	// returned by CreateRuleGroup and by ListRuleGroups.
	//
	// RuleGroupId is a required field
	RuleGroupId *string `min:"1" type:"string" required:"true"`

	// An array of RuleGroupUpdate objects that you want to insert into or delete
	// from a RuleGroup.
	//
	// You can only insert REGULAR rules into a rule group.
	//
	// ActivatedRule|OverrideAction applies only when updating or adding a RuleGroup
	// to a WebACL. In this case you do not use ActivatedRule|Action. For all other
	// update requests, ActivatedRule|Action is used instead of ActivatedRule|OverrideAction.
	//
	// Updates is a required field
	Updates []RuleGroupUpdate `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s UpdateRuleGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateRuleGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateRuleGroupInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.RuleGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleGroupId"))
	}
	if s.RuleGroupId != nil && len(*s.RuleGroupId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleGroupId", 1))
	}

	if s.Updates == nil {
		invalidParams.Add(aws.NewErrParamRequired("Updates"))
	}
	if s.Updates != nil && len(s.Updates) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Updates", 1))
	}
	if s.Updates != nil {
		for i, v := range s.Updates {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Updates", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateRuleGroupOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the UpdateRuleGroup request. You
	// can also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateRuleGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateRuleGroup = "UpdateRuleGroup"

// UpdateRuleGroupRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Inserts or deletes ActivatedRule objects in a RuleGroup.
//
// You can only insert REGULAR rules into a rule group.
//
// You can have a maximum of ten rules per rule group.
//
// To create and configure a RuleGroup, perform the following steps:
//
// Create and update the Rules that you want to include in the RuleGroup. See
// CreateRule.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateRuleGroup request.
//
// Submit an UpdateRuleGroup request to add Rules to the RuleGroup.
//
// Create and update a WebACL that contains the RuleGroup. See CreateWebACL.
//
// If you want to replace one Rule with another, you delete the existing one
// and add the new one.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
//
//    // Example sending a request using UpdateRuleGroupRequest.
//    req := client.UpdateRuleGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/UpdateRuleGroup
func (c *Client) UpdateRuleGroupRequest(input *UpdateRuleGroupInput) UpdateRuleGroupRequest {
	op := &aws.Operation{
		Name:       opUpdateRuleGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateRuleGroupInput{}
	}

	req := c.newRequest(op, input, &UpdateRuleGroupOutput{})
	return UpdateRuleGroupRequest{Request: req, Input: input, Copy: c.UpdateRuleGroupRequest}
}

// UpdateRuleGroupRequest is the request type for the
// UpdateRuleGroup API operation.
type UpdateRuleGroupRequest struct {
	*aws.Request
	Input *UpdateRuleGroupInput
	Copy  func(*UpdateRuleGroupInput) UpdateRuleGroupRequest
}

// Send marshals and sends the UpdateRuleGroup API request.
func (r UpdateRuleGroupRequest) Send(ctx context.Context) (*UpdateRuleGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateRuleGroupResponse{
		UpdateRuleGroupOutput: r.Request.Data.(*UpdateRuleGroupOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateRuleGroupResponse is the response type for the
// UpdateRuleGroup API operation.
type UpdateRuleGroupResponse struct {
	*UpdateRuleGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateRuleGroup request.
func (r *UpdateRuleGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
