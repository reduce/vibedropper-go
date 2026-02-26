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

// ListSubscriberService contains methods and other services that help with
// interacting with the vibedropper API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewListSubscriberService] method instead.
type ListSubscriberService struct {
	Options []option.RequestOption
}

// NewListSubscriberService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewListSubscriberService(opts ...option.RequestOption) (r ListSubscriberService) {
	r = ListSubscriberService{}
	r.Options = opts
	return
}

// List subscribers
func (r *ListSubscriberService) List(ctx context.Context, listID string, opts ...option.RequestOption) (res *ListSubscriberListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return
	}
	path := fmt.Sprintf("lists/%s/subscribers", url.PathEscape(listID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Add subscriber
func (r *ListSubscriberService) Add(ctx context.Context, listID string, body ListSubscriberAddParams, opts ...option.RequestOption) (res *ListSubscriberAddResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return
	}
	path := fmt.Sprintf("lists/%s/subscribers", url.PathEscape(listID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Remove subscriber
func (r *ListSubscriberService) Remove(ctx context.Context, subscriberID string, body ListSubscriberRemoveParams, opts ...option.RequestOption) (res *ListSubscriberRemoveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ListID == "" {
		err = errors.New("missing required listId parameter")
		return
	}
	if subscriberID == "" {
		err = errors.New("missing required subscriberId parameter")
		return
	}
	path := fmt.Sprintf("lists/%s/subscribers/%s", url.PathEscape(body.ListID), url.PathEscape(subscriberID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Subscriber struct {
	ID           string `json:"id"`
	Customer     any    `json:"customer" api:"nullable"`
	CustomFields any    `json:"customFields" api:"nullable"`
	Email        string `json:"email"`
	ListID       string `json:"listId"`
	Name         string `json:"name" api:"nullable"`
	// Any of "SUBSCRIBED", "UNSUBSCRIBED", "PENDING", "BOUNCED".
	Status       SubscriberStatus `json:"status"`
	SubscribedAt time.Time        `json:"subscribedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Customer     respjson.Field
		CustomFields respjson.Field
		Email        respjson.Field
		ListID       respjson.Field
		Name         respjson.Field
		Status       respjson.Field
		SubscribedAt respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Subscriber) RawJSON() string { return r.JSON.raw }
func (r *Subscriber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriberStatus string

const (
	SubscriberStatusSubscribed   SubscriberStatus = "SUBSCRIBED"
	SubscriberStatusUnsubscribed SubscriberStatus = "UNSUBSCRIBED"
	SubscriberStatusPending      SubscriberStatus = "PENDING"
	SubscriberStatusBounced      SubscriberStatus = "BOUNCED"
)

type ListSubscriberListResponse struct {
	Subscribers []Subscriber `json:"subscribers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Subscribers respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListSubscriberListResponse) RawJSON() string { return r.JSON.raw }
func (r *ListSubscriberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListSubscriberAddResponse struct {
	Subscriber Subscriber `json:"subscriber"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Subscriber  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListSubscriberAddResponse) RawJSON() string { return r.JSON.raw }
func (r *ListSubscriberAddResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListSubscriberRemoveResponse struct {
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListSubscriberRemoveResponse) RawJSON() string { return r.JSON.raw }
func (r *ListSubscriberRemoveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListSubscriberAddParams struct {
	Email            string            `json:"email" api:"required" format:"email"`
	Name             param.Opt[string] `json:"name,omitzero"`
	PickupLocationID param.Opt[string] `json:"pickupLocationId,omitzero"`
	RegionID         param.Opt[string] `json:"regionId,omitzero"`
	CustomFields     any               `json:"customFields,omitzero"`
	paramObj
}

func (r ListSubscriberAddParams) MarshalJSON() (data []byte, err error) {
	type shadow ListSubscriberAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListSubscriberAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListSubscriberRemoveParams struct {
	ListID string `path:"listId" api:"required" json:"-"`
	paramObj
}
