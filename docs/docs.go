// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "lion",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/community": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "查询社区列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "社区相关接口"
                ],
                "summary": "查询社区列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseCommunityList"
                        }
                    }
                }
            }
        },
        "/community/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "查询社区详情列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "社区相关接口"
                ],
                "summary": "查询社区详情列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "社区ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseCommunityDetail"
                        }
                    }
                }
            }
        },
        "/post/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "单个帖子详情接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子相关接口"
                ],
                "summary": "单个帖子详情接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "帖子id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponsePostDetail"
                        }
                    }
                }
            }
        },
        "/post2": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "可按社区按时间或分数排序查询帖子列表接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子相关接口"
                ],
                "summary": "升级版帖子列表接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "可以为空,如果为空，则查询所有帖子，否则查询相应社区的帖子",
                        "name": "community_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "score",
                        "description": "排序依据",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 10,
                        "description": "每页数据量",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponsePostList"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.ResCode": {
            "type": "integer",
            "enum": [
                1000,
                1001,
                1002,
                1003,
                1004,
                1005,
                1006,
                1007
            ],
            "x-enum-varnames": [
                "CodeSuccess",
                "CodeInvalidParam",
                "CodeUserExist",
                "CodeUserNotExist",
                "CodeInvalidPassword",
                "CodeServerBusy",
                "CodeNeedLogin",
                "CodeInvalidToken"
            ]
        },
        "controller._ResponseCommunityDetail": {
            "type": "object",
            "properties": {
                "code": {
                    "$ref": "#/definitions/controller.ResCode"
                },
                "data": {
                    "$ref": "#/definitions/models.CommunityDetail"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "controller._ResponseCommunityList": {
            "type": "object",
            "properties": {
                "code": {
                    "$ref": "#/definitions/controller.ResCode"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Community"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "controller._ResponsePostDetail": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务响应状态码",
                    "allOf": [
                        {
                            "$ref": "#/definitions/controller.ResCode"
                        }
                    ]
                },
                "data": {
                    "description": "数据",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.ApiPostDetail"
                        }
                    ]
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "controller._ResponsePostList": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务响应状态码",
                    "allOf": [
                        {
                            "$ref": "#/definitions/controller.ResCode"
                        }
                    ]
                },
                "data": {
                    "description": "数据",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ApiPostDetail"
                    }
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "models.ApiPostDetail": {
            "type": "object",
            "required": [
                "community_id",
                "content",
                "title"
            ],
            "properties": {
                "author_id": {
                    "description": "TODO:在请求中没有问题，但是会出现 响应类型不符的问题：在于json中的“string”",
                    "type": "integer"
                },
                "author_name": {
                    "description": "发帖作者的名称",
                    "type": "string"
                },
                "community_id": {
                    "type": "integer"
                },
                "content": {
                    "type": "string"
                },
                "create_time": {
                    "description": "这里如果要使用 time 类型的时间，注意连接库的时候，需要设置参数 parseTime=true",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "introduction": {
                    "description": "omitempty 是说如果这个字段为空，那么就不返回这个字段",
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "vote_num": {
                    "type": "integer"
                }
            }
        },
        "models.Community": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.CommunityDetail": {
            "type": "object",
            "properties": {
                "create_time": {
                    "description": "这里如果要使用 time 类型的时间，注意连接库的时候，需要设置参数 parseTime=true",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "introduction": {
                    "description": "omitempty 是说如果这个字段为空，那么就不返回这个字段",
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "这里写标题",
	Description:      "这里写描述信息",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
