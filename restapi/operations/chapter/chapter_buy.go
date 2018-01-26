// Code generated by go-swagger; DO NOT EDIT.

package chapter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ChapterBuyHandlerFunc turns a function with the right signature into a chapter buy handler
type ChapterBuyHandlerFunc func(ChapterBuyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ChapterBuyHandlerFunc) Handle(params ChapterBuyParams) middleware.Responder {
	return fn(params)
}

// ChapterBuyHandler interface for that can handle valid chapter buy params
type ChapterBuyHandler interface {
	Handle(ChapterBuyParams) middleware.Responder
}

// NewChapterBuy creates a new http.Handler for the chapter buy operation
func NewChapterBuy(ctx *middleware.Context, handler ChapterBuyHandler) *ChapterBuy {
	return &ChapterBuy{Context: ctx, Handler: handler}
}

/*ChapterBuy swagger:route GET /chapter/buy Chapter chapterBuy

购买章节

购买章节

*/
type ChapterBuy struct {
	Context *middleware.Context
	Handler ChapterBuyHandler
}

func (o *ChapterBuy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewChapterBuyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}