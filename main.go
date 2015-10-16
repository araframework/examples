package main

import (
    "fmt"
    "reflect"
    "net/http"
)

type T struct {}

func (t *T) HandleA(w http.ResponseWriter, r *http.Request) {
    fmt.Println("AAA! im HandleA")
    w.Write([]byte("dddddddddddddddddddddd"))
}

func (t *T) NotFound() {
    fmt.Println("404 Not Found")
}

var t T
func main() {
    http.HandleFunc("/", handle)

    http.HandleFunc("/aaa/", makeHandler("HandleA"))

    err := http.ListenAndServe(":10000", nil)
    if err != nil {
        fmt.Println("ListenAndServe:", err)
    }
}

func handle(w http.ResponseWriter, r *http.Request) {
    fmt.Println("im handle")
}

func makeHandler(tp string) http.HandlerFunc {
    fn := func(w http.ResponseWriter, r *http.Request) {
        fmt.Println(tp)

        method := reflect.ValueOf(&t).MethodByName(tp)

        in := make([]reflect.Value, 2)
        in[0] = reflect.ValueOf(w)
        in[1] = reflect.ValueOf(r)

        method.Call(in)
    }
    if fn == nil {
        return func(w http.ResponseWriter, r *http.Request) {
            fmt.Println("404")

            method := reflect.ValueOf(&t).MethodByName("NotFound")
            method.Call([]reflect.Value{})
        }
    }

    return fn
}
