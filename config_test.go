package auth

import (
	"context"
	"testing"
	_ "fmt"
)

func TestNewConfigWithCredentials(t *testing.T) {

	creds := []string{
		"anon:",
		"env:",
		"iam:",
		"static:key:secret:",
		"static:key:secret:token",
		"fixtures/credentials:default",
		"fixtures/credentials:example",
		"default",
	}

	ctx := context.Background()

	for _, str := range creds {

		cfg, err := NewConfigWithCredentialsString(ctx, str)

		if err != nil {
			t.Fatalf("Unable to create config with credentials '%s', %v", str, err)
		}
	}
}
