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
	"github.com/stainless-sdks/vibedropper-go/internal/requestconfig"
	"github.com/stainless-sdks/vibedropper-go/option"
	"github.com/stainless-sdks/vibedropper-go/packages/param"
	"github.com/stainless-sdks/vibedropper-go/packages/respjson"
)

// KnowledgeBaseService contains methods and other services that help with
// interacting with the vibedropper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKnowledgeBaseService] method instead.
type KnowledgeBaseService struct {
	Options  []option.RequestOption
	Articles KnowledgeBaseArticleService
}

// NewKnowledgeBaseService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewKnowledgeBaseService(opts ...option.RequestOption) (r KnowledgeBaseService) {
	r = KnowledgeBaseService{}
	r.Options = opts
	r.Articles = NewKnowledgeBaseArticleService(opts...)
	return
}

// Get a knowledge base
func (r *KnowledgeBaseService) Get(ctx context.Context, kbID string, opts ...option.RequestOption) (res *KnowledgeBaseGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if kbID == "" {
		err = errors.New("missing required kbId parameter")
		return
	}
	path := fmt.Sprintf("knowledge-bases/%s", url.PathEscape(kbID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a knowledge base
func (r *KnowledgeBaseService) Update(ctx context.Context, kbID string, body KnowledgeBaseUpdateParams, opts ...option.RequestOption) (res *KnowledgeBaseUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if kbID == "" {
		err = errors.New("missing required kbId parameter")
		return
	}
	path := fmt.Sprintf("knowledge-bases/%s", url.PathEscape(kbID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List knowledge bases
func (r *KnowledgeBaseService) List(ctx context.Context, opts ...option.RequestOption) (res *KnowledgeBaseListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "knowledge-bases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a knowledge base
func (r *KnowledgeBaseService) Delete(ctx context.Context, kbID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if kbID == "" {
		err = errors.New("missing required kbId parameter")
		return
	}
	path := fmt.Sprintf("knowledge-bases/%s", url.PathEscape(kbID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type KnowledgeBase struct {
	ID          string              `json:"id"`
	Count       KnowledgeBase_Count `json:"_count"`
	CreatedAt   time.Time           `json:"createdAt" format:"date-time"`
	Description string              `json:"description" api:"nullable"`
	Name        string              `json:"name"`
	OrgID       string              `json:"orgId"`
	Slug        string              `json:"slug"`
	SortOrder   int64               `json:"sortOrder"`
	UpdatedAt   time.Time           `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Count       respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Name        respjson.Field
		OrgID       respjson.Field
		Slug        respjson.Field
		SortOrder   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KnowledgeBase) RawJSON() string { return r.JSON.raw }
func (r *KnowledgeBase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KnowledgeBase_Count struct {
	Articles   int64 `json:"articles"`
	Categories int64 `json:"categories"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Articles    respjson.Field
		Categories  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KnowledgeBase_Count) RawJSON() string { return r.JSON.raw }
func (r *KnowledgeBase_Count) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KnowledgeBaseGetResponse struct {
	KnowledgeBase KnowledgeBase `json:"knowledgeBase"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		KnowledgeBase respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KnowledgeBaseGetResponse) RawJSON() string { return r.JSON.raw }
func (r *KnowledgeBaseGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KnowledgeBaseUpdateResponse struct {
	KnowledgeBase KnowledgeBase `json:"knowledgeBase"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		KnowledgeBase respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KnowledgeBaseUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *KnowledgeBaseUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KnowledgeBaseListResponse struct {
	KnowledgeBases []KnowledgeBase `json:"knowledgeBases"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		KnowledgeBases respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KnowledgeBaseListResponse) RawJSON() string { return r.JSON.raw }
func (r *KnowledgeBaseListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KnowledgeBaseUpdateParams struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Name        param.Opt[string] `json:"name,omitzero"`
	SortOrder   param.Opt[int64]  `json:"sortOrder,omitzero"`
	paramObj
}

func (r KnowledgeBaseUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow KnowledgeBaseUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KnowledgeBaseUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
