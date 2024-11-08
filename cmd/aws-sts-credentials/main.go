package main

// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/sts#Client.AssumeRole
// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/sts#AssumeRoleInput
// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/sts#AssumeRoleOutput

import (
	"context"
	"flag"
	"log"

	"github.com/aaronland/go-aws-auth"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func main() {

	var config_uri string

	var role_arn string
	var role_session string
	var role_duration int // update to use ISO8601 duration string

	var mfa_require bool
	var mfa_serial string
	var mfa_token string

	flag.StringVar(&config_uri, "config-uri", "", "...")

	flag.StringVar(&role_arn, "role-arn", "", "The AWS role ARN URI of the role you want to assume.")
	flag.StringVar(&role_session, "role-session", "", "...")
	flag.IntVar(&role_duration, "role-duration", 3600, "The duration, in seconds, of the role session.")

	flag.BoolVar(&mfa_require, "mfa_require", true, "...")
	flag.StringVar(&mfa_serial, "mfa-serial-number", "", "...")
	flag.StringVar(&mfa_token, "mfa-token", "", "...")

	flag.Parse()

	ctx := context.Background()

	cfg, err := auth.NewConfig(ctx, config_uri)

	if err != nil {
		log.Fatalf("Failed to create new config, %v", err)
	}

	cl := sts.NewFromConfig(cfg)

	assume_opts := &sts.AssumeRoleInput{
		RoleArn:         aws.String(role_arn),
		RoleSessionName: aws.String(role_session),
		DurationSeconds: aws.Int32(int32(role_duration)),
	}

	if mfa_require {
		assume_opts.SerialNumber = aws.String(mfa_serial)
		assume_opts.TokenCode = aws.String(mfa_token)
	}

	rsp, err := cl.AssumeRole(ctx, assume_opts)

	if err != nil {
		log.Fatalf("Failed to assume role, %v", err)
	}

	log.Println(rsp)
}
