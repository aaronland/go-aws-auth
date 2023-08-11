// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds or updates an inline policy document that is embedded in the specified IAM
// role. When you embed an inline policy in a role, the inline policy is used as
// part of the role's access (permissions) policy. The role's trust policy is
// created at the same time as the role, using CreateRole (https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html)
// . You can update a role's trust policy using UpdateAssumeRolePolicy (https://docs.aws.amazon.com/IAM/latest/APIReference/API_UpdateAssumeRolePolicy.html)
// . For more information about roles, see IAM roles (https://docs.aws.amazon.com/IAM/latest/UserGuide/roles-toplevel.html)
// in the IAM User Guide. A role can also have a managed policy attached to it. To
// attach a managed policy to a role, use AttachRolePolicy (https://docs.aws.amazon.com/IAM/latest/APIReference/API_AttachRolePolicy.html)
// . To create a new managed policy, use CreatePolicy (https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreatePolicy.html)
// . For information about policies, see Managed policies and inline policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide. For information about the maximum number of inline
// policies that you can embed with a role, see IAM and STS quotas (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html)
// in the IAM User Guide. Because policy documents can be large, you should use
// POST rather than GET when calling PutRolePolicy . For general information about
// using the Query API with IAM, see Making query requests (https://docs.aws.amazon.com/IAM/latest/UserGuide/IAM_UsingQueryAPI.html)
// in the IAM User Guide.
func (c *Client) PutRolePolicy(ctx context.Context, params *PutRolePolicyInput, optFns ...func(*Options)) (*PutRolePolicyOutput, error) {
	if params == nil {
		params = &PutRolePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRolePolicy", params, optFns, c.addOperationPutRolePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRolePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRolePolicyInput struct {

	// The policy document. You must provide policies in JSON format in IAM. However,
	// for CloudFormation templates formatted in YAML, you can provide the policy in
	// JSON or YAML format. CloudFormation always converts a YAML policy to JSON format
	// before submitting it to IAM. The regex pattern (http://wikipedia.org/wiki/regex)
	// used to validate this parameter is a string of characters consisting of the
	// following:
	//   - Any printable ASCII character ranging from the space character ( \u0020 )
	//   through the end of the ASCII character range
	//   - The printable characters in the Basic Latin and Latin-1 Supplement
	//   character set (through \u00FF )
	//   - The special characters tab ( \u0009 ), line feed ( \u000A ), and carriage
	//   return ( \u000D )
	//
	// This member is required.
	PolicyDocument *string

	// The name of the policy document. This parameter allows (through its regex
	// pattern (http://wikipedia.org/wiki/regex) ) a string of characters consisting of
	// upper and lowercase alphanumeric characters with no spaces. You can also include
	// any of the following characters: _+=,.@-
	//
	// This member is required.
	PolicyName *string

	// The name of the role to associate the policy with. This parameter allows
	// (through its regex pattern (http://wikipedia.org/wiki/regex) ) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	RoleName *string

	noSmithyDocumentSerde
}

type PutRolePolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutRolePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPutRolePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPutRolePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addPutRolePolicyResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPutRolePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRolePolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutRolePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "PutRolePolicy",
	}
}

type opPutRolePolicyResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opPutRolePolicyResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opPutRolePolicyResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "iam"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "iam"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("iam")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addPutRolePolicyResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opPutRolePolicyResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
