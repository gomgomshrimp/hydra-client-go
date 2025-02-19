/*
 * Ory Oathkeeper API
 *
 * Documentation for all of Ory Oathkeeper's APIs. 
 *
 * API version: v1.11.7
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CompletedRequest struct for CompletedRequest
type CompletedRequest struct {
	// RedirectURL is the URL which you should redirect the user to once the authentication process is completed.
	RedirectTo string `json:"redirect_to"`
}

// NewCompletedRequest instantiates a new CompletedRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompletedRequest(redirectTo string) *CompletedRequest {
	this := CompletedRequest{}
	this.RedirectTo = redirectTo
	return &this
}

// NewCompletedRequestWithDefaults instantiates a new CompletedRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompletedRequestWithDefaults() *CompletedRequest {
	this := CompletedRequest{}
	return &this
}

// GetRedirectTo returns the RedirectTo field value
func (o *CompletedRequest) GetRedirectTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectTo
}

// GetRedirectToOk returns a tuple with the RedirectTo field value
// and a boolean to check if the value has been set.
func (o *CompletedRequest) GetRedirectToOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RedirectTo, true
}

// SetRedirectTo sets field value
func (o *CompletedRequest) SetRedirectTo(v string) {
	o.RedirectTo = v
}

func (o CompletedRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["redirect_to"] = o.RedirectTo
	}
	return json.Marshal(toSerialize)
}

type NullableCompletedRequest struct {
	value *CompletedRequest
	isSet bool
}

func (v NullableCompletedRequest) Get() *CompletedRequest {
	return v.value
}

func (v *NullableCompletedRequest) Set(val *CompletedRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCompletedRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCompletedRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompletedRequest(val *CompletedRequest) *NullableCompletedRequest {
	return &NullableCompletedRequest{value: val, isSet: true}
}

func (v NullableCompletedRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompletedRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


