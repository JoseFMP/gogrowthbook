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

// checks if the ExperimentAnalysisSettings type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &ExperimentAnalysisSettings{}

// ExperimentAnalysisSettings struct for ExperimentAnalysisSettings
type ExperimentAnalysisSettings struct {
	DatasourceId          interface{}       `json:"datasourceId"`
	AssignmentQueryId     interface{}       `json:"assignmentQueryId"`
	ExperimentId          interface{}       `json:"experimentId"`
	SegmentId             interface{}       `json:"segmentId"`
	QueryFilter           interface{}       `json:"queryFilter"`
	InProgressConversions interface{}       `json:"inProgressConversions"`
	AttributionModel      interface{}       `json:"attributionModel"`
	StatsEngine           interface{}       `json:"statsEngine"`
	Goals                 interface{}       `json:"goals"`
	Guardrails            interface{}       `json:"guardrails"`
	ActivationMetric      *ExperimentMetric `json:"activationMetric,omitempty"`
}

type _ExperimentAnalysisSettings ExperimentAnalysisSettings

// NewExperimentAnalysisSettings instantiates a new ExperimentAnalysisSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExperimentAnalysisSettings(datasourceId interface{}, assignmentQueryId interface{}, experimentId interface{}, segmentId interface{}, queryFilter interface{}, inProgressConversions interface{}, attributionModel interface{}, statsEngine interface{}, goals interface{}, guardrails interface{}) *ExperimentAnalysisSettings {
	this := ExperimentAnalysisSettings{}
	this.DatasourceId = datasourceId
	this.AssignmentQueryId = assignmentQueryId
	this.ExperimentId = experimentId
	this.SegmentId = segmentId
	this.QueryFilter = queryFilter
	this.InProgressConversions = inProgressConversions
	this.AttributionModel = attributionModel
	this.StatsEngine = statsEngine
	this.Goals = goals
	this.Guardrails = guardrails
	return &this
}

// NewExperimentAnalysisSettingsWithDefaults instantiates a new ExperimentAnalysisSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExperimentAnalysisSettingsWithDefaults() *ExperimentAnalysisSettings {
	this := ExperimentAnalysisSettings{}
	return &this
}

// GetDatasourceId returns the DatasourceId field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExperimentAnalysisSettings) GetDatasourceId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.DatasourceId
}

// GetDatasourceIdOk returns a tuple with the DatasourceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExperimentAnalysisSettings) GetDatasourceIdOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.DatasourceId) {
		return nil, false
	}
	return &o.DatasourceId, true
}

// SetDatasourceId sets field value
func (o *ExperimentAnalysisSettings) SetDatasourceId(v interface{}) {
	o.DatasourceId = v
}

// GetAssignmentQueryId returns the AssignmentQueryId field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExperimentAnalysisSettings) GetAssignmentQueryId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.AssignmentQueryId
}

// GetAssignmentQueryIdOk returns a tuple with the AssignmentQueryId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExperimentAnalysisSettings) GetAssignmentQueryIdOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.AssignmentQueryId) {
		return nil, false
	}
	return &o.AssignmentQueryId, true
}

// SetAssignmentQueryId sets field value
func (o *ExperimentAnalysisSettings) SetAssignmentQueryId(v interface{}) {
	o.AssignmentQueryId = v
}

// GetExperimentId returns the ExperimentId field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExperimentAnalysisSettings) GetExperimentId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.ExperimentId
}

// GetExperimentIdOk returns a tuple with the ExperimentId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExperimentAnalysisSettings) GetExperimentIdOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.ExperimentId) {
		return nil, false
	}
	return &o.ExperimentId, true
}

// SetExperimentId sets field value
func (o *ExperimentAnalysisSettings) SetExperimentId(v interface{}) {
	o.ExperimentId = v
}

// GetSegmentId returns the SegmentId field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExperimentAnalysisSettings) GetSegmentId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.SegmentId
}

// GetSegmentIdOk returns a tuple with the SegmentId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExperimentAnalysisSettings) GetSegmentIdOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.SegmentId) {
		return nil, false
	}
	return &o.SegmentId, true
}

// SetSegmentId sets field value
func (o *ExperimentAnalysisSettings) SetSegmentId(v interface{}) {
	o.SegmentId = v
}

// GetQueryFilter returns the QueryFilter field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExperimentAnalysisSettings) GetQueryFilter() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.QueryFilter
}

// GetQueryFilterOk returns a tuple with the QueryFilter field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExperimentAnalysisSettings) GetQueryFilterOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.QueryFilter) {
		return nil, false
	}
	return &o.QueryFilter, true
}

// SetQueryFilter sets field value
func (o *ExperimentAnalysisSettings) SetQueryFilter(v interface{}) {
	o.QueryFilter = v
}

// GetInProgressConversions returns the InProgressConversions field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExperimentAnalysisSettings) GetInProgressConversions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.InProgressConversions
}

// GetInProgressConversionsOk returns a tuple with the InProgressConversions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExperimentAnalysisSettings) GetInProgressConversionsOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.InProgressConversions) {
		return nil, false
	}
	return &o.InProgressConversions, true
}

// SetInProgressConversions sets field value
func (o *ExperimentAnalysisSettings) SetInProgressConversions(v interface{}) {
	o.InProgressConversions = v
}

// GetAttributionModel returns the AttributionModel field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExperimentAnalysisSettings) GetAttributionModel() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.AttributionModel
}

