package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/go-chi/chi"
)

type Host struct {
    Servers         *Servers `json:"servers"`
    ServersChanged  bool   `json:"servers_changed"`
    SslGrade        string `json:"SslGrade"`
    PreviusSslGrade string `json:"previus_ssl_grade"`
    Logo            string `json:"logo"`
    IsDown          bool   `json:"is_down"`
}

//var host map[Host]Servers
var host = Host { Servers: &Servers {
        {Name: "hola"},
        {Name: "mundo"},
    }}

type Servers []struct {
    Name     string `json:"name"`
    Sslgrade string `json:"Sslgrade"`
    Country  string `json:"country"`
    Owner    string `json:"owner"`
}

func GetServicesEndpoint(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(host)
}

func main() {
    router := chi.NewRouter()

    router.Get("/services", GetServicesEndpoint)

    log.Fatal(http.ListenAndServe(":3000", router))
}
