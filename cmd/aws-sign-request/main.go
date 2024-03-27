package main

import (
	"context"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aaronland/go-aws-auth"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

func main() {

	var api_signing_name string
	var api_signing_region string

	var credentials_uri string
	var method string
	var uri string

	flag.StringVar(&api_signing_name, "api-signing-name", "", "")
	flag.StringVar(&api_signing_region, "api-signing-region", "", "")

	flag.StringVar(&method, "method", "GET", "")
	flag.StringVar(&uri, "uri", "", "")
	flag.StringVar(&credentials_uri, "credentials-uri", "", "")

	flag.Parse()

	ctx := context.Background()

	body_r := strings.NewReader(strings.Join(flag.Args(), " "))

	//

	cfg, err := auth.NewConfig(ctx, credentials_uri)

	if err != nil {
		log.Fatalf("Failed to create new config, %v", err)
	}

	creds, err := cfg.Credentials.Retrieve(ctx)

	if err != nil {
		log.Fatalf("Failed to derive credentials from config, %v", err)
	}

	//

	req, err := http.NewRequest(method, uri, body_r)

	if err != nil {
		log.Fatalf("Failed to create new HTTP request, %v", err)
	}

	// https://github.com/aws/aws-sdk-go-v2/blob/main/aws/signer/v4/v4.go#L287

	body_sha256 := ""

	if body_r.Len() > 0 {

		h := sha256.New()

		_, err := io.Copy(h, body_r)

		if err != nil {
			log.Fatalf("Failed to hash request body, %v", err)
		}

		body_sha256 = fmt.Sprintf("%x", h.Sum(nil))
	}

	signer := v4.NewSigner()

	err = signer.SignHTTP(ctx, creds, req, body_sha256, api_signing_name, api_signing_region, time.Now())

	if err != nil {
		log.Fatalf("Failed to sign request, %v", err)
	}

	err = req.Write(os.Stdout)

	if err != nil {
		log.Fatalf("Failed to write request, %v", err)
	}

	cl := http.Client{}
	rsp, err := cl.Do(req)

	if err != nil {
		log.Fatalf("Failed to execute request, %v", err)
	}

	defer rsp.Body.Close()

	_, err = io.Copy(os.Stdout, rsp.Body)

	if err != nil {
		log.Fatalf("Failed to read response, %v", err)
	}

}
