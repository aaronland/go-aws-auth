package main

import (
	"bufio"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"text/template"

	"github.com/aws/aws-sdk-go-v2/service/sts/types"
)

//go:embed credentials.ini
var credentials_t string

type CredentialsVars struct {
	Name         string
	Region       string
	KeyId        string
	KeySecret    string
	SessionToken string
}

func main() {

	var infile string
	var outfile string

	var name string
	var region string

	flag.StringVar(&infile, "json", "", "...")
	flag.StringVar(&outfile, "ini", "", "...")

	flag.StringVar(&name, "name", "default", "...")
	flag.StringVar(&region, "region", "us-east-1", "...")

	flag.Parse()

	var r io.ReadCloser
	var wr io.WriteCloser

	switch infile {
	case "-":
		br := bufio.NewReader(os.Stdin)
		r = io.NopCloser(br)
	default:

		_r, err := os.Open(infile)

		if err != nil {
			log.Fatalf("Failed to open %s for reading, %v", infile, err)
		}

		r = _r
	}

	defer r.Close()

	switch outfile {
	case "-":
		wr = os.Stdout
	default:
		_wr, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0600)

		if err != nil {
			log.Fatalf("Failed to open %s for writing, %v", outfile, err)
		}

		wr = _wr
	}

	err := Convert(r, wr, name, region)

	if err != nil {
		log.Fatalf("Failed to convert credentials, %v", err)
	}

	err = wr.Close()

	if err != nil {
		log.Fatalf("Failed to close %s after writing, %v", outfile, err)
	}

}

func Convert(r io.Reader, wr io.Writer, name string, region string) error {

	t, err := template.New("credentials").Parse(credentials_t)

	if err != nil {
		return fmt.Errorf("Failed to parse credentials template, %w", err)
	}

	var creds *types.Credentials

	dec := json.NewDecoder(r)
	err = dec.Decode(&creds)

	if err != nil {
		return fmt.Errorf("Failed to decode credentials reader, %w", err)
	}

	vars := CredentialsVars{
		Name:         name,
		Region:       region,
		KeyId:        *creds.AccessKeyId,
		KeySecret:    *creds.SecretAccessKey,
		SessionToken: *creds.SessionToken,
	}

	err = t.Execute(wr, vars)

	if err != nil {
		return fmt.Errorf("Failed to write credentials template, %w", err)
	}

	return nil
}
