/*
GrowthBook REST API

GrowthBook offers a full REST API for interacting with the GrowthBook application. This is currently in **beta** as we add more authenticated API routes and features.  Request data can use either JSON or Form data encoding (with proper `Content-Type` headers). All response bodies are JSON-encoded.  The API base URL for GrowthBook Cloud is `https://api.growthbook.io`. For self-hosted deployments, it is the same as your API_HOST environment variable (defaults to `http://localhost:3100`). The rest of these docs will assume you are using GrowthBook Cloud.  ## Authentication  We support both the HTTP Basic and Bearer authentication schemes for convenience.  You first need to generate a new API Key in GrowthBook. Different keys have different permissions:  - **Personal Access Tokens**: These are sensitive and provide the same level of access as the user has to an organization. These can be created by going to `Personal Access Tokens` under the your user menu. - **Secret Keys**: These are sensitive and provide the level of access for the role, which currently is either `admin` or `readonly`. Only Admins with the `manageApiKeys` permission can manage Secret Keys on behalf of an organization. These can be created by going to `Settings -> API Keys`  If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:  ```bash curl https://api.growthbook.io/api/v1 \\   -u secret_abc123DEF456: # The \":\" at the end stops curl from asking for a password ```  If using Bearer auth, pass the Secret Key as the token:  ```bash curl https://api.growthbook.io/api/v1 \\ -H \"Authorization: Bearer secret_abc123DEF456\" ```  ## Errors  The API may return the following error status codes:  - **400** - Bad Request - Often due to a missing required parameter - **401** - Unauthorized - No valid API key provided - **402** - Request Failed - The parameters are valid, but the request failed - **403** - Forbidden - Provided API key does not have the required access - **404** - Not Found - Unknown API route or requested resource - **429** - Too Many Requests - You exceeded the rate limit of 60 requests per minute. Try again later. - **5XX** - Server Error - Something went wrong on GrowthBook's end (these are rare)  The response body will be a JSON object with the following properties:  - **message** - Information about the error

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package models

import (
	"encoding/json"

	"github.com/JoseFMP/gogrowthbook/common"
)

// checks if the PaginationFields type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &PaginationFields{}

// PaginationFields struct for PaginationFields
type PaginationFields struct {
	Limit      int32  `json:"limit"`
	Offset     int32  `json:"offset"`
	Count      int32  `json:"count"`
	Total      int32  `json:"total"`
	HasMore    bool   `json:"hasMore"`
	NextOffset *int32 `json:"nextOffset"`
}

type _PaginationFields PaginationFields

// NewPaginationFields instantiates a new PaginationFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationFields(limit int32, offset int32, count int32, total int32, hasMore bool, nextOffset *int32) *PaginationFields {
	this := PaginationFields{}
	this.Limit = limit
	this.Offset = offset
	this.Count = count
	this.Total = total
	this.HasMore = hasMore
	this.NextOffset = nextOffset
	return &this
}

// NewPaginationFieldsWithDefaults instantiates a new PaginationFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationFieldsWithDefaults() *PaginationFields {
	this := PaginationFields{}
	return &this
}

// GetLimit returns the Limit field value
func (o *PaginationFields) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *PaginationFields) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *PaginationFields) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *PaginationFields) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *PaginationFields) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *PaginationFields) SetOffset(v int32) {
	o.Offset = v
}

// GetCount returns the Count field value
func (o *PaginationFields) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *PaginationFields) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *PaginationFields) SetCount(v int32) {
	o.Count = v
}

// GetTotal returns the Total field value
func (o *PaginationFields) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *PaginationFields) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *PaginationFields) SetTotal(v int32) {
	o.Total = v
}

// GetHasMore returns the HasMore field value
func (o *PaginationFields) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *PaginationFields) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *PaginationFields) SetHasMore(v bool) {
	o.HasMore = v
}

// GetNextOffset returns the NextOffset field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PaginationFields) GetNextOffset() int32 {
	if o == nil || o.NextOffset == nil {
		var ret int32
		return ret
	}

	return *o.NextOffset
}

// GetNextOffsetOk returns a tuple with the NextOffset field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginationFields) GetNextOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextOffset, o.NextOffset != nil
}

// SetNextOffset sets field value
func (o *PaginationFields) SetNextOffset(v *int32) {
	o.NextOffset = v
}

func (o PaginationFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["count"] = o.Count
	toSerialize["total"] = o.Total
	toSerialize["hasMore"] = o.HasMore
	toSerialize["nextOffset"] = o.NextOffset
	return toSerialize, nil
}

type NullablePaginationFields struct {
	value *PaginationFields
	isSet bool
}

func (v NullablePaginationFields) Get() *PaginationFields {
	return v.value
}

func (v *NullablePaginationFields) Set(val *PaginationFields) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationFields) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationFields(val *PaginationFields) *NullablePaginationFields {
	return &NullablePaginationFields{value: val, isSet: true}
}

func (v NullablePaginationFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
