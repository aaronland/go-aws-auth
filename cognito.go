package auth

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/sts/types"
)

type STSCredentialsForDeveloperIdentityOptions struct {
	IdentityPoolId  string
	Logins          map[string]string
	RoleArn         string
	RoleSessionName string
	Duration        int32
}

func STSCredentialsForDeveloperIdentity(ctx context.Context, aws_cfg aws.Config, opts *STSCredentialsForDeveloperIdentityOptions) (*types.Credentials, error) {

	cognito_client := cognitoidentity.NewFromConfig(aws_cfg)
	sts_client := sts.NewFromConfig(aws_cfg)

	// Get temporary OpenID token from Cognito

	token_opts := &cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput{
		IdentityPoolId: aws.String(opts.IdentityPoolId),
		Logins:         opts.Logins,
	}

	token_rsp, err := cognito_client.GetOpenIdTokenForDeveloperIdentity(ctx, token_opts)

	if err != nil {
		return nil, fmt.Errorf("Failed to derive token for developer identity, %w", err)
	}

	// Get temporary credentials from STS

	creds_opts := &sts.AssumeRoleWithWebIdentityInput{
		RoleArn:          aws.String(opts.RoleArn),
		RoleSessionName:  aws.String(opts.RoleSessionName),
		WebIdentityToken: token_rsp.Token,
		DurationSeconds:  aws.Int32(opts.Duration),
	}

	creds_rsp, err := sts_client.AssumeRoleWithWebIdentity(ctx, creds_opts)

	if err != nil {
		return nil, fmt.Errorf("Failed to assume role, %w", err)
	}

	return creds_rsp.Credentials, nil
}
