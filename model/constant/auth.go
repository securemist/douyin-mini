package constant

import "github.com/securemist/douyin-mini/model/resp"

const Secret = "dsaddsadqdadafasfafada"

var AUTH_SUCCESS = resp.Response{0, "权限验证成功"}

var AUTH_FAILED = resp.Response{300, "权限验证失败"}

var AUTH_EXPIRE = resp.Response{301, "token过期"}
