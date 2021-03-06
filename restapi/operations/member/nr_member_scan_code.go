// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	_"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	middleware "github.com/go-openapi/runtime/middleware"
	"tingtingapi/models"
	"fmt"
	"tingtingbackend/var"
)

// NrMemberScanCodeHandlerFunc turns a function with the right signature into a member scan code handler
type NrMemberScanCodeHandlerFunc func(NrMemberScanCodeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrMemberScanCodeHandlerFunc) Handle(params NrMemberScanCodeParams) middleware.Responder {
	return fn(params)
}

// NrMemberScanCodeHandler interface for that can handle valid member scan code params
type NrMemberScanCodeHandler interface {
	Handle(NrMemberScanCodeParams) middleware.Responder
}

// NewNrMemberScanCode creates a new http.Handler for the member scan code operation
func NewNrMemberScanCode(ctx *middleware.Context, handler NrMemberScanCodeHandler) *NrMemberScanCode {
	return &NrMemberScanCode{Context: ctx, Handler: handler}
}

/*NrMemberScanCode swagger:route GET /member/scanCode Member memberScanCode

扫码

扫码

*/
type NrMemberScanCode struct {
	Context *middleware.Context
	Handler NrMemberScanCodeHandler
}

func (o *NrMemberScanCode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrMemberScanCodeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok MemberScanCodeOK
	var response models.InlineResponse2009
	var status models.Response

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}

	var scan models.Scan
	fmt.Println("Params.CodeUrl = ",Params.CodeURL)
	db.Table("scan").Where("scanId=?",Params.CodeURL).Where(map[string]interface{}{"status":0}).Last(&scan)
    fmt.Println(scan)
	if(scan.BookId==0) {
		status.UnmarshalBinary([]byte(_var.Response200(201,"没找到对应书本")))
	}else{
		status.UnmarshalBinary([]byte(_var.Response200(200,"找到对应书本")))
		response.Data.BookID = scan.BookId
	}


	response.Status = &status

	ok.SetPayload(&response)
	o.Context.Respond(rw, r, route.Produces, route, ok)


}
