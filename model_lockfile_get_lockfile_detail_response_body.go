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

// LockfileGetLockfileDetailResponseBody struct for LockfileGetLockfileDetailResponseBody
type LockfileGetLockfileDetailResponseBody struct {
	// created time of lockfile
	CreatedAt string `json:"createdAt"`
	// FileContent of lockfile
	FileContent string `json:"fileContent"`
	// ID of lockfile
	Id int64 `json:"id"`
	// LibraryPkgs of lockfile
	LibraryPkgs []LibraryPkgChildResponseBody `json:"libraryPkgs,omitempty"`
	// Path of lockfile
	Path string `json:"path"`
	Server *ServerChildResponseBody `json:"server,omitempty"`
	// updated time of lockfile
	UpdatedAt string `json:"updatedAt"`
}

// NewLockfileGetLockfileDetailResponseBody instantiates a new LockfileGetLockfileDetailResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLockfileGetLockfileDetailResponseBody(createdAt string, fileContent string, id int64, path string, updatedAt string) *LockfileGetLockfileDetailResponseBody {
	this := LockfileGetLockfileDetailResponseBody{}
	this.CreatedAt = createdAt
	this.FileContent = fileContent
	this.Id = id
	this.Path = path
	this.UpdatedAt = updatedAt
	return &this
}

// NewLockfileGetLockfileDetailResponseBodyWithDefaults instantiates a new LockfileGetLockfileDetailResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLockfileGetLockfileDetailResponseBodyWithDefaults() *LockfileGetLockfileDetailResponseBody {
	this := LockfileGetLockfileDetailResponseBody{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *LockfileGetLockfileDetailResponseBody) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LockfileGetLockfileDetailResponseBody) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *LockfileGetLockfileDetailResponseBody) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetFileContent returns the FileContent field value
func (o *LockfileGetLockfileDetailResponseBody) GetFileContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileContent
}

// GetFileContentOk returns a tuple with the FileContent field value
// and a boolean to check if the value has been set.
func (o *LockfileGetLockfileDetailResponseBody) GetFileContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileContent, true
}

// SetFileContent sets field value
func (o *LockfileGetLockfileDetailResponseBody) SetFileContent(v string) {
	o.FileContent = v
}

// GetId returns the Id field value
func (o *LockfileGetLockfileDetailResponseBody) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LockfileGetLockfileDetailResponseBody) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LockfileGetLockfileDetailResponseBody) SetId(v int64) {
	o.Id = v
}

// GetLibraryPkgs returns the LibraryPkgs field value if set, zero value otherwise.
func (o *LockfileGetLockfileDetailResponseBody) GetLibraryPkgs() []LibraryPkgChildResponseBody {
	if o == nil || o.LibraryPkgs == nil {
		var ret []LibraryPkgChildResponseBody
		return ret
	}
	return o.LibraryPkgs
}

// GetLibraryPkgsOk returns a tuple with the LibraryPkgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockfileGetLockfileDetailResponseBody) GetLibraryPkgsOk() ([]LibraryPkgChildResponseBody, bool) {
	if o == nil || o.LibraryPkgs == nil {
		return nil, false
	}
	return o.LibraryPkgs, true
}

// HasLibraryPkgs returns a boolean if a field has been set.
func (o *LockfileGetLockfileDetailResponseBody) HasLibraryPkgs() bool {
	if o != nil && o.LibraryPkgs != nil {
		return true
	}

	return false
}

// SetLibraryPkgs gets a reference to the given []LibraryPkgChildResponseBody and assigns it to the LibraryPkgs field.
func (o *LockfileGetLockfileDetailResponseBody) SetLibraryPkgs(v []LibraryPkgChildResponseBody) {
	o.LibraryPkgs = v
}

// GetPath returns the Path field value
func (o *LockfileGetLockfileDetailResponseBody) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *LockfileGetLockfileDetailResponseBody) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *LockfileGetLockfileDetailResponseBody) SetPath(v string) {
	o.Path = v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *LockfileGetLockfileDetailResponseBody) GetServer() ServerChildResponseBody {
	if o == nil || o.Server == nil {
		var ret ServerChildResponseBody
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockfileGetLockfileDetailResponseBody) GetServerOk() (*ServerChildResponseBody, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *LockfileGetLockfileDetailResponseBody) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given ServerChildResponseBody and assigns it to the Server field.
func (o *LockfileGetLockfileDetailResponseBody) SetServer(v ServerChildResponseBody) {
	o.Server = &v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *LockfileGetLockfileDetailResponseBody) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *LockfileGetLockfileDetailResponseBody) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *LockfileGetLockfileDetailResponseBody) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o LockfileGetLockfileDetailResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["fileContent"] = o.FileContent
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.LibraryPkgs != nil {
		toSerialize["libraryPkgs"] = o.LibraryPkgs
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if o.Server != nil {
		toSerialize["server"] = o.Server
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableLockfileGetLockfileDetailResponseBody struct {
	value *LockfileGetLockfileDetailResponseBody
	isSet bool
}

func (v NullableLockfileGetLockfileDetailResponseBody) Get() *LockfileGetLockfileDetailResponseBody {
	return v.value
}

func (v *NullableLockfileGetLockfileDetailResponseBody) Set(val *LockfileGetLockfileDetailResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableLockfileGetLockfileDetailResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableLockfileGetLockfileDetailResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLockfileGetLockfileDetailResponseBody(val *LockfileGetLockfileDetailResponseBody) *NullableLockfileGetLockfileDetailResponseBody {
	return &NullableLockfileGetLockfileDetailResponseBody{value: val, isSet: true}
}

func (v NullableLockfileGetLockfileDetailResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLockfileGetLockfileDetailResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


