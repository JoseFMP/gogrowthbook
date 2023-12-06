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

// checks if the PutMetricRequestSql type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &PutMetricRequestSql{}

// PutMetricRequestSql Preferred way to define SQL. Only one of `sql`, `sqlBuilder` or `mixpanel` allowed.
type PutMetricRequestSql struct {
	IdentifierTypes []string `json:"identifierTypes,omitempty"`
	ConversionSQL *string `json:"conversionSQL,omitempty"`
	// Custom user level aggregation for your metric (default: `SUM(value)`)
	UserAggregationSQL *string `json:"userAggregationSQL,omitempty"`
	// The metric ID for a [denominator metric for funnel and ratio metrics](/app/metrics#denominator-ratio--funnel-metrics)
	DenominatorMetricId *string `json:"denominatorMetricId,omitempty"`
}

// NewPutMetricRequestSql instantiates a new PutMetricRequestSql object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutMetricRequestSql() *PutMetricRequestSql {
	this := PutMetricRequestSql{}
	return &this
}

// NewPutMetricRequestSqlWithDefaults instantiates a new PutMetricRequestSql object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutMetricRequestSqlWithDefaults() *PutMetricRequestSql {
	this := PutMetricRequestSql{}
	return &this
}

// GetIdentifierTypes returns the IdentifierTypes field value if set, zero value otherwise.
func (o *PutMetricRequestSql) GetIdentifierTypes() []string {
	if o == nil || common.IsNil(o.IdentifierTypes) {
		var ret []string
		return ret
	}
	return o.IdentifierTypes
}

// GetIdentifierTypesOk returns a tuple with the IdentifierTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutMetricRequestSql) GetIdentifierTypesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.IdentifierTypes) {
		return nil, false
	}
	return o.IdentifierTypes, true
}

// HasIdentifierTypes returns a boolean if a field has been set.
func (o *PutMetricRequestSql) HasIdentifierTypes() bool {
	if o != nil && !common.IsNil(o.IdentifierTypes) {
		return true
	}

	return false
}

// SetIdentifierTypes gets a reference to the given []string and assigns it to the IdentifierTypes field.
func (o *PutMetricRequestSql) SetIdentifierTypes(v []string) {
	o.IdentifierTypes = v
}

// GetConversionSQL returns the ConversionSQL field value if set, zero value otherwise.
func (o *PutMetricRequestSql) GetConversionSQL() string {
	if o == nil || common.IsNil(o.ConversionSQL) {
		var ret string
		return ret
	}
	return *o.ConversionSQL
}

// GetConversionSQLOk returns a tuple with the ConversionSQL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutMetricRequestSql) GetConversionSQLOk() (*string, bool) {
	if o == nil || common.IsNil(o.ConversionSQL) {
		return nil, false
	}
	return o.ConversionSQL, true
}

// HasConversionSQL returns a boolean if a field has been set.
func (o *PutMetricRequestSql) HasConversionSQL() bool {
	if o != nil && !common.IsNil(o.ConversionSQL) {
		return true
	}

	return false
}

// SetConversionSQL gets a reference to the given string and assigns it to the ConversionSQL field.
func (o *PutMetricRequestSql) SetConversionSQL(v string) {
	o.ConversionSQL = &v
}

// GetUserAggregationSQL returns the UserAggregationSQL field value if set, zero value otherwise.
func (o *PutMetricRequestSql) GetUserAggregationSQL() string {
	if o == nil || common.IsNil(o.UserAggregationSQL) {
		var ret string
		return ret
	}
	return *o.UserAggregationSQL
}

// GetUserAggregationSQLOk returns a tuple with the UserAggregationSQL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutMetricRequestSql) GetUserAggregationSQLOk() (*string, bool) {
	if o == nil || common.IsNil(o.UserAggregationSQL) {
		return nil, false
	}
	return o.UserAggregationSQL, true
}

// HasUserAggregationSQL returns a boolean if a field has been set.
func (o *PutMetricRequestSql) HasUserAggregationSQL() bool {
	if o != nil && !common.IsNil(o.UserAggregationSQL) {
		return true
	}

	return false
}

// SetUserAggregationSQL gets a reference to the given string and assigns it to the UserAggregationSQL field.
func (o *PutMetricRequestSql) SetUserAggregationSQL(v string) {
	o.UserAggregationSQL = &v
}

// GetDenominatorMetricId returns the DenominatorMetricId field value if set, zero value otherwise.
func (o *PutMetricRequestSql) GetDenominatorMetricId() string {
	if o == nil || common.IsNil(o.DenominatorMetricId) {
		var ret string
		return ret
	}
	return *o.DenominatorMetricId
}

// GetDenominatorMetricIdOk returns a tuple with the DenominatorMetricId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutMetricRequestSql) GetDenominatorMetricIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.DenominatorMetricId) {
		return nil, false
	}
	return o.DenominatorMetricId, true
}

// HasDenominatorMetricId returns a boolean if a field has been set.
func (o *PutMetricRequestSql) HasDenominatorMetricId() bool {
	if o != nil && !common.IsNil(o.DenominatorMetricId) {
		return true
	}

	return false
}

// SetDenominatorMetricId gets a reference to the given string and assigns it to the DenominatorMetricId field.
func (o *PutMetricRequestSql) SetDenominatorMetricId(v string) {
	o.DenominatorMetricId = &v
}

func (o PutMetricRequestSql) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutMetricRequestSql) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.IdentifierTypes) {
		toSerialize["identifierTypes"] = o.IdentifierTypes
	}
	if !common.IsNil(o.ConversionSQL) {
		toSerialize["conversionSQL"] = o.ConversionSQL
	}
	if !common.IsNil(o.UserAggregationSQL) {
		toSerialize["userAggregationSQL"] = o.UserAggregationSQL
	}
	if !common.IsNil(o.DenominatorMetricId) {
		toSerialize["denominatorMetricId"] = o.DenominatorMetricId
	}
	return toSerialize, nil
}

type NullablePutMetricRequestSql struct {
	value *PutMetricRequestSql
	isSet bool
}

func (v NullablePutMetricRequestSql) Get() *PutMetricRequestSql {
	return v.value
}

func (v *NullablePutMetricRequestSql) Set(val *PutMetricRequestSql) {
	v.value = val
	v.isSet = true
}

func (v NullablePutMetricRequestSql) IsSet() bool {
	return v.isSet
}

func (v *NullablePutMetricRequestSql) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutMetricRequestSql(val *PutMetricRequestSql) *NullablePutMetricRequestSql {
	return &NullablePutMetricRequestSql{value: val, isSet: true}
}

func (v NullablePutMetricRequestSql) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutMetricRequestSql) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

