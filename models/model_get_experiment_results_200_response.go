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

// checks if the GetExperimentResults200Response type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &GetExperimentResults200Response{}

// GetExperimentResults200Response struct for GetExperimentResults200Response
type GetExperimentResults200Response struct {
	Result *ExperimentResults `json:"result,omitempty"`
}

// NewGetExperimentResults200Response instantiates a new GetExperimentResults200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExperimentResults200Response() *GetExperimentResults200Response {
	this := GetExperimentResults200Response{}
	return &this
}

// NewGetExperimentResults200ResponseWithDefaults instantiates a new GetExperimentResults200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExperimentResults200ResponseWithDefaults() *GetExperimentResults200Response {
	this := GetExperimentResults200Response{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetExperimentResults200Response) GetResult() ExperimentResults {
	if o == nil || common.IsNil(o.Result) {
		var ret ExperimentResults
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExperimentResults200Response) GetResultOk() (*ExperimentResults, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetExperimentResults200Response) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ExperimentResults and assigns it to the Result field.
func (o *GetExperimentResults200Response) SetResult(v ExperimentResults) {
	o.Result = &v
}

func (o GetExperimentResults200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetExperimentResults200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetExperimentResults200Response struct {
	value *GetExperimentResults200Response
	isSet bool
}

func (v NullableGetExperimentResults200Response) Get() *GetExperimentResults200Response {
	return v.value
}

func (v *NullableGetExperimentResults200Response) Set(val *GetExperimentResults200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExperimentResults200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExperimentResults200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExperimentResults200Response(val *GetExperimentResults200Response) *NullableGetExperimentResults200Response {
	return &NullableGetExperimentResults200Response{value: val, isSet: true}
}

func (v NullableGetExperimentResults200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExperimentResults200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

