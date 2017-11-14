package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()
  m.Get("/hello/:name", func(params martini.Params) string {
    return "Hello " + params["name"]+"\n"
  })
  m.Run()
}
