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

// checks if the FeatureExperimentRuleNamespace type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &FeatureExperimentRuleNamespace{}

// FeatureExperimentRuleNamespace struct for FeatureExperimentRuleNamespace
type FeatureExperimentRuleNamespace struct {
	Enabled *bool       `json:"enabled"`
	Name    *string     `json:"name"`
	Range   interface{} `json:"range"`
}

type _FeatureExperimentRuleNamespace FeatureExperimentRuleNamespace

// NewFeatureExperimentRuleNamespace instantiates a new FeatureExperimentRuleNamespace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureExperimentRuleNamespace(enabled *bool, name *string, range_ interface{}) *FeatureExperimentRuleNamespace {
	this := FeatureExperimentRuleNamespace{}
	this.Enabled = enabled
	this.Name = name
	this.Range = range_
	return &this
}

// NewFeatureExperimentRuleNamespaceWithDefaults instantiates a new FeatureExperimentRuleNamespace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureExperimentRuleNamespaceWithDefaults() *FeatureExperimentRuleNamespace {
	this := FeatureExperimentRuleNamespace{}
	return &this
}

// GetEnabled returns the Enabled field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureExperimentRuleNamespace) GetEnabled() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureExperimentRuleNamespace) GetEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// SetEnabled sets field value
func (o *FeatureExperimentRuleNamespace) SetEnabled(v *bool) {
	o.Enabled = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureExperimentRuleNamespace) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureExperimentRuleNamespace) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *FeatureExperimentRuleNamespace) SetName(v *string) {
	o.Name = v
}

// GetRange returns the Range field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FeatureExperimentRuleNamespace) GetRange() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Range
}

// GetRangeOk returns a tuple with the Range field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureExperimentRuleNamespace) GetRangeOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Range) {
		return nil, false
	}
	return &o.Range, true
}

// SetRange sets field value
func (o *FeatureExperimentRuleNamespace) SetRange(v interface{}) {
	o.Range = v
}

func (o FeatureExperimentRuleNamespace) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureExperimentRuleNamespace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Range != nil {
		toSerialize["range"] = o.Range
	}
	return toSerialize, nil
}

func (o *FeatureExperimentRuleNamespace) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
		"name",
		"range",
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

	varFeatureExperimentRuleNamespace := _FeatureExperimentRuleNamespace{}

	err = json.Unmarshal(bytes, &varFeatureExperimentRuleNamespace)

	if err != nil {
		return err
	}

	*o = FeatureExperimentRuleNamespace(varFeatureExperimentRuleNamespace)

	return err
}

type NullableFeatureExperimentRuleNamespace struct {
	value *FeatureExperimentRuleNamespace
	isSet bool
}

func (v NullableFeatureExperimentRuleNamespace) Get() *FeatureExperimentRuleNamespace {
	return v.value
}

func (v *NullableFeatureExperimentRuleNamespace) Set(val *FeatureExperimentRuleNamespace) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureExperimentRuleNamespace) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureExperimentRuleNamespace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureExperimentRuleNamespace(val *FeatureExperimentRuleNamespace) *NullableFeatureExperimentRuleNamespace {
	return &NullableFeatureExperimentRuleNamespace{value: val, isSet: true}
}

func (v NullableFeatureExperimentRuleNamespace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureExperimentRuleNamespace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
