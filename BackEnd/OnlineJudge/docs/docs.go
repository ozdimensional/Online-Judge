// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/sort": {
            "post": {
                "description": "获取所有用户并按分数降序排序",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "根据用户分数降序排序",
                "parameters": [
                    {
                        "type": "string",
                        "description": "执行操作的用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功获取并排序用户列表",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseSort"
                        }
                    },
                    "403": {
                        "description": "Token 已超时",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseMsg"
                        }
                    },
                    "500": {
                        "description": "数据库查询错误",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseError"
                        }
                    }
                }
            }
        },
        "/changeScore": {
            "post": {
                "description": "通过用户名和新分数的表单数据，增加用户的总分数，并返回新的总分数",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "增加用户分数",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "要增加的分数",
                        "name": "newscore",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "增加分数成功，返回消息和新的总分数",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseAddScore"
                        }
                    },
                    "403": {
                        "description": "Token 已超时",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "数据库查询或保存出错",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseError"
                        }
                    }
                }
            }
        },
        "/check": {
            "post": {
                "description": "通过用户名的表单数据，检查该用户名是否已存在",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "检查用户名是否存在",
                "parameters": [
                    {
                        "type": "string",
                        "description": "要检查的用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用户名不存在",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseMsg"
                        }
                    },
                    "403": {
                        "description": "Token 已超时或用户名已存在",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseMsg"
                        }
                    }
                }
            }
        },
        "/get": {
            "post": {
                "description": "通过用户名的表单数据，获取该用户的昵称",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "获取用户昵称",
                "parameters": [
                    {
                        "type": "string",
                        "description": "要获取昵称的用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功获取用户昵称",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseMsg"
                        }
                    },
                    "403": {
                        "description": "Token 已超时",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseMsg"
                        }
                    },
                    "404": {
                        "description": "数据库查询出错或用户不存在",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseError"
                        }
                    }
                }
            }
        },
        "/judge": {
            "post": {
                "description": "通过用户身份、问题ID和代码内容，进行代码评测，并返回评测结果",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "代码评测",
                "parameters": [
                    {
                        "type": "string",
                        "description": "执行操作的用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "要评测的问题ID",
                        "name": "problem",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "要评测的代码内容",
                        "name": "code",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功进行代码评测",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseMsg"
                        }
                    },
                    "403": {
                        "description": "Token 已超时",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseMsg"
                        }
                    }
                }
            }
        },
        "/list": {
            "post": {
                "description": "获取所有问题的列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "获取问题列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "要获取问题列表的用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功获取问题列表",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseProblems"
                        }
                    },
                    "403": {
                        "description": "Token 已超时",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseMsg"
                        }
                    },
                    "500": {
                        "description": "获取问题列表出错",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseMsg"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "用户使用用户名和密码进行登录，成功返回Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功登录，返回Token",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseToken"
                        }
                    },
                    "403": {
                        "description": "用户名不存在或密码错误",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseMsg"
                        }
                    }
                }
            }
        },
        "/open": {
            "get": {
                "description": "通过表单数据中的用户名检查令牌的有效性，返回检查结果",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "检查令牌有效性",
                "parameters": [
                    {
                        "type": "string",
                        "description": "要检查令牌的用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "令牌有效，返回用户名",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseUsername"
                        }
                    },
                    "403": {
                        "description": "无效令牌",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseMsg"
                        }
                    }
                }
            }
        },
        "/problem/:id": {
            "post": {
                "description": "通过问题ID和管理员身份，获取指定问题的详细信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "获取问题详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "执行操作的管理员用户名",
                        "name": "_",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "要获取的问题ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功获取问题详情",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseQuestionDetail"
                        }
                    },
                    "403": {
                        "description": "获取问题详情出错",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseMsg"
                        }
                    }
                }
            }
        },
        "/problem/add": {
            "post": {
                "description": "通过管理员身份，使用表单数据添加新的问题，并返回操作结果",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "添加问题",
                "parameters": [
                    {
                        "type": "string",
                        "description": "执行操作的管理员用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "问题ID",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "问题标题",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "问题描述",
                        "name": "lore",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "问题标准输入",
                        "name": "input",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "问题标准输出",
                        "name": "output",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "问题提示",
                        "name": "tips",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功添加问题",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseAddProblems"
                        }
                    },
                    "403": {
                        "description": "解析表单数据出错",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseMsg"
                        }
                    }
                }
            }
        },
        "/problem/file/:id": {
            "post": {
                "description": "通过问题ID，获取指定问题的文件列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "获取问题文件列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "要获取文件列表的问题ID",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功获取问题文件列表",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseGetData"
                        }
                    },
                    "403": {
                        "description": "读取文件列表出错",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseDataString"
                        }
                    }
                }
            }
        },
        "/problem/file/add/:id": {
            "post": {
                "description": "通过管理员身份，使用表单数据上传问题的输入和输出文件，并返回操作结果",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "上传问题文件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "执行操作的管理员用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "要上传文件的问题ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "问题输入文件（.in）",
                        "name": "input",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "问题输出文件（.out）",
                        "name": "output",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功上传问题文件",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseMsg"
                        }
                    },
                    "400": {
                        "description": "接收或保存文件出错",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseMsg"
                        }
                    },
                    "403": {
                        "description": "Token 已超时或用户非管理员",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseMsg"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "用户使用用户名、密码和昵称进行注册，成功返回Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "昵称",
                        "name": "nickname",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功注册，返回Token和Token过期时间",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseReg"
                        }
                    },
                    "403": {
                        "description": "用户已存在或注册失败",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseRegErr"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models._ResponseAddProblems": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                }
            }
        },
        "models._ResponseAddScore": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                },
                "new_score": {
                    "type": "integer"
                }
            }
        },
        "models._ResponseDataString": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                }
            }
        },
        "models._ResponseError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "models._ResponseGetData": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "models._ResponseMsg": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                }
            }
        },
        "models._ResponseProblems": {
            "type": "object"
        },
        "models._ResponseQuestionDetail": {
            "type": "object"
        },
        "models._ResponseReg": {
            "type": "object",
            "properties": {
                "Token": {
                    "type": "string"
                },
                "end_time": {
                    "type": "string"
                }
            }
        },
        "models._ResponseRegErr": {
            "type": "object",
            "properties": {
                "err": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "models._ResponseSort": {
            "type": "object"
        },
        "models._ResponseToken": {
            "type": "object",
            "properties": {
                "Token": {
                    "type": "string"
                }
            }
        },
        "models._ResponseUsername": {
            "type": "object",
            "properties": {
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
