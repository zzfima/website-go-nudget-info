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

func versionsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Show all versions of Nuget MvvmCross")
	fmt.Fprintf(w, "\n")
	versions, _ := nugetInfo.GetNugetVersions("MvvmCross")
	for _, version := range versions {
		fmt.Fprintf(w, version)
		fmt.Fprintf(w, "\n")
	}
}

// HomePageMessage ...
type HomePageMessage struct {
	Message     string
	CurrentTime string
}

// VersionsPageMessage ...
type VersionsPageMessage struct {
	Message     string
	CurrentTime string
	Versions    []string
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	homePageTemplate, _ := template.ParseFiles("templates/home_page.html")
	homePageMsg := HomePageMessage{"Welcome to Nuget page Information", time.Now().Format("2006-01-02 15:04:05")}
	homePageTemplate.Execute(w, homePageMsg)

	versionsPageTemplate, _ := template.ParseFiles("templates/versions_page.html")
	versions, _ := nugetInfo.GetNugetVersions(r.FormValue("nugetName"))
	versionsPageMsg := VersionsPageMessage{
		"Welcome to Versions page Information", time.Now().Format("2006-01-02 15:04:05"),
		versions}

	versionsPageTemplate.Execute(w, versionsPageMsg)
}

func startServer() {
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/versions", versionsHandler)
	http.ListenAndServe(":8080", nil)
}
