// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetNamespacedNotebookLogParams creates a new GetNamespacedNotebookLogParams object
// with the default values initialized.
func NewGetNamespacedNotebookLogParams() GetNamespacedNotebookLogParams {

	var (
		// initialize parameters with default values

		ascDefault         = bool(false)
		currentPageDefault = int64(1)

		pageSizeDefault = int64(10)
	)

	return GetNamespacedNotebookLogParams{
		Asc: &ascDefault,

		CurrentPage: &currentPageDefault,

		PageSize: &pageSizeDefault,
	}
}

// GetNamespacedNotebookLogParams contains all the bound params for the get namespaced notebook log operation
// typically these are obtained from a http.Request
//
// swagger:parameters getNamespacedNotebookLog
type GetNamespacedNotebookLogParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*whether sort by time  in asc, default false (desc).
	  In: query
	  Default: false
	*/
	Asc *bool
	/*
	  In: query
	  Default: 1
	*/
	CurrentPage *int64
	/*
	  Required: true
	  In: path
	*/
	Namespace string
	/*
	  Required: true
	  In: path
	*/
	NotebookName string
	/*
	  In: query
	  Default: 10
	*/
	PageSize *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetNamespacedNotebookLogParams() beforehand.
func (o *GetNamespacedNotebookLogParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAsc, qhkAsc, _ := qs.GetOK("asc")
	if err := o.bindAsc(qAsc, qhkAsc, route.Formats); err != nil {
		res = append(res, err)
	}

	qCurrentPage, qhkCurrentPage, _ := qs.GetOK("currentPage")
	if err := o.bindCurrentPage(qCurrentPage, qhkCurrentPage, route.Formats); err != nil {
		res = append(res, err)
	}

	rNamespace, rhkNamespace, _ := route.Params.GetOK("namespace")
	if err := o.bindNamespace(rNamespace, rhkNamespace, route.Formats); err != nil {
		res = append(res, err)
	}

	rNotebookName, rhkNotebookName, _ := route.Params.GetOK("notebook_name")
	if err := o.bindNotebookName(rNotebookName, rhkNotebookName, route.Formats); err != nil {
		res = append(res, err)
	}

	qPageSize, qhkPageSize, _ := qs.GetOK("pageSize")
	if err := o.bindPageSize(qPageSize, qhkPageSize, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAsc binds and validates parameter Asc from query.
func (o *GetNamespacedNotebookLogParams) bindAsc(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetNamespacedNotebookLogParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("asc", "query", "bool", raw)
	}
	o.Asc = &value

	return nil
}

// bindCurrentPage binds and validates parameter CurrentPage from query.
func (o *GetNamespacedNotebookLogParams) bindCurrentPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetNamespacedNotebookLogParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("currentPage", "query", "int64", raw)
	}
	o.CurrentPage = &value

	return nil
}

// bindNamespace binds and validates parameter Namespace from path.
func (o *GetNamespacedNotebookLogParams) bindNamespace(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Namespace = raw

	return nil
}

// bindNotebookName binds and validates parameter NotebookName from path.
func (o *GetNamespacedNotebookLogParams) bindNotebookName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.NotebookName = raw

	return nil
}

// bindPageSize binds and validates parameter PageSize from query.
func (o *GetNamespacedNotebookLogParams) bindPageSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetNamespacedNotebookLogParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("pageSize", "query", "int64", raw)
	}
	o.PageSize = &value

	return nil
}
