# FileCoverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | This is the path of the file relative to the git repository | 
**Coverage** | **string** | Data consists of 3 semi-column separated parts: * &#x60;C&#x60; - list of fully covered lines * &#x60;P&#x60; - list of partly covered lines (line considered partly covered when it has multiple branches,   some of them covered and some not) * &#x60;U&#x60; - list of uncovered lines  | 

## Methods

### NewFileCoverage

`func NewFileCoverage(path string, coverage string, ) *FileCoverage`

NewFileCoverage instantiates a new FileCoverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileCoverageWithDefaults

`func NewFileCoverageWithDefaults() *FileCoverage`

NewFileCoverageWithDefaults instantiates a new FileCoverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *FileCoverage) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FileCoverage) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FileCoverage) SetPath(v string)`

SetPath sets Path field to given value.


### GetCoverage

`func (o *FileCoverage) GetCoverage() string`

GetCoverage returns the Coverage field if non-nil, zero value otherwise.

### GetCoverageOk

`func (o *FileCoverage) GetCoverageOk() (*string, bool)`

GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverage

`func (o *FileCoverage) SetCoverage(v string)`

SetCoverage sets Coverage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


