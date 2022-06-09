package auth

import (
	"fmt"
	"strings"
	"sort"
)

func XCredentials() []string {

	valid := []string{
		"anon:",
		"env:",
		"iam:",
		"{PROFILE}",
		"{PATH}:{PROFILE}",
		"static:{ID}:{KEY}:{SECRET}",
	}

	return valid
}

func XCredentialsString() string {

	creds := XCredentials()
	sort.Strings(creds)
	
	return fmt.Sprintf("Valid credential flags are: %s", strings.Join(creds, ", "))
}

