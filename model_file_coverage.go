/*
 * Bitbucket Server Code Coverage plugin API
 *
 * Bitbucket Server Code Coverage plugin API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package coverage

import (
	"encoding/json"
)

// FileCoverage File code coverage data
type FileCoverage struct {
	// This is the path of the file relative to the git repository
	Path string `json:"path"`
	// Data consists of 3 semi-column separated parts: * `C` - list of fully covered lines * `P` - list of partly covered lines (line considered partly covered when it has multiple branches,   some of them covered and some not) * `U` - list of uncovered lines
	Coverage string `json:"coverage"`
}

// NewFileCoverage instantiates a new FileCoverage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileCoverage(path string, coverage string) *FileCoverage {
	this := FileCoverage{}
	this.Path = path
	this.Coverage = coverage
	return &this
}

// NewFileCoverageWithDefaults instantiates a new FileCoverage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileCoverageWithDefaults() *FileCoverage {
	this := FileCoverage{}
	return &this
}

// GetPath returns the Path field value
func (o *FileCoverage) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *FileCoverage) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *FileCoverage) SetPath(v string) {
	o.Path = v
}

// GetCoverage returns the Coverage field value
func (o *FileCoverage) GetCoverage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Coverage
}

// GetCoverageOk returns a tuple with the Coverage field value
// and a boolean to check if the value has been set.
func (o *FileCoverage) GetCoverageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Coverage, true
}

// SetCoverage sets field value
func (o *FileCoverage) SetCoverage(v string) {
	o.Coverage = v
}

func (o FileCoverage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["coverage"] = o.Coverage
	}
	return json.Marshal(toSerialize)
}

type NullableFileCoverage struct {
	value *FileCoverage
	isSet bool
}

func (v NullableFileCoverage) Get() *FileCoverage {
	return v.value
}

func (v *NullableFileCoverage) Set(val *FileCoverage) {
	v.value = val
	v.isSet = true
}

func (v NullableFileCoverage) IsSet() bool {
	return v.isSet
}

func (v *NullableFileCoverage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileCoverage(val *FileCoverage) *NullableFileCoverage {
	return &NullableFileCoverage{value: val, isSet: true}
}

func (v NullableFileCoverage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileCoverage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}