package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/bloqhead/demigods.go/handler"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func RegisterRouter(r *gin.RouterGroup) {
	// prepare our store for route caching
	store := persistence.NewInMemoryStore(time.Second)

	// fetch all items
	r.GET("/api/all", cache.CachePage(store, time.Minute, handler.FetchAll))

	// fetch by ID
	r.GET("/api/id/:id", cache.CachePage(store, time.Minute, handler.FetchById))
	
	// fetch by category
	r.GET("/api/category/:category", cache.CachePage(store, time.Minute, handler.FetchByCategory))

	// fetch category list
	r.GET("/api/categories", cache.CachePage(store, time.Minute, handler.FetchCategories))
}

func ErrRouter(c *gin.Context) {
	c.String(http.StatusBadRequest, "url err")
}

// init gin app
func init() {
	app = gin.New()

	// Use the CORS middleware
	app.Use(handler.Cors)

	// Handling routing errors
	app.NoRoute(func(c *gin.Context) {
		sb := &strings.Builder{}
		sb.WriteString("Available routes:\n")
		for _, v := range app.Routes() {
			sb.WriteString(fmt.Sprintf("- %s %s\n", v.Method, v.Path))
		}
		c.String(http.StatusBadRequest, sb.String())
	})

	r := app.Group("/")

	// register route
	RegisterRouter(r)
}

func main() {
	// make sure our data is prepped
	handler.FetchData()

	// run the app
	app.Run(":8080")
}