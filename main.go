package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	nugetInfo "github.com/zzfima/Golang-Nuget-info"
)

func main() {
	fmt.Println("web page for Nuget")
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
		versions, _ := nugetInfo.GetNugetVersions(r.FormValue("nugetName"))
		versionsPageMsg := VersionsPageMessage{versions}
		versionsPageTemplate.Execute(w, versionsPageMsg)
	} else {
		versionsPageTemplate.Execute(w, nil)
	}
}

func metadataPageHandler(w http.ResponseWriter, r *http.Request) {
	versionsPageTemplate, _ := template.ParseFiles("templates/metadata_page.html")

	if r.ContentLength != 0 {
		metadata, _ := nugetInfo.GetNugetMetadata(r.FormValue("nugetName"), r.FormValue("nugetVersion"))
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
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/versions", versionsPageHandler)
	http.HandleFunc("/metadata", metadataPageHandler)
	http.ListenAndServe(":8080", nil)
}
