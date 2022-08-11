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

// ServerTagResponseBody ServerTag describes a server tag
type ServerTagResponseBody struct {
	// ID of server tag
	Id int64 `json:"id"`
	// Name of server tag
	Name string `json:"name"`
}

// NewServerTagResponseBody instantiates a new ServerTagResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerTagResponseBody(id int64, name string) *ServerTagResponseBody {
	this := ServerTagResponseBody{}
	this.Id = id
	this.Name = name
	return &this
}

// NewServerTagResponseBodyWithDefaults instantiates a new ServerTagResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerTagResponseBodyWithDefaults() *ServerTagResponseBody {
	this := ServerTagResponseBody{}
	return &this
}

// GetId returns the Id field value
func (o *ServerTagResponseBody) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServerTagResponseBody) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServerTagResponseBody) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ServerTagResponseBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServerTagResponseBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServerTagResponseBody) SetName(v string) {
	o.Name = v
}

func (o ServerTagResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableServerTagResponseBody struct {
	value *ServerTagResponseBody
	isSet bool
}

func (v NullableServerTagResponseBody) Get() *ServerTagResponseBody {
	return v.value
}

func (v *NullableServerTagResponseBody) Set(val *ServerTagResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableServerTagResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableServerTagResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerTagResponseBody(val *ServerTagResponseBody) *NullableServerTagResponseBody {
	return &NullableServerTagResponseBody{value: val, isSet: true}
}

func (v NullableServerTagResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerTagResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

