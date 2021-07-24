# CommitCoverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | [**[]FileCoverage**](FileCoverage.md) | Code coverage data for files included in commit | 

## Methods

### NewCommitCoverage

`func NewCommitCoverage(files []FileCoverage, ) *CommitCoverage`

NewCommitCoverage instantiates a new CommitCoverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitCoverageWithDefaults

`func NewCommitCoverageWithDefaults() *CommitCoverage`

NewCommitCoverageWithDefaults instantiates a new CommitCoverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *CommitCoverage) GetFiles() []FileCoverage`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *CommitCoverage) GetFilesOk() (*[]FileCoverage, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *CommitCoverage) SetFiles(v []FileCoverage)`

SetFiles sets Files field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


