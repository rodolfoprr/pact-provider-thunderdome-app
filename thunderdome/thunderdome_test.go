package thunderdome_test

import (
	"fmt"
	"os"
	"pact-provider-thunderdome-app/thunderdome"
	"path/filepath"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

var dir, _ = os.Getwd()
var pactDir = fmt.Sprintf("%s/../pacts", dir)

func TestProvider(t *testing.T) {
	go thunderdome.Start()

	pact := &dsl.Pact{
		Port:     6666,
		Consumer: "SpaceX",
		Provider: "Thunderdome",
	}

	pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://localhost:8000",
		PactURLs:        []string{filepath.ToSlash(fmt.Sprintf("%s/SpaceX-Thunderdome.json", pactDir))},
	})
}
