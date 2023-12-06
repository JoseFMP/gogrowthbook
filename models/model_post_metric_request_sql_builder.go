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

// checks if the PostMetricRequestSqlBuilder type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &PostMetricRequestSqlBuilder{}

// PostMetricRequestSqlBuilder An alternative way to specify a SQL metric, rather than a full query. Using `sql` is preferred to `sqlBuilder`. Only one of `sql`, `sqlBuilder` or `mixpanel` allowed, and at least one must be specified.
type PostMetricRequestSqlBuilder struct {
	IdentifierTypeColumns []PostMetricRequestSqlBuilderIdentifierTypeColumnsInner `json:"identifierTypeColumns"`
	TableName string `json:"tableName"`
	ValueColumnName *string `json:"valueColumnName,omitempty"`
	TimestampColumnName string `json:"timestampColumnName"`
	Conditions []PostMetricRequestSqlBuilderConditionsInner `json:"conditions,omitempty"`
}

type _PostMetricRequestSqlBuilder PostMetricRequestSqlBuilder

// NewPostMetricRequestSqlBuilder instantiates a new PostMetricRequestSqlBuilder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostMetricRequestSqlBuilder(identifierTypeColumns []PostMetricRequestSqlBuilderIdentifierTypeColumnsInner, tableName string, timestampColumnName string) *PostMetricRequestSqlBuilder {
	this := PostMetricRequestSqlBuilder{}
	this.IdentifierTypeColumns = identifierTypeColumns
	this.TableName = tableName
	this.TimestampColumnName = timestampColumnName
	return &this
}

// NewPostMetricRequestSqlBuilderWithDefaults instantiates a new PostMetricRequestSqlBuilder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostMetricRequestSqlBuilderWithDefaults() *PostMetricRequestSqlBuilder {
	this := PostMetricRequestSqlBuilder{}
	return &this
}

// GetIdentifierTypeColumns returns the IdentifierTypeColumns field value
func (o *PostMetricRequestSqlBuilder) GetIdentifierTypeColumns() []PostMetricRequestSqlBuilderIdentifierTypeColumnsInner {
	if o == nil {
		var ret []PostMetricRequestSqlBuilderIdentifierTypeColumnsInner
		return ret
	}

	return o.IdentifierTypeColumns
}

// GetIdentifierTypeColumnsOk returns a tuple with the IdentifierTypeColumns field value
// and a boolean to check if the value has been set.
func (o *PostMetricRequestSqlBuilder) GetIdentifierTypeColumnsOk() ([]PostMetricRequestSqlBuilderIdentifierTypeColumnsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdentifierTypeColumns, true
}

// SetIdentifierTypeColumns sets field value
func (o *PostMetricRequestSqlBuilder) SetIdentifierTypeColumns(v []PostMetricRequestSqlBuilderIdentifierTypeColumnsInner) {
	o.IdentifierTypeColumns = v
}

// GetTableName returns the TableName field value
func (o *PostMetricRequestSqlBuilder) GetTableName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value
// and a boolean to check if the value has been set.
func (o *PostMetricRequestSqlBuilder) GetTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableName, true
}

// SetTableName sets field value
func (o *PostMetricRequestSqlBuilder) SetTableName(v string) {
	o.TableName = v
}

// GetValueColumnName returns the ValueColumnName field value if set, zero value otherwise.
func (o *PostMetricRequestSqlBuilder) GetValueColumnName() string {
	if o == nil || common.IsNil(o.ValueColumnName) {
		var ret string
		return ret
	}
	return *o.ValueColumnName
}

// GetValueColumnNameOk returns a tuple with the ValueColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMetricRequestSqlBuilder) GetValueColumnNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.ValueColumnName) {
		return nil, false
	}
	return o.ValueColumnName, true
}

// HasValueColumnName returns a boolean if a field has been set.
func (o *PostMetricRequestSqlBuilder) HasValueColumnName() bool {
	if o != nil && !common.IsNil(o.ValueColumnName) {
		return true
	}

	return false
}

// SetValueColumnName gets a reference to the given string and assigns it to the ValueColumnName field.
func (o *PostMetricRequestSqlBuilder) SetValueColumnName(v string) {
	o.ValueColumnName = &v
}

// GetTimestampColumnName returns the TimestampColumnName field value
func (o *PostMetricRequestSqlBuilder) GetTimestampColumnName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimestampColumnName
}

// GetTimestampColumnNameOk returns a tuple with the TimestampColumnName field value
// and a boolean to check if the value has been set.
func (o *PostMetricRequestSqlBuilder) GetTimestampColumnNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimestampColumnName, true
}

// SetTimestampColumnName sets field value
func (o *PostMetricRequestSqlBuilder) SetTimestampColumnName(v string) {
	o.TimestampColumnName = v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *PostMetricRequestSqlBuilder) GetConditions() []PostMetricRequestSqlBuilderConditionsInner {
	if o == nil || common.IsNil(o.Conditions) {
		var ret []PostMetricRequestSqlBuilderConditionsInner
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMetricRequestSqlBuilder) GetConditionsOk() ([]PostMetricRequestSqlBuilderConditionsInner, bool) {
	if o == nil || common.IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *PostMetricRequestSqlBuilder) HasConditions() bool {
	if o != nil && !common.IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []PostMetricRequestSqlBuilderConditionsInner and assigns it to the Conditions field.
func (o *PostMetricRequestSqlBuilder) SetConditions(v []PostMetricRequestSqlBuilderConditionsInner) {
	o.Conditions = v
}

func (o PostMetricRequestSqlBuilder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostMetricRequestSqlBuilder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identifierTypeColumns"] = o.IdentifierTypeColumns
	toSerialize["tableName"] = o.TableName
	if !common.IsNil(o.ValueColumnName) {
		toSerialize["valueColumnName"] = o.ValueColumnName
	}
	toSerialize["timestampColumnName"] = o.TimestampColumnName
	if !common.IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	return toSerialize, nil
}

func (o *PostMetricRequestSqlBuilder) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"identifierTypeColumns",
		"tableName",
		"timestampColumnName",
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

	varPostMetricRequestSqlBuilder := _PostMetricRequestSqlBuilder{}

	err = json.Unmarshal(bytes, &varPostMetricRequestSqlBuilder)

	if err != nil {
		return err
	}

	*o = PostMetricRequestSqlBuilder(varPostMetricRequestSqlBuilder)

	return err
}

type NullablePostMetricRequestSqlBuilder struct {
	value *PostMetricRequestSqlBuilder
	isSet bool
}

func (v NullablePostMetricRequestSqlBuilder) Get() *PostMetricRequestSqlBuilder {
	return v.value
}

func (v *NullablePostMetricRequestSqlBuilder) Set(val *PostMetricRequestSqlBuilder) {
	v.value = val
	v.isSet = true
}

func (v NullablePostMetricRequestSqlBuilder) IsSet() bool {
	return v.isSet
}

func (v *NullablePostMetricRequestSqlBuilder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostMetricRequestSqlBuilder(val *PostMetricRequestSqlBuilder) *NullablePostMetricRequestSqlBuilder {
	return &NullablePostMetricRequestSqlBuilder{value: val, isSet: true}
}

func (v NullablePostMetricRequestSqlBuilder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostMetricRequestSqlBuilder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

