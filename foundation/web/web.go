// standard library web framework extension
package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

// HandlerFunc represents a function that handles a http request within our own framework
type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// App is entry point to our app and what configs our context
type App struct {
	*http.ServeMux
	shutdown chan os.Signal
	mw       []MidFunc
}

// NewApp creates an App value that handle a set of routes for the application.
func NewApp(shutdown chan os.Signal, mw ...MidFunc) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		shutdown: shutdown,
		mw:       mw,
	}
}

// HandlerFunc sets a handler function for a given HTTP method and path pair
// to the application server mux.
func (a *App) HandleFunc(pattern string, handlerFunc HandlerFunc, mw ...MidFunc) {
	handlerFunc = wrapMiddleware(mw, handlerFunc)
	handlerFunc = wrapMiddleware(a.mw, handlerFunc)

	h := func(w http.ResponseWriter, r *http.Request) {
		if err := handlerFunc(r.Context(), w, r); err != nil {
			// to be error handled proper
			fmt.Println(err)
			return
		}
	}

	a.ServeMux.HandleFunc(pattern, h)
}
