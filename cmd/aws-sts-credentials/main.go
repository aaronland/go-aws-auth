package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/aaronland/go-aws-auth"
)

func main() {

	var aws_config_uri string

	var identity_pool_id string
	var role_arn string
	var role_session_name string
	var duration int

	flag.StringVar(&aws_config_uri, "aws-config-uri", "", "...")

	flag.StringVar(&identity_pool_id, "identity-pool-id", "", "...")
	flag.StringVar(&role_arn, "role-arn", "", "...")
	flag.StringVar(&role_session_name, "role-session-name", "", "...")
	flag.IntVar(&duration, "duration", 300, "...")

	flag.Parse()

	ctx := context.Background()

	cfg, err := auth.NewConfig(ctx, aws_config_uri)

	if err != nil {
		log.Fatalf("Failed to derive AWS config, %v", err)
	}

	logins := make(map[string]string)

	opts := &auth.STSCredentialsForDeveloperIdentityOptions{
		RoleArn:         role_arn,
		RoleSessionName: role_session_name,
		Duration:        int32(duration),
		IdentityPoolId:  identity_pool_id,
		Logins:          logins,
	}

	creds, err := auth.STSCredentialsForDeveloperIdentity(ctx, cfg, opts)

	if err != nil {
		log.Fatalf("Failed to derive credentials, %v", err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(creds)

	if err != nil {
		log.Fatalf("Failed to encode credentials, %v", err)
	}
}
