package main

import (
 "net/http"
 "fmt"
 "github.com/alexflint/go-arg"
 "log"
 "net"
    "strings"
    "time"
 "strconv"
)

func main() {
 var args struct {
       Port   int
 }
 args.Port = 8080
 arg.MustParse(&args)
 fmt.Println("Listen on port", args.Port)
 srv := &http.Server{
    Addr: ":"+strconv.Itoa(args.Port),
    ReadTimeout: 5 * time.Second,
    WriteTimeout: 10 * time.Second,
}


 http.HandleFunc("/", rootPage)
 http.HandleFunc("/ip",ipPage)
 http.HandleFunc("/ua", uaPage)
 http.HandleFunc("/header", headerPage)
 http.HandleFunc("/health",getHealth)
 log.Fatal(srv.ListenAndServe())
}

func rootPage(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodHead:
        for name, values := range r.Header {
            // Loop over all values for the name.
                w.Header().Add("Request-"+name,strings.Join(values," "))
        }
        w.Header().Add("Your IP",getIp(r))
        w.Write([]byte(""))
    default:
        response := "Your IP is "+ getIp(r) + "\n"
        response += "Your user Agent is "+ r.Header.Get("User-Agent") + "\n"
        response += "Your request header is : \n"
        for name, values := range r.Header {
            // Loop over all values for the name.
            for _, value := range values {
                response += name+" : "+value+"\n"
            }
        }
        w.Write([]byte(response))
    }
}

func ipPage(w http.ResponseWriter, r *http.Request){

    switch r.Method {
    case http.MethodHead:
        w.Header().Add("Your IP",getIp(r))
    default:
        response := getIp(r) + "\n"
        w.Write([]byte(response))
    }
}

func uaPage(w http.ResponseWriter, r *http.Request){
 response := "Your user Agent is "+ r.Header.Get("User-Agent") + "\n"
 w.Write([]byte(response))
}

func headerPage(w http.ResponseWriter, r *http.Request){
 response := "Your request header is : \n"
 for name, values := range r.Header {
    // Loop over all values for the name.
    for _, value := range values {
        response += name+" : "+value+"\n"
    }
}
 w.Write([]byte(response))
}

func getIp(r *http.Request) (string){
  if r.Header.Get("X-Forwarded-For")!= ""{
    ip := r.Header.Get("X-Forwarded-For")
    return ip
  }else{
    ip,_,_ := net.SplitHostPort(r.RemoteAddr)
    return ip
  }
}

func getHealth(w http.ResponseWriter,_ *http.Request){
    w.Write([]byte("Ok"))
}