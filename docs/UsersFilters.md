# UsersFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | Pointer to **string** | The name of the user. | [optional] 
**UserType** | Pointer to **string** | The type of the user. | [optional] 
**ComponentUser** | Pointer to **bool** | Whether the user is a known component or not. | [optional] 

## Methods

### NewUsersFilters

`func NewUsersFilters() *UsersFilters`

NewUsersFilters instantiates a new UsersFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersFiltersWithDefaults

`func NewUsersFiltersWithDefaults() *UsersFilters`

NewUsersFiltersWithDefaults instantiates a new UsersFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *UsersFilters) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UsersFilters) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UsersFilters) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UsersFilters) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUserType

`func (o *UsersFilters) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UsersFilters) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UsersFilters) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *UsersFilters) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetComponentUser

`func (o *UsersFilters) GetComponentUser() bool`

GetComponentUser returns the ComponentUser field if non-nil, zero value otherwise.

### GetComponentUserOk

`func (o *UsersFilters) GetComponentUserOk() (*bool, bool)`

GetComponentUserOk returns a tuple with the ComponentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentUser

`func (o *UsersFilters) SetComponentUser(v bool)`

SetComponentUser sets ComponentUser field to given value.

### HasComponentUser

`func (o *UsersFilters) HasComponentUser() bool`

HasComponentUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


