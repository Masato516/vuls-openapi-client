/*
Future Vuls Public API

Future Vuls Public API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"encoding/json"
	"time"
)

// EnvMetricV2ResponseBody EnvMetricV2 describes a envMetricV2
type EnvMetricV2ResponseBody struct {
	// CDP of envMetricV2
	Cdp string `json:"cdp"`
	// created time
	CreatedAt time.Time `json:"createdAt"`
	// CveID of envMetricV2
	CveID string `json:"cveID"`
	// ServerRoleID of envMetricV2
	RoleID int64 `json:"roleID"`
	// ServerRoleName of envMetricV2
	RoleName string `json:"roleName"`
	// TD of envMetricV2
	Td string `json:"td"`
	// updated time
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewEnvMetricV2ResponseBody instantiates a new EnvMetricV2ResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvMetricV2ResponseBody(cdp string, createdAt time.Time, cveID string, roleID int64, roleName string, td string, updatedAt time.Time) *EnvMetricV2ResponseBody {
	this := EnvMetricV2ResponseBody{}
	this.Cdp = cdp
	this.CreatedAt = createdAt
	this.CveID = cveID
	this.RoleID = roleID
	this.RoleName = roleName
	this.Td = td
	this.UpdatedAt = updatedAt
	return &this
}

// NewEnvMetricV2ResponseBodyWithDefaults instantiates a new EnvMetricV2ResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvMetricV2ResponseBodyWithDefaults() *EnvMetricV2ResponseBody {
	this := EnvMetricV2ResponseBody{}
	return &this
}

// GetCdp returns the Cdp field value
func (o *EnvMetricV2ResponseBody) GetCdp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cdp
}

// GetCdpOk returns a tuple with the Cdp field value
// and a boolean to check if the value has been set.
func (o *EnvMetricV2ResponseBody) GetCdpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cdp, true
}

// SetCdp sets field value
func (o *EnvMetricV2ResponseBody) SetCdp(v string) {
	o.Cdp = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *EnvMetricV2ResponseBody) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EnvMetricV2ResponseBody) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EnvMetricV2ResponseBody) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCveID returns the CveID field value
func (o *EnvMetricV2ResponseBody) GetCveID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CveID
}

// GetCveIDOk returns a tuple with the CveID field value
// and a boolean to check if the value has been set.
func (o *EnvMetricV2ResponseBody) GetCveIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CveID, true
}

// SetCveID sets field value
func (o *EnvMetricV2ResponseBody) SetCveID(v string) {
	o.CveID = v
}

// GetRoleID returns the RoleID field value
func (o *EnvMetricV2ResponseBody) GetRoleID() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RoleID
}

// GetRoleIDOk returns a tuple with the RoleID field value
// and a boolean to check if the value has been set.
func (o *EnvMetricV2ResponseBody) GetRoleIDOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleID, true
}

// SetRoleID sets field value
func (o *EnvMetricV2ResponseBody) SetRoleID(v int64) {
	o.RoleID = v
}

// GetRoleName returns the RoleName field value
func (o *EnvMetricV2ResponseBody) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *EnvMetricV2ResponseBody) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *EnvMetricV2ResponseBody) SetRoleName(v string) {
	o.RoleName = v
}

// GetTd returns the Td field value
func (o *EnvMetricV2ResponseBody) GetTd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Td
}

// GetTdOk returns a tuple with the Td field value
// and a boolean to check if the value has been set.
func (o *EnvMetricV2ResponseBody) GetTdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Td, true
}

// SetTd sets field value
func (o *EnvMetricV2ResponseBody) SetTd(v string) {
	o.Td = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *EnvMetricV2ResponseBody) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *EnvMetricV2ResponseBody) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *EnvMetricV2ResponseBody) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o EnvMetricV2ResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cdp"] = o.Cdp
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["cveID"] = o.CveID
	}
	if true {
		toSerialize["roleID"] = o.RoleID
	}
	if true {
		toSerialize["roleName"] = o.RoleName
	}
	if true {
		toSerialize["td"] = o.Td
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableEnvMetricV2ResponseBody struct {
	value *EnvMetricV2ResponseBody
	isSet bool
}

func (v NullableEnvMetricV2ResponseBody) Get() *EnvMetricV2ResponseBody {
	return v.value
}

func (v *NullableEnvMetricV2ResponseBody) Set(val *EnvMetricV2ResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvMetricV2ResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvMetricV2ResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvMetricV2ResponseBody(val *EnvMetricV2ResponseBody) *NullableEnvMetricV2ResponseBody {
	return &NullableEnvMetricV2ResponseBody{value: val, isSet: true}
}

func (v NullableEnvMetricV2ResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvMetricV2ResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

