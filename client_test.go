package dizzy

import (
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/SomethingBot/dizzy/discordwebapi"
	"github.com/SomethingBot/dizzy/logging"
	"github.com/SomethingBot/dizzy/primitives"
)

func TestClient_OpenClose(t *testing.T) {
	if testing.Short() {
		t.Skipf("test short flag set, skipping integration tests")
	}

	var apikey string
	apikey = os.Getenv("discordapikey")
	if apikey == "" {
		apikeyBytes, err := os.ReadFile("apikeyfile")
		if err != nil {
			t.Fatalf("error on reading apikeyfile (%v)\n", err)
		}
		apikey = strings.ReplaceAll(string(apikeyBytes), "\n", "")
	}

	gatewayURI, err := discordwebapi.GetGatewayWebsocketURI(func() error { return nil }, "")
	if err != nil {
		t.Fatal(err)
	}

	client := NewClient(apikey, gatewayURI.String(), primitives.GatewayIntentAll, NewEventDistributor(), &logging.Standard{Logger: *log.Default()})
	err = client.Open()
	if err != nil {
		t.Fatalf("error on open (%v)\n", err)
	}

	time.Sleep(time.Second * 8)

	err = client.Close()
	if err != nil {
		t.Fatalf("error on close (%v)\n", err)
	}
}

//todo: we should test every event we can from discord for proper json parsing
