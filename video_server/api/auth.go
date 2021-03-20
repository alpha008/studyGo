package main

import (
	"log"
	"net/http"
	"video_server/api/defs"
	"video_server/api/session"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"

func validateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FIELD_SESSION) // session什么时候起作用
	log.Println("validateUserSession sid:", sid)
	if len(sid) == 0 {
		return false
	}

	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}

	//如果没有过期,那么就把用户名加入到HEADER_FIELD_UNAME里面
	r.Header.Add(HEADER_FIELD_UNAME, uname)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0 {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		log.Println("ErrorNotAuthUser")
		return false
	}

	return true
}
