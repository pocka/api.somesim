package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "# 概要\nMoE用染色シミュレータ「そめしむ」のバックエンドAPI.\n",
    "title": "Somesim",
    "version": "1.0.0"
  },
  "basePath": "/v1",
  "paths": {
    "/docs": {
      "get": {
        "description": "ReDocによるAPIドキュメントを取得する(HTML)\n",
        "produces": [
          "text/html"
        ],
        "tags": [
          "InternalAPI"
        ],
        "responses": {
          "200": {
            "description": "APIドキュメント(ReDoc)\n",
            "schema": {
              "type": "string"
            }
          },
          "404": {
            "description": "ファイルが存在しない\n"
          }
        }
      }
    },
    "/docs/swagger.yml": {
      "get": {
        "description": "Swaggerによるスキーマを取得する(YAML)\n",
        "produces": [
          "text/plain"
        ],
        "tags": [
          "InternalAPI"
        ],
        "responses": {
          "200": {
            "description": "スキーマ(Swagger)\n",
            "schema": {
              "type": "string"
            }
          },
          "404": {
            "description": "ファイルが存在しない\n"
          }
        }
      }
    },
    "/flowers": {
      "get": {
        "description": "花びら定義を取得する\n\n花びらのカテゴリによる絞り込みが可能(e.g. ラーファン, 課金花びら)\n",
        "tags": [
          "Flower"
        ],
        "summary": "花びら定義の取得",
        "parameters": [
          {
            "type": "string",
            "description": "花びらのカテゴリ",
            "name": "category",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "花びら定義",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Flower"
              }
            }
          }
        }
      },
      "post": {
        "description": "花びらの定義を追加する\n指定されたカテゴリが存在しなければカテゴリを新規作成する\n",
        "tags": [
          "Flower"
        ],
        "summary": "花びら定義の追加",
        "parameters": [
          {
            "description": "追加する花びら定義",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Flower"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "追加に成功"
          },
          "409": {
            "description": "指定された名前の花びらが既に存在する"
          }
        }
      }
    },
    "/flowers/{name}": {
      "get": {
        "description": "名前を指定して花びら定義を取得する\n",
        "tags": [
          "Flower"
        ],
        "summary": "指定された名前の花びら定義を取得",
        "parameters": [
          {
            "type": "string",
            "description": "花びらの名前",
            "name": "name",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "花びら定義",
            "schema": {
              "$ref": "#/definitions/Flower"
            }
          },
          "404": {
            "description": "指定された花びら定義が存在しない"
          }
        }
      },
      "put": {
        "description": "指定した名前の花びら定義を更新する\n",
        "tags": [
          "Flower"
        ],
        "summary": "花びら定義を更新",
        "parameters": [
          {
            "type": "string",
            "description": "更新したい花びらの名前",
            "name": "name",
            "in": "path",
            "required": true
          },
          {
            "description": "上書きする内容",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Flower"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "更新に成功"
          },
          "404": {
            "description": "指定された名前の花びらが存在しない"
          }
        }
      }
    },
    "/items": {
      "get": {
        "description": "装備情報を取得する\n\nカテゴリによる絞り込みが可能(e.g. グラフィック装備, 生産装備(裁縫))\n",
        "tags": [
          "Item"
        ],
        "summary": "装備情報を取得",
        "parameters": [
          {
            "type": "string",
            "description": "装備のカテゴリ",
            "name": "category",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "装備情報",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Item"
              }
            }
          }
        }
      },
      "post": {
        "description": "装備情報を新たに追加する\n",
        "tags": [
          "Item"
        ],
        "summary": "装備情報を追加",
        "parameters": [
          {
            "description": "追加したい装備情報",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Item"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "作成に成功"
          },
          "409": {
            "description": "同じ名前の装備情報が既に存在する"
          }
        }
      }
    },
    "/items/{name}": {
      "get": {
        "description": "名前を指定して装備情報を取得する\n",
        "tags": [
          "Item"
        ],
        "summary": "指定した名前の装備情報を取得",
        "parameters": [
          {
            "type": "string",
            "description": "取得したい装備情報の名前",
            "name": "name",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "指定した名前の装備情報",
            "schema": {
              "$ref": "#/definitions/Item"
            }
          },
          "404": {
            "description": "指定された名前の装備情報が存在しない"
          }
        }
      },
      "put": {
        "description": "指定された名前の装備情報を更新する\n",
        "tags": [
          "Item"
        ],
        "summary": "装備情報を更新",
        "parameters": [
          {
            "type": "string",
            "description": "更新したい装備の名前",
            "name": "name",
            "in": "path",
            "required": true
          },
          {
            "description": "更新したい内容",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Item"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "更新に成功した"
          },
          "404": {
            "description": "指定された名前の装備情報が存在しない"
          }
        }
      }
    },
    "/items/{name}/images": {
      "get": {
        "description": "装備名を指定して染色用画像を取得する\n",
        "tags": [
          "Item"
        ],
        "summary": "染色用画像を取得",
        "parameters": [
          {
            "type": "string",
            "description": "装備名",
            "name": "name",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "その装備に登録されている全ての染色用画像セット\n",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/ImageSet"
              }
            }
          }
        }
      },
      "post": {
        "description": "指定された名前の装備に新しい装備画像セットを追加\n\ne.g. ハイキャスター装備にNFの画像セットを追加\n",
        "tags": [
          "Item"
        ],
        "summary": "染色用画像を追加",
        "parameters": [
          {
            "type": "string",
            "description": "装備名",
            "name": "name",
            "in": "path",
            "required": true
          },
          {
            "description": "装備画像セット",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ImageSet"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "追加に成功した"
          },
          "409": {
            "description": "指定された名前の装備画像セットが既に存在する"
          }
        }
      }
    },
    "/items/{name}/images/{image_name}": {
      "get": {
        "description": "指定された名前の装備の画像セットの中から、\nさらに名前を指定して画像セットを取得する\n",
        "tags": [
          "Item"
        ],
        "summary": "名前を指定して染色用画像を取得",
        "parameters": [
          {
            "type": "string",
            "description": "装備の名前",
            "name": "name",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "画像セットの名前",
            "name": "image_name",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "画像セット\n",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/ImageSet"
              }
            }
          },
          "404": {
            "description": "指定された名前の画像セットが存在しない"
          }
        }
      },
      "put": {
        "description": "染色用画像を更新する\n",
        "tags": [
          "Item"
        ],
        "summary": "染色用画像を更新",
        "parameters": [
          {
            "type": "string",
            "description": "更新する染色用画像がある装備名",
            "name": "name",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "更新したい染色用画像の名前",
            "name": "image_name",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "更新に成功"
          },
          "404": {
            "description": "指定された名前の染色用画像が存在しない"
          }
        }
      }
    }
  },
  "definitions": {
    "Flower": {
      "description": "花びら定義",
      "type": "object",
      "properties": {
        "blue": {
          "description": "青色の値(%)",
          "type": "integer"
        },
        "category": {
          "description": "花びらの所属するカテゴリ",
          "type": "string"
        },
        "green": {
          "description": "緑色の値(%)",
          "type": "integer"
        },
        "name": {
          "description": "花びらの名前",
          "type": "string"
        },
        "red": {
          "description": "赤色の値(%)",
          "type": "integer"
        }
      }
    },
    "ImageSet": {
      "description": "装備画像セット(ベース画像とマスク画像のペア)",
      "type": "object",
      "properties": {
        "base_image": {
          "description": "Base64エンコードされたベース画像",
          "type": "string"
        },
        "mask_image": {
          "description": "Base64エンコードされたマスク画像",
          "type": "string"
        },
        "name": {
          "description": "装備画像セットの名前",
          "type": "string"
        }
      }
    },
    "Item": {
      "description": "装備情報",
      "type": "object",
      "properties": {
        "images": {
          "description": "装備画像セット",
          "type": "array",
          "items": {
            "$ref": "#/definitions/ImageSet"
          }
        },
        "name": {
          "description": "装備の名前",
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "花びらの色情報を取得・管理する",
      "name": "Flower"
    },
    {
      "description": "染色対象の装備の情報や画像を取得・管理する",
      "name": "Item"
    },
    {
      "description": "スキーマやドキュメントを参照する",
      "name": "InternalAPI"
    }
  ]
}`))
}
