// Code generated by go-swagger; DO NOT EDIT.

package banner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewNrBannerDetailParams creates a new NrBannerDetailParams object
// with the default values initialized.
func NewNrBannerDetailParams() NrBannerDetailParams {
	var ()
	return NrBannerDetailParams{}
}

// NrBannerDetailParams contains all the bound params for the banner detail operation
// typically these are obtained from a http.Request
//
// swagger:parameters /banner/detail
type NrBannerDetailParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*bannerId
	  In: query
	*/
	BannerID *string
	/*客户端类型
	  In: query
	*/
	Client *string
	/*唯一识别号
	  In: query
	*/
	Imei *string
	/*用户id
	  In: query
	*/
	MemberID *string
	/*版本号
	  In: query
	*/
	Version *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *NrBannerDetailParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qBannerID, qhkBannerID, _ := qs.GetOK("bannerId")
	if err := o.bindBannerID(qBannerID, qhkBannerID, route.Formats); err != nil {
		res = append(res, err)
	}

	qClient, qhkClient, _ := qs.GetOK("client")
	if err := o.bindClient(qClient, qhkClient, route.Formats); err != nil {
		res = append(res, err)
	}

	qImei, qhkImei, _ := qs.GetOK("imei")
	if err := o.bindImei(qImei, qhkImei, route.Formats); err != nil {
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

func (o *NrBannerDetailParams) bindBannerID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.BannerID = &raw

	return nil
}

func (o *NrBannerDetailParams) bindClient(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrBannerDetailParams) bindImei(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrBannerDetailParams) bindMemberID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrBannerDetailParams) bindVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
