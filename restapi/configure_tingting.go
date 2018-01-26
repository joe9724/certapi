// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"tingtingapi/restapi/operations"
	"tingtingapi/restapi/operations/album"
	"tingtingapi/restapi/operations/banner"
	"tingtingapi/restapi/operations/book"
	"tingtingapi/restapi/operations/category"
	"tingtingapi/restapi/operations/chapter"
	"tingtingapi/restapi/operations/icon"
	"tingtingapi/restapi/operations/member"
	"tingtingapi/restapi/operations/recommend"
	"tingtingapi/restapi/operations/search"
	"github.com/didip/tollbooth"
	_"time"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name  --spec ../swagger.json

func configureFlags(api *operations.TingtingAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TingtingAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	api.AlbumNrAlbumFavHandler = album.NrAlbumFavHandlerFunc(func(params album.NrAlbumFavParams) middleware.Responder {
		return middleware.NotImplemented("operation album.NrAlbumFav has not yet been implemented")
	})
	api.BannerNrBannerDetailHandler = banner.NrBannerDetailHandlerFunc(func(params banner.NrBannerDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation banner.NrBannerDetail has not yet been implemented")
	})
	api.BannerNrBannerListHandler = banner.NrBannerListHandlerFunc(func(params banner.NrBannerListParams) middleware.Responder {
		return middleware.NotImplemented("operation banner.NrBannerList has not yet been implemented")
	})
	api.BookNrBookFavHandler = book.NrBookFavHandlerFunc(func(params book.NrBookFavParams) middleware.Responder {
		return middleware.NotImplemented("operation book.NrBookFav has not yet been implemented")
	})
	api.CategoryNrCategoryListHandler = category.NrCategoryListHandlerFunc(func(params category.NrCategoryListParams) middleware.Responder {
		return middleware.NotImplemented("operation category.NrCategoryList has not yet been implemented")
	})
	api.CategoryNrCategorySubListHandler = category.NrCategorySubListHandlerFunc(func(params category.NrCategorySubListParams) middleware.Responder {
		return middleware.NotImplemented("operation category.NrCategorySubList has not yet been implemented")
	})
	api.IconNrIconListHandler = icon.NrIconListHandlerFunc(func(params icon.NrIconListParams) middleware.Responder {
		return middleware.NotImplemented("operation icon.NrIconList has not yet been implemented")
	})
	api.MemberNrMemberCheckRechargeHandler = member.NrMemberCheckRechargeHandlerFunc(func(params member.NrMemberCheckRechargeParams) middleware.Responder {
		return middleware.NotImplemented("operation member.NrMemberCheckRecharge has not yet been implemented")
	})
	api.MemberNrMemberInitHandler = member.NrMemberInitHandlerFunc(func(params member.NrMemberInitParams) middleware.Responder {
		return middleware.NotImplemented("operation member.NrMemberInit has not yet been implemented")
	})
	api.MemberNrMemberLoginByThirdPartyHandler = member.NrMemberLoginByThirdPartyHandlerFunc(func(params member.NrMemberLoginByThirdPartyParams) middleware.Responder {
		return middleware.NotImplemented("operation member.NrMemberLoginByThirdParty has not yet been implemented")
	})
	api.MemberNrMemberRegisterSendSmsHandler = member.NrMemberRegisterSendSmsHandlerFunc(func(params member.NrMemberRegisterSendSmsParams) middleware.Responder {
		return middleware.NotImplemented("operation member.NrMemberRegisterSendSms has not yet been implemented")
	})
	api.MemberNrMemberRegisterMemberHandler = member.NrMemberRegisterMemberHandlerFunc(func(params member.NrMemberRegisterMemberParams) middleware.Responder {
		return middleware.NotImplemented("operation member.NrMemberRegisterMember has not yet been implemented")
	})
	api.MemberNrMemberReportErrHandler = member.NrMemberReportErrHandlerFunc(func(params member.NrMemberReportErrParams) middleware.Responder {
		return middleware.NotImplemented("operation member.NrMemberReportErr has not yet been implemented")
	})
	api.MemberNrMemberScanCodeHandler = member.NrMemberScanCodeHandlerFunc(func(params member.NrMemberScanCodeParams) middleware.Responder {
		return middleware.NotImplemented("operation member.NrMemberScanCode has not yet been implemented")
	})
	api.MemberNrMemberUploadRecordHandler = member.NrMemberUploadRecordHandlerFunc(func(params member.NrMemberUploadRecordParams) middleware.Responder {
		return middleware.NotImplemented("operation member.NrMemberUploadRecord has not yet been implemented")
	})
	api.MemberNrOrderSerialNuberHandler = member.NrOrderSerialNuberHandlerFunc(func(params member.NrOrderSerialNuberParams) middleware.Responder {
		return middleware.NotImplemented("operation member.NrOrderSerialNuber has not yet been implemented")
	})
	api.BookAlbumBookListHandler = book.AlbumBookListHandlerFunc(func(params book.AlbumBookListParams) middleware.Responder {
		return middleware.NotImplemented("operation book.AlbumBookList has not yet been implemented")
	})
	api.AlbumAlbumBuyHandler = album.AlbumBuyHandlerFunc(func(params album.AlbumBuyParams) middleware.Responder {
		return middleware.NotImplemented("operation album.AlbumBuy has not yet been implemented")
	})
	api.AlbumAlbumDetailHandler = album.AlbumDetailHandlerFunc(func(params album.AlbumDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation album.AlbumDetail has not yet been implemented")
	})
	api.AlbumAlbumListHandler = album.AlbumListHandlerFunc(func(params album.AlbumListParams) middleware.Responder {
		return middleware.NotImplemented("operation album.AlbumList has not yet been implemented")
	})
	api.AlbumAlbumListFavHandler = album.AlbumListFavHandlerFunc(func(params album.AlbumListFavParams) middleware.Responder {
		return middleware.NotImplemented("operation album.AlbumListFav has not yet been implemented")
	})
	api.AlbumAlbumListMatchHandler = album.AlbumListMatchHandlerFunc(func(params album.AlbumListMatchParams) middleware.Responder {
		return middleware.NotImplemented("operation album.AlbumListMatch has not yet been implemented")
	})
	api.BookBookBuyHandler = book.BookBuyHandlerFunc(func(params book.BookBuyParams) middleware.Responder {
		return middleware.NotImplemented("operation book.BookBuy has not yet been implemented")
	})
	api.BookBookChapListHandler = book.BookChapListHandlerFunc(func(params book.BookChapListParams) middleware.Responder {
		return middleware.NotImplemented("operation book.BookChapList has not yet been implemented")
	})
	api.BookBookDetailHandler = book.BookDetailHandlerFunc(func(params book.BookDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation book.BookDetail has not yet been implemented")
	})
	api.BookBookListHandler = book.BookListHandlerFunc(func(params book.BookListParams) middleware.Responder {
		return middleware.NotImplemented("operation book.BookList has not yet been implemented")
	})
	api.BookBookListFavHandler = book.BookListFavHandlerFunc(func(params book.BookListFavParams) middleware.Responder {
		return middleware.NotImplemented("operation book.BookListFav has not yet been implemented")
	})
	api.CategoryCategoryDetailHandler = category.CategoryDetailHandlerFunc(func(params category.CategoryDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation category.CategoryDetail has not yet been implemented")
	})
	api.ChapterChapterBuyHandler = chapter.ChapterBuyHandlerFunc(func(params chapter.ChapterBuyParams) middleware.Responder {
		return middleware.NotImplemented("operation chapter.ChapterBuy has not yet been implemented")
	})
	api.ChapterChapterDetailHandler = chapter.ChapterDetailHandlerFunc(func(params chapter.ChapterDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation chapter.ChapterDetail has not yet been implemented")
	})
	api.ChapterChapterFavHandler = chapter.ChapterFavHandlerFunc(func(params chapter.ChapterFavParams) middleware.Responder {
		return middleware.NotImplemented("operation chapter.ChapterFav has not yet been implemented")
	})
	api.ChapterChapterFavListHandler = chapter.ChapterFavListHandlerFunc(func(params chapter.ChapterFavListParams) middleware.Responder {
		return middleware.NotImplemented("operation chapter.ChapterFavList has not yet been implemented")
	})
	api.ChapterChapterHistoryListHandler = chapter.ChapterHistoryListHandlerFunc(func(params chapter.ChapterHistoryListParams) middleware.Responder {
		return middleware.NotImplemented("operation chapter.ChapterHistoryList has not yet been implemented")
	})
	api.MemberFeedbackHandler = member.FeedbackHandlerFunc(func(params member.FeedbackParams) middleware.Responder {
		return middleware.NotImplemented("operation member.Feedback has not yet been implemented")
	})
	api.MemberFindPassEditPassHandler = member.FindPassEditPassHandlerFunc(func(params member.FindPassEditPassParams) middleware.Responder {
		return middleware.NotImplemented("operation member.FindPassEditPass has not yet been implemented")
	})
	api.MemberFindPassSendSmsHandler = member.FindPassSendSmsHandlerFunc(func(params member.FindPassSendSmsParams) middleware.Responder {
		return middleware.NotImplemented("operation member.FindPassSendSms has not yet been implemented")
	})
	api.MemberLoginHandler = member.LoginHandlerFunc(func(params member.LoginParams) middleware.Responder {
		return middleware.NotImplemented("operation member.Login has not yet been implemented")
	})
	api.MemberMsgListHandler = member.MsgListHandlerFunc(func(params member.MsgListParams) middleware.Responder {
		return middleware.NotImplemented("operation member.MsgList has not yet been implemented")
	})
	api.MemberOrderListHandler = member.OrderListHandlerFunc(func(params member.OrderListParams) middleware.Responder {
		return middleware.NotImplemented("operation member.OrderList has not yet been implemented")
	})
	api.RecommendRecommendHandler = recommend.RecommendHandlerFunc(func(params recommend.RecommendParams) middleware.Responder {
		return middleware.NotImplemented("operation recommend.Recommend has not yet been implemented")
	})
	api.MemberRecordListHandler = member.RecordListHandlerFunc(func(params member.RecordListParams) middleware.Responder {
		return middleware.NotImplemented("operation member.RecordList has not yet been implemented")
	})
	api.MemberRegisterHandler = member.RegisterHandlerFunc(func(params member.RegisterParams) middleware.Responder {
		return middleware.NotImplemented("operation member.Register has not yet been implemented")
	})
	api.SearchSearchHandler = search.SearchHandlerFunc(func(params search.SearchParams) middleware.Responder {
		return middleware.NotImplemented("operation search.Search has not yet been implemented")
	})
	api.MemberMemberDetailHandler = member.MemberDetailHandlerFunc(func(params member.MemberDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation member.MemberDetail has not yet been implemented")
	})
	api.MemberNrMemberEditHandler = member.NrMemberEditHandlerFunc(func(params member.NrMemberEditParams) middleware.Responder {
		return middleware.NotImplemented("operation member.NrMemberEdit has not yet been implemented")
	})
	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	limiter := tollbooth.NewLimiter(1000, nil)
	limiter.SetMessage("api 每1000秒只能请求一次")
	limiter.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"})
	return tollbooth.LimitHandler(limiter,handler)
	//return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
