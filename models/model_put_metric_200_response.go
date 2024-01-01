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

// checks if the PutMetric200Response type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &PutMetric200Response{}

// PutMetric200Response struct for PutMetric200Response
type PutMetric200Response struct {
	UpdatedId string `json:"updatedId"`
}

type _PutMetric200Response PutMetric200Response

// NewPutMetric200Response instantiates a new PutMetric200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutMetric200Response(updatedId string) *PutMetric200Response {
	this := PutMetric200Response{}
	this.UpdatedId = updatedId
	return &this
}

// NewPutMetric200ResponseWithDefaults instantiates a new PutMetric200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutMetric200ResponseWithDefaults() *PutMetric200Response {
	this := PutMetric200Response{}
	return &this
}

// GetUpdatedId returns the UpdatedId field value
func (o *PutMetric200Response) GetUpdatedId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedId
}

// GetUpdatedIdOk returns a tuple with the UpdatedId field value
// and a boolean to check if the value has been set.
func (o *PutMetric200Response) GetUpdatedIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedId, true
}

// SetUpdatedId sets field value
func (o *PutMetric200Response) SetUpdatedId(v string) {
	o.UpdatedId = v
}

func (o PutMetric200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutMetric200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["updatedId"] = o.UpdatedId
	return toSerialize, nil
}

func (o *PutMetric200Response) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"updatedId",
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

	varPutMetric200Response := _PutMetric200Response{}

	err = json.Unmarshal(bytes, &varPutMetric200Response)

	if err != nil {
		return err
	}

	*o = PutMetric200Response(varPutMetric200Response)

	return err
}

type NullablePutMetric200Response struct {
	value *PutMetric200Response
	isSet bool
}

func (v NullablePutMetric200Response) Get() *PutMetric200Response {
	return v.value
}

func (v *NullablePutMetric200Response) Set(val *PutMetric200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePutMetric200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePutMetric200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutMetric200Response(val *PutMetric200Response) *NullablePutMetric200Response {
	return &NullablePutMetric200Response{value: val, isSet: true}
}

func (v NullablePutMetric200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutMetric200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
