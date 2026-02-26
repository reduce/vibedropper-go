// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vibedropper

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/vibedropper-go/internal/apijson"
	"github.com/stainless-sdks/vibedropper-go/internal/apiquery"
	"github.com/stainless-sdks/vibedropper-go/internal/requestconfig"
	"github.com/stainless-sdks/vibedropper-go/option"
	"github.com/stainless-sdks/vibedropper-go/packages/param"
	"github.com/stainless-sdks/vibedropper-go/packages/respjson"
)

// CustomerService contains methods and other services that help with interacting
// with the vibedropper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerService] method instead.
type CustomerService struct {
	Options []option.RequestOption
}

// NewCustomerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerService(opts ...option.RequestOption) (r CustomerService) {
	r = CustomerService{}
	r.Options = opts
	return
}

// Get customer
func (r *CustomerService) Get(ctx context.Context, customerID string, opts ...option.RequestOption) (res *CustomerGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if customerID == "" {
		err = errors.New("missing required customerId parameter")
		return
	}
	path := fmt.Sprintf("customers/%s", url.PathEscape(customerID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update customer
func (r *CustomerService) Update(ctx context.Context, customerID string, body CustomerUpdateParams, opts ...option.RequestOption) (res *CustomerUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if customerID == "" {
		err = errors.New("missing required customerId parameter")
		return
	}
	path := fmt.Sprintf("customers/%s", url.PathEscape(customerID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List customers
func (r *CustomerService) List(ctx context.Context, query CustomerListParams, opts ...option.RequestOption) (res *CustomerListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Customer struct {
	ID                string    `json:"id"`
	AverageOrderValue float64   `json:"averageOrderValue"`
	Email             string    `json:"email"`
	LastPurchaseDate  time.Time `json:"lastPurchaseDate" api:"nullable" format:"date-time"`
	Lists             []any     `json:"lists"`
	Name              string    `json:"name" api:"nullable"`
	OrgID             string    `json:"orgId"`
	PickupLocation    any       `json:"pickupLocation" api:"nullable"`
	PurchaseCount     int64     `json:"purchaseCount"`
	Region            any       `json:"region" api:"nullable"`
	Roles             []any     `json:"roles"`
	TotalSpent        float64   `json:"totalSpent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AverageOrderValue respjson.Field
		Email             respjson.Field
		LastPurchaseDate  respjson.Field
		Lists             respjson.Field
		Name              respjson.Field
		OrgID             respjson.Field
		PickupLocation    respjson.Field
		PurchaseCount     respjson.Field
		Region            respjson.Field
		Roles             respjson.Field
		TotalSpent        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Customer) RawJSON() string { return r.JSON.raw }
func (r *Customer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerGetResponse struct {
	Customer Customer `json:"customer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Customer    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerGetResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomerGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerUpdateResponse struct {
	Customer Customer `json:"customer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Customer    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomerUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerListResponse struct {
	Customers  []Customer                     `json:"customers"`
	Pagination CustomerListResponsePagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Customers   respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerListResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomerListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerListResponsePagination struct {
	Limit      int64 `json:"limit"`
	Page       int64 `json:"page"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"totalPages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Page        respjson.Field
		Total       respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *CustomerListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerUpdateParams struct {
	PickupLocationID param.Opt[string] `json:"pickupLocationId,omitzero"`
	RegionID         param.Opt[string] `json:"regionId,omitzero"`
	Name             param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r CustomerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerListParams struct {
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Page   param.Opt[int64]  `query:"page,omitzero" json:"-"`
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CustomerListParams]'s query parameters as `url.Values`.
func (r CustomerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
