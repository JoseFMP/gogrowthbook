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

// checks if the Experiment type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &Experiment{}

// Experiment struct for Experiment
type Experiment struct {
	Id            string                     `json:"id"`
	DateCreated   interface{}                `json:"dateCreated"`
	DateUpdated   interface{}                `json:"dateUpdated"`
	Name          string                     `json:"name"`
	Project       string                     `json:"project"`
	Hypothesis    string                     `json:"hypothesis"`
	Description   string                     `json:"description"`
	Tags          interface{}                `json:"tags"`
	Owner         string                     `json:"owner"`
	Archived      bool                       `json:"archived"`
	Status        string                     `json:"status"`
	AutoRefresh   bool                       `json:"autoRefresh"`
	HashAttribute string                     `json:"hashAttribute"`
	HashVersion   float32                    `json:"hashVersion"`
	Variations    interface{}                `json:"variations"`
	Phases        interface{}                `json:"phases"`
	Settings      ExperimentAnalysisSettings `json:"settings"`
	ResultSummary *ExperimentResultSummary   `json:"resultSummary,omitempty"`
}

type _Experiment Experiment

// NewExperiment instantiates a new Experiment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExperiment(id string, dateCreated interface{}, dateUpdated interface{}, name string, project string, hypothesis string, description string, tags interface{}, owner string, archived bool, status string, autoRefresh bool, hashAttribute string, hashVersion float32, variations interface{}, phases interface{}, settings ExperimentAnalysisSettings) *Experiment {
	this := Experiment{}
	this.Id = id
	this.DateCreated = dateCreated
	this.DateUpdated = dateUpdated
	this.Name = name
	this.Project = project
	this.Hypothesis = hypothesis
	this.Description = description
	this.Tags = tags
	this.Owner = owner
	this.Archived = archived
	this.Status = status
	this.AutoRefresh = autoRefresh
	this.HashAttribute = hashAttribute
	this.HashVersion = hashVersion
	this.Variations = variations
	this.Phases = phases
	this.Settings = settings
	return &this
}

// NewExperimentWithDefaults instantiates a new Experiment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExperimentWithDefaults() *Experiment {
	this := Experiment{}
	return &this
}

// GetId returns the Id field value
func (o *Experiment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Experiment) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Experiment) SetId(v string) {
	o.Id = v
}

// GetDateCreated returns the DateCreated field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Experiment) GetDateCreated() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Experiment) GetDateCreatedOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.DateCreated) {
		return nil, false
	}
	return &o.DateCreated, true
}

// SetDateCreated sets field value
func (o *Experiment) SetDateCreated(v interface{}) {
	o.DateCreated = v
}

// GetDateUpdated returns the DateUpdated field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Experiment) GetDateUpdated() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.DateUpdated
}

// GetDateUpdatedOk returns a tuple with the DateUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Experiment) GetDateUpdatedOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.DateUpdated) {
		return nil, false
	}
	return &o.DateUpdated, true
}

// SetDateUpdated sets field value
func (o *Experiment) SetDateUpdated(v interface{}) {
	o.DateUpdated = v
}

// GetName returns the Name field value
func (o *Experiment) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Experiment) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Experiment) SetName(v string) {
	o.Name = v
}

// GetProject returns the Project field value
func (o *Experiment) GetProject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *Experiment) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *Experiment) SetProject(v string) {
	o.Project = v
}

// GetHypothesis returns the Hypothesis field value
func (o *Experiment) GetHypothesis() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hypothesis
}

// GetHypothesisOk returns a tuple with the Hypothesis field value
// and a boolean to check if the value has been set.
func (o *Experiment) GetHypothesisOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hypothesis, true
}

// SetHypothesis sets field value
func (o *Experiment) SetHypothesis(v string) {
	o.Hypothesis = v
}

// GetDescription returns the Description field value
func (o *Experiment) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Experiment) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Experiment) SetDescription(v string) {
	o.Description = v
}

// GetTags returns the Tags field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Experiment) GetTags() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Experiment) GetTagsOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Tags) {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value
func (o *Experiment) SetTags(v interface{}) {
	o.Tags = v
}

// GetOwner returns the Owner field value
func (o *Experiment) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *Experiment) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *Experiment) SetOwner(v string) {
	o.Owner = v
}

// GetArchived returns the Archived field value
func (o *Experiment) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *Experiment) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *Experiment) SetArchived(v bool) {
	o.Archived = v
}

// GetStatus returns the Status field value
func (o *Experiment) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Experiment) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Experiment) SetStatus(v string) {
	o.Status = v
}

// GetAutoRefresh returns the AutoRefresh field value
func (o *Experiment) GetAutoRefresh() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoRefresh
}

// GetAutoRefreshOk returns a tuple with the AutoRefresh field value
// and a boolean to check if the value has been set.
func (o *Experiment) GetAutoRefreshOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoRefresh, true
}

// SetAutoRefresh sets field value
func (o *Experiment) SetAutoRefresh(v bool) {
	o.AutoRefresh = v
}

