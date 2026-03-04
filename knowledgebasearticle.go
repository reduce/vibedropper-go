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

// KnowledgeBaseArticleService contains methods and other services that help with
// interacting with the vibedropper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKnowledgeBaseArticleService] method instead.
type KnowledgeBaseArticleService struct {
	Options []option.RequestOption
}

// NewKnowledgeBaseArticleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewKnowledgeBaseArticleService(opts ...option.RequestOption) (r KnowledgeBaseArticleService) {
	r = KnowledgeBaseArticleService{}
	r.Options = opts
	return
}

// Create an article
func (r *KnowledgeBaseArticleService) New(ctx context.Context, kbID string, body KnowledgeBaseArticleNewParams, opts ...option.RequestOption) (res *KnowledgeBaseArticleNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if kbID == "" {
		err = errors.New("missing required kbId parameter")
		return
	}
	path := fmt.Sprintf("knowledge-bases/%s/articles", url.PathEscape(kbID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List articles in a knowledge base
func (r *KnowledgeBaseArticleService) List(ctx context.Context, kbID string, query KnowledgeBaseArticleListParams, opts ...option.RequestOption) (res *KnowledgeBaseArticleListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if kbID == "" {
		err = errors.New("missing required kbId parameter")
		return
	}
	path := fmt.Sprintf("knowledge-bases/%s/articles", url.PathEscape(kbID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type KnowledgeBaseArticle struct {
	ID         string                       `json:"id"`
	Category   KnowledgeBaseArticleCategory `json:"category" api:"nullable"`
	CategoryID string                       `json:"categoryId" api:"nullable"`
	// HTML content
	Content         string    `json:"content"`
	CreatedAt       time.Time `json:"createdAt" format:"date-time"`
	Excerpt         string    `json:"excerpt" api:"nullable"`
	KnowledgeBaseID string    `json:"knowledgeBaseId"`
	Published       bool      `json:"published"`
	Slug            string    `json:"slug"`
	SortOrder       int64     `json:"sortOrder"`
	Title           string    `json:"title"`
	UpdatedAt       time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Category        respjson.Field
		CategoryID      respjson.Field
		Content         respjson.Field
		CreatedAt       respjson.Field
		Excerpt         respjson.Field
		KnowledgeBaseID respjson.Field
		Published       respjson.Field
		Slug            respjson.Field
		SortOrder       respjson.Field
		Title           respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KnowledgeBaseArticle) RawJSON() string { return r.JSON.raw }
func (r *KnowledgeBaseArticle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KnowledgeBaseArticleCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KnowledgeBaseArticleCategory) RawJSON() string { return r.JSON.raw }
func (r *KnowledgeBaseArticleCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KnowledgeBaseArticleNewResponse struct {
	Article KnowledgeBaseArticle `json:"article"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Article     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KnowledgeBaseArticleNewResponse) RawJSON() string { return r.JSON.raw }
func (r *KnowledgeBaseArticleNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KnowledgeBaseArticleListResponse struct {
	Articles   []KnowledgeBaseArticle `json:"articles"`
	Pagination Pagination             `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Articles    respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KnowledgeBaseArticleListResponse) RawJSON() string { return r.JSON.raw }
func (r *KnowledgeBaseArticleListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KnowledgeBaseArticleNewParams struct {
	Title      string            `json:"title" api:"required"`
	CategoryID param.Opt[string] `json:"categoryId,omitzero"`
	Excerpt    param.Opt[string] `json:"excerpt,omitzero"`
	// HTML content
	Content   param.Opt[string] `json:"content,omitzero"`
	Published param.Opt[bool]   `json:"published,omitzero"`
	paramObj
}

func (r KnowledgeBaseArticleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow KnowledgeBaseArticleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KnowledgeBaseArticleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KnowledgeBaseArticleListParams struct {
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Page  param.Opt[int64] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [KnowledgeBaseArticleListParams]'s query parameters as
// `url.Values`.
func (r KnowledgeBaseArticleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
