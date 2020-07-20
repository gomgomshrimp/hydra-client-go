// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra-client-go/models"
)

// NewRejectLoginRequestParams creates a new RejectLoginRequestParams object
// with the default values initialized.
func NewRejectLoginRequestParams() *RejectLoginRequestParams {
	var ()
	return &RejectLoginRequestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRejectLoginRequestParamsWithTimeout creates a new RejectLoginRequestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRejectLoginRequestParamsWithTimeout(timeout time.Duration) *RejectLoginRequestParams {
	var ()
	return &RejectLoginRequestParams{

		timeout: timeout,
	}
}

// NewRejectLoginRequestParamsWithContext creates a new RejectLoginRequestParams object
// with the default values initialized, and the ability to set a context for a request
func NewRejectLoginRequestParamsWithContext(ctx context.Context) *RejectLoginRequestParams {
	var ()
	return &RejectLoginRequestParams{

		Context: ctx,
	}
}

// NewRejectLoginRequestParamsWithHTTPClient creates a new RejectLoginRequestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRejectLoginRequestParamsWithHTTPClient(client *http.Client) *RejectLoginRequestParams {
	var ()
	return &RejectLoginRequestParams{
		HTTPClient: client,
	}
}

/*RejectLoginRequestParams contains all the parameters to send to the API endpoint
for the reject login request operation typically these are written to a http.Request
*/
type RejectLoginRequestParams struct {

	/*Body*/
	Body *models.RejectRequest
	/*LoginChallenge*/
	LoginChallenge string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the reject login request params
func (o *RejectLoginRequestParams) WithTimeout(timeout time.Duration) *RejectLoginRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reject login request params
func (o *RejectLoginRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reject login request params
func (o *RejectLoginRequestParams) WithContext(ctx context.Context) *RejectLoginRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reject login request params
func (o *RejectLoginRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reject login request params
func (o *RejectLoginRequestParams) WithHTTPClient(client *http.Client) *RejectLoginRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reject login request params
func (o *RejectLoginRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the reject login request params
func (o *RejectLoginRequestParams) WithBody(body *models.RejectRequest) *RejectLoginRequestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the reject login request params
func (o *RejectLoginRequestParams) SetBody(body *models.RejectRequest) {
	o.Body = body
}

// WithLoginChallenge adds the loginChallenge to the reject login request params
func (o *RejectLoginRequestParams) WithLoginChallenge(loginChallenge string) *RejectLoginRequestParams {
	o.SetLoginChallenge(loginChallenge)
	return o
}

// SetLoginChallenge adds the loginChallenge to the reject login request params
func (o *RejectLoginRequestParams) SetLoginChallenge(loginChallenge string) {
	o.LoginChallenge = loginChallenge
}

// WriteToRequest writes these params to a swagger request
func (o *RejectLoginRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// query param login_challenge
	qrLoginChallenge := o.LoginChallenge
	qLoginChallenge := qrLoginChallenge
	if qLoginChallenge != "" {
		if err := r.SetQueryParam("login_challenge", qLoginChallenge); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
