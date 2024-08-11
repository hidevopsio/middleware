package main

import (
	"github.com/hidevopsio/iris"

	// $ go get github.com/hidevopsio/middleware/...
	ravenIris "github.com/hidevopsio/middleware/raven"
)

// https://docs.sentry.io/clients/go/integrations/http/
func init() {
	raven.SetDSN("https://<key>:<secret>@sentry.io/<project>")
}

func main() {
	app := iris.New()
	app.Use(ravenIris.RecoveryHandler)

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("Hi")
	})

	app.Run(iris.Addr(":8080"))
}
