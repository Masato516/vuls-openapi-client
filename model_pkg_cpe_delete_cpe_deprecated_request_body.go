/*
Future Vuls Public API

Future Vuls Public API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"encoding/json"
)

// PkgCpeDeleteCpeDeprecatedRequestBody struct for PkgCpeDeleteCpeDeprecatedRequestBody
type PkgCpeDeleteCpeDeprecatedRequestBody struct {
	// cpe ID
	CpeID int64 `json:"cpeID"`
}

// NewPkgCpeDeleteCpeDeprecatedRequestBody instantiates a new PkgCpeDeleteCpeDeprecatedRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkgCpeDeleteCpeDeprecatedRequestBody(cpeID int64) *PkgCpeDeleteCpeDeprecatedRequestBody {
	this := PkgCpeDeleteCpeDeprecatedRequestBody{}
	this.CpeID = cpeID
	return &this
}

// NewPkgCpeDeleteCpeDeprecatedRequestBodyWithDefaults instantiates a new PkgCpeDeleteCpeDeprecatedRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkgCpeDeleteCpeDeprecatedRequestBodyWithDefaults() *PkgCpeDeleteCpeDeprecatedRequestBody {
	this := PkgCpeDeleteCpeDeprecatedRequestBody{}
	return &this
}

// GetCpeID returns the CpeID field value
func (o *PkgCpeDeleteCpeDeprecatedRequestBody) GetCpeID() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CpeID
}

// GetCpeIDOk returns a tuple with the CpeID field value
// and a boolean to check if the value has been set.
func (o *PkgCpeDeleteCpeDeprecatedRequestBody) GetCpeIDOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpeID, true
}

// SetCpeID sets field value
func (o *PkgCpeDeleteCpeDeprecatedRequestBody) SetCpeID(v int64) {
	o.CpeID = v
}

func (o PkgCpeDeleteCpeDeprecatedRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cpeID"] = o.CpeID
	}
	return json.Marshal(toSerialize)
}

type NullablePkgCpeDeleteCpeDeprecatedRequestBody struct {
	value *PkgCpeDeleteCpeDeprecatedRequestBody
	isSet bool
}

func (v NullablePkgCpeDeleteCpeDeprecatedRequestBody) Get() *PkgCpeDeleteCpeDeprecatedRequestBody {
	return v.value
}

func (v *NullablePkgCpeDeleteCpeDeprecatedRequestBody) Set(val *PkgCpeDeleteCpeDeprecatedRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullablePkgCpeDeleteCpeDeprecatedRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullablePkgCpeDeleteCpeDeprecatedRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkgCpeDeleteCpeDeprecatedRequestBody(val *PkgCpeDeleteCpeDeprecatedRequestBody) *NullablePkgCpeDeleteCpeDeprecatedRequestBody {
	return &NullablePkgCpeDeleteCpeDeprecatedRequestBody{value: val, isSet: true}
}

func (v NullablePkgCpeDeleteCpeDeprecatedRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkgCpeDeleteCpeDeprecatedRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


