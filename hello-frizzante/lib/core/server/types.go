package server

import (
	"embed"
	"log"
	"net/http"

	"main/lib/core/guard"
	"main/lib/core/route"
	_view "main/lib/core/view"
)

type Server struct {
	*http.Server
	Guards      []guard.Guard
	Routes      []route.Route
	PublicRoot  string
	SecureAddr  string
	Certificate string
	Key         string
	Channels    Channels
	InfoLog     *log.Logger
	Efs         embed.FS
	Render      func(view _view.View) (html string, err error)
}

type Channels struct {
	Stop chan any
}
