package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// 编写单元测试：基于对http请求的单元测试，一般为了避免形成循环依赖（循环引用），直接自己写然后测试

func TestCreatePostHandler(t *testing.T){
	// 设置gin框架的模式为测试模式
	gin.SetMode(gin.TestMode)
	// 创建一个gin.Engine实例
	r := gin.Default()
	// 创建一个路由
	r.POST("/api/v1/post", CreatePostHandler)

	// 创建一个请求
	reqBody := strings.NewReader(
		`{
			"title":"test title",
			"content":"test content",
			"community_id":2
		}`,
	)
	// 创建一个请求对象: http包调用新建请求函数，参数：请求的方法、请求路径、请求体
	request, _ := http.NewRequest(http.MethodPost, "/api/v1/post", reqBody)
	
	// 创建一个响应对象： http包新建一个响应对象，其实就是用来“记录”http请求响应的
	response := httptest.NewRecorder()
	// 发送并接收请求
	r.ServeHTTP(response, request)

	// 使用断言函数，判断 状态码、响应头、响应体等是否正确: 
	// 参数： 测试参数、希望值、实际值
	assert.Equal(t, http.StatusOK, response.Code)// 响应体

	// 方法一：判断响应内容中是否包含某字符串
	//assert.Contains(t, response.Body.String(), "需要登录")

	// 方法二：
	res := new(ResponseData)
	err := json.Unmarshal(response.Body.Bytes(), res)
	if err != nil {
		t.Fatal(err)
	}
	// 参数： 测试参数、希望值、实际值
	assert.Equal(t, CodeNeedLogin, res.Code)
}