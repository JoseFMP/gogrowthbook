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

// checks if the FeatureExperimentRefRule type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &FeatureExperimentRefRule{}

// FeatureExperimentRefRule struct for FeatureExperimentRefRule
type FeatureExperimentRefRule struct {
	Description interface{} `json:"description"`
	Id interface{} `json:"id"`
	Enabled interface{} `json:"enabled"`
	Type interface{} `json:"type"`
	Condition interface{} `json:"condition,omitempty"`
	Variations interface{} `json:"variations"`
	ExperimentId interface{} `json:"experimentId"`
}

type _FeatureExperimentRefRule FeatureExperimentRefRule

// NewFeatureExperimentRefRule instantiates a new FeatureExperimentRefRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureExperimentRefRule(description interface{}, id interface{}, enabled interface{}, type_ interface{}, variations interface{}, experimentId interface{}) *FeatureExperimentRefRule {
	this := FeatureExperimentRefRule{}
	this.Description = description
	this.Id = id
	this.Enabled = enabled
	this.Type = type_
	this.Variations = variations
	this.ExperimentId = experimentId
	return &this
}

// NewFeatureExperimentRefRuleWithDefaults instantiates a new FeatureExperimentRefRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureExperimentRefRuleWithDefaults() *FeatureExperimentRefRule {
	this := FeatureExperimentRefRule{}
	return &this
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureExperimentRefRule) GetDescription() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureExperimentRefRule) GetDescriptionOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *FeatureExperimentRefRule) SetDescription(v interface{}) {
	o.Description = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureExperimentRefRule) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureExperimentRefRule) GetIdOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FeatureExperimentRefRule) SetId(v interface{}) {
	o.Id = v
}

// GetEnabled returns the Enabled field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureExperimentRefRule) GetEnabled() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureExperimentRefRule) GetEnabledOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Enabled) {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *FeatureExperimentRefRule) SetEnabled(v interface{}) {
	o.Enabled = v
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureExperimentRefRule) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureExperimentRefRule) GetTypeOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FeatureExperimentRefRule) SetType(v interface{}) {
	o.Type = v
}

// GetCondition returns the Condition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureExperimentRefRule) GetCondition() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureExperimentRefRule) GetConditionOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Condition) {
		return nil, false
	}
	return &o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *FeatureExperimentRefRule) HasCondition() bool {
	if o != nil && common.IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given interface{} and assigns it to the Condition field.
func (o *FeatureExperimentRefRule) SetCondition(v interface{}) {
	o.Condition = v
}

// GetVariations returns the Variations field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureExperimentRefRule) GetVariations() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Variations
}

// GetVariationsOk returns a tuple with the Variations field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureExperimentRefRule) GetVariationsOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Variations) {
		return nil, false
	}
	return &o.Variations, true
}

// SetVariations sets field value
func (o *FeatureExperimentRefRule) SetVariations(v interface{}) {
	o.Variations = v
}

// GetExperimentId returns the ExperimentId field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureExperimentRefRule) GetExperimentId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.ExperimentId
}

// GetExperimentIdOk returns a tuple with the ExperimentId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureExperimentRefRule) GetExperimentIdOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.ExperimentId) {
		return nil, false
	}
	return &o.ExperimentId, true
}

// SetExperimentId sets field value
func (o *FeatureExperimentRefRule) SetExperimentId(v interface{}) {
	o.ExperimentId = v
}

func (o FeatureExperimentRefRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureExperimentRefRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
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
	if o.Condition != nil {
		toSerialize["condition"] = o.Condition
	}
	if o.Variations != nil {
		toSerialize["variations"] = o.Variations
	}
	if o.ExperimentId != nil {
		toSerialize["experimentId"] = o.ExperimentId
	}
	return toSerialize, nil
}

func (o *FeatureExperimentRefRule) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"id",
		"enabled",
		"type",
		"variations",
		"experimentId",
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

	varFeatureExperimentRefRule := _FeatureExperimentRefRule{}

	err = json.Unmarshal(bytes, &varFeatureExperimentRefRule)

	if err != nil {
		return err
	}

	*o = FeatureExperimentRefRule(varFeatureExperimentRefRule)

	return err
}

type NullableFeatureExperimentRefRule struct {
	value *FeatureExperimentRefRule
	isSet bool
}

func (v NullableFeatureExperimentRefRule) Get() *FeatureExperimentRefRule {
	return v.value
}

func (v *NullableFeatureExperimentRefRule) Set(val *FeatureExperimentRefRule) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureExperimentRefRule) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureExperimentRefRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureExperimentRefRule(val *FeatureExperimentRefRule) *NullableFeatureExperimentRefRule {
	return &NullableFeatureExperimentRefRule{value: val, isSet: true}
}

func (v NullableFeatureExperimentRefRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureExperimentRefRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


