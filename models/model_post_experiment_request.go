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

// checks if the PostExperimentRequest type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &PostExperimentRequest{}

// PostExperimentRequest struct for PostExperimentRequest
type PostExperimentRequest struct {
	// ID for the [DataSource](#tag/DataSource_model)
	DatasourceId string `json:"datasourceId"`
	// The ID property of one of the assignment query objects associated with the datasource
	AssignmentQueryId string `json:"assignmentQueryId"`
	TrackingKey string `json:"trackingKey"`
	// Name of the experiment
	Name string `json:"name"`
	// Project ID which the experiment belongs to
	Project *string `json:"project,omitempty"`
	// Hypothesis of the experiment
	Hypothesis *string `json:"hypothesis,omitempty"`
	// Description of the experiment
	Description *string `json:"description,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Metrics []string `json:"metrics,omitempty"`
	GuardrailMetrics interface{} `json:"guardrailMetrics,omitempty"`
	// Email of the person who owns this experiment
	Owner string `json:"owner"`
	Archived *bool `json:"archived,omitempty"`
	Status *string `json:"status,omitempty"`
	AutoRefresh *bool `json:"autoRefresh,omitempty"`
	HashAttribute *string `json:"hashAttribute,omitempty"`
	HashVersion *float32 `json:"hashVersion,omitempty"`
	ReleasedVariationId *string `json:"releasedVariationId,omitempty"`
	ExcludeFromPayload *bool `json:"excludeFromPayload,omitempty"`
	Variations []PostExperimentRequestVariationsInner `json:"variations"`
	Phases []PostExperimentRequestPhasesInner `json:"phases,omitempty"`
}

type _PostExperimentRequest PostExperimentRequest

// NewPostExperimentRequest instantiates a new PostExperimentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostExperimentRequest(datasourceId string, assignmentQueryId string, trackingKey string, name string, owner string, variations []PostExperimentRequestVariationsInner) *PostExperimentRequest {
	this := PostExperimentRequest{}
	this.DatasourceId = datasourceId
	this.AssignmentQueryId = assignmentQueryId
	this.TrackingKey = trackingKey
	this.Name = name
	this.Owner = owner
	this.Variations = variations
	return &this
}

// NewPostExperimentRequestWithDefaults instantiates a new PostExperimentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostExperimentRequestWithDefaults() *PostExperimentRequest {
	this := PostExperimentRequest{}
	return &this
}

// GetDatasourceId returns the DatasourceId field value
func (o *PostExperimentRequest) GetDatasourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatasourceId
}

// GetDatasourceIdOk returns a tuple with the DatasourceId field value
// and a boolean to check if the value has been set.
func (o *PostExperimentRequest) GetDatasourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasourceId, true
}

// SetDatasourceId sets field value
func (o *PostExperimentRequest) SetDatasourceId(v string) {
	o.DatasourceId = v
}

// GetAssignmentQueryId returns the AssignmentQueryId field value
func (o *PostExperimentRequest) GetAssignmentQueryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignmentQueryId
}

// GetAssignmentQueryIdOk returns a tuple with the AssignmentQueryId field value
// and a boolean to check if the value has been set.
func (o *PostExperimentRequest) GetAssignmentQueryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignmentQueryId, true
}

// SetAssignmentQueryId sets field value
func (o *PostExperimentRequest) SetAssignmentQueryId(v string) {
	o.AssignmentQueryId = v
}

// GetTrackingKey returns the TrackingKey field value
func (o *PostExperimentRequest) GetTrackingKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrackingKey
}

// GetTrackingKeyOk returns a tuple with the TrackingKey field value
// and a boolean to check if the value has been set.
func (o *PostExperimentRequest) GetTrackingKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrackingKey, true
}

// SetTrackingKey sets field value
func (o *PostExperimentRequest) SetTrackingKey(v string) {
	o.TrackingKey = v
}

// GetName returns the Name field value
func (o *PostExperimentRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PostExperimentRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PostExperimentRequest) SetName(v string) {
	o.Name = v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *PostExperimentRequest) GetProject() string {
	if o == nil || common.IsNil(o.Project) {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostExperimentRequest) GetProjectOk() (*string, bool) {
	if o == nil || common.IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *PostExperimentRequest) HasProject() bool {
	if o != nil && !common.IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *PostExperimentRequest) SetProject(v string) {
	o.Project = &v
}

// GetHypothesis returns the Hypothesis field value if set, zero value otherwise.
func (o *PostExperimentRequest) GetHypothesis() string {
	if o == nil || common.IsNil(o.Hypothesis) {
		var ret string
		return ret
	}
	return *o.Hypothesis
}

// GetHypothesisOk returns a tuple with the Hypothesis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostExperimentRequest) GetHypothesisOk() (*string, bool) {
	if o == nil || common.IsNil(o.Hypothesis) {
		return nil, false
	}
	return o.Hypothesis, true
}

// HasHypothesis returns a boolean if a field has been set.
func (o *PostExperimentRequest) HasHypothesis() bool {
	if o != nil && !common.IsNil(o.Hypothesis) {
		return true
	}

	return false
}

// SetHypothesis gets a reference to the given string and assigns it to the Hypothesis field.
func (o *PostExperimentRequest) SetHypothesis(v string) {
	o.Hypothesis = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PostExperimentRequest) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostExperimentRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PostExperimentRequest) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PostExperimentRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PostExperimentRequest) GetTags() []string {
	if o == nil || common.IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostExperimentRequest) GetTagsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PostExperimentRequest) HasTags() bool {
	if o != nil && !common.IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *PostExperimentRequest) SetTags(v []string) {
	o.Tags = v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *PostExperimentRequest) GetMetrics() []string {
	if o == nil || common.IsNil(o.Metrics) {
		var ret []string
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostExperimentRequest) GetMetricsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Metrics) {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *PostExperimentRequest) HasMetrics() bool {
	if o != nil && !common.IsNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given []string and assigns it to the Metrics field.
func (o *PostExperimentRequest) SetMetrics(v []string) {
	o.Metrics = v
}

// GetGuardrailMetrics returns the GuardrailMetrics field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostExperimentRequest) GetGuardrailMetrics() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.GuardrailMetrics
}

// GetGuardrailMetricsOk returns a tuple with the GuardrailMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostExperimentRequest) GetGuardrailMetricsOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.GuardrailMetrics) {
		return nil, false
	}
	return &o.GuardrailMetrics, true
}

// HasGuardrailMetrics returns a boolean if a field has been set.
func (o *PostExperimentRequest) HasGuardrailMetrics() bool {
	if o != nil && common.IsNil(o.GuardrailMetrics) {
		return true
	}

	return false
}

// SetGuardrailMetrics gets a reference to the given interface{} and assigns it to the GuardrailMetrics field.
func (o *PostExperimentRequest) SetGuardrailMetrics(v interface{}) {
	o.GuardrailMetrics = v
}

// GetOwner returns the Owner field value
func (o *PostExperimentRequest) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *PostExperimentRequest) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *PostExperimentRequest) SetOwner(v string) {
	o.Owner = v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *PostExperimentRequest) GetArchived() bool {
	if o == nil || common.IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostExperimentRequest) GetArchivedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *PostExperimentRequest) HasArchived() bool {
	if o != nil && !common.IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *PostExperimentRequest) SetArchived(v bool) {
	o.Archived = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PostExperimentRequest) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostExperimentRequest) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PostExperimentRequest) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PostExperimentRequest) SetStatus(v string) {
	o.Status = &v
}

// GetAutoRefresh returns the AutoRefresh field value if set, zero value otherwise.
func (o *PostExperimentRequest) GetAutoRefresh() bool {
	if o == nil || common.IsNil(o.AutoRefresh) {
		var ret bool
		return ret
	}
	return *o.AutoRefresh
}

// GetAutoRefreshOk returns a tuple with the AutoRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostExperimentRequest) GetAutoRefreshOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AutoRefresh) {
		return nil, false
	}
	return o.AutoRefresh, true
}

// HasAutoRefresh returns a boolean if a field has been set.
func (o *PostExperimentRequest) HasAutoRefresh() bool {
	if o != nil && !common.IsNil(o.AutoRefresh) {
		return true
	}

	return false
}

// SetAutoRefresh gets a reference to the given bool and assigns it to the AutoRefresh field.
func (o *PostExperimentRequest) SetAutoRefresh(v bool) {
	o.AutoRefresh = &v
}

// GetHashAttribute returns the HashAttribute field value if set, zero value otherwise.
func (o *PostExperimentRequest) GetHashAttribute() string {
	if o == nil || common.IsNil(o.HashAttribute) {
		var ret string
		return ret
	}
	return *o.HashAttribute
}

// GetHashAttributeOk returns a tuple with the HashAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostExperimentRequest) GetHashAttributeOk() (*string, bool) {
	if o == nil || common.IsNil(o.HashAttribute) {
		return nil, false
	}
	return o.HashAttribute, true
}

// HasHashAttribute returns a boolean if a field has been set.
func (o *PostExperimentRequest) HasHashAttribute() bool {
	if o != nil && !common.IsNil(o.HashAttribute) {
		return true
	}

	return false
}

// SetHashAttribute gets a reference to the given string and assigns it to the HashAttribute field.
func (o *PostExperimentRequest) SetHashAttribute(v string) {
	o.HashAttribute = &v
}

// GetHashVersion returns the HashVersion field value if set, zero value otherwise.
func (o *PostExperimentRequest) GetHashVersion() float32 {
	if o == nil || common.IsNil(o.HashVersion) {
		var ret float32
		return ret
	}
	return *o.HashVersion
}

// GetHashVersionOk returns a tuple with the HashVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostExperimentRequest) GetHashVersionOk() (*float32, bool) {
	if o == nil || common.IsNil(o.HashVersion) {
		return nil, false
	}
	return o.HashVersion, true
}

// HasHashVersion returns a boolean if a field has been set.
func (o *PostExperimentRequest) HasHashVersion() bool {
	if o != nil && !common.IsNil(o.HashVersion) {
		return true
	}

	return false
}

// SetHashVersion gets a reference to the given float32 and assigns it to the HashVersion field.
func (o *PostExperimentRequest) SetHashVersion(v float32) {
	o.HashVersion = &v
}

// GetReleasedVariationId returns the ReleasedVariationId field value if set, zero value otherwise.
func (o *PostExperimentRequest) GetReleasedVariationId() string {
	if o == nil || common.IsNil(o.ReleasedVariationId) {
		var ret string
		return ret
	}
	return *o.ReleasedVariationId
}

// GetReleasedVariationIdOk returns a tuple with the ReleasedVariationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostExperimentRequest) GetReleasedVariationIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ReleasedVariationId) {
		return nil, false
	}
	return o.ReleasedVariationId, true
}

// HasReleasedVariationId returns a boolean if a field has been set.
func (o *PostExperimentRequest) HasReleasedVariationId() bool {
	if o != nil && !common.IsNil(o.ReleasedVariationId) {
		return true
	}

	return false
}

// SetReleasedVariationId gets a reference to the given string and assigns it to the ReleasedVariationId field.
func (o *PostExperimentRequest) SetReleasedVariationId(v string) {
	o.ReleasedVariationId = &v
}

// GetExcludeFromPayload returns the ExcludeFromPayload field value if set, zero value otherwise.
func (o *PostExperimentRequest) GetExcludeFromPayload() bool {
	if o == nil || common.IsNil(o.ExcludeFromPayload) {
		var ret bool
		return ret
	}
	return *o.ExcludeFromPayload
}

// GetExcludeFromPayloadOk returns a tuple with the ExcludeFromPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostExperimentRequest) GetExcludeFromPayloadOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ExcludeFromPayload) {
		return nil, false
	}
	return o.ExcludeFromPayload, true
}

// HasExcludeFromPayload returns a boolean if a field has been set.
func (o *PostExperimentRequest) HasExcludeFromPayload() bool {
	if o != nil && !common.IsNil(o.ExcludeFromPayload) {
		return true
	}

	return false
}

// SetExcludeFromPayload gets a reference to the given bool and assigns it to the ExcludeFromPayload field.
func (o *PostExperimentRequest) SetExcludeFromPayload(v bool) {
	o.ExcludeFromPayload = &v
}

// GetVariations returns the Variations field value
func (o *PostExperimentRequest) GetVariations() []PostExperimentRequestVariationsInner {
	if o == nil {
		var ret []PostExperimentRequestVariationsInner
		return ret
	}

	return o.Variations
}

// GetVariationsOk returns a tuple with the Variations field value
// and a boolean to check if the value has been set.
func (o *PostExperimentRequest) GetVariationsOk() ([]PostExperimentRequestVariationsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Variations, true
}

// SetVariations sets field value
func (o *PostExperimentRequest) SetVariations(v []PostExperimentRequestVariationsInner) {
	o.Variations = v
}

// GetPhases returns the Phases field value if set, zero value otherwise.
func (o *PostExperimentRequest) GetPhases() []PostExperimentRequestPhasesInner {
	if o == nil || common.IsNil(o.Phases) {
		var ret []PostExperimentRequestPhasesInner
		return ret
	}
	return o.Phases
}

// GetPhasesOk returns a tuple with the Phases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostExperimentRequest) GetPhasesOk() ([]PostExperimentRequestPhasesInner, bool) {
	if o == nil || common.IsNil(o.Phases) {
		return nil, false
	}
	return o.Phases, true
}

// HasPhases returns a boolean if a field has been set.
func (o *PostExperimentRequest) HasPhases() bool {
	if o != nil && !common.IsNil(o.Phases) {
		return true
	}

	return false
}

// SetPhases gets a reference to the given []PostExperimentRequestPhasesInner and assigns it to the Phases field.
func (o *PostExperimentRequest) SetPhases(v []PostExperimentRequestPhasesInner) {
	o.Phases = v
}

func (o PostExperimentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostExperimentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datasourceId"] = o.DatasourceId
	toSerialize["assignmentQueryId"] = o.AssignmentQueryId
	toSerialize["trackingKey"] = o.TrackingKey
	toSerialize["name"] = o.Name
	if !common.IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !common.IsNil(o.Hypothesis) {
		toSerialize["hypothesis"] = o.Hypothesis
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !common.IsNil(o.Metrics) {
		toSerialize["metrics"] = o.Metrics
	}
	if o.GuardrailMetrics != nil {
		toSerialize["guardrailMetrics"] = o.GuardrailMetrics
	}
	toSerialize["owner"] = o.Owner
	if !common.IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.AutoRefresh) {
		toSerialize["autoRefresh"] = o.AutoRefresh
	}
	if !common.IsNil(o.HashAttribute) {
		toSerialize["hashAttribute"] = o.HashAttribute
	}
	if !common.IsNil(o.HashVersion) {
		toSerialize["hashVersion"] = o.HashVersion
	}
	if !common.IsNil(o.ReleasedVariationId) {
		toSerialize["releasedVariationId"] = o.ReleasedVariationId
	}
	if !common.IsNil(o.ExcludeFromPayload) {
		toSerialize["excludeFromPayload"] = o.ExcludeFromPayload
	}
	toSerialize["variations"] = o.Variations
	if !common.IsNil(o.Phases) {
		toSerialize["phases"] = o.Phases
	}
	return toSerialize, nil
}

func (o *PostExperimentRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datasourceId",
		"assignmentQueryId",
		"trackingKey",
		"name",
		"owner",
		"variations",
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

	varPostExperimentRequest := _PostExperimentRequest{}

	err = json.Unmarshal(bytes, &varPostExperimentRequest)

	if err != nil {
		return err
	}

	*o = PostExperimentRequest(varPostExperimentRequest)

	return err
}

type NullablePostExperimentRequest struct {
	value *PostExperimentRequest
	isSet bool
}

func (v NullablePostExperimentRequest) Get() *PostExperimentRequest {
	return v.value
}

func (v *NullablePostExperimentRequest) Set(val *PostExperimentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostExperimentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostExperimentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostExperimentRequest(val *PostExperimentRequest) *NullablePostExperimentRequest {
	return &NullablePostExperimentRequest{value: val, isSet: true}
}

func (v NullablePostExperimentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostExperimentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

