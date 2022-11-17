package main

import (
	"fmt"
	"net/http"

	nugetInfo "github.com/zzfima/Golang-Nuget-info"
)

func main() {
	fmt.Println("web page for Nuget")

	startServer()
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Nuget page. Show all versions of Nuget MvvmCross")
	fmt.Fprintf(w, "\n")
	versions, _ := nugetInfo.GetNugetVersions("MvvmCross")
	for _, version := range versions {
		fmt.Fprintf(w, version)
		fmt.Fprintf(w, "\n")
	}
}

func startServer() {
	http.HandleFunc("/", welcomeHandler)
	http.ListenAndServe(":8080", nil)
}
