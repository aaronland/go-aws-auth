// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Modifies an existing patch baseline. Fields not specified in the request are
// left unchanged.
//
// For information about valid key-value pairs in PatchFilters for each supported
// operating system type, see PatchFilter.
func (c *Client) UpdatePatchBaseline(ctx context.Context, params *UpdatePatchBaselineInput, optFns ...func(*Options)) (*UpdatePatchBaselineOutput, error) {
	if params == nil {
		params = &UpdatePatchBaselineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePatchBaseline", params, optFns, c.addOperationUpdatePatchBaselineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePatchBaselineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePatchBaselineInput struct {

	// The ID of the patch baseline to update.
	//
	// This member is required.
	BaselineId *string

	// A set of rules used to include patches in the baseline.
	ApprovalRules *types.PatchRuleGroup

	// A list of explicitly approved patches for the baseline.
	//
	// For information about accepted formats for lists of approved patches and
	// rejected patches, see [Package name formats for approved and rejected patch lists]in the Amazon Web Services Systems Manager User Guide.
	//
	// [Package name formats for approved and rejected patch lists]: https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager-approved-rejected-package-name-formats.html
	ApprovedPatches []string

	// Assigns a new compliance severity level to an existing patch baseline.
	ApprovedPatchesComplianceLevel types.PatchComplianceLevel

	// Indicates whether the list of approved patches includes non-security updates
	// that should be applied to the managed nodes. The default value is false .
	// Applies to Linux managed nodes only.
	ApprovedPatchesEnableNonSecurity *bool

	// Indicates the status to be assigned to security patches that are available but
	// not approved because they don't meet the installation criteria specified in the
	// patch baseline.
	//
	// Example scenario: Security patches that you might want installed can be skipped
	// if you have specified a long period to wait after a patch is released before
	// installation. If an update to the patch is released during your specified
	// waiting period, the waiting period for installing the patch starts over. If the
	// waiting period is too long, multiple versions of the patch could be released but
	// never installed.
	//
	// Supported for Windows Server managed nodes only.
	AvailableSecurityUpdatesComplianceStatus types.PatchComplianceStatus

	// A description of the patch baseline.
	Description *string

	// A set of global filters used to include patches in the baseline.
	//
	// The GlobalFilters parameter can be configured only by using the CLI or an
	// Amazon Web Services SDK. It can't be configured from the Patch Manager console,
	// and its value isn't displayed in the console.
	GlobalFilters *types.PatchFilterGroup

	// The name of the patch baseline.
	Name *string

	// A list of explicitly rejected patches for the baseline.
	//
	// For information about accepted formats for lists of approved patches and
	// rejected patches, see [Package name formats for approved and rejected patch lists]in the Amazon Web Services Systems Manager User Guide.
	//
	// [Package name formats for approved and rejected patch lists]: https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager-approved-rejected-package-name-formats.html
	RejectedPatches []string

	// The action for Patch Manager to take on patches included in the RejectedPackages
	// list.
	//
	// ALLOW_AS_DEPENDENCY  Linux and macOS: A package in the rejected patches list is
	// installed only if it is a dependency of another package. It is considered
	// compliant with the patch baseline, and its status is reported as INSTALLED_OTHER
	// . This is the default action if no option is specified.
	//
	// Windows Server: Windows Server doesn't support the concept of package
	// dependencies. If a package in the rejected patches list and already installed on
	// the node, its status is reported as INSTALLED_OTHER . Any package not already
	// installed on the node is skipped. This is the default action if no option is
	// specified.
	//
	// BLOCK  All OSs: Packages in the rejected patches list, and packages that
	// include them as dependencies, aren't installed by Patch Manager under any
	// circumstances. If a package was installed before it was added to the rejected
	// patches list, or is installed outside of Patch Manager afterward, it's
	// considered noncompliant with the patch baseline and its status is reported as
	// INSTALLED_REJECTED .
	RejectedPatchesAction types.PatchAction

	// If True, then all fields that are required by the CreatePatchBaseline operation are also required
	// for this API request. Optional fields that aren't specified are set to null.
	Replace *bool

	// Information about the patches to use to update the managed nodes, including
	// target operating systems and source repositories. Applies to Linux managed nodes
	// only.
	Sources []types.PatchSource

	noSmithyDocumentSerde
}

type UpdatePatchBaselineOutput struct {

	// A set of rules used to include patches in the baseline.
	ApprovalRules *types.PatchRuleGroup

	// A list of explicitly approved patches for the baseline.
	ApprovedPatches []string

	// The compliance severity level assigned to the patch baseline after the update
	// completed.
	ApprovedPatchesComplianceLevel types.PatchComplianceLevel

	// Indicates whether the list of approved patches includes non-security updates
	// that should be applied to the managed nodes. The default value is false .
	// Applies to Linux managed nodes only.
	ApprovedPatchesEnableNonSecurity *bool

	// Indicates the compliance status of managed nodes for which security-related
	// patches are available but were not approved. This preference is specified when
	// the CreatePatchBaseline or UpdatePatchBaseline commands are run.
	//
	// Applies to Windows Server managed nodes only.
	AvailableSecurityUpdatesComplianceStatus types.PatchComplianceStatus

	// The ID of the deleted patch baseline.
	BaselineId *string

	// The date when the patch baseline was created.
	CreatedDate *time.Time

	// A description of the patch baseline.
	Description *string

	// A set of global filters used to exclude patches from the baseline.
	GlobalFilters *types.PatchFilterGroup

	// The date when the patch baseline was last modified.
	ModifiedDate *time.Time

	// The name of the patch baseline.
	Name *string

	// The operating system rule used by the updated patch baseline.
	OperatingSystem types.OperatingSystem

	// A list of explicitly rejected patches for the baseline.
	RejectedPatches []string

	// The action specified to take on patches included in the RejectedPatches list. A
	// patch can be allowed only if it is a dependency of another package, or blocked
	// entirely along with packages that include it as a dependency.
	RejectedPatchesAction types.PatchAction

	// Information about the patches to use to update the managed nodes, including
	// target operating systems and source repositories. Applies to Linux managed nodes
	// only.
	Sources []types.PatchSource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePatchBaselineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdatePatchBaseline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdatePatchBaseline{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdatePatchBaseline"); err != nil {
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
	if err = addOpUpdatePatchBaselineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePatchBaseline(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePatchBaseline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdatePatchBaseline",
	}
}
