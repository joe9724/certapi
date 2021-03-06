// Code generated by go-swagger; DO NOT EDIT.

package category

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

// NrCategoryListHandlerFunc turns a function with the right signature into a category list handler
type NrCategoryListHandlerFunc func(NrCategoryListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrCategoryListHandlerFunc) Handle(params NrCategoryListParams) middleware.Responder {
	return fn(params)
}

// NrCategoryListHandler interface for that can handle valid category list params
type NrCategoryListHandler interface {
	Handle(NrCategoryListParams) middleware.Responder
}

// NewNrCategoryList creates a new http.Handler for the category list operation
func NewNrCategoryList(ctx *middleware.Context, handler NrCategoryListHandler) *NrCategoryList {
	return &NrCategoryList{Context: ctx, Handler: handler}
}

/*NrCategoryList swagger:route GET /category/list Category categoryList

获取分类列表

获取分类列表

*/
type NrCategoryList struct {
	Context *middleware.Context
	Handler NrCategoryListHandler
}

func (o *NrCategoryList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrCategoryListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok CategoryListOK
	var response models.InlineResponse20023
	var categoryList models.InlineResponse20023Data

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	if (*(Params.ParentID) == int64(-1)){
		db.Table("sub_category_items").Where(map[string]interface{}{"status":0}).Find(&categoryList).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex)*(*(Params.PageSize)))
	}else{
		db.Table("sub_category_items").Where(map[string]interface{}{"status":0}).Where("category_id=?",Params.ParentID).Find(&categoryList).Limit(*(Params.PageSize)).Offset(*(Params.PageIndex)*(*(Params.PageSize)))
	}
	//query

	//data
	fmt.Println(categoryList)
	response.Data = categoryList

	//status
	var status models.Response
	status.UnmarshalBinary([]byte(_var.Response200(200,"ok")))
	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}
