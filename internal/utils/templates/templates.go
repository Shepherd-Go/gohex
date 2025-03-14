package templates

const DiFile = `package providers

import (
	"{{.module}}/cmd/api/router"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() *dig.Container {
	Container = dig.New()

	_ = Container.Provide(func() *echo.Echo {
		return echo.New()
	})

	_ = Container.Provide(router.New)

	return Container
}`

const MainFile = `package main

import (
	"fmt"
	"log"

	"{{.module}}/cmd/providers"
	"{{.module}}/config"
	"{{.module}}/cmd/api/router"
	"github.com/labstack/echo/v4"
)

func main() {

	container := provider.BuildContainer()

	if err := container.Invoke(func(router *router.Router, server *echo.Echo, config config.Config) {

		router.Init()

		server.Logger.Fatal(server.Start(fmt.Sprintf(":%d", config.Server.Port)))

	}); err != nil {

		log.Fatal(err)
	}

}
`

const ConfigFile = `package config

import (
	"log"
	"sync"

	"github.com/andresxlp/gosuite/config"
)

var (
	Once sync.Once
	cfg  *Config
)

func Get() *Config {
	if cfg == nil {
		log.Panic("Configuration has not yet been initialized")
	}
	return cfg
}

type Config struct {
	Server   Server   env:"server"
}

type Server struct {
	Port int env:"port"
}


func Environments() {
	Once.Do(func() {
		cfg = new(Config)
		if err := config.GetConfigFromEnv(cfg); err != nil {
			log.Panicf("error parsing enviroment vars \n%v", err)
		}
	})
}
`

const HealthFile = `package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Health struct {
	Code    int    ` + "`json:" + `"status"` + "`" + `
	Message string ` + "`json:" + `"message"` + "`" + `
}

func HealthCheck(context echo.Context) error {
	response := &Health{
		Code:    http.StatusOK,
		Message: "Active!",
	}

	return context.JSON(http.StatusOK, response)
}
`

const RouterFile = `package router

import (
	"{{.module}}/cmd/api/handler"

	"github.com/labstack/echo/v4"
)

type Router struct {
	server   *echo.Echo
}

func New(server *echo.Echo) *Router {
	return &Router{
		server,
	}
}

func (r *Router) Init() {
	basePath := r.server.Group("/api/microservice") //customize your basePath 
	basePath.GET("/health", handler.HealthCheck)
}
`

const Launch = `
{
    // Use IntelliSense to learn about possible attributes.
	// Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Go Program with .env",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}/cmd/main.go",
            "envFile": "${workspaceFolder}/.env",
            "args": [],
            "trace": "verbose"
        }
    ]
}
`

const GitHubActionsIntegration = `name: Continuous Integration

on:
  pull_request:
    branches:  ["main"] 

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v4
      with: 
        fetch-depth: 0

    - name: Set up Go
      uses: actions/setup-go@v5
      with:
        go-version: '1.24'

    - name: Install dependencies
      run: go get ./...

    - name: Build
      run: go build -v ./...

  test:
    runs-on: ubuntu-latest
    needs: build
    steps:

      - uses: actions/checkout@v4
        with: 
          fetch-depth: 0

      - name: Load .env file
        uses: aarcangeli/load-dotenv@v1
        with:
            filenames: |
              .env.testing
    
      - name: Test with the Go CLI
        run: go test -race -covermode atomic -coverprofile=covprofile ./...    

  docker_validate:
    runs-on: ubuntu-latest
    needs: build
    steps:
      - uses: actions/checkout@v4

      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v2

      - name: Validate Dockerfile build
        run: docker build --no-cache .
`

const EnvFile = `SERVER_PORT=3000`
