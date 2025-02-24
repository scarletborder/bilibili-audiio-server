package ctx

import (
	"strings"

	bilibili "github.com/CuteReimu/bilibili/v2"
)

type SrvCtx struct {
	Version string

	B23_client *bilibili.Client
}

func NewSrvCtx() *SrvCtx {
	b23_client := bilibili.NewAnonymousClient()

	return &SrvCtx{
		B23_client: b23_client,
	}
}

func (ctx *SrvCtx) SetBiliCookies(cookies []interface{}) {
	var cookieString []string
	for _, cookie := range cookies {
		if cookie, ok := cookie.(map[string]interface{}); ok {
			if raw, ok := cookie["Raw"].(string); ok {
				cookieString = append(cookieString, raw)
			}
		}

		// 将 cookie 字符串数组连接为一个单独的字符串
		finalCookieString := strings.Join(cookieString, "\n")
		ctx.B23_client.SetCookiesString(finalCookieString)
	}
}
