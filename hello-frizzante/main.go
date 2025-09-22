//go:build !types

package main

import (
	"embed"
	"os"

	"main/lib/config"
	"main/lib/core/guard"
	"main/lib/core/tag"
	"main/lib/core/route"
	"main/lib/core/server"
	"main/lib/core/view/ssr"
	"main/lib/routes/handlers/auth"
	"main/lib/routes/handlers/lessons"
	"main/lib/routes/handlers/todos"
	"main/lib/routes/handlers/welcome"
	authguard "main/lib/guards/auth"
)

//go:generate make clean configure
//go:generate make package
//go:generate make types
//go:embed app/dist
var efs embed.FS
var srv = server.New()
var dev = os.Getenv("DEV") == "1"
var render = ssr.New(ssr.Config{Efs: efs, Disk: dev})

func mustEnv() {
    if _, err := config.SupabaseJWKSURL(); err != nil {
        panic(err)
    }
    if _, err := config.AuthCookieName(); err != nil {
        panic(err)
    }
}

func main() {
	mustEnv()
	defer server.Start(srv)

	srv.Efs = efs
	srv.Render = render
	// Register guards
	srv.Guards = []guard.Guard{
		authguard.New(),
	}
	srv.Routes = []route.Route{
		{Pattern: "GET /", Handler: welcome.View},
		{Pattern: "GET /welcome", Handler: welcome.View},
		{Pattern: "GET /login", Handler: auth.Login},
		{Pattern: "POST /auth/session", Handler: auth.Session},
		{Pattern: "GET /auth/session", Handler: auth.Session},
		{Pattern: "GET /auth/logout", Handler: auth.Logout},
		{Pattern: "GET /auth/me", Handler: auth.Me},
		{Pattern: "GET /lessons", Handler: lessons.View, Tags: []tag.Tag{authguard.Tag}},
		{Pattern: "GET /lessons/book", Handler: lessons.Book, Tags: []tag.Tag{authguard.Tag}},
		{Pattern: "GET /lessons/cancel", Handler: lessons.Cancel, Tags: []tag.Tag{authguard.Tag}},
		{Pattern: "GET /todos", Handler: todos.View},
		{Pattern: "GET /check", Handler: todos.Check},
		{Pattern: "GET /uncheck", Handler: todos.Uncheck},
		{Pattern: "GET /add", Handler: todos.Add},
		{Pattern: "GET /remove", Handler: todos.Remove},
	}
}
