// Exemplo 07: Echo com adaptador de middleware stdlib.
package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"

	"rate-limiting-study/internal/ratelimit"
)

func wrapStd(mw func(http.Handler) http.Handler) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var err error
			h := mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				err = next(c)
			}))
			h.ServeHTTP(c.Response(), c.Request())
			return err
		}
	}
}

func main() {
	e := echo.New()
	l := rate.NewLimiter(2, 4)
	e.Use(wrapStd(ratelimit.Middleware(l)))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "echo + rate limit")
	})

	log.Println("echo bridge listening :8080")
	log.Fatal(e.Start(":8080"))
}
