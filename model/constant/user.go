package constant

import "github.com/securemist/douyin-mini/model/resp"

var USERNAME_EXISTED = resp.Response{102, "用户名已存在"}

var USER_REGISTER_SUCCESS = resp.Response{0, "用户注册成功"}

var USER_LOGIN_FAILED = resp.Response{103, "用户名或密码错误"}

var GENERAL_ERROR = resp.Response{100, "系统错误"}

var GENERAL_SUCCESS = resp.Response{0, "请求成功"}

var FILE_UPLOAD_SUCCESS = resp.Response{0, "文件上传成功"}

var FILE_UPLOAD_FAILED = resp.Response{104, "文件上传失败"}

var REQUEST_PARAM_ERROR = resp.Response{101, "请求参数异常"}
