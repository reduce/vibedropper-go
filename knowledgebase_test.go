// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vibedropper_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/vibedropper-go"
	"github.com/stainless-sdks/vibedropper-go/internal/testutil"
	"github.com/stainless-sdks/vibedropper-go/option"
)

func TestKnowledgeBaseGet(t *testing.T) {
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
	_, err := client.KnowledgeBases.Get(context.TODO(), "kbId")
	if err != nil {
		var apierr *vibedropper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestKnowledgeBaseUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.KnowledgeBases.Update(
		context.TODO(),
		"kbId",
		vibedropper.KnowledgeBaseUpdateParams{
			Description: vibedropper.String("description"),
			Name:        vibedropper.String("name"),
			SortOrder:   vibedropper.Int(0),
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

func TestKnowledgeBaseList(t *testing.T) {
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
	_, err := client.KnowledgeBases.List(context.TODO())
	if err != nil {
		var apierr *vibedropper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestKnowledgeBaseDelete(t *testing.T) {
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
	err := client.KnowledgeBases.Delete(context.TODO(), "kbId")
	if err != nil {
		var apierr *vibedropper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
