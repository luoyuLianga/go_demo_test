package controller

import "github.com/gin-gonic/gin"

type JSONStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}

func ReturnSuccess(c *gin.Context, code int, msg interface{}, data interface{}, count int64) {
	json := &JSONStruct{Code: code, Msg: msg, Data: data, Count: count}
	c.JSON(200, json)
}

func ReturnError(c *gin.Context, code int, msg interface{}) {
	json := &JSONStruct{Code: code, Msg: msg}
	c.JSON(200, json)
}
