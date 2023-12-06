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

// checks if the PostMetricRequestMixpanel type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &PostMetricRequestMixpanel{}

// PostMetricRequestMixpanel Only use for MixPanel (non-SQL) Data Sources. Only one of `sql`, `sqlBuilder` or `mixpanel` allowed, and at least one must be specified.
type PostMetricRequestMixpanel struct {
	EventName string `json:"eventName"`
	EventValue *string `json:"eventValue,omitempty"`
	UserAggregation string `json:"userAggregation"`
	Conditions []PostMetricRequestMixpanelConditionsInner `json:"conditions,omitempty"`
}

type _PostMetricRequestMixpanel PostMetricRequestMixpanel

// NewPostMetricRequestMixpanel instantiates a new PostMetricRequestMixpanel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostMetricRequestMixpanel(eventName string, userAggregation string) *PostMetricRequestMixpanel {
	this := PostMetricRequestMixpanel{}
	this.EventName = eventName
	this.UserAggregation = userAggregation
	return &this
}

// NewPostMetricRequestMixpanelWithDefaults instantiates a new PostMetricRequestMixpanel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostMetricRequestMixpanelWithDefaults() *PostMetricRequestMixpanel {
	this := PostMetricRequestMixpanel{}
	return &this
}

// GetEventName returns the EventName field value
func (o *PostMetricRequestMixpanel) GetEventName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value
// and a boolean to check if the value has been set.
func (o *PostMetricRequestMixpanel) GetEventNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventName, true
}

// SetEventName sets field value
func (o *PostMetricRequestMixpanel) SetEventName(v string) {
	o.EventName = v
}

// GetEventValue returns the EventValue field value if set, zero value otherwise.
func (o *PostMetricRequestMixpanel) GetEventValue() string {
	if o == nil || common.IsNil(o.EventValue) {
		var ret string
		return ret
	}
	return *o.EventValue
}

// GetEventValueOk returns a tuple with the EventValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMetricRequestMixpanel) GetEventValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.EventValue) {
		return nil, false
	}
	return o.EventValue, true
}

// HasEventValue returns a boolean if a field has been set.
func (o *PostMetricRequestMixpanel) HasEventValue() bool {
	if o != nil && !common.IsNil(o.EventValue) {
		return true
	}

	return false
}

// SetEventValue gets a reference to the given string and assigns it to the EventValue field.
func (o *PostMetricRequestMixpanel) SetEventValue(v string) {
	o.EventValue = &v
}

// GetUserAggregation returns the UserAggregation field value
func (o *PostMetricRequestMixpanel) GetUserAggregation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserAggregation
}

// GetUserAggregationOk returns a tuple with the UserAggregation field value
// and a boolean to check if the value has been set.
func (o *PostMetricRequestMixpanel) GetUserAggregationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserAggregation, true
}

// SetUserAggregation sets field value
func (o *PostMetricRequestMixpanel) SetUserAggregation(v string) {
	o.UserAggregation = v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *PostMetricRequestMixpanel) GetConditions() []PostMetricRequestMixpanelConditionsInner {
	if o == nil || common.IsNil(o.Conditions) {
		var ret []PostMetricRequestMixpanelConditionsInner
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMetricRequestMixpanel) GetConditionsOk() ([]PostMetricRequestMixpanelConditionsInner, bool) {
	if o == nil || common.IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *PostMetricRequestMixpanel) HasConditions() bool {
	if o != nil && !common.IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []PostMetricRequestMixpanelConditionsInner and assigns it to the Conditions field.
func (o *PostMetricRequestMixpanel) SetConditions(v []PostMetricRequestMixpanelConditionsInner) {
	o.Conditions = v
}

func (o PostMetricRequestMixpanel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostMetricRequestMixpanel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventName"] = o.EventName
	if !common.IsNil(o.EventValue) {
		toSerialize["eventValue"] = o.EventValue
	}
	toSerialize["userAggregation"] = o.UserAggregation
	if !common.IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	return toSerialize, nil
}

func (o *PostMetricRequestMixpanel) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"eventName",
		"userAggregation",
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

	varPostMetricRequestMixpanel := _PostMetricRequestMixpanel{}

	err = json.Unmarshal(bytes, &varPostMetricRequestMixpanel)

	if err != nil {
		return err
	}

	*o = PostMetricRequestMixpanel(varPostMetricRequestMixpanel)

	return err
}

type NullablePostMetricRequestMixpanel struct {
	value *PostMetricRequestMixpanel
	isSet bool
}

func (v NullablePostMetricRequestMixpanel) Get() *PostMetricRequestMixpanel {
	return v.value
}

func (v *NullablePostMetricRequestMixpanel) Set(val *PostMetricRequestMixpanel) {
	v.value = val
	v.isSet = true
}

func (v NullablePostMetricRequestMixpanel) IsSet() bool {
	return v.isSet
}

func (v *NullablePostMetricRequestMixpanel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostMetricRequestMixpanel(val *PostMetricRequestMixpanel) *NullablePostMetricRequestMixpanel {
	return &NullablePostMetricRequestMixpanel{value: val, isSet: true}
}

func (v NullablePostMetricRequestMixpanel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostMetricRequestMixpanel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

