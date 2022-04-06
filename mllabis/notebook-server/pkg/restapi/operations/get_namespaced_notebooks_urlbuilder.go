// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// GetNamespacedNotebooksURL generates an URL for the get namespaced notebooks operation
type GetNamespacedNotebooksURL struct {
	Namespace string

	ClusterName *string
	Page        *string
	Size        *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetNamespacedNotebooksURL) WithBasePath(bp string) *GetNamespacedNotebooksURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetNamespacedNotebooksURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetNamespacedNotebooksURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/aide/v1/namespaces/{namespace}/notebooks"

	namespace := o.Namespace
	if namespace != "" {
		_path = strings.Replace(_path, "{namespace}", namespace, -1)
	} else {
		return nil, errors.New("namespace is required on GetNamespacedNotebooksURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var clusterNameQ string
	if o.ClusterName != nil {
		clusterNameQ = *o.ClusterName
	}
	if clusterNameQ != "" {
		qs.Set("clusterName", clusterNameQ)
	}

	var pageQ string
	if o.Page != nil {
		pageQ = *o.Page
	}
	if pageQ != "" {
		qs.Set("page", pageQ)
	}

	var sizeQ string
	if o.Size != nil {
		sizeQ = *o.Size
	}
	if sizeQ != "" {
		qs.Set("size", sizeQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetNamespacedNotebooksURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetNamespacedNotebooksURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetNamespacedNotebooksURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetNamespacedNotebooksURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetNamespacedNotebooksURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetNamespacedNotebooksURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
