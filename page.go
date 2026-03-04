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

	"github.com/reduce/vibedropper-go/internal/apijson"
	"github.com/reduce/vibedropper-go/internal/apiquery"
	"github.com/reduce/vibedropper-go/internal/requestconfig"
	"github.com/reduce/vibedropper-go/option"
	"github.com/reduce/vibedropper-go/packages/param"
	"github.com/reduce/vibedropper-go/packages/respjson"
)

// Manage landing pages
//
// PageService contains methods and other services that help with interacting with
// the vibedropper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageService] method instead.
type PageService struct {
	Options []option.RequestOption
}

// NewPageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPageService(opts ...option.RequestOption) (r PageService) {
	r = PageService{}
	r.Options = opts
	return
}

// Get a page
func (r *PageService) Get(ctx context.Context, pageID string, opts ...option.RequestOption) (res *PageGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if pageID == "" {
		err = errors.New("missing required pageId parameter")
		return
	}
	path := fmt.Sprintf("pages/%s", url.PathEscape(pageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a page
func (r *PageService) Update(ctx context.Context, pageID string, body PageUpdateParams, opts ...option.RequestOption) (res *PageUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if pageID == "" {
		err = errors.New("missing required pageId parameter")
		return
	}
	path := fmt.Sprintf("pages/%s", url.PathEscape(pageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List pages
func (r *PageService) List(ctx context.Context, query PageListParams, opts ...option.RequestOption) (res *PageListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a page
func (r *PageService) Delete(ctx context.Context, pageID string, opts ...option.RequestOption) (res *PageDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if pageID == "" {
		err = errors.New("missing required pageId parameter")
		return
	}
	path := fmt.Sprintf("pages/%s", url.PathEscape(pageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Page struct {
	ID          string     `json:"id"`
	Count       Page_Count `json:"_count"`
	CreatedAt   time.Time  `json:"createdAt" format:"date-time"`
	Description string     `json:"description" api:"nullable"`
	Name        string     `json:"name"`
	OrgID       string     `json:"orgId"`
	Slug        string     `json:"slug"`
	// Any of "DRAFT", "ACTIVE", "ENDED", "ARCHIVED".
	Status    PageStatus `json:"status"`
	UpdatedAt time.Time  `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Count       respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Name        respjson.Field
		OrgID       respjson.Field
		Slug        respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Page) RawJSON() string { return r.JSON.raw }
func (r *Page) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Page_Count struct {
	Items int64 `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Page_Count) RawJSON() string { return r.JSON.raw }
func (r *Page_Count) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageStatus string

const (
	PageStatusDraft    PageStatus = "DRAFT"
	PageStatusActive   PageStatus = "ACTIVE"
	PageStatusEnded    PageStatus = "ENDED"
	PageStatusArchived PageStatus = "ARCHIVED"
)

type PageGetResponse struct {
	Page Page `json:"page"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageUpdateResponse struct {
	Page Page `json:"page"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PageUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *PageUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageListResponse struct {
	Pages      []Page     `json:"pages"`
	Pagination Pagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pages       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PageListResponse) RawJSON() string { return r.JSON.raw }
func (r *PageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageDeleteResponse struct {
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PageDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *PageDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageUpdateParams struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Name        param.Opt[string] `json:"name,omitzero"`
	// Any of "DRAFT", "ACTIVE", "ENDED", "ARCHIVED".
	Status PageUpdateParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r PageUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PageUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PageUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageUpdateParamsStatus string

const (
	PageUpdateParamsStatusDraft    PageUpdateParamsStatus = "DRAFT"
	PageUpdateParamsStatusActive   PageUpdateParamsStatus = "ACTIVE"
	PageUpdateParamsStatusEnded    PageUpdateParamsStatus = "ENDED"
	PageUpdateParamsStatusArchived PageUpdateParamsStatus = "ARCHIVED"
)

type PageListParams struct {
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Page  param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Filter by status. Omit or use "all" to return all pages.
	//
	// Any of "DRAFT", "ACTIVE", "ENDED", "ARCHIVED", "all".
	Status PageListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PageListParams]'s query parameters as `url.Values`.
func (r PageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by status. Omit or use "all" to return all pages.
type PageListParamsStatus string

const (
	PageListParamsStatusDraft    PageListParamsStatus = "DRAFT"
	PageListParamsStatusActive   PageListParamsStatus = "ACTIVE"
	PageListParamsStatusEnded    PageListParamsStatus = "ENDED"
	PageListParamsStatusArchived PageListParamsStatus = "ARCHIVED"
	PageListParamsStatusAll      PageListParamsStatus = "all"
)
