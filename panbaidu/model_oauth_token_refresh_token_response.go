/*
xpan

xpanapi

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package panbaidu

import (
	"encoding/json"
)

// OauthTokenRefreshTokenResponse struct for OauthTokenRefreshTokenResponse
type OauthTokenRefreshTokenResponse struct {
	ExpiresIn *int32 `json:"expires_in,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	AccessToken *string `json:"access_token,omitempty"`
	SessionSecret *string `json:"session_secret,omitempty"`
	SessionKey *string `json:"session_key,omitempty"`
	Scope *string `json:"scope,omitempty"`
}

// NewOauthTokenRefreshTokenResponse instantiates a new OauthTokenRefreshTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOauthTokenRefreshTokenResponse() *OauthTokenRefreshTokenResponse {
	this := OauthTokenRefreshTokenResponse{}
	return &this
}

// NewOauthTokenRefreshTokenResponseWithDefaults instantiates a new OauthTokenRefreshTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOauthTokenRefreshTokenResponseWithDefaults() *OauthTokenRefreshTokenResponse {
	this := OauthTokenRefreshTokenResponse{}
	return &this
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *OauthTokenRefreshTokenResponse) GetExpiresIn() int32 {
	if o == nil || o.ExpiresIn == nil {
		var ret int32
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthTokenRefreshTokenResponse) GetExpiresInOk() (*int32, bool) {
	if o == nil || o.ExpiresIn == nil {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *OauthTokenRefreshTokenResponse) HasExpiresIn() bool {
	if o != nil && o.ExpiresIn != nil {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int32 and assigns it to the ExpiresIn field.
func (o *OauthTokenRefreshTokenResponse) SetExpiresIn(v int32) {
	o.ExpiresIn = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *OauthTokenRefreshTokenResponse) GetRefreshToken() string {
	if o == nil || o.RefreshToken == nil {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthTokenRefreshTokenResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil || o.RefreshToken == nil {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *OauthTokenRefreshTokenResponse) HasRefreshToken() bool {
	if o != nil && o.RefreshToken != nil {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *OauthTokenRefreshTokenResponse) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *OauthTokenRefreshTokenResponse) GetAccessToken() string {
	if o == nil || o.AccessToken == nil {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthTokenRefreshTokenResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *OauthTokenRefreshTokenResponse) HasAccessToken() bool {
	if o != nil && o.AccessToken != nil {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *OauthTokenRefreshTokenResponse) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetSessionSecret returns the SessionSecret field value if set, zero value otherwise.
func (o *OauthTokenRefreshTokenResponse) GetSessionSecret() string {
	if o == nil || o.SessionSecret == nil {
		var ret string
		return ret
	}
	return *o.SessionSecret
}

// GetSessionSecretOk returns a tuple with the SessionSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthTokenRefreshTokenResponse) GetSessionSecretOk() (*string, bool) {
	if o == nil || o.SessionSecret == nil {
		return nil, false
	}
	return o.SessionSecret, true
}

// HasSessionSecret returns a boolean if a field has been set.
func (o *OauthTokenRefreshTokenResponse) HasSessionSecret() bool {
	if o != nil && o.SessionSecret != nil {
		return true
	}

	return false
}

// SetSessionSecret gets a reference to the given string and assigns it to the SessionSecret field.
func (o *OauthTokenRefreshTokenResponse) SetSessionSecret(v string) {
	o.SessionSecret = &v
}

// GetSessionKey returns the SessionKey field value if set, zero value otherwise.
func (o *OauthTokenRefreshTokenResponse) GetSessionKey() string {
	if o == nil || o.SessionKey == nil {
		var ret string
		return ret
	}
	return *o.SessionKey
}

// GetSessionKeyOk returns a tuple with the SessionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthTokenRefreshTokenResponse) GetSessionKeyOk() (*string, bool) {
	if o == nil || o.SessionKey == nil {
		return nil, false
	}
	return o.SessionKey, true
}

// HasSessionKey returns a boolean if a field has been set.
func (o *OauthTokenRefreshTokenResponse) HasSessionKey() bool {
	if o != nil && o.SessionKey != nil {
		return true
	}

	return false
}

// SetSessionKey gets a reference to the given string and assigns it to the SessionKey field.
func (o *OauthTokenRefreshTokenResponse) SetSessionKey(v string) {
	o.SessionKey = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *OauthTokenRefreshTokenResponse) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthTokenRefreshTokenResponse) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *OauthTokenRefreshTokenResponse) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *OauthTokenRefreshTokenResponse) SetScope(v string) {
	o.Scope = &v
}

func (o OauthTokenRefreshTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpiresIn != nil {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if o.RefreshToken != nil {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	if o.AccessToken != nil {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.SessionSecret != nil {
		toSerialize["session_secret"] = o.SessionSecret
	}
	if o.SessionKey != nil {
		toSerialize["session_key"] = o.SessionKey
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableOauthTokenRefreshTokenResponse struct {
	value *OauthTokenRefreshTokenResponse
	isSet bool
}

func (v NullableOauthTokenRefreshTokenResponse) Get() *OauthTokenRefreshTokenResponse {
	return v.value
}

func (v *NullableOauthTokenRefreshTokenResponse) Set(val *OauthTokenRefreshTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOauthTokenRefreshTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOauthTokenRefreshTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOauthTokenRefreshTokenResponse(val *OauthTokenRefreshTokenResponse) *NullableOauthTokenRefreshTokenResponse {
	return &NullableOauthTokenRefreshTokenResponse{value: val, isSet: true}
}

func (v NullableOauthTokenRefreshTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOauthTokenRefreshTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


