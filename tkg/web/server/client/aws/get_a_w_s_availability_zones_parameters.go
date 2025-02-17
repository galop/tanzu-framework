// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAWSAvailabilityZonesParams creates a new GetAWSAvailabilityZonesParams object
// with the default values initialized.
func NewGetAWSAvailabilityZonesParams() *GetAWSAvailabilityZonesParams {

	return &GetAWSAvailabilityZonesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAWSAvailabilityZonesParamsWithTimeout creates a new GetAWSAvailabilityZonesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAWSAvailabilityZonesParamsWithTimeout(timeout time.Duration) *GetAWSAvailabilityZonesParams {

	return &GetAWSAvailabilityZonesParams{

		timeout: timeout,
	}
}

// NewGetAWSAvailabilityZonesParamsWithContext creates a new GetAWSAvailabilityZonesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAWSAvailabilityZonesParamsWithContext(ctx context.Context) *GetAWSAvailabilityZonesParams {

	return &GetAWSAvailabilityZonesParams{

		Context: ctx,
	}
}

// NewGetAWSAvailabilityZonesParamsWithHTTPClient creates a new GetAWSAvailabilityZonesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAWSAvailabilityZonesParamsWithHTTPClient(client *http.Client) *GetAWSAvailabilityZonesParams {

	return &GetAWSAvailabilityZonesParams{
		HTTPClient: client,
	}
}

/*
GetAWSAvailabilityZonesParams contains all the parameters to send to the API endpoint
for the get a w s availability zones operation typically these are written to a http.Request
*/
type GetAWSAvailabilityZonesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get a w s availability zones params
func (o *GetAWSAvailabilityZonesParams) WithTimeout(timeout time.Duration) *GetAWSAvailabilityZonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get a w s availability zones params
func (o *GetAWSAvailabilityZonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get a w s availability zones params
func (o *GetAWSAvailabilityZonesParams) WithContext(ctx context.Context) *GetAWSAvailabilityZonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get a w s availability zones params
func (o *GetAWSAvailabilityZonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get a w s availability zones params
func (o *GetAWSAvailabilityZonesParams) WithHTTPClient(client *http.Client) *GetAWSAvailabilityZonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get a w s availability zones params
func (o *GetAWSAvailabilityZonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAWSAvailabilityZonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
