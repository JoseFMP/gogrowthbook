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

// checks if the FeatureRolloutRule type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &FeatureRolloutRule{}

// FeatureRolloutRule struct for FeatureRolloutRule
type FeatureRolloutRule struct {
	Description *string `json:"description"`
	Condition *string `json:"condition"`
	SavedGroupTargeting []FeatureRuleSavedGroupTargeting `json:"savedGroupTargeting,omitempty"`
	Id *string `json:"id"`
	Enabled *bool `json:"enabled"`
	Type *FeatureRuleType `json:"type"`
	Value *string `json:"value"`
	Coverage *float64 `json:"coverage"`
	HashAttribute *string `json:"hashAttribute"`
}

type _FeatureRolloutRule FeatureRolloutRule

// NewFeatureRolloutRule instantiates a new FeatureRolloutRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureRolloutRule(description *string, condition *string, id *string, enabled *bool, type_ *FeatureRuleType, value *string, coverage *float64, hashAttribute *string) *FeatureRolloutRule {
	this := FeatureRolloutRule{}
	this.Description = description
	this.Condition = condition
	this.Id = id
	this.Enabled = enabled
	this.Type = type_
	this.Value = value
	this.Coverage = coverage
	this.HashAttribute = hashAttribute
	return &this
}

// NewFeatureRolloutRuleWithDefaults instantiates a new FeatureRolloutRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureRolloutRuleWithDefaults() *FeatureRolloutRule {
	this := FeatureRolloutRule{}
	return &this
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureRolloutRule) GetDescription() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureRolloutRule) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// SetDescription sets field value
func (o *FeatureRolloutRule) SetDescription(v *string) {
	o.Description = v
}

// GetCondition returns the Condition field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureRolloutRule) GetCondition() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureRolloutRule) GetConditionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// SetCondition sets field value
func (o *FeatureRolloutRule) SetCondition(v *string) {
	o.Condition = v
}

// GetSavedGroupTargeting returns the SavedGroupTargeting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureRolloutRule) GetSavedGroupTargeting() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SavedGroupTargeting
}

// GetSavedGroupTargetingOk returns a tuple with the SavedGroupTargeting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureRolloutRule) GetSavedGroupTargetingOk() ([]FeatureRuleSavedGroupTargeting, bool) {
	if o == nil || common.IsNil(o.SavedGroupTargeting) {
		return nil, false
	}
	return o.SavedGroupTargeting, true
}

// HasSavedGroupTargeting returns a boolean if a field has been set.
func (o *FeatureRolloutRule) HasSavedGroupTargeting() bool {
	if o != nil && common.IsNil(o.SavedGroupTargeting) {
		return true
	}

	return false
}

// SetSavedGroupTargeting gets a reference to the given interface{} and assigns it to the SavedGroupTargeting field.
func (o *FeatureRolloutRule) SetSavedGroupTargeting(v []FeatureRuleSavedGroupTargeting) {
	o.SavedGroupTargeting = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureRolloutRule) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureRolloutRule) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// SetId sets field value
func (o *FeatureRolloutRule) SetId(v *string) {
	o.Id = v
}

// GetEnabled returns the Enabled field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureRolloutRule) GetEnabled() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureRolloutRule) GetEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// SetEnabled sets field value
func (o *FeatureRolloutRule) SetEnabled(v *bool) {
	o.Enabled = v
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureRolloutRule) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureRolloutRule) GetTypeOk() (*FeatureRuleType, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *FeatureRolloutRule) SetType(v *FeatureRuleType) {
	o.Type = v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureRolloutRule) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureRolloutRule) GetValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *FeatureRolloutRule) SetValue(v *string) {
	o.Value = v
}

// GetCoverage returns the Coverage field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureRolloutRule) GetCoverage() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Coverage
}

// GetCoverageOk returns a tuple with the Coverage field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureRolloutRule) GetCoverageOk() (*float64, bool) {
	if o == nil || common.IsNil(o.Coverage) {
		return nil, false
	}
	return o.Coverage, true
}

// SetCoverage sets field value
func (o *FeatureRolloutRule) SetCoverage(v *float64) {
	o.Coverage = v
}

// GetHashAttribute returns the HashAttribute field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureRolloutRule) GetHashAttribute() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.HashAttribute
}

// GetHashAttributeOk returns a tuple with the HashAttribute field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureRolloutRule) GetHashAttributeOk() (*string, bool) {
	if o == nil || common.IsNil(o.HashAttribute) {
		return nil, false
	}
	return o.HashAttribute, true
}

// SetHashAttribute sets field value
func (o *FeatureRolloutRule) SetHashAttribute(v *string) {
	o.HashAttribute = v
}

func (o FeatureRolloutRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureRolloutRule) ToMap() (map[string]interface{}, error) {
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
	if o.Coverage != nil {
		toSerialize["coverage"] = o.Coverage
	}
	if o.HashAttribute != nil {
		toSerialize["hashAttribute"] = o.HashAttribute
	}
	return toSerialize, nil
}

func (o *FeatureRolloutRule) UnmarshalJSON(bytes []byte) (err error) {
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
		"coverage",
		"hashAttribute",
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

	varFeatureRolloutRule := _FeatureRolloutRule{}

	err = json.Unmarshal(bytes, &varFeatureRolloutRule)

	if err != nil {
		return err
	}

	*o = FeatureRolloutRule(varFeatureRolloutRule)

	return err
}

type NullableFeatureRolloutRule struct {
	value *FeatureRolloutRule
	isSet bool
}

func (v NullableFeatureRolloutRule) Get() *FeatureRolloutRule {
	return v.value
}

func (v *NullableFeatureRolloutRule) Set(val *FeatureRolloutRule) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureRolloutRule) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureRolloutRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureRolloutRule(val *FeatureRolloutRule) *NullableFeatureRolloutRule {
	return &NullableFeatureRolloutRule{value: val, isSet: true}
}

func (v NullableFeatureRolloutRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureRolloutRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


