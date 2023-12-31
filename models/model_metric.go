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

// checks if the Metric type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &Metric{}

// Metric struct for Metric
type Metric struct {
	Id           string            `json:"id"`
	DateCreated  string            `json:"dateCreated"`
	DateUpdated  string            `json:"dateUpdated"`
	Owner        string            `json:"owner"`
	DatasourceId string            `json:"datasourceId"`
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	Type         string            `json:"type"`
	Tags         []string          `json:"tags"`
	Projects     []string          `json:"projects"`
	Archived     bool              `json:"archived"`
	Behavior     MetricBehavior    `json:"behavior"`
	Sql          *MetricSql        `json:"sql,omitempty"`
	SqlBuilder   *MetricSqlBuilder `json:"sqlBuilder,omitempty"`
	Mixpanel     *MetricMixpanel   `json:"mixpanel,omitempty"`
}

type _Metric Metric

// NewMetric instantiates a new Metric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetric(id string, dateCreated string, dateUpdated string, owner string, datasourceId string, name string, description string, type_ string, tags []string, projects []string, archived bool, behavior MetricBehavior) *Metric {
	this := Metric{}
	this.Id = id
	this.DateCreated = dateCreated
	this.DateUpdated = dateUpdated
	this.Owner = owner
	this.DatasourceId = datasourceId
	this.Name = name
	this.Description = description
	this.Type = type_
	this.Tags = tags
	this.Projects = projects
	this.Archived = archived
	this.Behavior = behavior
	return &this
}

// NewMetricWithDefaults instantiates a new Metric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricWithDefaults() *Metric {
	this := Metric{}
	return &this
}

// GetId returns the Id field value
func (o *Metric) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Metric) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Metric) SetId(v string) {
	o.Id = v
}

// GetDateCreated returns the DateCreated field value
func (o *Metric) GetDateCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value
// and a boolean to check if the value has been set.
func (o *Metric) GetDateCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateCreated, true
}

// SetDateCreated sets field value
func (o *Metric) SetDateCreated(v string) {
	o.DateCreated = v
}

// GetDateUpdated returns the DateUpdated field value
func (o *Metric) GetDateUpdated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateUpdated
}

// GetDateUpdatedOk returns a tuple with the DateUpdated field value
// and a boolean to check if the value has been set.
func (o *Metric) GetDateUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateUpdated, true
}

// SetDateUpdated sets field value
func (o *Metric) SetDateUpdated(v string) {
	o.DateUpdated = v
}

// GetOwner returns the Owner field value
func (o *Metric) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *Metric) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *Metric) SetOwner(v string) {
	o.Owner = v
}

// GetDatasourceId returns the DatasourceId field value
func (o *Metric) GetDatasourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatasourceId
}

// GetDatasourceIdOk returns a tuple with the DatasourceId field value
// and a boolean to check if the value has been set.
func (o *Metric) GetDatasourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasourceId, true
}

// SetDatasourceId sets field value
func (o *Metric) SetDatasourceId(v string) {
	o.DatasourceId = v
}

// GetName returns the Name field value
func (o *Metric) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Metric) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Metric) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *Metric) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Metric) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Metric) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value
func (o *Metric) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Metric) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Metric) SetType(v string) {
	o.Type = v
}

// GetTags returns the Tags field value
func (o *Metric) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *Metric) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *Metric) SetTags(v []string) {
	o.Tags = v
}

// GetProjects returns the Projects field value
func (o *Metric) GetProjects() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value
// and a boolean to check if the value has been set.
func (o *Metric) GetProjectsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Projects, true
}

// SetProjects sets field value
func (o *Metric) SetProjects(v []string) {
	o.Projects = v
}

// GetArchived returns the Archived field value
func (o *Metric) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *Metric) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *Metric) SetArchived(v bool) {
	o.Archived = v
}

// GetBehavior returns the Behavior field value
func (o *Metric) GetBehavior() MetricBehavior {
	if o == nil {
		var ret MetricBehavior
		return ret
	}

	return o.Behavior
}

// GetBehaviorOk returns a tuple with the Behavior field value
// and a boolean to check if the value has been set.
func (o *Metric) GetBehaviorOk() (*MetricBehavior, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Behavior, true
}

// SetBehavior sets field value
func (o *Metric) SetBehavior(v MetricBehavior) {
	o.Behavior = v
}

