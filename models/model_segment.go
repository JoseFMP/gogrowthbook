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

// checks if the Segment type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &Segment{}

// Segment struct for Segment
type Segment struct {
	Id             string `json:"id"`
	DateCreated    string `json:"dateCreated"`
	DateUpdated    string `json:"dateUpdated"`
	Owner          string `json:"owner"`
	DatasourceId   string `json:"datasourceId"`
	IdentifierType string `json:"identifierType"`
	Name           string `json:"name"`
	Query          string `json:"query"`
}

type _Segment Segment

// NewSegment instantiates a new Segment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegment(id string, dateCreated string, dateUpdated string, owner string, datasourceId string, identifierType string, name string, query string) *Segment {
	this := Segment{}
	this.Id = id
	this.DateCreated = dateCreated
	this.DateUpdated = dateUpdated
	this.Owner = owner
	this.DatasourceId = datasourceId
	this.IdentifierType = identifierType
	this.Name = name
	this.Query = query
	return &this
}

// NewSegmentWithDefaults instantiates a new Segment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentWithDefaults() *Segment {
	this := Segment{}
	return &this
}

// GetId returns the Id field value
func (o *Segment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Segment) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Segment) SetId(v string) {
	o.Id = v
}

// GetDateCreated returns the DateCreated field value
func (o *Segment) GetDateCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value
// and a boolean to check if the value has been set.
func (o *Segment) GetDateCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateCreated, true
}

// SetDateCreated sets field value
func (o *Segment) SetDateCreated(v string) {
	o.DateCreated = v
}

// GetDateUpdated returns the DateUpdated field value
func (o *Segment) GetDateUpdated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateUpdated
}

// GetDateUpdatedOk returns a tuple with the DateUpdated field value
// and a boolean to check if the value has been set.
func (o *Segment) GetDateUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateUpdated, true
}

// SetDateUpdated sets field value
func (o *Segment) SetDateUpdated(v string) {
	o.DateUpdated = v
}

// GetOwner returns the Owner field value
func (o *Segment) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *Segment) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *Segment) SetOwner(v string) {
	o.Owner = v
}

// GetDatasourceId returns the DatasourceId field value
func (o *Segment) GetDatasourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatasourceId
}

// GetDatasourceIdOk returns a tuple with the DatasourceId field value
// and a boolean to check if the value has been set.
func (o *Segment) GetDatasourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasourceId, true
}

// SetDatasourceId sets field value
func (o *Segment) SetDatasourceId(v string) {
	o.DatasourceId = v
}

// GetIdentifierType returns the IdentifierType field value
func (o *Segment) GetIdentifierType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentifierType
}

// GetIdentifierTypeOk returns a tuple with the IdentifierType field value
// and a boolean to check if the value has been set.
func (o *Segment) GetIdentifierTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentifierType, true
}

// SetIdentifierType sets field value
func (o *Segment) SetIdentifierType(v string) {
	o.IdentifierType = v
}

// GetName returns the Name field value
func (o *Segment) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Segment) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Segment) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value
func (o *Segment) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *Segment) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *Segment) SetQuery(v string) {
	o.Query = v
}

func (o Segment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Segment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["dateCreated"] = o.DateCreated
	toSerialize["dateUpdated"] = o.DateUpdated
	toSerialize["owner"] = o.Owner
	toSerialize["datasourceId"] = o.DatasourceId
	toSerialize["identifierType"] = o.IdentifierType
	toSerialize["name"] = o.Name
	toSerialize["query"] = o.Query
	return toSerialize, nil
}

func (o *Segment) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"dateCreated",
		"dateUpdated",
		"owner",
		"datasourceId",
		"identifierType",
		"name",
		"query",
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

	varSegment := _Segment{}

	err = json.Unmarshal(bytes, &varSegment)

	if err != nil {
		return err
	}

	*o = Segment(varSegment)

	return err
}

type NullableSegment struct {
	value *Segment
	isSet bool
}

func (v NullableSegment) Get() *Segment {
	return v.value
}

func (v *NullableSegment) Set(val *Segment) {
	v.value = val
	v.isSet = true
}

func (v NullableSegment) IsSet() bool {
	return v.isSet
}

func (v *NullableSegment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegment(val *Segment) *NullableSegment {
	return &NullableSegment{value: val, isSet: true}
}

func (v NullableSegment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
