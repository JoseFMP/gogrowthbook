/*
GrowthBook REST API

GrowthBook offers a full REST API for interacting with the GrowthBook application. This is currently in **beta** as we add more authenticated API routes and features.  Request data can use either JSON or Form data encoding (with proper `Content-Type` headers). All response bodies are JSON-encoded.  The API base URL for GrowthBook Cloud is `https://api.growthbook.io`. For self-hosted deployments, it is the same as your API_HOST environment variable (defaults to `http://localhost:3100`). The rest of these docs will assume you are using GrowthBook Cloud.  ## Authentication  We support both the HTTP Basic and Bearer authentication schemes for convenience.  You first need to generate a new API Key in GrowthBook. Different keys have different permissions:  - **Personal Access Tokens**: These are sensitive and provide the same level of access as the user has to an organization. These can be created by going to `Personal Access Tokens` under the your user menu. - **Secret Keys**: These are sensitive and provide the level of access for the role, which currently is either `admin` or `readonly`. Only Admins with the `manageApiKeys` permission can manage Secret Keys on behalf of an organization. These can be created by going to `Settings -> API Keys`  If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:  ```bash curl https://api.growthbook.io/api/v1 \\   -u secret_abc123DEF456: # The \":\" at the end stops curl from asking for a password ```  If using Bearer auth, pass the Secret Key as the token:  ```bash curl https://api.growthbook.io/api/v1 \\ -H \"Authorization: Bearer secret_abc123DEF456\" ```  ## Errors  The API may return the following error status codes:  - **400** - Bad Request - Often due to a missing required parameter - **401** - Unauthorized - No valid API key provided - **402** - Request Failed - The parameters are valid, but the request failed - **403** - Forbidden - Provided API key does not have the required access - **404** - Not Found - Unknown API route or requested resource - **429** - Too Many Requests - You exceeded the rate limit of 60 requests per minute. Try again later. - **5XX** - Server Error - Something went wrong on GrowthBook's end (these are rare)  The response body will be a JSON object with the following properties:  - **message** - Information about the error

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package models

import (
	"encoding/json"
	"fmt"

	"github.com/JoseFMP/gogrowthbook/common"
)

// checks if the ListSegments200Response type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &ListSegments200Response{}

// ListSegments200Response struct for ListSegments200Response
type ListSegments200Response struct {
	PaginationFields
	Segments []Segment `json:"segments"`
}

type _ListSegments200Response ListSegments200Response

// NewListSegments200Response instantiates a new ListSegments200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSegments200Response(limit int32, offset int32, count int32, total int32, hasMore bool, nextOffset *int32, segments []Segment) *ListSegments200Response {
	this := ListSegments200Response{}
	this.Limit = limit
	this.Offset = offset
	this.Count = count
	this.Total = total
	this.HasMore = hasMore
	this.NextOffset = nextOffset
	this.Segments = segments
	return &this
}

// NewListSegments200ResponseWithDefaults instantiates a new ListSegments200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSegments200ResponseWithDefaults() *ListSegments200Response {
	this := ListSegments200Response{}
	return &this
}

// GetLimit returns the Limit field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ListSegments200Response) GetLimit() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Limit
}

// GetSegments returns the Segments field value
func (o *ListSegments200Response) GetSegments() []Segment {
	if o == nil {
		var ret []Segment
		return ret
	}

	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value
// and a boolean to check if the value has been set.
func (o *ListSegments200Response) GetSegmentsOk() ([]Segment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Segments, true
}

// SetSegments sets field value
func (o *ListSegments200Response) SetSegments(v []Segment) {
	o.Segments = v
}

func (o ListSegments200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSegments200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit

	toSerialize["offset"] = o.Offset

	toSerialize["count"] = o.Count

	toSerialize["total"] = o.Total

	toSerialize["hasMore"] = o.HasMore

	toSerialize["nextOffset"] = o.NextOffset

	toSerialize["segments"] = o.Segments
	return toSerialize, nil
}

func (o *ListSegments200Response) UnmarshalJSON(bytes []byte) (err error) {
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
		"segments",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varListSegments200Response := _ListSegments200Response{}

	err = json.Unmarshal(bytes, &varListSegments200Response)

	if err != nil {
		return err
	}

	*o = ListSegments200Response(varListSegments200Response)

	return err
}

type NullableListSegments200Response struct {
	value *ListSegments200Response
	isSet bool
}

func (v NullableListSegments200Response) Get() *ListSegments200Response {
	return v.value
}

func (v *NullableListSegments200Response) Set(val *ListSegments200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListSegments200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListSegments200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSegments200Response(val *ListSegments200Response) *NullableListSegments200Response {
	return &NullableListSegments200Response{value: val, isSet: true}
}

func (v NullableListSegments200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSegments200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
