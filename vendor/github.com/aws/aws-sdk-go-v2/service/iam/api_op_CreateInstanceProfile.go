// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Creates a new instance profile. For information about instance profiles, see [Using roles for applications on Amazon EC2]
//
// in the IAM User Guide, and [Instance profiles]in the Amazon EC2 User Guide.
//
// For information about the number of instance profiles you can create, see [IAM object quotas] in
// the IAM User Guide.
//
// [Instance profiles]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/iam-roles-for-amazon-ec2.html#ec2-instance-profile
// [IAM object quotas]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html
//
// [Using roles for applications on Amazon EC2]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2.html
func (c *Client) CreateInstanceProfile(ctx context.Context, params *CreateInstanceProfileInput, optFns ...func(*Options)) (*CreateInstanceProfileOutput, error) {
	if params == nil {
		params = &CreateInstanceProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateInstanceProfile", params, optFns, c.addOperationCreateInstanceProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateInstanceProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateInstanceProfileInput struct {

	// The name of the instance profile to create.
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	InstanceProfileName *string

	//  The path to the instance profile. For more information about paths, see [IAM Identifiers] in
	// the IAM User Guide.
	//
	// This parameter is optional. If it is not included, it defaults to a slash (/).
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of
	// either a forward slash (/) by itself or a string that must begin and end with
	// forward slashes. In addition, it can contain any ASCII character from the ! (
	// \u0021 ) through the DEL character ( \u007F ), including most punctuation
	// characters, digits, and upper and lowercased letters.
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	// [IAM Identifiers]: https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html
	Path *string

	// A list of tags that you want to attach to the newly created IAM instance
	// profile. Each tag consists of a key name and an associated value. For more
	// information about tagging, see [Tagging IAM resources]in the IAM User Guide.
	//
	// If any one of the tags is invalid or if you exceed the allowed maximum number
	// of tags, then the entire request fails and the resource is not created.
	//
	// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

// Contains the response to a successful CreateInstanceProfile request.
type CreateInstanceProfileOutput struct {

	// A structure containing details about the new instance profile.
	//
	// This member is required.
	InstanceProfile *types.InstanceProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateInstanceProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateInstanceProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateInstanceProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateInstanceProfile"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addCredentialSource(stack, options); err != nil {
		return err
	}
	if err = addOpCreateInstanceProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateInstanceProfile(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateInstanceProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateInstanceProfile",
	}
}
