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

// checks if the FeatureForceRule type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &FeatureForceRule{}

// FeatureForceRule struct for FeatureForceRule
type FeatureForceRule struct {
	Description interface{} `json:"description"`
	Condition interface{} `json:"condition"`
	SavedGroupTargeting interface{} `json:"savedGroupTargeting,omitempty"`
	Id interface{} `json:"id"`
	Enabled interface{} `json:"enabled"`
	Type interface{} `json:"type"`
	Value interface{} `json:"value"`
}

type _FeatureForceRule FeatureForceRule

// NewFeatureForceRule instantiates a new FeatureForceRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureForceRule(description interface{}, condition interface{}, id interface{}, enabled interface{}, type_ interface{}, value interface{}) *FeatureForceRule {
	this := FeatureForceRule{}
	this.Description = description
	this.Condition = condition
	this.Id = id
	this.Enabled = enabled
	this.Type = type_
	this.Value = value
	return &this
}

// NewFeatureForceRuleWithDefaults instantiates a new FeatureForceRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureForceRuleWithDefaults() *FeatureForceRule {
	this := FeatureForceRule{}
	return &this
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureForceRule) GetDescription() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureForceRule) GetDescriptionOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *FeatureForceRule) SetDescription(v interface{}) {
	o.Description = v
}

// GetCondition returns the Condition field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureForceRule) GetCondition() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureForceRule) GetConditionOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Condition) {
		return nil, false
	}
	return &o.Condition, true
}

// SetCondition sets field value
func (o *FeatureForceRule) SetCondition(v interface{}) {
	o.Condition = v
}

// GetSavedGroupTargeting returns the SavedGroupTargeting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureForceRule) GetSavedGroupTargeting() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SavedGroupTargeting
}

// GetSavedGroupTargetingOk returns a tuple with the SavedGroupTargeting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureForceRule) GetSavedGroupTargetingOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.SavedGroupTargeting) {
		return nil, false
	}
	return &o.SavedGroupTargeting, true
}

// HasSavedGroupTargeting returns a boolean if a field has been set.
func (o *FeatureForceRule) HasSavedGroupTargeting() bool {
	if o != nil && common.IsNil(o.SavedGroupTargeting) {
		return true
	}

	return false
}

// SetSavedGroupTargeting gets a reference to the given interface{} and assigns it to the SavedGroupTargeting field.
func (o *FeatureForceRule) SetSavedGroupTargeting(v interface{}) {
	o.SavedGroupTargeting = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureForceRule) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureForceRule) GetIdOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FeatureForceRule) SetId(v interface{}) {
	o.Id = v
}

// GetEnabled returns the Enabled field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureForceRule) GetEnabled() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureForceRule) GetEnabledOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Enabled) {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *FeatureForceRule) SetEnabled(v interface{}) {
	o.Enabled = v
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureForceRule) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureForceRule) GetTypeOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FeatureForceRule) SetType(v interface{}) {
	o.Type = v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureForceRule) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureForceRule) GetValueOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *FeatureForceRule) SetValue(v interface{}) {
	o.Value = v
}

func (o FeatureForceRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureForceRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Condition != nil {
		toSerialize["condition"] = o.Condition
	}
	if o.SavedGroupTargeting != nil {
		toSerialize["savedGroupTargeting"] = o.SavedGroupTargeting
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

func (o *FeatureForceRule) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"condition",
		"id",
		"enabled",
		"type",
		"value",
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

	varFeatureForceRule := _FeatureForceRule{}

	err = json.Unmarshal(bytes, &varFeatureForceRule)

	if err != nil {
		return err
	}

	*o = FeatureForceRule(varFeatureForceRule)

	return err
}

type NullableFeatureForceRule struct {
	value *FeatureForceRule
	isSet bool
}

func (v NullableFeatureForceRule) Get() *FeatureForceRule {
	return v.value
}

func (v *NullableFeatureForceRule) Set(val *FeatureForceRule) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureForceRule) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureForceRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureForceRule(val *FeatureForceRule) *NullableFeatureForceRule {
	return &NullableFeatureForceRule{value: val, isSet: true}
}

func (v NullableFeatureForceRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureForceRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

