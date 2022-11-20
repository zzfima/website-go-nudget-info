package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	nugetInfo "github.com/zzfima/Golang-Nuget-info"
)

var router = mux.NewRouter()

func main() {
	fmt.Println("web page for Nuget")

	router.HandleFunc("/", homePageHandler).Methods("GET")
	router.HandleFunc("/versions", versionsPageHandler).Methods("GET")
	router.HandleFunc("/metadata", metadataPageHandler).Methods("GET")

	startServer()
}

// HomePageMessage ...
type HomePageMessage struct {
	Message     string
	CurrentTime string
}

// VersionsPageMessage ...
type VersionsPageMessage struct {
	Versions []string
}

func versionsPageHandler(w http.ResponseWriter, r *http.Request) {
	versionsPageTemplate, _ := template.ParseFiles("templates/versions_page.html")

	if r.ContentLength != 0 {
		vars := mux.Vars(r)
		versions, _ := nugetInfo.GetNugetVersions(vars["nugetName"])
		versionsPageMsg := VersionsPageMessage{versions}
		versionsPageTemplate.Execute(w, versionsPageMsg)
	} else {
		versionsPageTemplate.Execute(w, nil)
	}
}

func metadataPageHandler(w http.ResponseWriter, r *http.Request) {
	versionsPageTemplate, _ := template.ParseFiles("templates/metadata_page.html")

	if r.ContentLength != 0 {
		vars := mux.Vars(r)
		metadata, _ := nugetInfo.GetNugetMetadata(vars["nugetName"], vars["nugetVersion"])
		versionsPageTemplate.Execute(w, metadata)
	} else {
		versionsPageTemplate.Execute(w, nil)
	}
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	homePageTemplate, _ := template.ParseFiles("templates/home_page.html")

	homePageMsg := HomePageMessage{
		"Welcome to Nuget page Information",
		time.Now().Format("2006-01-02 15:04:05")}
	homePageTemplate.Execute(w, homePageMsg)
}

func startServer() {
	http.ListenAndServe(":8080", router)
}
