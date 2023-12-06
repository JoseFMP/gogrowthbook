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

// checks if the FeatureDefinition type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &FeatureDefinition{}

// FeatureDefinition struct for FeatureDefinition
type FeatureDefinition struct {
	DefaultValue NullableFeatureDefinitionDefaultValue `json:"defaultValue"`
	Rules interface{} `json:"rules,omitempty"`
}

type _FeatureDefinition FeatureDefinition

// NewFeatureDefinition instantiates a new FeatureDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureDefinition(defaultValue NullableFeatureDefinitionDefaultValue) *FeatureDefinition {
	this := FeatureDefinition{}
	this.DefaultValue = defaultValue
	return &this
}

// NewFeatureDefinitionWithDefaults instantiates a new FeatureDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureDefinitionWithDefaults() *FeatureDefinition {
	this := FeatureDefinition{}
	return &this
}

// GetDefaultValue returns the DefaultValue field value
// If the value is explicit nil, the zero value for FeatureDefinitionDefaultValue will be returned
func (o *FeatureDefinition) GetDefaultValue() FeatureDefinitionDefaultValue {
	if o == nil || o.DefaultValue.Get() == nil {
		var ret FeatureDefinitionDefaultValue
		return ret
	}

	return *o.DefaultValue.Get()
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureDefinition) GetDefaultValueOk() (*FeatureDefinitionDefaultValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultValue.Get(), o.DefaultValue.IsSet()
}

// SetDefaultValue sets field value
func (o *FeatureDefinition) SetDefaultValue(v FeatureDefinitionDefaultValue) {
	o.DefaultValue.Set(&v)
}

// GetRules returns the Rules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureDefinition) GetRules() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureDefinition) GetRulesOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Rules) {
		return nil, false
	}
	return &o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *FeatureDefinition) HasRules() bool {
	if o != nil && common.IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given interface{} and assigns it to the Rules field.
func (o *FeatureDefinition) SetRules(v interface{}) {
	o.Rules = v
}

func (o FeatureDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["defaultValue"] = o.DefaultValue.Get()
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

func (o *FeatureDefinition) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"defaultValue",
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

	varFeatureDefinition := _FeatureDefinition{}

	err = json.Unmarshal(bytes, &varFeatureDefinition)

	if err != nil {
		return err
	}

	*o = FeatureDefinition(varFeatureDefinition)

	return err
}

type NullableFeatureDefinition struct {
	value *FeatureDefinition
	isSet bool
}

func (v NullableFeatureDefinition) Get() *FeatureDefinition {
	return v.value
}

func (v *NullableFeatureDefinition) Set(val *FeatureDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureDefinition(val *FeatureDefinition) *NullableFeatureDefinition {
	return &NullableFeatureDefinition{value: val, isSet: true}
}

func (v NullableFeatureDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


