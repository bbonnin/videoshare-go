package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/json"
)

var resourceDir string = "resources"


// Handle the HTTP request returning the list of the video files
func handlerListVideos(writer http.ResponseWriter, req *http.Request) {
    resourceNames := []string {}
    resources, _ := ioutil.ReadDir(resourceDir)
    for _, resource := range resources {
        if !resource.IsDir() {
            resourceNames = append(resourceNames, resource.Name()); 
        }
    }

    jsonResp, err := json.Marshal(resourceNames)
    if err != nil {
        fmt.Fprintf(writer, "Boom !")
        http.Error(writer, err.Error(), http.StatusInternalServerError)
    }
    fmt.Fprintf(writer, string(jsonResp[:]))
}


func handleGetVideo(writer http.ResponseWriter, req *http.Request) {
    resource := resourceDir + "/" + req.URL.Path[len("/video/"):]
    http.ServeFile(writer, req, resource)
}

func handlePutVideo(writer http.ResponseWriter, req *http.Request) {
    http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
}

func handleDeleteVideo(writer http.ResponseWriter, req *http.Request) {
    http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
}

// Handle the HTTP request accessing a video file (DELETE/PUT/GET)
func handlerVideo(writer http.ResponseWriter, req *http.Request) {
    switch req.Method {
    case "GET":
        handleGetVideo(writer, req)
    case "PUT":
        handlePutVideo(writer, req)
    case "DELETE":
        handleDeleteVideo(writer, req)
    default:
        fmt.Fprintf(writer, "Unable to process method %s", req.Method)
    }
}

// And the main function
func main() {
    http.HandleFunc("/videos", handlerListVideos)
    http.HandleFunc("/video/", handlerVideo)
    http.ListenAndServe(":8080", nil)
}