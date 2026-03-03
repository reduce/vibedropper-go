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

func TestCustomerGet(t *testing.T) {
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
	_, err := client.Customers.Get(context.TODO(), "customerId")
	if err != nil {
		var apierr *vibedropper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Update(
		context.TODO(),
		"customerId",
		vibedropper.CustomerUpdateParams{
			AddressLine1:     vibedropper.String("addressLine1"),
			AddressLine2:     vibedropper.String("addressLine2"),
			City:             vibedropper.String("city"),
			Country:          vibedropper.String("country"),
			FirstName:        vibedropper.String("firstName"),
			LastName:         vibedropper.String("lastName"),
			Name:             vibedropper.String("name"),
			PickupLocationID: vibedropper.String("pickupLocationId"),
			PostalCode:       vibedropper.String("postalCode"),
			RegionID:         vibedropper.String("regionId"),
			State:            vibedropper.String("state"),
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

func TestCustomerListWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.List(context.TODO(), vibedropper.CustomerListParams{
		Limit:  vibedropper.Int(100),
		Page:   vibedropper.Int(0),
		Search: vibedropper.String("search"),
	})
	if err != nil {
		var apierr *vibedropper.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
