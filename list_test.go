// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vibedropper_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/reduce/vibedropper-go"
	"github.com/reduce/vibedropper-go/internal/testutil"
	"github.com/reduce/vibedropper-go/option"
)

func TestListGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Lists.Get(context.TODO(), "listId")
	if err != nil {
		var apierr *vibedropper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestListListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Lists.List(context.TODO(), vibedropper.ListListParams{
		Limit: vibedropper.Int(0),
		Page:  vibedropper.Int(0),
	})
	if err != nil {
		var apierr *vibedropper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
