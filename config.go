package auth

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"os/user"
	"path/filepath"
	"strings"
)

func CredentialsStrings() []string {

	valid := []string{
		"anon:",
		"env:",
		"iam:",
		"{PROFILE}",
		"{PATH}:{PROFILE}",
		"static:{KEY}:{SECRET}:{TOKEN}",
	}

	return valid
}

func NewConfigWithCredentialsString(ctx context.Context, str_creds string) (config.Config, error) {

	if strings.HasPrefix(str_creds, "anon:") {

		provider := aws.AnonymousCredentials{}

		return config.LoadDefaultConfig(ctx,
			config.WithCredentialsProvider(provider),
		)

	} else if strings.HasPrefix(str_creds, "static:") {

		details := strings.Split(str_creds, ":")

		if len(details) != 4 {
			return nil, fmt.Errorf("Expected (4) components for 'static:' credentials URI but got %d", len(details))
		}

		key := details[1]
		secret := details[2]
		token := details[3]

		provider := credentials.NewStaticCredentialsProvider(key, secret, token)

		return config.LoadDefaultConfig(ctx,
			config.WithCredentialsProvider(provider),
		)

	} else if str_creds == "iam:" || str_creds == "env:" {

		return config.LoadDefaultConfig(ctx)

	} else if str_creds != "" {

		details := strings.Split(str_creds, ":")

		var creds_file string
		var profile string

		if len(details) == 1 {

			whoami, err := user.Current()

			if err != nil {
				return nil, fmt.Errorf("Failed to determine current user, %w", err)
			}

			dotaws := filepath.Join(whoami.HomeDir, ".aws")
			creds_file = filepath.Join(dotaws, "credentials")

			profile = details[0]

		} else {

			path, err := filepath.Abs(details[0])

			if err != nil {
				return nil, fmt.Errorf("Failed to derive absolute path for %s, %w", details[0], err)
			}

			creds_file = path
			profile = details[1]
		}

		return config.LoadDefaultConfig(ctx,
			config.WithSharedCredentialsFiles([]string{creds_file}),
			config.WithSharedConfigProfile(profile),
		)

	} else {

		return nil, fmt.Errorf("Invalid or unsupported credentials string")
	}

}
