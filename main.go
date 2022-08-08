package main

import (
	"UltimateExplorer-Server/internal/config"
	"UltimateExplorer-Server/internal/fileManager"
	"UltimateExplorer-Server/internal/router"
	"fmt"
	"log"
	"mime"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"github.com/nanmu42/gzip"
)

func main() {
	// MIME 設定
	mime.AddExtensionType(".js", "application/javascript")

	configPath, err := config.GetConfigPath()
	if err != nil {
		log.Fatal(err)
	}

	config, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	router.Init(config.WebFolder)
	fileManager.Init(config.FilePaths, config.DatabasePath)

	//process color
	gin.DefaultWriter = colorable.NewColorableStdout()
	gin.ForceConsoleColor()

	//release mode
	gin.SetMode(gin.ReleaseMode)

	//create server
	server := gin.Default()

	//Use Gzip
	server.Use(gzip.DefaultHandler().Gin)

	//router
	router.MapRouter(server)

	//favicon
	// server.Use(favicon.New(config.WebFolder + "favicon.ico"))
	// fmt.Println("favicon", config.WebFolder+"favicon.ico")

	//run server
	//server.Run(config.Server.Host + ":" + config.Server.Port)
	server.Run(":" + config.Server.Port)

	//ctrl c -> exit
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()
}
