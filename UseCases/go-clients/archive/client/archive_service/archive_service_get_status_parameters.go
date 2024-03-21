// Code generated by go-swagger; DO NOT EDIT.

package archive_service

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
)

// NewArchiveServiceGetStatusParams creates a new ArchiveServiceGetStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewArchiveServiceGetStatusParams() *ArchiveServiceGetStatusParams {
	return &ArchiveServiceGetStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewArchiveServiceGetStatusParamsWithTimeout creates a new ArchiveServiceGetStatusParams object
// with the ability to set a timeout on a request.
func NewArchiveServiceGetStatusParamsWithTimeout(timeout time.Duration) *ArchiveServiceGetStatusParams {
	return &ArchiveServiceGetStatusParams{
		timeout: timeout,
	}
}

// NewArchiveServiceGetStatusParamsWithContext creates a new ArchiveServiceGetStatusParams object
// with the ability to set a context for a request.
func NewArchiveServiceGetStatusParamsWithContext(ctx context.Context) *ArchiveServiceGetStatusParams {
	return &ArchiveServiceGetStatusParams{
		Context: ctx,
	}
}

// NewArchiveServiceGetStatusParamsWithHTTPClient creates a new ArchiveServiceGetStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewArchiveServiceGetStatusParamsWithHTTPClient(client *http.Client) *ArchiveServiceGetStatusParams {
	return &ArchiveServiceGetStatusParams{
		HTTPClient: client,
	}
}

/*
ArchiveServiceGetStatusParams contains all the parameters to send to the API endpoint

	for the archive service get status operation.

	Typically these are written to a http.Request.
*/
type ArchiveServiceGetStatusParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the archive service get status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ArchiveServiceGetStatusParams) WithDefaults() *ArchiveServiceGetStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the archive service get status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ArchiveServiceGetStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the archive service get status params
func (o *ArchiveServiceGetStatusParams) WithTimeout(timeout time.Duration) *ArchiveServiceGetStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the archive service get status params
func (o *ArchiveServiceGetStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the archive service get status params
func (o *ArchiveServiceGetStatusParams) WithContext(ctx context.Context) *ArchiveServiceGetStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the archive service get status params
func (o *ArchiveServiceGetStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the archive service get status params
func (o *ArchiveServiceGetStatusParams) WithHTTPClient(client *http.Client) *ArchiveServiceGetStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the archive service get status params
func (o *ArchiveServiceGetStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ArchiveServiceGetStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}