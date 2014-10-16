package spares

import (
	"appengine"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"net/http"
)

const ()

func init() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	AccountHandlers(m)
	AdministationHandlers(m)

	m.Get("/", indexHandler)

	http.Handle("/", m)
}

func AppEngine(c martini.Context, r *http.Request) {
	c.Map(appengine.NewContext(r))
}

func indexHandler(r render.Render) {
	r.HTML(200, "index", nil)
}
