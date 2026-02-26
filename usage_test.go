// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vibedropper_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/vibedropper-go"
	"github.com/stainless-sdks/vibedropper-go/internal/testutil"
	"github.com/stainless-sdks/vibedropper-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := vibedropper.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	t.Skip("Mock server tests are disabled")
	lists, err := client.Lists.List(context.TODO(), vibedropper.ListListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", lists.Lists)
}
