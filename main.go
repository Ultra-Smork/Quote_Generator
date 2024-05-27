package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type quote struct {
    Quote_text string `json:"quote"`
    Author string `json:"author"`
    Category string `json:"category"`
}

func Connector(w http.ResponseWriter, r *http.Request){
    sanedata, err := Quote_parse(w, r)
    if err != nil{
        log.Fatal("Blood is a fuel")
    }
    bestdata, err := json.Marshal(sanedata)
    if err != nil{
        log.Fatal("Hell is full")
    }
    w.Header().Set("content-type", "application/json; charset=utf-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Write(bestdata)
}

func Quote_parse(w http.ResponseWriter, r *http.Request) ([]quote, error) {
    url := "https://api.api-ninjas.com/v1/quotes"
    request, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatal("Problem with creating GET request")
        return nil, err
    }
    request.Header.Add("X-Api-Key", "WvjVzEZMj11DzG3q9LiaDg==wUc7uZiIEppoIFP1")
    response, err := http.DefaultClient.Do(request)
    if err != nil {
        log.Fatal("Couldn't send the request")
        return nil, err
    }
    defer response.Body.Close()
    var quote_info []quote
    err = json.NewDecoder(response.Body).Decode(&quote_info)
    if err != nil {
        log.Fatal("Problem with last step")
        return nil, err
    }
    return quote_info, err
}

func main() {
    http.HandleFunc("/quote", Connector)
    http.HandleFunc("/CHEBUREK", Connector)
    http.ListenAndServe(":8080", nil)
}
