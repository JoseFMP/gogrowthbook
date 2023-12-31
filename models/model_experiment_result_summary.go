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

// checks if the ExperimentResultSummary type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &ExperimentResultSummary{}

// ExperimentResultSummary struct for ExperimentResultSummary
type ExperimentResultSummary struct {
	Status              interface{} `json:"status"`
	Winner              interface{} `json:"winner"`
	Conclusions         interface{} `json:"conclusions"`
	ReleasedVariationId interface{} `json:"releasedVariationId"`
	ExcludeFromPayload  interface{} `json:"excludeFromPayload"`
}

type _ExperimentResultSummary ExperimentResultSummary

// NewExperimentResultSummary instantiates a new ExperimentResultSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExperimentResultSummary(status interface{}, winner interface{}, conclusions interface{}, releasedVariationId interface{}, excludeFromPayload interface{}) *ExperimentResultSummary {
	this := ExperimentResultSummary{}
	this.Status = status
	this.Winner = winner
	this.Conclusions = conclusions
	this.ReleasedVariationId = releasedVariationId
	this.ExcludeFromPayload = excludeFromPayload
	return &this
}

// NewExperimentResultSummaryWithDefaults instantiates a new ExperimentResultSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExperimentResultSummaryWithDefaults() *ExperimentResultSummary {
	this := ExperimentResultSummary{}
	return &this
}

// GetStatus returns the Status field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExperimentResultSummary) GetStatus() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExperimentResultSummary) GetStatusOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ExperimentResultSummary) SetStatus(v interface{}) {
	o.Status = v
}

// GetWinner returns the Winner field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExperimentResultSummary) GetWinner() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Winner
}

// GetWinnerOk returns a tuple with the Winner field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExperimentResultSummary) GetWinnerOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Winner) {
		return nil, false
	}
	return &o.Winner, true
}

// SetWinner sets field value
func (o *ExperimentResultSummary) SetWinner(v interface{}) {
	o.Winner = v
}

// GetConclusions returns the Conclusions field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExperimentResultSummary) GetConclusions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Conclusions
}

// GetConclusionsOk returns a tuple with the Conclusions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExperimentResultSummary) GetConclusionsOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Conclusions) {
		return nil, false
	}
	return &o.Conclusions, true
}

// SetConclusions sets field value
func (o *ExperimentResultSummary) SetConclusions(v interface{}) {
	o.Conclusions = v
}

// GetReleasedVariationId returns the ReleasedVariationId field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExperimentResultSummary) GetReleasedVariationId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.ReleasedVariationId
}

// GetReleasedVariationIdOk returns a tuple with the ReleasedVariationId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExperimentResultSummary) GetReleasedVariationIdOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.ReleasedVariationId) {
		return nil, false
	}
	return &o.ReleasedVariationId, true
}

// SetReleasedVariationId sets field value
func (o *ExperimentResultSummary) SetReleasedVariationId(v interface{}) {
	o.ReleasedVariationId = v
}

// GetExcludeFromPayload returns the ExcludeFromPayload field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExperimentResultSummary) GetExcludeFromPayload() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.ExcludeFromPayload
}

// GetExcludeFromPayloadOk returns a tuple with the ExcludeFromPayload field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExperimentResultSummary) GetExcludeFromPayloadOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.ExcludeFromPayload) {
		return nil, false
	}
	return &o.ExcludeFromPayload, true
}

// SetExcludeFromPayload sets field value
func (o *ExperimentResultSummary) SetExcludeFromPayload(v interface{}) {
	o.ExcludeFromPayload = v
}

func (o ExperimentResultSummary) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExperimentResultSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Winner != nil {
		toSerialize["winner"] = o.Winner
	}
	if o.Conclusions != nil {
		toSerialize["conclusions"] = o.Conclusions
	}
	if o.ReleasedVariationId != nil {
		toSerialize["releasedVariationId"] = o.ReleasedVariationId
	}
	if o.ExcludeFromPayload != nil {
		toSerialize["excludeFromPayload"] = o.ExcludeFromPayload
	}
	return toSerialize, nil
}

func (o *ExperimentResultSummary) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"winner",
		"conclusions",
		"releasedVariationId",
		"excludeFromPayload",
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

	varExperimentResultSummary := _ExperimentResultSummary{}

	err = json.Unmarshal(bytes, &varExperimentResultSummary)

	if err != nil {
		return err
	}

	*o = ExperimentResultSummary(varExperimentResultSummary)

	return err
}

type NullableExperimentResultSummary struct {
	value *ExperimentResultSummary
	isSet bool
}

func (v NullableExperimentResultSummary) Get() *ExperimentResultSummary {
	return v.value
}

func (v *NullableExperimentResultSummary) Set(val *ExperimentResultSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableExperimentResultSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableExperimentResultSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExperimentResultSummary(val *ExperimentResultSummary) *NullableExperimentResultSummary {
	return &NullableExperimentResultSummary{value: val, isSet: true}
}

func (v NullableExperimentResultSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExperimentResultSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
