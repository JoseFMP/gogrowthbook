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

// checks if the FeatureExperimentRule type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &FeatureExperimentRule{}

// FeatureExperimentRule struct for FeatureExperimentRule
type FeatureExperimentRule struct {
	Description   *string                         `json:"description"`
	Condition     *string                         `json:"condition"`
	Id            *string                         `json:"id"`
	Enabled       *bool                           `json:"enabled"`
	Type          *FeatureRuleType                `json:"type"`
	TrackingKey   *string                         `json:"trackingKey,omitempty"`
	HashAttribute *string                         `json:"hashAttribute,omitempty"`
	Namespace     *FeatureExperimentRuleNamespace `json:"namespace,omitempty"`
	Coverage      *float64                        `json:"coverage,omitempty"`
	Value         []FeatureExperimentRuleValue    `json:"value,omitempty"`
}

type FeatureExperimentRuleValue struct {
	Value  *string  `json:"value"`
	Weight *float64 `json:"weight"`
	Name   *string  `json:"name"`
}

type _FeatureExperimentRule FeatureExperimentRule

// NewFeatureExperimentRule instantiates a new FeatureExperimentRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureExperimentRule(description *string, condition *string, id *string, enabled *bool, type_ *FeatureRuleType) *FeatureExperimentRule {
	this := FeatureExperimentRule{}
	this.Description = description
	this.Condition = condition
	this.Id = id
	this.Enabled = enabled
	this.Type = type_
	return &this
}

// NewFeatureExperimentRuleWithDefaults instantiates a new FeatureExperimentRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureExperimentRuleWithDefaults() *FeatureExperimentRule {
	this := FeatureExperimentRule{}
	return &this
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureExperimentRule) GetDescription() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureExperimentRule) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// SetDescription sets field value
func (o *FeatureExperimentRule) SetDescription(v *string) {
	o.Description = v
}

// GetCondition returns the Condition field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureExperimentRule) GetCondition() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureExperimentRule) GetConditionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// SetCondition sets field value
func (o *FeatureExperimentRule) SetCondition(v *string) {
	o.Condition = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureExperimentRule) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureExperimentRule) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// SetId sets field value
func (o *FeatureExperimentRule) SetId(v *string) {
	o.Id = v
}

// GetEnabled returns the Enabled field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureExperimentRule) GetEnabled() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureExperimentRule) GetEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// SetEnabled sets field value
func (o *FeatureExperimentRule) SetEnabled(v *bool) {
	o.Enabled = v
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureExperimentRule) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureExperimentRule) GetTypeOk() (*FeatureRuleType, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *FeatureExperimentRule) SetType(v *FeatureRuleType) {
	o.Type = v
}

// GetTrackingKey returns the TrackingKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureExperimentRule) GetTrackingKey() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TrackingKey
}

// GetTrackingKeyOk returns a tuple with the TrackingKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureExperimentRule) GetTrackingKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.TrackingKey) {
		return nil, false
	}
	return o.TrackingKey, true
}

// HasTrackingKey returns a boolean if a field has been set.
func (o *FeatureExperimentRule) HasTrackingKey() bool {
	if o != nil && common.IsNil(o.TrackingKey) {
		return true
	}

	return false
}

// SetTrackingKey gets a reference to the given interface{} and assigns it to the TrackingKey field.
func (o *FeatureExperimentRule) SetTrackingKey(v *string) {
	o.TrackingKey = v
}

// GetHashAttribute returns the HashAttribute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureExperimentRule) GetHashAttribute() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.HashAttribute
}

// GetHashAttributeOk returns a tuple with the HashAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureExperimentRule) GetHashAttributeOk() (*string, bool) {
	if o == nil || common.IsNil(o.HashAttribute) {
		return nil, false
	}
	return o.HashAttribute, true
}

// HasHashAttribute returns a boolean if a field has been set.
func (o *FeatureExperimentRule) HasHashAttribute() bool {
	if o != nil && common.IsNil(o.HashAttribute) {
		return true
	}

	return false
}

// SetHashAttribute gets a reference to the given interface{} and assigns it to the HashAttribute field.
func (o *FeatureExperimentRule) SetHashAttribute(v *string) {
	o.HashAttribute = v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *FeatureExperimentRule) GetNamespace() FeatureExperimentRuleNamespace {
	if o == nil || common.IsNil(o.Namespace) {
		var ret FeatureExperimentRuleNamespace
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureExperimentRule) GetNamespaceOk() (*FeatureExperimentRuleNamespace, bool) {
	if o == nil || common.IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *FeatureExperimentRule) HasNamespace() bool {
	if o != nil && !common.IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given FeatureExperimentRuleNamespace and assigns it to the Namespace field.
func (o *FeatureExperimentRule) SetNamespace(v FeatureExperimentRuleNamespace) {
	o.Namespace = &v
}

// GetCoverage returns the Coverage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureExperimentRule) GetCoverage() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Coverage
}

// GetCoverageOk returns a tuple with the Coverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureExperimentRule) GetCoverageOk() (*float64, bool) {
	if o == nil || common.IsNil(o.Coverage) {
		return nil, false
	}
	return o.Coverage, true
}

// HasCoverage returns a boolean if a field has been set.
func (o *FeatureExperimentRule) HasCoverage() bool {
	if o != nil && common.IsNil(o.Coverage) {
		return true
	}

	return false
}

// SetCoverage gets a reference to the given interface{} and assigns it to the Coverage field.
func (o *FeatureExperimentRule) SetCoverage(v *float64) {
	o.Coverage = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureExperimentRule) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureExperimentRule) GetValueOk() ([]FeatureExperimentRuleValue, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FeatureExperimentRule) HasValue() bool {
	if o != nil && common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *FeatureExperimentRule) SetValue(v []FeatureExperimentRuleValue) {
	o.Value = v
}

func (o FeatureExperimentRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureExperimentRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Condition != nil {
		toSerialize["condition"] = o.Condition
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
	if o.TrackingKey != nil {
		toSerialize["trackingKey"] = o.TrackingKey
	}
	if o.HashAttribute != nil {
		toSerialize["hashAttribute"] = o.HashAttribute
	}
	if !common.IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Coverage != nil {
		toSerialize["coverage"] = o.Coverage
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

func (o *FeatureExperimentRule) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"condition",
		"id",
		"enabled",
		"type",
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

	varFeatureExperimentRule := _FeatureExperimentRule{}

	err = json.Unmarshal(bytes, &varFeatureExperimentRule)

	if err != nil {
		return err
	}

	*o = FeatureExperimentRule(varFeatureExperimentRule)

	return err
}

type NullableFeatureExperimentRule struct {
	value *FeatureExperimentRule
	isSet bool
}

func (v NullableFeatureExperimentRule) Get() *FeatureExperimentRule {
	return v.value
}

func (v *NullableFeatureExperimentRule) Set(val *FeatureExperimentRule) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureExperimentRule) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureExperimentRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureExperimentRule(val *FeatureExperimentRule) *NullableFeatureExperimentRule {
	return &NullableFeatureExperimentRule{value: val, isSet: true}
}

func (v NullableFeatureExperimentRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureExperimentRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
