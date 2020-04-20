module main

go 1.13

require (
	github.com/gin-gonic/gin v1.6.2
	github.com/lestrrat-go/file-rotatelogs v2.3.0+incompatible
	github.com/src/controller v0.0.0-incompatible
	github.com/src/functions v0.0.0-incompatible
	github.com/src/models v0.0.0-incompatible
	github.com/src/mylog v0.0.0-incompatible
)

replace github.com/src/models => ./src/models

replace github.com/src/controller => ./src/controller

replace github.com/src/functions => ./src/functions

replace github.com/src/mylog => ./src/log
