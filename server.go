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


 http.HandleFunc("/", rootPage)
 http.HandleFunc("/ip",ipPage)
 // server_info := fmt.Sprintf(":%s", strconv.Itoa(args.Port))
 log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", args.Port), nil))
}

func rootPage(w http.ResponseWriter, r *http.Request) {
 response := "Your IP is "+ getIp(r) + "\n"
 response += "Your user Agent is "+ r.Header.Get("User-Agent") + "\n"
 w.Write([]byte(response))
}
func ipPage(w http.ResponseWriter, r *http.Request){
 ip,_,_ := net.SplitHostPort(r.RemoteAddr)
 w.Write([]byte(ip))
}

func getIp(r *http.Request) (string){
  ip,_,_ := net.SplitHostPort(r.RemoteAddr)
  return ip
}