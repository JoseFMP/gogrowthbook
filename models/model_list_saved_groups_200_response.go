/*
GrowthBook REST API

GrowthBook offers a full REST API for interacting with the GrowthBook application. This is currently in **beta** as we add more authenticated API routes and features.  Request data can use either JSON or Form data encoding (with proper `Content-Type` headers). All response bodies are JSON-encoded.  The API base URL for GrowthBook Cloud is `https://api.growthbook.io`. For self-hosted deployments, it is the same as your API_HOST environment variable (defaults to `http://localhost:3100`). The rest of these docs will assume you are using GrowthBook Cloud.  ## Authentication  We support both the HTTP Basic and Bearer authentication schemes for convenience.  You first need to generate a new API Key in GrowthBook. Different keys have different permissions:  - **Personal Access Tokens**: These are sensitive and provide the same level of access as the user has to an organization. These can be created by going to `Personal Access Tokens` under the your user menu. - **Secret Keys**: These are sensitive and provide the level of access for the role, which currently is either `admin` or `readonly`. Only Admins with the `manageApiKeys` permission can manage Secret Keys on behalf of an organization. These can be created by going to `Settings -> API Keys`  If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:  ```bash curl https://api.growthbook.io/api/v1 \\   -u secret_abc123DEF456: # The \":\" at the end stops curl from asking for a password ```  If using Bearer auth, pass the Secret Key as the token:  ```bash curl https://api.growthbook.io/api/v1 \\ -H \"Authorization: Bearer secret_abc123DEF456\" ```  ## Errors  The API may return the following error status codes:  - **400** - Bad Request - Often due to a missing required parameter - **401** - Unauthorized - No valid API key provided - **402** - Request Failed - The parameters are valid, but the request failed - **403** - Forbidden - Provided API key does not have the required access - **404** - Not Found - Unknown API route or requested resource - **429** - Too Many Requests - You exceeded the rate limit of 60 requests per minute. Try again later. - **5XX** - Server Error - Something went wrong on GrowthBook's end (these are rare)  The response body will be a JSON object with the following properties:  - **message** - Information about the error 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
"github.com/JoseFMP/gogrowthbook/common"
)

// checks if the ListSavedGroups200Response type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &ListSavedGroups200Response{}

// ListSavedGroups200Response struct for ListSavedGroups200Response
type ListSavedGroups200Response struct {
	Limit interface{} `json:"limit"`
	Offset interface{} `json:"offset"`
	Count interface{} `json:"count"`
	Total interface{} `json:"total"`
	HasMore interface{} `json:"hasMore"`
	NextOffset interface{} `json:"nextOffset"`
	SavedGroups []SavedGroup `json:"savedGroups"`
}

type _ListSavedGroups200Response ListSavedGroups200Response

// NewListSavedGroups200Response instantiates a new ListSavedGroups200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSavedGroups200Response(limit interface{}, offset interface{}, count interface{}, total interface{}, hasMore interface{}, nextOffset interface{}, savedGroups []SavedGroup) *ListSavedGroups200Response {
	this := ListSavedGroups200Response{}
	this.Limit = limit
	this.Offset = offset
	this.Count = count
	this.Total = total
	this.HasMore = hasMore
	this.NextOffset = nextOffset
	this.SavedGroups = savedGroups
	return &this
}

// NewListSavedGroups200ResponseWithDefaults instantiates a new ListSavedGroups200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSavedGroups200ResponseWithDefaults() *ListSavedGroups200Response {
	this := ListSavedGroups200Response{}
	return &this
}

// GetLimit returns the Limit field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ListSavedGroups200Response) GetLimit() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListSavedGroups200Response) GetLimitOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Limit) {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *ListSavedGroups200Response) SetLimit(v interface{}) {
	o.Limit = v
}

// GetOffset returns the Offset field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ListSavedGroups200Response) GetOffset() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListSavedGroups200Response) GetOffsetOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Offset) {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *ListSavedGroups200Response) SetOffset(v interface{}) {
	o.Offset = v
}

// GetCount returns the Count field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ListSavedGroups200Response) GetCount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListSavedGroups200Response) GetCountOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Count) {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *ListSavedGroups200Response) SetCount(v interface{}) {
	o.Count = v
}

// GetTotal returns the Total field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ListSavedGroups200Response) GetTotal() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListSavedGroups200Response) GetTotalOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ListSavedGroups200Response) SetTotal(v interface{}) {
	o.Total = v
}

// GetHasMore returns the HasMore field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ListSavedGroups200Response) GetHasMore() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListSavedGroups200Response) GetHasMoreOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.HasMore) {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *ListSavedGroups200Response) SetHasMore(v interface{}) {
	o.HasMore = v
}

// GetNextOffset returns the NextOffset field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ListSavedGroups200Response) GetNextOffset() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.NextOffset
}

// GetNextOffsetOk returns a tuple with the NextOffset field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListSavedGroups200Response) GetNextOffsetOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.NextOffset) {
		return nil, false
	}
	return &o.NextOffset, true
}

// SetNextOffset sets field value
func (o *ListSavedGroups200Response) SetNextOffset(v interface{}) {
	o.NextOffset = v
}

// GetSavedGroups returns the SavedGroups field value
func (o *ListSavedGroups200Response) GetSavedGroups() []SavedGroup {
	if o == nil {
		var ret []SavedGroup
		return ret
	}

	return o.SavedGroups
}

// GetSavedGroupsOk returns a tuple with the SavedGroups field value
// and a boolean to check if the value has been set.
func (o *ListSavedGroups200Response) GetSavedGroupsOk() ([]SavedGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.SavedGroups, true
}

// SetSavedGroups sets field value
func (o *ListSavedGroups200Response) SetSavedGroups(v []SavedGroup) {
	o.SavedGroups = v
}

func (o ListSavedGroups200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSavedGroups200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.HasMore != nil {
		toSerialize["hasMore"] = o.HasMore
	}
	if o.NextOffset != nil {
		toSerialize["nextOffset"] = o.NextOffset
	}
	toSerialize["savedGroups"] = o.SavedGroups
	return toSerialize, nil
}

func (o *ListSavedGroups200Response) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"limit",
		"offset",
		"count",
		"total",
		"hasMore",
		"nextOffset",
		"savedGroups",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varListSavedGroups200Response := _ListSavedGroups200Response{}

	err = json.Unmarshal(bytes, &varListSavedGroups200Response)

	if err != nil {
		return err
	}

	*o = ListSavedGroups200Response(varListSavedGroups200Response)

	return err
}

type NullableListSavedGroups200Response struct {
	value *ListSavedGroups200Response
	isSet bool
}

func (v NullableListSavedGroups200Response) Get() *ListSavedGroups200Response {
	return v.value
}

func (v *NullableListSavedGroups200Response) Set(val *ListSavedGroups200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListSavedGroups200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListSavedGroups200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSavedGroups200Response(val *ListSavedGroups200Response) *NullableListSavedGroups200Response {
	return &NullableListSavedGroups200Response{value: val, isSet: true}
}

func (v NullableListSavedGroups200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSavedGroups200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


