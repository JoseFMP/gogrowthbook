/*
GrowthBook REST API

GrowthBook offers a full REST API for interacting with the GrowthBook application. This is currently in **beta** as we add more authenticated API routes and features.  Request data can use either JSON or Form data encoding (with proper `Content-Type` headers). All response bodies are JSON-encoded.  The API base URL for GrowthBook Cloud is `https://api.growthbook.io`. For self-hosted deployments, it is the same as your API_HOST environment variable (defaults to `http://localhost:3100`). The rest of these docs will assume you are using GrowthBook Cloud.  ## Authentication  We support both the HTTP Basic and Bearer authentication schemes for convenience.  You first need to generate a new API Key in GrowthBook. Different keys have different permissions:  - **Personal Access Tokens**: These are sensitive and provide the same level of access as the user has to an organization. These can be created by going to `Personal Access Tokens` under the your user menu. - **Secret Keys**: These are sensitive and provide the level of access for the role, which currently is either `admin` or `readonly`. Only Admins with the `manageApiKeys` permission can manage Secret Keys on behalf of an organization. These can be created by going to `Settings -> API Keys`  If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:  ```bash curl https://api.growthbook.io/api/v1 \\   -u secret_abc123DEF456: # The \":\" at the end stops curl from asking for a password ```  If using Bearer auth, pass the Secret Key as the token:  ```bash curl https://api.growthbook.io/api/v1 \\ -H \"Authorization: Bearer secret_abc123DEF456\" ```  ## Errors  The API may return the following error status codes:  - **400** - Bad Request - Often due to a missing required parameter - **401** - Unauthorized - No valid API key provided - **402** - Request Failed - The parameters are valid, but the request failed - **403** - Forbidden - Provided API key does not have the required access - **404** - Not Found - Unknown API route or requested resource - **429** - Too Many Requests - You exceeded the rate limit of 60 requests per minute. Try again later. - **5XX** - Server Error - Something went wrong on GrowthBook's end (these are rare)  The response body will be a JSON object with the following properties:  - **message** - Information about the error

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package models

import (
	"encoding/json"

	"github.com/JoseFMP/gogrowthbook/common"
)

// checks if the UpdateFeatureRequest type satisfies the common.MappedNullable interface at compile time
var _ common.MappedNullable = &UpdateFeatureRequest{}

// UpdateFeatureRequest struct for UpdateFeatureRequest
type UpdateFeatureRequest struct {
	// Description of the feature
	Description *string `json:"description,omitempty"`
	Archived    *bool   `json:"archived,omitempty"`
	// An associated project ID
	Project      *string                                        `json:"project,omitempty"`
	Owner        *string                                        `json:"owner,omitempty"`
	DefaultValue *string                                        `json:"defaultValue,omitempty"`
	Tags         []string                                       `json:"tags,omitempty"`
	Environments map[string]PostFeatureRequestEnvironmentsValue `json:"environments,omitempty"`
	// Use JSON schema to validate the payload of a JSON-type feature value (enterprise only).
	JsonSchema *string `json:"jsonSchema,omitempty"`
}

// NewUpdateFeatureRequest instantiates a new UpdateFeatureRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFeatureRequest() *UpdateFeatureRequest {
	this := UpdateFeatureRequest{}
	return &this
}

// NewUpdateFeatureRequestWithDefaults instantiates a new UpdateFeatureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFeatureRequestWithDefaults() *UpdateFeatureRequest {
	this := UpdateFeatureRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateFeatureRequest) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFeatureRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateFeatureRequest) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateFeatureRequest) SetDescription(v string) {
	o.Description = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *UpdateFeatureRequest) GetArchived() bool {
	if o == nil || common.IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFeatureRequest) GetArchivedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *UpdateFeatureRequest) HasArchived() bool {
	if o != nil && !common.IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *UpdateFeatureRequest) SetArchived(v bool) {
	o.Archived = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *UpdateFeatureRequest) GetProject() string {
	if o == nil || common.IsNil(o.Project) {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFeatureRequest) GetProjectOk() (*string, bool) {
	if o == nil || common.IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *UpdateFeatureRequest) HasProject() bool {
	if o != nil && !common.IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *UpdateFeatureRequest) SetProject(v string) {
	o.Project = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *UpdateFeatureRequest) GetOwner() string {
	if o == nil || common.IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFeatureRequest) GetOwnerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *UpdateFeatureRequest) HasOwner() bool {
	if o != nil && !common.IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *UpdateFeatureRequest) SetOwner(v string) {
	o.Owner = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *UpdateFeatureRequest) GetDefaultValue() string {
	if o == nil || common.IsNil(o.DefaultValue) {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFeatureRequest) GetDefaultValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *UpdateFeatureRequest) HasDefaultValue() bool {
	if o != nil && !common.IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *UpdateFeatureRequest) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateFeatureRequest) GetTags() []string {
	if o == nil || common.IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFeatureRequest) GetTagsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateFeatureRequest) HasTags() bool {
	if o != nil && !common.IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateFeatureRequest) SetTags(v []string) {
	o.Tags = v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateFeatureRequest) GetEnvironments() map[string]PostFeatureRequestEnvironmentsValue {
	if o == nil {
		var ret map[string]PostFeatureRequestEnvironmentsValue
		return ret
	}
	return o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFeatureRequest) GetEnvironmentsOk() (*map[string]PostFeatureRequestEnvironmentsValue, bool) {
	if o == nil || common.IsNil(o.Environments) {
		return nil, false
	}
	return &o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *UpdateFeatureRequest) HasEnvironments() bool {
	if o != nil && common.IsNil(o.Environments) {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given map[string]PostFeatureRequestEnvironmentsValue and assigns it to the Environments field.
func (o *UpdateFeatureRequest) SetEnvironments(v map[string]PostFeatureRequestEnvironmentsValue) {
	o.Environments = v
}

// GetJsonSchema returns the JsonSchema field value if set, zero value otherwise.
func (o *UpdateFeatureRequest) GetJsonSchema() string {
	if o == nil || common.IsNil(o.JsonSchema) {
		var ret string
		return ret
	}
	return *o.JsonSchema
}

// GetJsonSchemaOk returns a tuple with the JsonSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFeatureRequest) GetJsonSchemaOk() (*string, bool) {
	if o == nil || common.IsNil(o.JsonSchema) {
		return nil, false
	}
	return o.JsonSchema, true
}

// HasJsonSchema returns a boolean if a field has been set.
func (o *UpdateFeatureRequest) HasJsonSchema() bool {
	if o != nil && !common.IsNil(o.JsonSchema) {
		return true
	}

	return false
}

// SetJsonSchema gets a reference to the given string and assigns it to the JsonSchema field.
func (o *UpdateFeatureRequest) SetJsonSchema(v string) {
	o.JsonSchema = &v
}

func (o UpdateFeatureRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFeatureRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	if !common.IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !common.IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !common.IsNil(o.DefaultValue) {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if !common.IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if o.Environments != nil {
		toSerialize["environments"] = o.Environments
	}
	if !common.IsNil(o.JsonSchema) {
		toSerialize["jsonSchema"] = o.JsonSchema
	}
	return toSerialize, nil
}

type NullableUpdateFeatureRequest struct {
	value *UpdateFeatureRequest
	isSet bool
}

func (v NullableUpdateFeatureRequest) Get() *UpdateFeatureRequest {
	return v.value
}

func (v *NullableUpdateFeatureRequest) Set(val *UpdateFeatureRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFeatureRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFeatureRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFeatureRequest(val *UpdateFeatureRequest) *NullableUpdateFeatureRequest {
	return &NullableUpdateFeatureRequest{value: val, isSet: true}
}

func (v NullableUpdateFeatureRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFeatureRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
