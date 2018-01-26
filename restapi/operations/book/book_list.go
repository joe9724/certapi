// Code generated by go-swagger; DO NOT EDIT.

package book

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// BookListHandlerFunc turns a function with the right signature into a book list handler
type BookListHandlerFunc func(BookListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn BookListHandlerFunc) Handle(params BookListParams) middleware.Responder {
	return fn(params)
}

// BookListHandler interface for that can handle valid book list params
type BookListHandler interface {
	Handle(BookListParams) middleware.Responder
}

// NewBookList creates a new http.Handler for the book list operation
func NewBookList(ctx *middleware.Context, handler BookListHandler) *BookList {
	return &BookList{Context: ctx, Handler: handler}
}

/*BookList swagger:route GET /book/list Book bookList

根据用户信息匹配书本

根据用户信息匹配书本

*/
type BookList struct {
	Context *middleware.Context
	Handler BookListHandler
}

func (o *BookList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewBookListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}