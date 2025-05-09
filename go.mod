module github.com/aaronland/go-aws-auth/v2

go 1.24.0

require (
	github.com/aws/aws-sdk-go-v2 v1.36.3
	github.com/aws/aws-sdk-go-v2/config v1.29.11
	github.com/aws/aws-sdk-go-v2/credentials v1.17.64
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.30
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.29.3
	github.com/aws/aws-sdk-go-v2/service/iam v1.40.2
	github.com/aws/aws-sdk-go-v2/service/ssm v1.58.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.33.17
	github.com/aws/smithy-go v1.22.3
	github.com/go-ini/ini v1.67.0
	github.com/sfomuseum/go-flags v0.10.0
	github.com/sfomuseum/iso8601duration v1.1.0
)

require (
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.34 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.34 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.12.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.12.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.25.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.29.2 // indirect
)
