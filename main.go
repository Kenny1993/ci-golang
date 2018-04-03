package main

import(
  "net/http"
  "io"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
  io.WriteString(w, "Hello World!")
}

func main(){
  http.HandleFunc("/hello",helloHandler)
  http.ListenAndServe(":8080",nil)
}
