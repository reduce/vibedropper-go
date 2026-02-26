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

// ListService contains methods and other services that help with interacting with
// the vibedropper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewListService] method instead.
type ListService struct {
	Options     []option.RequestOption
	Subscribers ListSubscriberService
}

// NewListService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewListService(opts ...option.RequestOption) (r ListService) {
	r = ListService{}
	r.Options = opts
	r.Subscribers = NewListSubscriberService(opts...)
	return
}

// Get a list
func (r *ListService) Get(ctx context.Context, listID string, opts ...option.RequestOption) (res *ListGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return
	}
	path := fmt.Sprintf("lists/%s", url.PathEscape(listID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List lists
func (r *ListService) List(ctx context.Context, query ListListParams, opts ...option.RequestOption) (res *ListListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "lists"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type List struct {
	ID          string     `json:"id"`
	Count       List_Count `json:"_count"`
	CreatedAt   time.Time  `json:"createdAt" format:"date-time"`
	Description string     `json:"description" api:"nullable"`
	Name        string     `json:"name"`
	OrgID       string     `json:"orgId"`
	UpdatedAt   time.Time  `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Count       respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Name        respjson.Field
		OrgID       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r List) RawJSON() string { return r.JSON.raw }
func (r *List) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type List_Count struct {
	Subscribers int64 `json:"subscribers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Subscribers respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r List_Count) RawJSON() string { return r.JSON.raw }
func (r *List_Count) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListGetResponse struct {
	List List `json:"list"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		List        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ListGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListResponse struct {
	Lists      []List                     `json:"lists"`
	Pagination ListListResponsePagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lists       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListListResponse) RawJSON() string { return r.JSON.raw }
func (r *ListListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListResponsePagination struct {
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
func (r ListListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ListListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListParams struct {
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Page  param.Opt[int64] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListListParams]'s query parameters as `url.Values`.
func (r ListListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
