/*
GrowthBook REST API

GrowthBook offers a full REST API for interacting with the GrowthBook application. This is currently in **beta** as we add more authenticated API routes and features.  Request data can use either JSON or Form data encoding (with proper `Content-Type` headers). All response bodies are JSON-encoded.  The API base URL for GrowthBook Cloud is `https://api.growthbook.io`. For self-hosted deployments, it is the same as your API_HOST environment variable (defaults to `http://localhost:3100`). The rest of these docs will assume you are using GrowthBook Cloud.  ## Authentication  We support both the HTTP Basic and Bearer authentication schemes for convenience.  You first need to generate a new API Key in GrowthBook. Different keys have different permissions:  - **Personal Access Tokens**: These are sensitive and provide the same level of access as the user has to an organization. These can be created by going to `Personal Access Tokens` under the your user menu. - **Secret Keys**: These are sensitive and provide the level of access for the role, which currently is either `admin` or `readonly`. Only Admins with the `manageApiKeys` permission can manage Secret Keys on behalf of an organization. These can be created by going to `Settings -> API Keys`  If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:  ```bash curl https://api.growthbook.io/api/v1 \\   -u secret_abc123DEF456: # The \":\" at the end stops curl from asking for a password ```  If using Bearer auth, pass the Secret Key as the token:  ```bash curl https://api.growthbook.io/api/v1 \\ -H \"Authorization: Bearer secret_abc123DEF456\" ```  ## Errors  The API may return the following error status codes:  - **400** - Bad Request - Often due to a missing required parameter - **401** - Unauthorized - No valid API key provided - **402** - Request Failed - The parameters are valid, but the request failed - **403** - Forbidden - Provided API key does not have the required access - **404** - Not Found - Unknown API route or requested resource - **429** - Too Many Requests - You exceeded the rate limit of 60 requests per minute. Try again later. - **5XX** - Server Error - Something went wrong on GrowthBook's end (these are rare)  The response body will be a JSON object with the following properties:  - **message** - Information about the error 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
"github.com/JoseFMP/gogrowthbook/common"
)

// checks if the UpdateSavedGroupRequest type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &UpdateSavedGroupRequest{}

// UpdateSavedGroupRequest struct for UpdateSavedGroupRequest
type UpdateSavedGroupRequest struct {
	// The display name of the Saved Group
	Name *string `json:"name,omitempty"`
	Values []string `json:"values,omitempty"`
	// The person or team that owns this Saved Group. If no owner, you can pass an empty string.
	Owner *string `json:"owner,omitempty"`
	// (Runtime groups only) The key used to reference the Saved Group in the SDK
	AttributeKey *string `json:"attributeKey,omitempty"`
}

// NewUpdateSavedGroupRequest instantiates a new UpdateSavedGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSavedGroupRequest() *UpdateSavedGroupRequest {
	this := UpdateSavedGroupRequest{}
	return &this
}

// NewUpdateSavedGroupRequestWithDefaults instantiates a new UpdateSavedGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSavedGroupRequestWithDefaults() *UpdateSavedGroupRequest {
	this := UpdateSavedGroupRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateSavedGroupRequest) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSavedGroupRequest) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateSavedGroupRequest) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateSavedGroupRequest) SetName(v string) {
	o.Name = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *UpdateSavedGroupRequest) GetValues() []string {
	if o == nil || common.IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSavedGroupRequest) GetValuesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *UpdateSavedGroupRequest) HasValues() bool {
	if o != nil && !common.IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *UpdateSavedGroupRequest) SetValues(v []string) {
	o.Values = v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *UpdateSavedGroupRequest) GetOwner() string {
	if o == nil || common.IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSavedGroupRequest) GetOwnerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *UpdateSavedGroupRequest) HasOwner() bool {
	if o != nil && !common.IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *UpdateSavedGroupRequest) SetOwner(v string) {
	o.Owner = &v
}

// GetAttributeKey returns the AttributeKey field value if set, zero value otherwise.
func (o *UpdateSavedGroupRequest) GetAttributeKey() string {
	if o == nil || common.IsNil(o.AttributeKey) {
		var ret string
		return ret
	}
	return *o.AttributeKey
}

// GetAttributeKeyOk returns a tuple with the AttributeKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSavedGroupRequest) GetAttributeKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.AttributeKey) {
		return nil, false
	}
	return o.AttributeKey, true
}

// HasAttributeKey returns a boolean if a field has been set.
func (o *UpdateSavedGroupRequest) HasAttributeKey() bool {
	if o != nil && !common.IsNil(o.AttributeKey) {
		return true
	}

	return false
}

// SetAttributeKey gets a reference to the given string and assigns it to the AttributeKey field.
func (o *UpdateSavedGroupRequest) SetAttributeKey(v string) {
	o.AttributeKey = &v
}

func (o UpdateSavedGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSavedGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	if !common.IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !common.IsNil(o.AttributeKey) {
		toSerialize["attributeKey"] = o.AttributeKey
	}
	return toSerialize, nil
}

type NullableUpdateSavedGroupRequest struct {
	value *UpdateSavedGroupRequest
	isSet bool
}

func (v NullableUpdateSavedGroupRequest) Get() *UpdateSavedGroupRequest {
	return v.value
}

func (v *NullableUpdateSavedGroupRequest) Set(val *UpdateSavedGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSavedGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSavedGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSavedGroupRequest(val *UpdateSavedGroupRequest) *NullableUpdateSavedGroupRequest {
	return &NullableUpdateSavedGroupRequest{value: val, isSet: true}
}

func (v NullableUpdateSavedGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSavedGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


