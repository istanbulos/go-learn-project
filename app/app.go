/*
	Copyright 2018 istanbulos

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/
package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type App struct {
	Chan chan os.Signal
	Router *mux.Router
	Routes []Router
}

func (a App) Start() App {

	a.Routes = Routes()

	log.Println("Application started...")

	r := mux.NewRouter()

	for _, route := range a.Routes {
		r.HandleFunc(route.Path, route.Func)
	}

	a.Router = r

	return a
}

func (a *App) HomeController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		a.Post(w)
		break
	case "GET":
		a.Get(w)
		break
	}
}

func (a *App) Get(w http.ResponseWriter) {
	_, _ = w.Write([]byte("GEt Methodu"))

}

func (a *App) Post(w http.ResponseWriter) {
	_, _ = w.Write([]byte("Post Methodu"))
}