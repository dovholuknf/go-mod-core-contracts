//
// Copyright (C) 2022 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"net/http"
)

// AuthenticationInjector defines an interface to obtain a JWT and secure transport for remote service calls
type AuthenticationInjector interface {
	// AddAuthenticationData mutates an HTTP request to add authentication data
	// (suth as an Authorization: header) to an outbound HTTP request
	AddAuthenticationData(_ *http.Request) error
	// SecureTransportProvider Returns the configured *http.Transport or a fallback *http.Transport to use when making the request
	SecureTransportProvider
}

// SecureTransportProvider defines an interface to obtain a secure http.Transport to use when making http requests.
// If a fallback RoundTripper is configured it will be used
type SecureTransportProvider interface {
	// RoundTripperWithFallback Returns the configured *http.Transport to use when making the request and a fallback if needed
	RoundTripperWithFallback() (http.RoundTripper, http.RoundTripper)
}
