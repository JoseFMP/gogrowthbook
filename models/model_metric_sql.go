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

// checks if the MetricSql type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &MetricSql{}

// MetricSql struct for MetricSql
type MetricSql struct {
	IdentifierTypes []string `json:"identifierTypes"`
	ConversionSQL string `json:"conversionSQL"`
	UserAggregationSQL string `json:"userAggregationSQL"`
	DenominatorMetricId string `json:"denominatorMetricId"`
}

type _MetricSql MetricSql

// NewMetricSql instantiates a new MetricSql object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricSql(identifierTypes []string, conversionSQL string, userAggregationSQL string, denominatorMetricId string) *MetricSql {
	this := MetricSql{}
	this.IdentifierTypes = identifierTypes
	this.ConversionSQL = conversionSQL
	this.UserAggregationSQL = userAggregationSQL
	this.DenominatorMetricId = denominatorMetricId
	return &this
}

// NewMetricSqlWithDefaults instantiates a new MetricSql object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricSqlWithDefaults() *MetricSql {
	this := MetricSql{}
	return &this
}

// GetIdentifierTypes returns the IdentifierTypes field value
func (o *MetricSql) GetIdentifierTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.IdentifierTypes
}

// GetIdentifierTypesOk returns a tuple with the IdentifierTypes field value
// and a boolean to check if the value has been set.
func (o *MetricSql) GetIdentifierTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdentifierTypes, true
}

// SetIdentifierTypes sets field value
func (o *MetricSql) SetIdentifierTypes(v []string) {
	o.IdentifierTypes = v
}

// GetConversionSQL returns the ConversionSQL field value
func (o *MetricSql) GetConversionSQL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConversionSQL
}

// GetConversionSQLOk returns a tuple with the ConversionSQL field value
// and a boolean to check if the value has been set.
func (o *MetricSql) GetConversionSQLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversionSQL, true
}

// SetConversionSQL sets field value
func (o *MetricSql) SetConversionSQL(v string) {
	o.ConversionSQL = v
}

// GetUserAggregationSQL returns the UserAggregationSQL field value
func (o *MetricSql) GetUserAggregationSQL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserAggregationSQL
}

// GetUserAggregationSQLOk returns a tuple with the UserAggregationSQL field value
// and a boolean to check if the value has been set.
func (o *MetricSql) GetUserAggregationSQLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserAggregationSQL, true
}

// SetUserAggregationSQL sets field value
func (o *MetricSql) SetUserAggregationSQL(v string) {
	o.UserAggregationSQL = v
}

// GetDenominatorMetricId returns the DenominatorMetricId field value
func (o *MetricSql) GetDenominatorMetricId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DenominatorMetricId
}

// GetDenominatorMetricIdOk returns a tuple with the DenominatorMetricId field value
// and a boolean to check if the value has been set.
func (o *MetricSql) GetDenominatorMetricIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DenominatorMetricId, true
}

// SetDenominatorMetricId sets field value
func (o *MetricSql) SetDenominatorMetricId(v string) {
	o.DenominatorMetricId = v
}

func (o MetricSql) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricSql) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identifierTypes"] = o.IdentifierTypes
	toSerialize["conversionSQL"] = o.ConversionSQL
	toSerialize["userAggregationSQL"] = o.UserAggregationSQL
	toSerialize["denominatorMetricId"] = o.DenominatorMetricId
	return toSerialize, nil
}

func (o *MetricSql) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"identifierTypes",
		"conversionSQL",
		"userAggregationSQL",
		"denominatorMetricId",
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

	varMetricSql := _MetricSql{}

	err = json.Unmarshal(bytes, &varMetricSql)

	if err != nil {
		return err
	}

	*o = MetricSql(varMetricSql)

	return err
}

type NullableMetricSql struct {
	value *MetricSql
	isSet bool
}

func (v NullableMetricSql) Get() *MetricSql {
	return v.value
}

func (v *NullableMetricSql) Set(val *MetricSql) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricSql) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricSql) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricSql(val *MetricSql) *NullableMetricSql {
	return &NullableMetricSql{value: val, isSet: true}
}

func (v NullableMetricSql) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricSql) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


