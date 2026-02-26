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

func TestListSubscriberList(t *testing.T) {
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
	_, err := client.Lists.Subscribers.List(context.TODO(), "listId")
	if err != nil {
		var apierr *vibedropper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestListSubscriberAddWithOptionalParams(t *testing.T) {
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
	_, err := client.Lists.Subscribers.Add(
		context.TODO(),
		"listId",
		vibedropper.ListSubscriberAddParams{
			Email:            "dev@stainless.com",
			CustomFields:     map[string]any{},
			Name:             vibedropper.String("name"),
			PickupLocationID: vibedropper.String("pickupLocationId"),
			RegionID:         vibedropper.String("regionId"),
		},
	)
	if err != nil {
		var apierr *vibedropper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestListSubscriberRemove(t *testing.T) {
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
	_, err := client.Lists.Subscribers.Remove(
		context.TODO(),
		"subscriberId",
		vibedropper.ListSubscriberRemoveParams{
			ListID: "listId",
		},
	)
	if err != nil {
		var apierr *vibedropper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
