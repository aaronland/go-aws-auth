package main

import (
	"context"
	"log"
	
	"github.com/aaronland/go-aws-auth"
	"github.com/sfomuseum/go-flags/flagset"
)

func main() {

	var aws_credentials_uri string

	fs := flagset.NewFlagSet("aws")

	fs.StringVar(&aws_credentials_uri, "aws-credentials-uri", "", "...")

	flagset.Parse(fs)

	ctx := context.Background()
	
	aws_cfg, err := auth.NewConfig(ctx, aws_credential_uri)
	
	if err != nil {
		log.Fatalf("Failed to create new AWS config, %w", err)
	}

	aws_sess, err := session.NewSessionWithOptions(aws_cfg)

	if err != nil {
		log.Fatalf("Failed to create session, %v", err)
	}
	
	v4.NewSigner(s.session.Config.Credentials)
}
