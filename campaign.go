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
	"github.com/reduce/vibedropper-go/internal/requestconfig"
	"github.com/reduce/vibedropper-go/option"
	"github.com/reduce/vibedropper-go/packages/respjson"
)

// CampaignService contains methods and other services that help with interacting
// with the vibedropper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCampaignService] method instead.
type CampaignService struct {
	Options []option.RequestOption
}

// NewCampaignService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCampaignService(opts ...option.RequestOption) (r CampaignService) {
	r = CampaignService{}
	r.Options = opts
	return
}

// Get campaign
func (r *CampaignService) Get(ctx context.Context, campaignID string, opts ...option.RequestOption) (res *CampaignGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return
	}
	path := fmt.Sprintf("campaigns/%s", url.PathEscape(campaignID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List campaigns
func (r *CampaignService) List(ctx context.Context, opts ...option.RequestOption) (res *CampaignListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "campaigns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Campaign struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	ListIDs   []string  `json:"listIds"`
	Name      string    `json:"name"`
	OrgID     string    `json:"orgId"`
	SentAt    time.Time `json:"sentAt" api:"nullable" format:"date-time"`
	// Any of "DRAFT", "SCHEDULED", "SENDING", "SENT".
	Status  CampaignStatus `json:"status"`
	Subject string         `json:"subject"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		ListIDs     respjson.Field
		Name        respjson.Field
		OrgID       respjson.Field
		SentAt      respjson.Field
		Status      respjson.Field
		Subject     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Campaign) RawJSON() string { return r.JSON.raw }
func (r *Campaign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CampaignStatus string

const (
	CampaignStatusDraft     CampaignStatus = "DRAFT"
	CampaignStatusScheduled CampaignStatus = "SCHEDULED"
	CampaignStatusSending   CampaignStatus = "SENDING"
	CampaignStatusSent      CampaignStatus = "SENT"
)

type CampaignGetResponse struct {
	Campaign Campaign `json:"campaign"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Campaign    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CampaignGetResponse) RawJSON() string { return r.JSON.raw }
func (r *CampaignGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CampaignListResponse struct {
	Campaigns []Campaign `json:"campaigns"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Campaigns   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CampaignListResponse) RawJSON() string { return r.JSON.raw }
func (r *CampaignListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
