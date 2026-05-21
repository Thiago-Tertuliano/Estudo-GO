// Exemplo 10: adaptador de middleware stdlib (http.Handler) para Echo.
package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"middleware-study/internal/httpmw"
)

// wrapStd converte func(http.Handler) http.Handler no formato echo.MiddlewareFunc.
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
	e.Use(wrapStd(httpmw.Recover))
	e.Use(wrapStd(httpmw.RequestID))
	e.Use(wrapStd(httpmw.Logging))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "echo + httpmw adapter")
	})
	e.GET("/panic", func(c echo.Context) error {
		panic("echo panic test")
	})

	log.Println("echo bridge listening :8080")
	log.Fatal(e.Start(":8080"))
}
