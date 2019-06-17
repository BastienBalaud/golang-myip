package main

import (
 "net/http"
 "fmt"
 "github.com/alexflint/go-arg"
 "log"
 "net"
)


func main() {
 var args struct {
       Port   int
 }
 args.Port = 8080
 arg.MustParse(&args)
 fmt.Println("Listen on port", args.Port)

 IndexPage := func(w http.ResponseWriter, r *http.Request) {
   ip,_,_ := net.SplitHostPort(r.RemoteAddr)
   fmt.Fprintf(w,ip+"\n")
   fmt.Println("Connection : " +ip)
  }
 http.HandleFunc("/", IndexPage)
 // server_info := fmt.Sprintf(":%s", strconv.Itoa(args.Port))
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", args.Port), nil))
}
