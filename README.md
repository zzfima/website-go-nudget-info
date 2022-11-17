# Secure-Coding-in-Go
Secure Coding in Go

based on https://www.linkedin.com/learning/secure-coding-in-go

Install Go
Install VSCode
Get sources from: https://github.com/zzfima/Go-Essential-Training.git

go module maintenance initialization: go mod init go_essential_training

Check what go tools installed: Ctrl Shift P-> Go: Locate Configured Go tools

Install / update tools: Ctrl Shift P-> Go: Install/Update tools. Select all. After successfully update see message: All tools successfully installed. You are ready to Go. :)

Install dependencies: go mod tidy

Build: go build .

Run: go run .

Help: go help


Upload docker file to docker.io:
make sure Docker file exists
make build
sudo docker login -u 'xxx' -p 'xxx' docker.io
sudo docker build -t zzfima/docker-website-golang-nuget-info:latest -f Dockerfile .
sudo docker push zzfima/docker-website-golang-nuget-info:latest

Upload Docker to Azure:
az group create --name AzureGoNugetRG --location "South Central US"
az appservice plan create --name AzureGoNugetSP --resource-group AzureGoNugetRG --sku S1 --is-linux
az webapp create --resource-group AzureGoNugetRG --plan AzureGoNugetSP  --name AzureGoNugetApp --deployment-container-image-name zzfima/docker-website-golang-nuget-info:latest