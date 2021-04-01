// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/?Action=GetCMDBInstance": {
            "post": {
                "description": "获取数据实例",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Instance"
                ],
                "summary": "获取数据实例",
                "parameters": [
                    {
                        "description": "body",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.GetInstanceRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    }
                }
            }
        },
        "/?Action=UpdateCMDBInstance": {
            "post": {
                "description": "修改数据实例，只根据uuid进行更新",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Instance"
                ],
                "summary": "修改数据实例",
                "parameters": [
                    {
                        "type": "string",
                        "description": "元数据表名",
                        "name": "category",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "需要修改的数据实例",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.UpdateInstanceBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    }
                }
            }
        },
        "/category/create": {
            "post": {
                "description": "创建模型",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "创建模型",
                "parameters": [
                    {
                        "description": "创建模型",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysCategory"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    }
                }
            }
        },
        "/category/delete": {
            "delete": {
                "description": "删除模型",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "删除模型",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "表模型ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    }
                }
            }
        },
        "/category/get": {
            "get": {
                "description": "获取模型",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "获取模型",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "表模型ID, 暂不支持多个查询",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "关键词",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序, 支持desc和asc",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "分页条数",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "分页偏移量",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    }
                }
            }
        },
        "/field/create": {
            "post": {
                "description": "创建字段",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Field"
                ],
                "summary": "创建字段",
                "parameters": [
                    {
                        "description": "创建字段",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CreateFieldBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    }
                }
            }
        },
        "/field/get": {
            "get": {
                "description": "获取字段",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Field"
                ],
                "summary": "获取字段",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "字段元数据ID, 暂不支持多个查询",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "关键词",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序, 支持desc和asc",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "分页条数",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "分页偏移量",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    }
                }
            }
        },
        "/field/update": {
            "put": {
                "description": "修改字段，只能修改字段的业务属性",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Field"
                ],
                "summary": "修改字段",
                "parameters": [
                    {
                        "description": "修改字段结构体",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.UpdateFieldBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    }
                }
            }
        },
        "/field/{id}": {
            "delete": {
                "description": "删除字段",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Field"
                ],
                "summary": "删除字段",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "字段元数据ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    }
                }
            }
        },
        "/fieldType/": {
            "get": {
                "description": "获取字段",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "FieldType"
                ],
                "summary": "获取字段",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.GetFieldTypeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    }
                }
            }
        },
        "/instance/{category}": {
            "post": {
                "description": "创建实例",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Instance"
                ],
                "summary": "创建实例",
                "parameters": [
                    {
                        "type": "string",
                        "description": "元数据表名",
                        "name": "category",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "创建实例",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CreateInstanceBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除数据（软删除，如果有唯一约束，在重新插入时会报错）",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Instance"
                ],
                "summary": "删除数据",
                "parameters": [
                    {
                        "type": "string",
                        "description": "元数据表名",
                        "name": "category",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "数据实例的uuid，多个值以半角逗号分隔",
                        "name": "uuid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    }
                }
            }
        },
        "/system/": {
            "get": {
                "description": "init",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System"
                ],
                "summary": "init",
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    }
                }
            }
        },
        "/system/migrate": {
            "put": {
                "description": "迁移",
                "tags": [
                    "System"
                ],
                "summary": "migrate",
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/myhttp.Responds"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.CreateFieldBody": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SysField"
                    }
                }
            }
        },
        "controller.CreateInstanceBody": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "additionalProperties": true
                    }
                }
            }
        },
        "controller.FieldTypeForGet": {
            "type": "object",
            "properties": {
                "can_unique": {
                    "description": "字段是否可以设置唯一，与isUnique一起用于判断是否要创建唯一索引",
                    "type": "boolean"
                },
                "cmdb_name": {
                    "description": "内置CMDB字段名(仅用于展示)",
                    "type": "string"
                },
                "index": {
                    "description": "cmdb内置字段序号",
                    "type": "integer"
                },
                "is_string": {
                    "description": "是否以字符串形式设置默认值",
                    "type": "boolean"
                },
                "sql_type": {
                    "description": "对应mysql中的字段类型",
                    "type": "string"
                }
            }
        },
        "controller.GetFieldTypeResponse": {
            "type": "object",
            "properties": {
                "Data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controller.FieldTypeForGet"
                    }
                },
                "Meta": {
                    "$ref": "#/definitions/myhttp.Meta"
                },
                "Total": {
                    "type": "integer"
                }
            }
        },
        "controller.GetInstanceRequestBody": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "id_list": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "keyword": {
                    "type": "string"
                },
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "order": {
                    "type": "string"
                },
                "query": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                }
            }
        },
        "controller.UpdateFieldBody": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_editable": {
                    "type": "boolean"
                },
                "is_required": {
                    "type": "boolean"
                },
                "label": {
                    "type": "string"
                }
            }
        },
        "controller.UpdateInstanceBody": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "additionalProperties": true
                    }
                }
            }
        },
        "model.SysCategory": {
            "type": "object",
            "properties": {
                "create_by": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "description": "模型描述",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "label": {
                    "description": "模型名(仅用于展示)",
                    "type": "string"
                },
                "name": {
                    "description": "模型名(建表时用)",
                    "type": "string"
                },
                "update_at": {
                    "type": "string"
                },
                "update_by": {
                    "type": "string"
                }
            }
        },
        "model.SysField": {
            "type": "object",
            "properties": {
                "default": {
                    "description": "默认值",
                    "type": "string"
                },
                "description": {
                    "description": "字段描述",
                    "type": "string"
                },
                "is_editable": {
                    "description": "字段是否可编辑(会在实例更新时对不可编辑的字段进行过滤)",
                    "type": "boolean"
                },
                "is_required": {
                    "description": "是否必填(会在创建实例时进行校验)",
                    "type": "boolean"
                },
                "label": {
                    "description": "字段名(仅用于展示)",
                    "type": "string"
                },
                "name": {
                    "description": "字段名(建表时用)",
                    "type": "string"
                },
                "not_null": {
                    "description": "是否不可为NULL",
                    "type": "boolean"
                },
                "type": {
                    "description": "字段对应的CMDB内置字段类型序号",
                    "type": "integer"
                }
            }
        },
        "myhttp.Meta": {
            "type": "object",
            "properties": {
                "Code": {
                    "type": "integer"
                },
                "Message": {
                    "type": "string"
                }
            }
        },
        "myhttp.Responds": {
            "type": "object",
            "properties": {
                "Data": {
                    "type": "object"
                },
                "Meta": {
                    "$ref": "#/definitions/myhttp.Meta"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.0.1",
	Host:        "",
	BasePath:    "/v1/",
	Schemes:     []string{},
	Title:       "CMDB-API文档",
	Description: "CMDB.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}