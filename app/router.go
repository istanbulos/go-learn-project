package app

import "net/http"

type Router struct {
	Path string
	Func func(w http.ResponseWriter, r *http.Request)
}

func Routes() []Router {
	return []Router{
		{Path: "/home", Func: func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Home"))
		}},
	}
}