// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag/v2"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},"swagger":"2.0","info":{"description":"{{escape .Description}}","title":"{{.Title}}","termsOfService":"http://swagger.io/terms/","contact":{"name":"API Support","url":"http://www.swagger.io/support","email":"support@swagger.io"},"license":{"name":"Apache 2.0","url":"http://www.apache.org/licenses/LICENSE-2.0.html"},"version":"{{.Version}}"},"host":"{{.Host}}","basePath":"{{.BasePath}}","paths":{"/authentication/user":{"post":{"description":"Registers a user","consumes":["application/json"],"produces":["application/json"],"tags":["authentication"],"summary":"Registers a user","parameters":[{"description":"User credentials","name":"payload","in":"body","required":true,"schema":{"$ref":"#/definitions/main.RegisterUserPayload"}}],"responses":{"201":{"description":"User registered","schema":{"$ref":"#/definitions/main.UserWithToken"}},"400":{"description":"Bad Request","schema":{}},"500":{"description":"Internal Server Error","schema":{}}}}},"/users/{id}":{"get":{"security":[{"ApiKeyAuth":[]}],"description":"Fetches a user profile by ID","consumes":["application/json"],"produces":["application/json"],"tags":["users"],"summary":"Fetches a user profile","parameters":[{"type":"integer","description":"User ID","name":"id","in":"path","required":true}],"responses":{"200":{"description":"OK","schema":{"$ref":"#/definitions/store.User"}},"400":{"description":"Bad Request","schema":{}},"404":{"description":"Not Found","schema":{}},"500":{"description":"Internal Server Error","schema":{}}}}},"/users/{userID}/follow":{"put":{"security":[{"ApiKeyAuth":[]}],"description":"Follows a user by ID","consumes":["application/json"],"produces":["application/json"],"tags":["users"],"summary":"Follows a user","parameters":[{"type":"integer","description":"User ID","name":"userID","in":"path","required":true}],"responses":{"204":{"description":"User followed","schema":{"type":"string"}},"400":{"description":"User payload missing","schema":{}},"404":{"description":"User not found","schema":{}}}}}},"definitions":{"main.RegisterUserPayload":{"type":"object","required":["email","password","username"],"properties":{"email":{"type":"string","maxLength":255},"password":{"type":"string","maxLength":72,"minLength":3},"username":{"type":"string","maxLength":100}}},"main.UserWithToken":{"type":"object","properties":{"created_at":{"type":"string"},"email":{"type":"string"},"id":{"type":"integer"},"is_active":{"type":"boolean"},"token":{"type":"string"},"username":{"type":"string"}}},"store.User":{"type":"object","properties":{"created_at":{"type":"string"},"email":{"type":"string"},"id":{"type":"integer"},"is_active":{"type":"boolean"},"username":{"type":"string"}}}},"securityDefinitions":{"ApiKeyAuth":{"type":"apiKey","name":"Authorization","in":"header"}}}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "GopherSocial API",
	Description:      "API for GopherSocial, a social network for gohpers",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
