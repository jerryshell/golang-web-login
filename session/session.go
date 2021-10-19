package session

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
	"net/http"
)

// Session 会话
type Session struct {
	attrMap map[string]interface{}
}

// SetAttr 设置属性
func (session *Session) SetAttr(name string, value interface{}) {
	session.attrMap[name] = value
}

// GetAttr 获得属性
func (session *Session) GetAttr(name string) (value interface{}, exist bool) {
	value = session.attrMap[name]
	if value != nil {
		exist = true
	} else {
		exist = false
	}
	return
}

// DelAttr 删除属性
func (session *Session) DelAttr(name string) {
	delete(session.attrMap, name)
}

var sessionMap = make(map[string]*Session)

// GetSession 从 http.Request 中获取 Session
func GetSession(w http.ResponseWriter, r *http.Request) (session *Session) {
	cookie, _ := r.Cookie("SESSION_ID")
	session = sessionMap[cookie.String()]
	if session == nil {
		// session 不存在，创建新的 Session 和 Cookie
		cookie = &http.Cookie{Name: "SESSION_ID", Value: generateSessionID()}
		http.SetCookie(w, cookie)
		session = &Session{attrMap: make(map[string]interface{})}
		sessionMap[cookie.String()] = session
	}
	return
}

func generateSessionID() (sessionID string) {
	buf := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, buf)
	checkError(err)

	sessionID = base64.URLEncoding.EncodeToString(buf)
	return
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