// GetHashAttribute returns the HashAttribute field value
func (o *Experiment) GetHashAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HashAttribute
}

// GetHashAttributeOk returns a tuple with the HashAttribute field value
// and a boolean to check if the value has been set.
func (o *Experiment) GetHashAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HashAttribute, true
}

// SetHashAttribute sets field value
func (o *Experiment) SetHashAttribute(v string) {
	o.HashAttribute = v
}

// GetHashVersion returns the HashVersion field value
func (o *Experiment) GetHashVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.HashVersion
}

// GetHashVersionOk returns a tuple with the HashVersion field value
// and a boolean to check if the value has been set.
func (o *Experiment) GetHashVersionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HashVersion, true
}

// SetHashVersion sets field value
func (o *Experiment) SetHashVersion(v float32) {
	o.HashVersion = v
}

// GetVariations returns the Variations field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Experiment) GetVariations() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Variations
}

// GetVariationsOk returns a tuple with the Variations field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Experiment) GetVariationsOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Variations) {
		return nil, false
	}
	return &o.Variations, true
}

// SetVariations sets field value
func (o *Experiment) SetVariations(v interface{}) {
	o.Variations = v
}

// GetPhases returns the Phases field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Experiment) GetPhases() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Phases
}

// GetPhasesOk returns a tuple with the Phases field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Experiment) GetPhasesOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Phases) {
		return nil, false
	}
	return &o.Phases, true
}

// SetPhases sets field value
func (o *Experiment) SetPhases(v interface{}) {
	o.Phases = v
}

// GetSettings returns the Settings field value
func (o *Experiment) GetSettings() ExperimentAnalysisSettings {
	if o == nil {
		var ret ExperimentAnalysisSettings
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *Experiment) GetSettingsOk() (*ExperimentAnalysisSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *Experiment) SetSettings(v ExperimentAnalysisSettings) {
	o.Settings = v
}

// GetResultSummary returns the ResultSummary field value if set, zero value otherwise.
func (o *Experiment) GetResultSummary() ExperimentResultSummary {
	if o == nil || common.IsNil(o.ResultSummary) {
		var ret ExperimentResultSummary
		return ret
	}
	return *o.ResultSummary
}

// GetResultSummaryOk returns a tuple with the ResultSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Experiment) GetResultSummaryOk() (*ExperimentResultSummary, bool) {
	if o == nil || common.IsNil(o.ResultSummary) {
		return nil, false
	}
	return o.ResultSummary, true
}

// HasResultSummary returns a boolean if a field has been set.
func (o *Experiment) HasResultSummary() bool {
	if o != nil && !common.IsNil(o.ResultSummary) {
		return true
	}

	return false
}

// SetResultSummary gets a reference to the given ExperimentResultSummary and assigns it to the ResultSummary field.
func (o *Experiment) SetResultSummary(v ExperimentResultSummary) {
	o.ResultSummary = &v
}

func (o Experiment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Experiment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.DateCreated != nil {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if o.DateUpdated != nil {
		toSerialize["dateUpdated"] = o.DateUpdated
	}
	toSerialize["name"] = o.Name
	toSerialize["project"] = o.Project
	toSerialize["hypothesis"] = o.Hypothesis
	toSerialize["description"] = o.Description
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["owner"] = o.Owner
	toSerialize["archived"] = o.Archived
	toSerialize["status"] = o.Status
	toSerialize["autoRefresh"] = o.AutoRefresh
	toSerialize["hashAttribute"] = o.HashAttribute
	toSerialize["hashVersion"] = o.HashVersion
	if o.Variations != nil {
		toSerialize["variations"] = o.Variations
	}
	if o.Phases != nil {
		toSerialize["phases"] = o.Phases
	}
	toSerialize["settings"] = o.Settings
	if !common.IsNil(o.ResultSummary) {
		toSerialize["resultSummary"] = o.ResultSummary
	}
	return toSerialize, nil
}

func (o *Experiment) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"dateCreated",
		"dateUpdated",
		"name",
		"project",
		"hypothesis",
		"description",
		"tags",
		"owner",
		"archived",
		"status",
		"autoRefresh",
		"hashAttribute",
		"hashVersion",
		"variations",
		"phases",
		"settings",
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

	varExperiment := _Experiment{}

	err = json.Unmarshal(bytes, &varExperiment)

	if err != nil {
		return err
	}

	*o = Experiment(varExperiment)

	return err
}

type NullableExperiment struct {
	value *Experiment
	isSet bool
}

func (v NullableExperiment) Get() *Experiment {
	return v.value
}

func (v *NullableExperiment) Set(val *Experiment) {
	v.value = val
	v.isSet = true
}

func (v NullableExperiment) IsSet() bool {
	return v.isSet
}

func (v *NullableExperiment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExperiment(val *Experiment) *NullableExperiment {
	return &NullableExperiment{value: val, isSet: true}
}

func (v NullableExperiment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExperiment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