// GetSql returns the Sql field value if set, zero value otherwise.
func (o *Metric) GetSql() MetricSql {
	if o == nil || common.IsNil(o.Sql) {
		var ret MetricSql
		return ret
	}
	return *o.Sql
}

// GetSqlOk returns a tuple with the Sql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metric) GetSqlOk() (*MetricSql, bool) {
	if o == nil || common.IsNil(o.Sql) {
		return nil, false
	}
	return o.Sql, true
}

// HasSql returns a boolean if a field has been set.
func (o *Metric) HasSql() bool {
	if o != nil && !common.IsNil(o.Sql) {
		return true
	}

	return false
}

// SetSql gets a reference to the given MetricSql and assigns it to the Sql field.
func (o *Metric) SetSql(v MetricSql) {
	o.Sql = &v
}

// GetSqlBuilder returns the SqlBuilder field value if set, zero value otherwise.
func (o *Metric) GetSqlBuilder() MetricSqlBuilder {
	if o == nil || common.IsNil(o.SqlBuilder) {
		var ret MetricSqlBuilder
		return ret
	}
	return *o.SqlBuilder
}

// GetSqlBuilderOk returns a tuple with the SqlBuilder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metric) GetSqlBuilderOk() (*MetricSqlBuilder, bool) {
	if o == nil || common.IsNil(o.SqlBuilder) {
		return nil, false
	}
	return o.SqlBuilder, true
}

// HasSqlBuilder returns a boolean if a field has been set.
func (o *Metric) HasSqlBuilder() bool {
	if o != nil && !common.IsNil(o.SqlBuilder) {
		return true
	}

	return false
}

// SetSqlBuilder gets a reference to the given MetricSqlBuilder and assigns it to the SqlBuilder field.
func (o *Metric) SetSqlBuilder(v MetricSqlBuilder) {
	o.SqlBuilder = &v
}

// GetMixpanel returns the Mixpanel field value if set, zero value otherwise.
func (o *Metric) GetMixpanel() MetricMixpanel {
	if o == nil || common.IsNil(o.Mixpanel) {
		var ret MetricMixpanel
		return ret
	}
	return *o.Mixpanel
}

// GetMixpanelOk returns a tuple with the Mixpanel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metric) GetMixpanelOk() (*MetricMixpanel, bool) {
	if o == nil || common.IsNil(o.Mixpanel) {
		return nil, false
	}
	return o.Mixpanel, true
}

// HasMixpanel returns a boolean if a field has been set.
func (o *Metric) HasMixpanel() bool {
	if o != nil && !common.IsNil(o.Mixpanel) {
		return true
	}

	return false
}

// SetMixpanel gets a reference to the given MetricMixpanel and assigns it to the Mixpanel field.
func (o *Metric) SetMixpanel(v MetricMixpanel) {
	o.Mixpanel = &v
}

func (o Metric) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Metric) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["dateCreated"] = o.DateCreated
	toSerialize["dateUpdated"] = o.DateUpdated
	toSerialize["owner"] = o.Owner
	toSerialize["datasourceId"] = o.DatasourceId
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["type"] = o.Type
	toSerialize["tags"] = o.Tags
	toSerialize["projects"] = o.Projects
	toSerialize["archived"] = o.Archived
	toSerialize["behavior"] = o.Behavior
	if !common.IsNil(o.Sql) {
		toSerialize["sql"] = o.Sql
	}
	if !common.IsNil(o.SqlBuilder) {
		toSerialize["sqlBuilder"] = o.SqlBuilder
	}
	if !common.IsNil(o.Mixpanel) {
		toSerialize["mixpanel"] = o.Mixpanel
	}
	return toSerialize, nil
}

func (o *Metric) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"dateCreated",
		"dateUpdated",
		"owner",
		"datasourceId",
		"name",
		"description",
		"type",
		"tags",
		"projects",
		"archived",
		"behavior",
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

	varMetric := _Metric{}

	err = json.Unmarshal(bytes, &varMetric)

	if err != nil {
		return err
	}

	*o = Metric(varMetric)

	return err
}

type NullableMetric struct {
	value *Metric
	isSet bool
}

func (v NullableMetric) Get() *Metric {
	return v.value
}

func (v *NullableMetric) Set(val *Metric) {
	v.value = val
	v.isSet = true
}

func (v NullableMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetric(val *Metric) *NullableMetric {
	return &NullableMetric{value: val, isSet: true}
}

func (v NullableMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
