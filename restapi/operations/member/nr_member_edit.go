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
	"tingtingapi/var"
	"io/ioutil"
	"time"
	"strings"
	"runtime"
	"strconv"
)

// NrMemberEditHandlerFunc turns a function with the right signature into a member edit handler
type NrMemberEditHandlerFunc func(NrMemberEditParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrMemberEditHandlerFunc) Handle(params NrMemberEditParams) middleware.Responder {
	return fn(params)
}

// NrMemberEditHandler interface for that can handle valid member edit params
type NrMemberEditHandler interface {
	Handle(NrMemberEditParams) middleware.Responder
}

// NewNrMemberEdit creates a new http.Handler for the member edit operation
func NewNrMemberEdit(ctx *middleware.Context, handler NrMemberEditHandler) *NrMemberEdit {
	return &NrMemberEdit{Context: ctx, Handler: handler}
}

/*NrMemberEdit swagger:route POST /member/edit Member memberEdit

修改用户资料

修改用户资料

*/
type NrMemberEdit struct {
	Context *middleware.Context
	Handler NrMemberEditHandler
}

func (o *NrMemberEdit) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrMemberEditParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	var ok MemberEditOK
	var response models.InlineResponse20017
	var msg string
	var code int64
	//var loginRet models.LoginRet
	//var msg string
	//var code int64
	var status models.Response
	msg = "修改成功"

	if (Params.Val == nil){
		status.UnmarshalBinary([]byte(_var.Response200(401,"missing encrypt str")))
		response.Status = &status
		ok.SetPayload(&response)
		o.Context.Respond(rw, r, route.Produces, route, ok)

		return
	}

	db,err := _var.OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	//query
	fmt.Println("Params.MemberID=",Params.MemberID)
	//判断验证码是否正确
	var member models.Member
	db.Table("members").Where("id=?",Params.MemberID).Find(&member)
	if (member.ID == 0){
		status.UnmarshalBinary([]byte(_var.Response200(402,"无此用户")))
		response.Status = &status
		ok.SetPayload(&response)
		o.Context.Respond(rw, r, route.Produces, route, ok)

		return
	}

	//判断member里是有已有phone对应的记录，没有的话插入，有的话更新
	var filename string
	var birth string
	var grade string

	filename = strconv.FormatInt((time.Now().Unix()),10)

	db.Table("members").Where("phone=?",Params.PhoneNumber).Find(&member)

	if(Params.Address!=nil) {
		member.Area = *Params.Address
	}

	if(Params.Membername!=nil) {
		member.Name = *Params.Membername
	}
	if(Params.Gender!=nil){
		member.Gender = *Params.Gender
	}

	member.Ts = time.Now().Unix()
	if (Params.BirthYear != nil) {
		birth = strconv.FormatInt(*(Params.BirthYear),10) + "-" + strconv.FormatInt(*(Params.BirthMonth),10) + "-" + strconv.FormatInt(*(Params.BirthDay),10)
		member.Birth = birth
	}
	if (Params.Grade != nil) {
		grade = *(Params.Grade)
		member.Grade = grade
	}

	if(member.ID==int64(0)){
		code = 203
		msg = "用户不存在"
	}else{
		code = 200
		msg = "修改成功"
		has,_ := HasEditAvatar(Params,filename)
		if has == true{
			member.Avatar = _var.GetResourceDomain("avatar")+filename+".jpg"
		}
		db.Save(&member)
	}

	status.UnmarshalBinary([]byte(_var.Response200(code,msg)))
	response.Status = &status

	ok.SetPayload(&response)

	o.Context.Respond(rw, r, route.Produces, route, ok)

}

func HasEditAvatar(Params  NrMemberEditParams , filename string)  (r bool,saveFileName string) {
	//如果有cover
	if (Params.Avatar != nil) {
		avatar, err := ioutil.ReadAll(Params.Avatar)
		if err != nil {
			fmt.Println("err upload:", err.Error())
		}
		contentType := http.DetectContentType(avatar)
		var lower string
		lower = strings.ToLower(contentType)
		if (strings.Contains(lower, "jp") || (strings.Contains(lower, "pn"))) {
			if (runtime.GOOS == "windows") {
				err1 := ioutil.WriteFile(filename+".jpg", avatar, 0644)
				if err1 != nil {
					fmt.Println(err1.Error())
				}
			} else {
				err1 := ioutil.WriteFile("/root/go/src/resource/image/avatar/"+filename+".jpg", avatar, 0644)
				if err1 != nil {
					fmt.Println(err1.Error())
				}
			}
			return true, filename + ".jpg"
		} else {
			return false, ""
		}

	} else {
		return false, ""
	}
}
