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

// checks if the VisualChangeset type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &VisualChangeset{}

// VisualChangeset struct for VisualChangeset
type VisualChangeset struct {
	Id            *string     `json:"id,omitempty"`
	UrlPatterns   interface{} `json:"urlPatterns"`
	EditorUrl     string      `json:"editorUrl"`
	Experiment    string      `json:"experiment"`
	VisualChanges interface{} `json:"visualChanges"`
}

type _VisualChangeset VisualChangeset

// NewVisualChangeset instantiates a new VisualChangeset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVisualChangeset(urlPatterns interface{}, editorUrl string, experiment string, visualChanges interface{}) *VisualChangeset {
	this := VisualChangeset{}
	this.UrlPatterns = urlPatterns
	this.EditorUrl = editorUrl
	this.Experiment = experiment
	this.VisualChanges = visualChanges
	return &this
}

// NewVisualChangesetWithDefaults instantiates a new VisualChangeset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVisualChangesetWithDefaults() *VisualChangeset {
	this := VisualChangeset{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VisualChangeset) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisualChangeset) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VisualChangeset) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VisualChangeset) SetId(v string) {
	o.Id = &v
}

// GetUrlPatterns returns the UrlPatterns field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *VisualChangeset) GetUrlPatterns() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.UrlPatterns
}

// GetUrlPatternsOk returns a tuple with the UrlPatterns field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VisualChangeset) GetUrlPatternsOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.UrlPatterns) {
		return nil, false
	}
	return &o.UrlPatterns, true
}

// SetUrlPatterns sets field value
func (o *VisualChangeset) SetUrlPatterns(v interface{}) {
	o.UrlPatterns = v
}

// GetEditorUrl returns the EditorUrl field value
func (o *VisualChangeset) GetEditorUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EditorUrl
}

// GetEditorUrlOk returns a tuple with the EditorUrl field value
// and a boolean to check if the value has been set.
func (o *VisualChangeset) GetEditorUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EditorUrl, true
}

// SetEditorUrl sets field value
func (o *VisualChangeset) SetEditorUrl(v string) {
	o.EditorUrl = v
}

// GetExperiment returns the Experiment field value
func (o *VisualChangeset) GetExperiment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Experiment
}

// GetExperimentOk returns a tuple with the Experiment field value
// and a boolean to check if the value has been set.
func (o *VisualChangeset) GetExperimentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Experiment, true
}

// SetExperiment sets field value
func (o *VisualChangeset) SetExperiment(v string) {
	o.Experiment = v
}

// GetVisualChanges returns the VisualChanges field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *VisualChangeset) GetVisualChanges() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.VisualChanges
}

// GetVisualChangesOk returns a tuple with the VisualChanges field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VisualChangeset) GetVisualChangesOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.VisualChanges) {
		return nil, false
	}
	return &o.VisualChanges, true
}

// SetVisualChanges sets field value
func (o *VisualChangeset) SetVisualChanges(v interface{}) {
	o.VisualChanges = v
}

func (o VisualChangeset) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VisualChangeset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.UrlPatterns != nil {
		toSerialize["urlPatterns"] = o.UrlPatterns
	}
	toSerialize["editorUrl"] = o.EditorUrl
	toSerialize["experiment"] = o.Experiment
	if o.VisualChanges != nil {
		toSerialize["visualChanges"] = o.VisualChanges
	}
	return toSerialize, nil
}

func (o *VisualChangeset) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"urlPatterns",
		"editorUrl",
		"experiment",
		"visualChanges",
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

	varVisualChangeset := _VisualChangeset{}

	err = json.Unmarshal(bytes, &varVisualChangeset)

	if err != nil {
		return err
	}

	*o = VisualChangeset(varVisualChangeset)

	return err
}

type NullableVisualChangeset struct {
	value *VisualChangeset
	isSet bool
}

func (v NullableVisualChangeset) Get() *VisualChangeset {
	return v.value
}

func (v *NullableVisualChangeset) Set(val *VisualChangeset) {
	v.value = val
	v.isSet = true
}

func (v NullableVisualChangeset) IsSet() bool {
	return v.isSet
}

func (v *NullableVisualChangeset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisualChangeset(val *VisualChangeset) *NullableVisualChangeset {
	return &NullableVisualChangeset{value: val, isSet: true}
}

func (v NullableVisualChangeset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisualChangeset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
