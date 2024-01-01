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

// checks if the PostFeatureRequestEnvironmentsValue type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &PostFeatureRequestEnvironmentsValue{}

// PostFeatureRequestEnvironmentsValue struct for PostFeatureRequestEnvironmentsValue
type PostFeatureRequestEnvironmentsValue struct {
	Enabled bool           `json:"enabled"`
	Rules   []*FeatureRule `json:"rules"`
	// A JSON stringified [FeatureDefinition](#tag/FeatureDefinition_model)
	Definition *string                                   `json:"definition,omitempty"`
	Draft      *PostFeatureRequestEnvironmentsValueDraft `json:"draft,omitempty"`
}

type _PostFeatureRequestEnvironmentsValue PostFeatureRequestEnvironmentsValue

// NewPostFeatureRequestEnvironmentsValue instantiates a new PostFeatureRequestEnvironmentsValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostFeatureRequestEnvironmentsValue(enabled bool, rules []*FeatureRule) *PostFeatureRequestEnvironmentsValue {
	this := PostFeatureRequestEnvironmentsValue{}
	this.Enabled = enabled
	this.Rules = rules
	return &this
}

// NewPostFeatureRequestEnvironmentsValueWithDefaults instantiates a new PostFeatureRequestEnvironmentsValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostFeatureRequestEnvironmentsValueWithDefaults() *PostFeatureRequestEnvironmentsValue {
	this := PostFeatureRequestEnvironmentsValue{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *PostFeatureRequestEnvironmentsValue) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *PostFeatureRequestEnvironmentsValue) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *PostFeatureRequestEnvironmentsValue) SetEnabled(v bool) {
	o.Enabled = v
}

// GetRules returns the Rules field value
func (o *PostFeatureRequestEnvironmentsValue) GetRules() []*FeatureRule {
	if o == nil {
		var ret []*FeatureRule
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *PostFeatureRequestEnvironmentsValue) GetRulesOk() ([]*FeatureRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *PostFeatureRequestEnvironmentsValue) SetRules(v []*FeatureRule) {
	o.Rules = v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *PostFeatureRequestEnvironmentsValue) GetDefinition() string {
	if o == nil || common.IsNil(o.Definition) {
		var ret string
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostFeatureRequestEnvironmentsValue) GetDefinitionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Definition) {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *PostFeatureRequestEnvironmentsValue) HasDefinition() bool {
	if o != nil && !common.IsNil(o.Definition) {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given string and assigns it to the Definition field.
func (o *PostFeatureRequestEnvironmentsValue) SetDefinition(v string) {
	o.Definition = &v
}

// GetDraft returns the Draft field value if set, zero value otherwise.
func (o *PostFeatureRequestEnvironmentsValue) GetDraft() PostFeatureRequestEnvironmentsValueDraft {
	if o == nil || common.IsNil(o.Draft) {
		var ret PostFeatureRequestEnvironmentsValueDraft
		return ret
	}
	return *o.Draft
}

// GetDraftOk returns a tuple with the Draft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostFeatureRequestEnvironmentsValue) GetDraftOk() (*PostFeatureRequestEnvironmentsValueDraft, bool) {
	if o == nil || common.IsNil(o.Draft) {
		return nil, false
	}
	return o.Draft, true
}

// HasDraft returns a boolean if a field has been set.
func (o *PostFeatureRequestEnvironmentsValue) HasDraft() bool {
	if o != nil && !common.IsNil(o.Draft) {
		return true
	}

	return false
}

// SetDraft gets a reference to the given PostFeatureRequestEnvironmentsValueDraft and assigns it to the Draft field.
func (o *PostFeatureRequestEnvironmentsValue) SetDraft(v PostFeatureRequestEnvironmentsValueDraft) {
	o.Draft = &v
}

func (o PostFeatureRequestEnvironmentsValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostFeatureRequestEnvironmentsValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	toSerialize["rules"] = o.Rules
	if !common.IsNil(o.Definition) {
		toSerialize["definition"] = o.Definition
	}
	if !common.IsNil(o.Draft) {
		toSerialize["draft"] = o.Draft
	}
	return toSerialize, nil
}

func (o *PostFeatureRequestEnvironmentsValue) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
		"rules",
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

	varPostFeatureRequestEnvironmentsValue := _PostFeatureRequestEnvironmentsValue{}

	err = json.Unmarshal(bytes, &varPostFeatureRequestEnvironmentsValue)

	if err != nil {
		return err
	}

	*o = PostFeatureRequestEnvironmentsValue(varPostFeatureRequestEnvironmentsValue)

	return err
}

type NullablePostFeatureRequestEnvironmentsValue struct {
	value *PostFeatureRequestEnvironmentsValue
	isSet bool
}

func (v NullablePostFeatureRequestEnvironmentsValue) Get() *PostFeatureRequestEnvironmentsValue {
	return v.value
}

func (v *NullablePostFeatureRequestEnvironmentsValue) Set(val *PostFeatureRequestEnvironmentsValue) {
	v.value = val
	v.isSet = true
}

func (v NullablePostFeatureRequestEnvironmentsValue) IsSet() bool {
	return v.isSet
}

func (v *NullablePostFeatureRequestEnvironmentsValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostFeatureRequestEnvironmentsValue(val *PostFeatureRequestEnvironmentsValue) *NullablePostFeatureRequestEnvironmentsValue {
	return &NullablePostFeatureRequestEnvironmentsValue{value: val, isSet: true}
}

func (v NullablePostFeatureRequestEnvironmentsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostFeatureRequestEnvironmentsValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
