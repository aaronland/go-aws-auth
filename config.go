package auth

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"path/filepath"
	"strings"
	"os/user"
)

func NewConfigWithCredentials(ctx context.Context, str_creds string) (*config.Config, error) {

	cfg, err := config.LoadDefaultConfig(ctx)

	if err != nil {
		return nil, fmt.Errorf("Failed to load default config, %w", err)
	}

	if strings.HasPrefix(str_creds, "anon:") {

		creds := credentials.AnonymousCredentials
		cfg.WithCredentials(creds)

	} else if strings.HasPrefix(str_creds, "env:") {

		creds := credentials.NewEnvCredentials()
		cfg.WithCredentials(creds)

	} else if strings.HasPrefix(str_creds, "iam:") {

		// assume an IAM role suffient for doing whatever

	} else if strings.HasPrefix(str_creds, "static:") {

		details := strings.Split(str_creds, ":")

		if len(details) != 4 {
			return nil, fmt.Errorf("Expected (4) components for 'static:' credentials URI but got %d", len(details))
		}

		id := details[1]
		key := details[2]
		secret := details[3]

		creds := credentials.NewStaticCredentials(id, key, secret)
		cfg.WithCredentials(creds)

	} else if str_creds != "" {

		details := strings.Split(str_creds, ":")

		var creds_file string
		var profile string

		if len(details) == 1 {

			whoami, err := user.Current()

			if err != nil {
				return nil, err
			}

			dotaws := filepath.Join(whoami.HomeDir, ".aws")
			creds_file = filepath.Join(dotaws, "credentials")

			profile = details[0]

		} else {

			path, err := filepath.Abs(details[0])

			if err != nil {
				return nil, err
			}

			creds_file = path
			profile = details[1]
		}

		creds := credentials.NewSharedCredentials(creds_file, profile)
		cfg.WithCredentials(creds)

	} else {

		// for backwards compatibility as of 05a6042dc5956c13513bdc5ab4969877013f795c
		// (20161203/thisisaaronland)

		creds := credentials.NewEnvCredentials()
		cfg.WithCredentials(creds)
	}

	return cfg, nil
}

/*

func NewConfig(uri string) (*config.Config, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URL, %w", err)
	}

	q := u.Query()

	creds := q.Get("credentials")
	region := q.Get("region")

	if creds == "" {
		return nil, fmt.Errorf("Missing ?credentials parameter")
	}

	if region == "" {
		return nil, fmt.Errorf("Missing ?region parameter")
	}

	return NewSessionWithCredentials(creds, region)
}

func NewSessionWithDSN(dsn_str string) (*config.Config, error) {

	dsn_map, err := dsn.StringToDSNWithKeys(dsn_str, "credentials", "region")

	if err != nil {
		return nil, err
	}

	return NewSessionWithCredentials(dsn_map["credentials"], dsn_map["region"])
}

func NewSessionWithCredentials(str_creds string, region string) (*config.Config, error) {

	cfg, err := NewConfigWithCredentialsAndRegion(str_creds, region)

	if err != nil {
		return nil, err
	}

	sess := aws_session.New(cfg)

	_, err = sess.Config.Credentials.Get()

	if err != nil {
		return nil, err
	}

	return sess, nil
}

*/
