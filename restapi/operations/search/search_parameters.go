// Code generated by go-swagger; DO NOT EDIT.

package search

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

// NewSearchParams creates a new SearchParams object
// with the default values initialized.
func NewSearchParams() SearchParams {
	var ()
	return SearchParams{}
}

// SearchParams contains all the bound params for the search operation
// typically these are obtained from a http.Request
//
// swagger:parameters search
type SearchParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*客户端类型
	  In: query
	*/
	Client *string
	/*唯一识别号
	  In: query
	*/
	Imei *string
	/*关键词
	  In: query
	*/
	Keyword *string
	/*用户id
	  In: query
	*/
	MemberID *string
	/*版本号
	  In: query
	*/
	Version *string

	Action *string

	PageIndex *int64
	/*分页尺寸
	  In: query
	*/
	PageSize *int64
	/*标签
	  In: query
	*/


}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *SearchParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qClient, qhkClient, _ := qs.GetOK("client")
	if err := o.bindClient(qClient, qhkClient, route.Formats); err != nil {
		res = append(res, err)
	}

	qImei, qhkImei, _ := qs.GetOK("imei")
	if err := o.bindImei(qImei, qhkImei, route.Formats); err != nil {
		res = append(res, err)
	}

	qPageIndex, qhkPageIndex, _ := qs.GetOK("pageIndex")
	if err := o.bindPageIndex(qPageIndex, qhkPageIndex, route.Formats); err != nil {
		res = append(res, err)
	}

	qPageSize, qhkPageSize, _ := qs.GetOK("pageSize")
	if err := o.bindPageSize(qPageSize, qhkPageSize, route.Formats); err != nil {
		res = append(res, err)
	}

	qAction, qhkAction, _ := qs.GetOK("action")
	if err := o.bindAction(qAction, qhkAction, route.Formats); err != nil {
		res = append(res, err)
	}

	qKeyword, qhkKeyword, _ := qs.GetOK("keyword")
	if err := o.bindKeyword(qKeyword, qhkKeyword, route.Formats); err != nil {
		res = append(res, err)
	}

	qMemberID, qhkMemberID, _ := qs.GetOK("memberId")
	if err := o.bindMemberID(qMemberID, qhkMemberID, route.Formats); err != nil {
		res = append(res, err)
	}

	qVersion, qhkVersion, _ := qs.GetOK("version")
	if err := o.bindVersion(qVersion, qhkVersion, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SearchParams) bindClient(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Client = &raw

	return nil
}

func (o *SearchParams) bindPageIndex(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("pageIndex", "query", "int64", raw)
	}
	o.PageIndex = &value

	return nil
}

func (o *SearchParams) bindPageSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("pageSize", "query", "int64", raw)
	}
	o.PageSize = &value

	return nil
}

func (o *SearchParams) bindAction(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Action = &raw

	return nil
}

func (o *SearchParams) bindImei(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Imei = &raw

	return nil
}

func (o *SearchParams) bindKeyword(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Keyword = &raw

	return nil
}

func (o *SearchParams) bindMemberID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.MemberID = &raw

	return nil
}

func (o *SearchParams) bindVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Version = &raw

	return nil
}