// GetAttributionModelOk returns a tuple with the AttributionModel field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExperimentAnalysisSettings) GetAttributionModelOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.AttributionModel) {
		return nil, false
	}
	return &o.AttributionModel, true
}

// SetAttributionModel sets field value
func (o *ExperimentAnalysisSettings) SetAttributionModel(v interface{}) {
	o.AttributionModel = v
}

// GetStatsEngine returns the StatsEngine field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExperimentAnalysisSettings) GetStatsEngine() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.StatsEngine
}

// GetStatsEngineOk returns a tuple with the StatsEngine field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExperimentAnalysisSettings) GetStatsEngineOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.StatsEngine) {
		return nil, false
	}
	return &o.StatsEngine, true
}

// SetStatsEngine sets field value
func (o *ExperimentAnalysisSettings) SetStatsEngine(v interface{}) {
	o.StatsEngine = v
}

// GetGoals returns the Goals field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExperimentAnalysisSettings) GetGoals() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Goals
}

// GetGoalsOk returns a tuple with the Goals field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExperimentAnalysisSettings) GetGoalsOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Goals) {
		return nil, false
	}
	return &o.Goals, true
}

// SetGoals sets field value
func (o *ExperimentAnalysisSettings) SetGoals(v interface{}) {
	o.Goals = v
}

// GetGuardrails returns the Guardrails field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExperimentAnalysisSettings) GetGuardrails() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Guardrails
}

// GetGuardrailsOk returns a tuple with the Guardrails field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExperimentAnalysisSettings) GetGuardrailsOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.Guardrails) {
		return nil, false
	}
	return &o.Guardrails, true
}

// SetGuardrails sets field value
func (o *ExperimentAnalysisSettings) SetGuardrails(v interface{}) {
	o.Guardrails = v
}

// GetActivationMetric returns the ActivationMetric field value if set, zero value otherwise.
func (o *ExperimentAnalysisSettings) GetActivationMetric() ExperimentMetric {
	if o == nil || common.IsNil(o.ActivationMetric) {
		var ret ExperimentMetric
		return ret
	}
	return *o.ActivationMetric
}

// GetActivationMetricOk returns a tuple with the ActivationMetric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentAnalysisSettings) GetActivationMetricOk() (*ExperimentMetric, bool) {
	if o == nil || common.IsNil(o.ActivationMetric) {
		return nil, false
	}
	return o.ActivationMetric, true
}

// HasActivationMetric returns a boolean if a field has been set.
func (o *ExperimentAnalysisSettings) HasActivationMetric() bool {
	if o != nil && !common.IsNil(o.ActivationMetric) {
		return true
	}

	return false
}

// SetActivationMetric gets a reference to the given ExperimentMetric and assigns it to the ActivationMetric field.
func (o *ExperimentAnalysisSettings) SetActivationMetric(v ExperimentMetric) {
	o.ActivationMetric = &v
}

func (o ExperimentAnalysisSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExperimentAnalysisSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DatasourceId != nil {
		toSerialize["datasourceId"] = o.DatasourceId
	}
	if o.AssignmentQueryId != nil {
		toSerialize["assignmentQueryId"] = o.AssignmentQueryId
	}
	if o.ExperimentId != nil {
		toSerialize["experimentId"] = o.ExperimentId
	}
	if o.SegmentId != nil {
		toSerialize["segmentId"] = o.SegmentId
	}
	if o.QueryFilter != nil {
		toSerialize["queryFilter"] = o.QueryFilter
	}
	if o.InProgressConversions != nil {
		toSerialize["inProgressConversions"] = o.InProgressConversions
	}
	if o.AttributionModel != nil {
		toSerialize["attributionModel"] = o.AttributionModel
	}
	if o.StatsEngine != nil {
		toSerialize["statsEngine"] = o.StatsEngine
	}
	if o.Goals != nil {
		toSerialize["goals"] = o.Goals
	}
	if o.Guardrails != nil {
		toSerialize["guardrails"] = o.Guardrails
	}
	if !common.IsNil(o.ActivationMetric) {
		toSerialize["activationMetric"] = o.ActivationMetric
	}
	return toSerialize, nil
}

func (o *ExperimentAnalysisSettings) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datasourceId",
		"assignmentQueryId",
		"experimentId",
		"segmentId",
		"queryFilter",
		"inProgressConversions",
		"attributionModel",
		"statsEngine",
		"goals",
		"guardrails",
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

	varExperimentAnalysisSettings := _ExperimentAnalysisSettings{}

	err = json.Unmarshal(bytes, &varExperimentAnalysisSettings)

	if err != nil {
		return err
	}

	*o = ExperimentAnalysisSettings(varExperimentAnalysisSettings)

	return err
}

type NullableExperimentAnalysisSettings struct {
	value *ExperimentAnalysisSettings
	isSet bool
}

func (v NullableExperimentAnalysisSettings) Get() *ExperimentAnalysisSettings {
	return v.value
}

func (v *NullableExperimentAnalysisSettings) Set(val *ExperimentAnalysisSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableExperimentAnalysisSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableExperimentAnalysisSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExperimentAnalysisSettings(val *ExperimentAnalysisSettings) *NullableExperimentAnalysisSettings {
	return &NullableExperimentAnalysisSettings{value: val, isSet: true}
}

func (v NullableExperimentAnalysisSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExperimentAnalysisSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
