package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	nugetInfo "github.com/zzfima/Golang-Nuget-info"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePageHandler).Methods("GET")

	router.HandleFunc("/versions", versionsPageHandler).Methods("GET", "POST")
	router.HandleFunc("/versions/{nugetName}", versionsHandler).Methods("GET")

	router.HandleFunc("/metadata", metadataPageHandler).Methods("GET", "POST")
	router.HandleFunc("/metadata/{nugetName}/{version}", metadataHandler).Methods("GET")

	fmt.Println("web page for Nuget")
	startServer(router)
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

var (
	metadataPageTemplate = template.Must(template.ParseFiles("templates/metadata_page.html"))
	versionsPageTemplate = template.Must(template.ParseFiles("templates/versions_page.html"))
	homePageTemplate     = template.Must(template.ParseFiles("templates/home_page.html"))
)

func versionsPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		versionsPageTemplate.Execute(w, nil)
		return
	}

	versions, _ := nugetInfo.GetNugetVersions(r.FormValue("nugetName"))
	versionsPageMsg := VersionsPageMessage{versions}
	versionsPageTemplate.Execute(w, versionsPageMsg)
}

func versionsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nugetName := vars["nugetName"]
	versions, _ := nugetInfo.GetNugetVersions(nugetName)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(versions)
}

func metadataPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		metadataPageTemplate.Execute(w, nil)
		return
	}

	metadata, _ := nugetInfo.GetNugetMetadata(r.FormValue("nugetName"), r.FormValue("nugetVersion"))
	metadataPageTemplate.Execute(w, metadata)
}

func metadataHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nugetName := vars["nugetName"]
	version := vars["version"]

	metadata, _ := nugetInfo.GetNugetMetadata(nugetName, version)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(metadata)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	homePageMsg := HomePageMessage{
		"Welcome to Nuget page Information",
		time.Now().Format("2006-01-02 15:04:05")}
	homePageTemplate.Execute(w, homePageMsg)
}

func startServer(router *mux.Router) {
	http.ListenAndServe(":8080", router)
}
