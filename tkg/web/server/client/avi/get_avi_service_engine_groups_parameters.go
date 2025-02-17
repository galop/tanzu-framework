// Code generated by go-swagger; DO NOT EDIT.

package avi

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

// NewGetAviServiceEngineGroupsParams creates a new GetAviServiceEngineGroupsParams object
// with the default values initialized.
func NewGetAviServiceEngineGroupsParams() *GetAviServiceEngineGroupsParams {

	return &GetAviServiceEngineGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAviServiceEngineGroupsParamsWithTimeout creates a new GetAviServiceEngineGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAviServiceEngineGroupsParamsWithTimeout(timeout time.Duration) *GetAviServiceEngineGroupsParams {

	return &GetAviServiceEngineGroupsParams{

		timeout: timeout,
	}
}

// NewGetAviServiceEngineGroupsParamsWithContext creates a new GetAviServiceEngineGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAviServiceEngineGroupsParamsWithContext(ctx context.Context) *GetAviServiceEngineGroupsParams {

	return &GetAviServiceEngineGroupsParams{

		Context: ctx,
	}
}

// NewGetAviServiceEngineGroupsParamsWithHTTPClient creates a new GetAviServiceEngineGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAviServiceEngineGroupsParamsWithHTTPClient(client *http.Client) *GetAviServiceEngineGroupsParams {

	return &GetAviServiceEngineGroupsParams{
		HTTPClient: client,
	}
}

/*
GetAviServiceEngineGroupsParams contains all the parameters to send to the API endpoint
for the get avi service engine groups operation typically these are written to a http.Request
*/
type GetAviServiceEngineGroupsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get avi service engine groups params
func (o *GetAviServiceEngineGroupsParams) WithTimeout(timeout time.Duration) *GetAviServiceEngineGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get avi service engine groups params
func (o *GetAviServiceEngineGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get avi service engine groups params
func (o *GetAviServiceEngineGroupsParams) WithContext(ctx context.Context) *GetAviServiceEngineGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get avi service engine groups params
func (o *GetAviServiceEngineGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get avi service engine groups params
func (o *GetAviServiceEngineGroupsParams) WithHTTPClient(client *http.Client) *GetAviServiceEngineGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get avi service engine groups params
func (o *GetAviServiceEngineGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAviServiceEngineGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
