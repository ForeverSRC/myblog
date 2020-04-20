module controller

go 1.13

require (
	github.com/gin-gonic/gin v1.6.2
	github.com/russross/blackfriday v2.0.0+incompatible
	github.com/src/models v0.0.0-incompatible
	github.com/src/mylog v0.0.0-incompatible
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.29.1 // indirect

)

replace github.com/src/models => ../models

replace github.com/src/functions => ../functions

replace github.com/src/mylog => ../log
