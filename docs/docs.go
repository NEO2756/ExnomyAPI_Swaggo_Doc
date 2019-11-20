// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-11-20 17:22:40.09147 +0530 IST m=+0.079311751

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://www.exnomy.com/terms-and-privacy",
        "contact": {
            "name": "Exnomy Support",
            "url": "https://www.exnomy.com/terms-and-privacy",
            "email": "info@dexhigh.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "https://github.com/blockbankcoltd/Exnomy/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/markets": {
            "get": {
                "description": "Find published markets from db and attach last 24hr market status for each market.\n{\n\"status\": int,    [Status Code, 0 for Success , Greater than 0 for ERROR]\n\"desc\": \"string\", [Description String for Success and Failue]\n\"data\": \"Array Of Objects\",    [Array of Market Object]\n\"markets\": [   [Market Model]\n{\n\"id\": \"string\",  [Market ID, Symbol Pair ex - XRP-BTC]\n\"baseToken\": \"string\", [SYMBOL of the Base Token, ex - XRP]\n\"baseTokenProjectUrl\": \"URI\", [OPTIONAL, URL of the BaseToken Projet, ex - www.ripple.com]\n\"baseTokenName\": \"string\", [NAME of the Base Token, ex - Riple]\n\"baseTokenDecimals\": int, [DECIMAL of the Base Token, ex - 8]\n\"baseTokenAddress\": [ADDRESS of the Base Token],\n\"quoteToken\": \"string\", [SYMBOL of the Quote Token, ex - BTC]\n\"quoteTokenDecimals\": int, [DECIMAL of the Quote Token, ex - 8]\n\"quoteTokenAddress\": string, [ADDRESS of the Base Token]\n\"minOrderSize\": \"string\", [MINIMUM order size for BASE-QUOTE token pair, ex- 0.0001]\n\"pricePrecision\": int, [PRECISION FOR PRICE for BASE-QUOTE token pair, ex- 5]\n\"priceDecimals\": int, [DECIMAL FOR PRICE for BASE-QUOTE token pair, ex- 8]\n\"amountDecimals\": int, [DECIMALS FOR AMOUNT for BASE-QUOTE token pair, ex- 6]\n\"asMakerFeeRate\": \"string\", [MAKER fee rate string for BASE-QUOTE token pair]\n\"asTakerFeeRate\": \"string\", [TAKER fee rate string for BASE-QUOTE token pair]\n\"gasFeeAmount\": \"string\",   [NOT applicable]\n\"supportedOrderTypes\": [   [only LIMIT is supported right now]\n\"limit\",\n\"market\"\n],\n\"marketOrderMaxSlippage\": \"string\", [FUTURE USE, Not applicable right now]\n\"lastPriceIncrease\": \"string\",      [LAST PRICE INCREASE value]\n\"lastPrice\": \"string\",              [Last Traded Price]\n\"price24h\": \"string\",\t\t\t\t[PRICE CHANGE in last 24hrs, ex - +0.3 % OR -0.4 %]\n\"amount24h\": \"string\",\t\t\t\t[LAST 24hr TOTAL AMOUNT CHANGE]\n\"quoteTokenVolume24h\": \"246789\"     [VOLUME of LAST 24Hr Traded QUOTE TOKEN]\n}\n]\n}\n}",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show all Published markets",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.Response"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            }
        },
        "/markets/{marketID}/candles": {
            "get": {
                "description": "Get high, low, open, close and volume data for an interval defined bt from to to param.\n{\n\"status\": int,    [STATUS CODE, 0 for Success , Greater than 0 for ERROR]\n\"desc\": \"string\", [Description String for Success and Failue]\n\"data\": {\n\"candles\": [  [Candle Information]\n{\n\"time\": int,    [Unix timestamp of Candle, ex - 1574232000]\n\"open\": \"string\",  [Open Price in this TimeStamp]\n\"close\": \"string\", [Close Price in this Timestamp]\n\"low\": \"string\",   [Lowest Trade Price in this timestamp]\n\"high\": \"string\",  [Highest timetsamp in this timestamp]\n\"volume\": \"string\" [Total traded volume in this timestamp]\n}\n]\n}\n}",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get Bar data from 'from' to 'to'.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Market ID, ex XRP-BTC",
                        "name": "marketID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "from",
                        "name": "from",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "to",
                        "name": "to",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "granularity",
                        "name": "granularity",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.Bar"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            }
        },
        "/markets/{marketID}/orderbook": {
            "get": {
                "description": "Find the latest snapshot version 2 corresponding to MarketID\n{\n\"status\": int,    [Status Code, 0 for Success , Greater than 0 for ERROR]\n\"desc\": \"string\", [Description String for Success and Failue]\n\"data\": {\n\"orderBook\": {\n\"sequence\": int,  [Database speacific Sequence Number]\n\"bids\": [   \"Array of Objects\", [Array bids Object]\n[\n\"string\", [Bid Price, ex - \"0.00003108\"]\n\"string\", [Bif Amount, ex - \"330\"]\n]\n],\n\"asks\": [ \"Array of Objects\", [Array asks Object]\n[\n\"string\", [Ask price, ex - \"0.00003108\"]\n\"string\", [Ask Amount, ex - \"330\"]\n]\n]\n}\n}\n}",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get OrderBook for a particular market ID.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "MarketID (ex: XRP-BTC)",
                        "name": "marketID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.SnapshotV2"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            }
        },
        "/markets/{marketID}/trades": {
            "get": {
                "description": "Find all the trades for a particular market\n{\n\"status\": int,    [STATUS CODE, 0 for Success , Greater than 0 for ERROR]\n\"desc\": \"string\", [Description String for Success and Failue]\n\"data\": {\n\"count\": int,   [Size of the trades array object]\n\"trades\": [     [Array of Objects]\n{\n\"id\": int,   [DB Speacific Trade ID ]\n\"transactionID\": 36671, [DB Speacific Transaction ID]\n\"transactionHash\": [Hash of Trade Transaction on Blockchain ex - 0xfba0a551880a7a9458fdcd9e55078645da4a11c399be91efdd19a6a82014e731],\n\"status\": \"string\", [Status of the Trade transaction, ex - successful, pending, pending]\n\"marketID\": \"string\", [Market ID, ex - XRP-BTC]\n\"maker\": \"string\", [Maker Account Address ex - 0x31ebd457b999bf99759602f5ece5aa5033cb56b3]\n\"taker\": \"string\", [Taker Account Address ex - 0x31ebd457b999bf99759602f5ece5aa5033cb56b3],\n\"takerSide\": \"string\", [Side of Taker order ex - but OR sell]\n\"makerOrderID\": \"string, [OrderID of Maker, ex - 0xbf9abcaa08bb820f791206564d261e3e36309ac16f5b091e507f0b14fb369656],\n\"takerOrderID\": \"string, [OrderID of Taker, ex - 0xbf9abcaa08bb820f791206564d261e3e36309ac16f5b091e507f0b14fb369656],\n\"sequence\": uint, [DB speacific sequence number]\n\"amount\": \"uint\", [Amount/Quantity of trade, ex - 230]\n\"price\": \"string\", [Price of Trade, ex - 0.00003096]\n\"executedAt\": \"time\", [Time When Trade executed,  ex - 2019-11-20T08:36:26Z]\n\"createdAt\": \"time\",  [Time When Trade created,  ex - 2019-11-20T08:36:26Z]\n\"updatedAt\": \"time\",  [Time When Trade executed,  ex - 2019-11-20T08:36:26Z]\n}\n]\n}\n}",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show all trades for a market",
                "parameters": [
                    {
                        "type": "string",
                        "description": "MarketID (ex XRP-BTC)",
                        "name": "marketID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.QueryTradeResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            }
        },
        "/orders": {
            "get": {
                "description": "Find latest 'N' speacified orders, 20 if 'N' is not speacifies corresponding to particular MarketID and user address.\n\"status\": int,    [STATUS CODE, 0 for Success , Greater than 0 for ERROR]\n\"desc\": \"string\", [Description String for Success and Failue]\n\"data\": {\n\"count\": int, \t\t\t\t\t   [Size of the Order Array]\n\"orders\": [\t\t\t\t\t       [Order Array]\n{\n\"id\": \"string\",                [Order ID of the Requested Order ex - 0x8d9cafdbed805fb495f7e29c9cf7a269e140f326701f388e77d840228134b103]\n\"traderAddress\": \"string\",     [Address of the Trader used for Loging In],\n\"marketID\": \"string\",          [Market ID of the order, ex - XRP-BTC]\n\"side\": \"string\",              [Side of the Order, ex - buy OR sell]\n\"price\": \"string\",             [Price of the Order, ex - 0.00003081]\n\"amount\": \"string\",            [Amount of the Order, ex - 160]\n\"status\": \"string\",            [Status of the order, canceled, pending, partial_filled OR full_filled]\n\"type\": \"string\",              [Order type Limit or Market, Limit is only supported right now]\n\"version\": \"string\",           [verison of the platform]\n\"availableAmount\": \"string\",   [Amount available fto Trade]\n\"confirmedAmount\": \"string\",   [Traded amount, confirmed on blockchain]\n\"canceledAmount\": \"string\",    [canceled Amount]\n\"pendingAmount\": \"string\",     [Left Over Amount for trading, Incase of partial fills]\n\"makerFeeRate\": \"string\",      [makerFee Rate]\n\"takerFeeRate\": \"string\",      [Taker fee Rate]\n\"makerRebateRate\": \"string\",   [makerRebateRate ]\n\"gasFeeAmount\": \"string\",      [Gas fee Amount in terms of speacified tokens, NOT APPLICABLE in EXNOMY]\n\"json\": \"string\",              [josn string of this order]\n\"createdAt\": \"time\", \t\t   [time at which has been created, ex - 2019-11-20T10:07:52.523606Z]\n\"updatedAt\": \t\t\t\t   [time at which has been updated, ex - 2019-11-20T10:07:52.523606Z]\n}\n]\n}\n}",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show orders, reverse-chrono order.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "MarketID ex - XRP-BTC",
                        "name": "marketID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "DEX-AUTH-TOKEN, ex - 0x31ebd457b999bf99759602f5ece5aa5033cb56b3#DEX-EXONOMY-AUTH#0xe99ac942cc4e126e62e1a98c756bf3e3e09de1dc5cbead6b02992063084af399120a45642a8e2876c608d43f11aaa07d7511d338e7d7d4acdc91e1e90c05e3581c",
                        "name": "DEX-EXONOMY-AUTH",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.QueryOrderResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            }
        },
        "/orders/{OrderID}": {
            "delete": {
                "description": "Find the order with ID speacified as OrderID and Mark it cancel.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Cancel particular order.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "OrderID",
                        "name": "OrderID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "DEX-AUTH-TOKEN, ex - 0x31ebd457b999bf99759602f5ece5aa5033cb56b3#DEX-EXONOMY-AUTH#0xe99ac942cc4e126e62e1a98c756bf3e3e09de1dc5cbead6b02992063084af399120a45642a8e2876c608d43f11aaa07d7511d338e7d7d4acdc91e1e90c05e3581c",
                        "name": "DEX-EXONOMY-AUTH",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.CancelOrderReq"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            }
        },
        "/orders/{orderID}": {
            "get": {
                "description": "Get order corresponding to a particular order ID\n{\n\"status\": int,    [STATUS CODE, 0 for Success , Greater than 0 for ERROR]\n\"desc\": \"string\", [Description String for Success and Failue]\n\"data\": {\n\"order\": {\n\"id\": \"string\",                [Order ID of the Requested Order ex - 0x8d9cafdbed805fb495f7e29c9cf7a269e140f326701f388e77d840228134b103]\n\"traderAddress\": \"string\",     [Address of the Trader used for Loging In],\n\"marketID\": \"string\",          [Market ID of the order, ex - XRP-BTC]\n\"side\": \"string\",              [Side of the Order, ex - buy OR sell]\n\"price\": \"string\",             [Price of the Order, ex - 0.00003081]\n\"amount\": \"string\",            [Amount of the Order, ex - 160]\n\"status\": \"string\",            [Status of the order, canceled, pending, partial_filled OR full_filled]\n\"type\": \"string\",              [Order type Limit or Market, Limit is only supported right now]\n\"version\": \"string\",           [verison of the platform]\n\"availableAmount\": \"string\",   [Amount available fto Trade]\n\"confirmedAmount\": \"string\",   [Traded amount, confirmed on blockchain]\n\"canceledAmount\": \"string\",    [canceled Amount]\n\"pendingAmount\": \"string\",     [Left Over Amount for trading, Incase of partial fills]\n\"makerFeeRate\": \"string\",      [makerFee Rate]\n\"takerFeeRate\": \"string\",      [Taker fee Rate]\n\"makerRebateRate\": \"string\",   [makerRebateRate ]\n\"gasFeeAmount\": \"string\",      [Gas fee Amount in terms of speacified tokens, NOT APPLICABLE in EXNOMY]\n\"json\": \"string\",              [josn string of this order]\n\"createdAt\": \"time\", \t\t   [time at which has been created, ex - 2019-11-20T10:07:52.523606Z]\n\"updatedAt\": \t\t\t\t   [time at which has been updated, ex - 2019-11-20T10:07:52.523606Z]\n}\n}\n}",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get Single order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "OrderID",
                        "name": "orderID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "DEX-AUTH-TOKEN, ex - 0x31ebd457b999bf99759602f5ece5aa5033cb56b3#DEX-EXONOMY-AUTH#0xe99ac942cc4e126e62e1a98c756bf3e3e09de1dc5cbead6b02992063084af399120a45642a8e2876c608d43f11aaa07d7511d338e7d7d4acdc91e1e90c05e3581c",
                        "name": "DEX-EXONOMY-AUTH",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.QuerySingleOrderResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.ApiError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "desc": {
                    "type": "string"
                }
            }
        },
        "api.Bar": {
            "type": "object",
            "properties": {
                "close": {
                    "type": "number"
                },
                "high": {
                    "type": "number"
                },
                "low": {
                    "type": "number"
                },
                "open": {
                    "type": "number"
                },
                "time": {
                    "type": "integer"
                },
                "volume": {
                    "type": "number"
                }
            }
        },
        "api.CancelOrderReq": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "address": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "api.QueryOrderResp": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "orders": {
                    "type": "string"
                }
            }
        },
        "api.QuerySingleOrderResp": {
            "type": "object",
            "properties": {
                "order": {
                    "type": "string"
                }
            }
        },
        "api.QueryTradeResp": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "trades": {
                    "type": "string"
                }
            }
        },
        "api.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "desc": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "api.SnapshotV2": {
            "type": "object",
            "properties": {
                "asks": {
                    "type": "array",
                    "items": {
                        "type": "\u0026{%!s(token.Pos=365) %!s(*ast.BasicLit=\u0026{366 5 2}) string}"
                    }
                },
                "bids": {
                    "type": "array",
                    "items": {
                        "type": "\u0026{%!s(token.Pos=329) %!s(*ast.BasicLit=\u0026{330 5 2}) string}"
                    }
                },
                "sequence": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ExnomyAuthToken": {
            "type": "apiKey",
            "name": "Authorization (DEX-EXONOMY-AUTH should be like {address}#DEX-EXONOMY-AUTH@{time}#{signature})",
            "in": "header"
        }
    },
    "x-extension-openapi": {
        "example": "value on a json format"
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
	Version:     "1.0",
	Host:        "api.exnomy.com",
	BasePath:    "",
	Schemes:     []string{"https"},
	Title:       "Exnomy API Swagger Documentation",
	Description: "This is Exnomy API Server.",
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
