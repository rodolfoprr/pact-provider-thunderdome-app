package thunderdome_test

import (
	"pact-provider-thunderdome-app/thunderdome"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestProvider(t *testing.T) {
	go thunderdome.Start()

	pact := &dsl.Pact{
		Port:     6666,
		Consumer: "SpaceX",
		Provider: "Thunderdome",
	}

	pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://localhost:8000",
		PactURLs:        []string{"https://raw.githubusercontent.com/rodolfoprr/pact-consumer-spacex-app/master/target/pacts/SpaceX-Thunderdome.json"},
	})
}
