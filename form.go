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

// FormService contains methods and other services that help with interacting with
// the vibedropper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFormService] method instead.
type FormService struct {
	Options []option.RequestOption
}

// NewFormService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFormService(opts ...option.RequestOption) (r FormService) {
	r = FormService{}
	r.Options = opts
	return
}

// Get a form
func (r *FormService) Get(ctx context.Context, formID string, opts ...option.RequestOption) (res *FormGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if formID == "" {
		err = errors.New("missing required formId parameter")
		return
	}
	path := fmt.Sprintf("forms/%s", url.PathEscape(formID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a form
func (r *FormService) Update(ctx context.Context, formID string, body FormUpdateParams, opts ...option.RequestOption) (res *FormUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if formID == "" {
		err = errors.New("missing required formId parameter")
		return
	}
	path := fmt.Sprintf("forms/%s", url.PathEscape(formID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List forms
func (r *FormService) List(ctx context.Context, query FormListParams, opts ...option.RequestOption) (res *FormListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "forms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a form
func (r *FormService) Delete(ctx context.Context, formID string, opts ...option.RequestOption) (res *FormDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if formID == "" {
		err = errors.New("missing required formId parameter")
		return
	}
	path := fmt.Sprintf("forms/%s", url.PathEscape(formID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// List form submissions
func (r *FormService) ListSubmissions(ctx context.Context, formID string, query FormListSubmissionsParams, opts ...option.RequestOption) (res *FormListSubmissionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if formID == "" {
		err = errors.New("missing required formId parameter")
		return
	}
	path := fmt.Sprintf("forms/%s/submissions", url.PathEscape(formID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Form struct {
	ID          string     `json:"id"`
	Count       Form_Count `json:"_count"`
	CreatedAt   time.Time  `json:"createdAt" format:"date-time"`
	Description string     `json:"description" api:"nullable"`
	ListID      string     `json:"listId" api:"nullable"`
	OrgID       string     `json:"orgId"`
	Slug        string     `json:"slug"`
	// Any of "DRAFT", "ACTIVE", "ARCHIVED".
	Status         FormStatus `json:"status"`
	StoreBlocks    []any      `json:"storeBlocks" api:"nullable"`
	SuccessMessage string     `json:"successMessage" api:"nullable"`
	Title          string     `json:"title"`
	UpdatedAt      time.Time  `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Count          respjson.Field
		CreatedAt      respjson.Field
		Description    respjson.Field
		ListID         respjson.Field
		OrgID          respjson.Field
		Slug           respjson.Field
		Status         respjson.Field
		StoreBlocks    respjson.Field
		SuccessMessage respjson.Field
		Title          respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Form) RawJSON() string { return r.JSON.raw }
func (r *Form) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Form_Count struct {
	FormSubmissions int64 `json:"formSubmissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FormSubmissions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Form_Count) RawJSON() string { return r.JSON.raw }
func (r *Form_Count) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FormStatus string

const (
	FormStatusDraft    FormStatus = "DRAFT"
	FormStatusActive   FormStatus = "ACTIVE"
	FormStatusArchived FormStatus = "ARCHIVED"
)

type FormGetResponse struct {
	Form Form `json:"form"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Form        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FormGetResponse) RawJSON() string { return r.JSON.raw }
func (r *FormGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FormUpdateResponse struct {
	Form Form `json:"form"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Form        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FormUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *FormUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FormListResponse struct {
	Forms      []Form     `json:"forms"`
	Pagination Pagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Forms       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FormListResponse) RawJSON() string { return r.JSON.raw }
func (r *FormListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FormDeleteResponse struct {
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FormDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *FormDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FormListSubmissionsResponse struct {
	Pagination  Pagination                              `json:"pagination"`
	Submissions []FormListSubmissionsResponseSubmission `json:"submissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pagination  respjson.Field
		Submissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FormListSubmissionsResponse) RawJSON() string { return r.JSON.raw }
func (r *FormListSubmissionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FormListSubmissionsResponseSubmission struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Key-value pairs of submitted form fields
	Data       any                                             `json:"data"`
	FormID     string                                          `json:"formId"`
	Subscriber FormListSubmissionsResponseSubmissionSubscriber `json:"subscriber" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Data        respjson.Field
		FormID      respjson.Field
		Subscriber  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FormListSubmissionsResponseSubmission) RawJSON() string { return r.JSON.raw }
func (r *FormListSubmissionsResponseSubmission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FormListSubmissionsResponseSubmissionSubscriber struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FormListSubmissionsResponseSubmissionSubscriber) RawJSON() string { return r.JSON.raw }
func (r *FormListSubmissionsResponseSubmissionSubscriber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FormUpdateParams struct {
	Description param.Opt[string] `json:"description,omitzero"`
	// List to subscribe form submitters to
	ListID         param.Opt[string] `json:"listId,omitzero"`
	SuccessMessage param.Opt[string] `json:"successMessage,omitzero"`
	Title          param.Opt[string] `json:"title,omitzero"`
	// Any of "DRAFT", "ACTIVE", "ARCHIVED".
	Status FormUpdateParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r FormUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FormUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FormUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FormUpdateParamsStatus string

const (
	FormUpdateParamsStatusDraft    FormUpdateParamsStatus = "DRAFT"
	FormUpdateParamsStatusActive   FormUpdateParamsStatus = "ACTIVE"
	FormUpdateParamsStatusArchived FormUpdateParamsStatus = "ARCHIVED"
)

type FormListParams struct {
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Page  param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Filter by status. Omit or use "all" to return all forms.
	//
	// Any of "DRAFT", "ACTIVE", "ARCHIVED", "all".
	Status FormListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FormListParams]'s query parameters as `url.Values`.
func (r FormListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by status. Omit or use "all" to return all forms.
type FormListParamsStatus string

const (
	FormListParamsStatusDraft    FormListParamsStatus = "DRAFT"
	FormListParamsStatusActive   FormListParamsStatus = "ACTIVE"
	FormListParamsStatusArchived FormListParamsStatus = "ARCHIVED"
	FormListParamsStatusAll      FormListParamsStatus = "all"
)

type FormListSubmissionsParams struct {
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Page  param.Opt[int64] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FormListSubmissionsParams]'s query parameters as
// `url.Values`.
func (r FormListSubmissionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
