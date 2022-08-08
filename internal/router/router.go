package router

import (
	"UltimateExplorer-Server/internal/alert"
	"UltimateExplorer-Server/internal/fileManager"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mandrigin/gin-spa/spa"
)

var webFolder string

func Init(a string) {
	webFolder = a
}

func Cors(c *gin.Context) {
	header := c.Writer.Header()
	if origin, ok := c.Request.Header["Origin"]; ok {
		header.Set("Access-Control-Allow-Origin", origin[0])
	}
	header.Set("Access-Control-Allow-Credentials", "true")
	header.Set("Access-Control-Allow-Headers", "Content-Type") //, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With
	// header.Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")

	c.Next()
}

func GET_RootFolder_Handler(c *gin.Context) {

	c.JSON(http.StatusOK, fileManager.GetRootFolders())
}

func GET_Folder_Handler(c *gin.Context) {

	folderId_str := c.Param("folderid")

	folderId, err := strconv.Atoi(folderId_str)

	if err != nil {
		panic(&alert.DataFormatError)
	}

	c.JSON(http.StatusOK, fileManager.GetFolderFiles(uint32(folderId)))
}

func GET_File_Handler(c *gin.Context) {

	fileId_str := c.Param("fileid")

	fileId, err := strconv.Atoi(fileId_str)

	if err != nil {
		panic(&alert.DataFormatError)
	}

	c.File(fileManager.GetFile(uint32(fileId)).FilePath)
}

// // GET => /poster/:filename
// func GET_Poster_Handler(c *gin.Context) {
// 	log.Println("GET => /poster/:filename")

// 	fileName := c.Param("filename")
// 	err := fileManager.Poster(c, fileName)
// 	if err != nil {
// 		c.JSON(err.Code, err)
// 	}

// }

// // GET => /exit/:password
// func GET_Exit_Handler(c *gin.Context) {
// 	log.Println("GET => /exit/:password")

// 	if c.Param("password") == "Aa0978540895" {
// 		os.Exit(0)
// 	}

// 	err := APIException.NewWarn("Invaild Password")
// 	c.JSON(err.Code, err)

// }

// // POST => /createFolder/
// func POST_CreateFolder_Handler(c *gin.Context) {
// 	log.Println("POST => /createfolder/")

// 	var m map[string]interface{}
// 	c.Bind(&m)

// 	bytes, err := fileManager.CreateFolder(int(m["id"].(float64)), m["name"].(string))
// 	if err != nil {
// 		c.JSON(err.Code, err)
// 	} else {
// 		c.Data(http.StatusOK, "application/json", bytes)
// 	}
// }

// // POST => /createFile/
// func POST_CreateFile_Handler(c *gin.Context) {
// 	log.Println("POST => /createfile/")
// 	var m map[string]interface{}
// 	c.Bind(&m)
// 	bytes, err := fileManager.CreateFile(int(m["id"].(float64)), m["name"].(string))
// 	if err != nil {
// 		c.JSON(err.Code, err)
// 	} else {
// 		c.Data(http.StatusOK, "application/json", bytes)
// 	}
// }

// // POST => /deletefolder/
// func POST_DeleteFolder_Handler(c *gin.Context) {
// 	log.Println("POST => /deletefolder/")
// 	var m map[string]interface{}
// 	c.Bind(&m)
// 	bytes, err := fileManager.DeleteFolder(int(m["id"].(float64)))
// 	if err != nil {
// 		c.JSON(err.Code, err)
// 	} else {
// 		c.Data(http.StatusOK, "application/json", bytes)
// 	}
// }

// // POST => /deletefile/
// func POST_DeleteFile_Handler(c *gin.Context) {
// 	log.Println("POST => /deletefile/")
// 	var m map[string]interface{}
// 	c.Bind(&m)
// 	bytes, err := fileManager.DeleteFile(int(m["id"].(float64)), m["name"].(string))
// 	if err != nil {
// 		c.JSON(err.Code, err)
// 	} else {
// 		c.Data(http.StatusOK, "application/json", bytes)
// 	}
// }

// // POST => /renamefolder/
// func POST_RenameFolder_Handler(c *gin.Context) {
// 	log.Println("POST => /renamefolder/")

// 	var m map[string]interface{}
// 	c.Bind(&m)
// 	log.Println(m)
// 	//TODO
// }

// // POST => /renamefile/
// func POST_RenameFile_Handler(c *gin.Context) {
// 	log.Println("POST => /renamefile/")

// 	var m map[string]interface{}
// 	c.Bind(&m)
// 	log.Println(m)
// 	//TODO
// }

// // POST => /move/
// func POST_Move_Handler(c *gin.Context) {
// 	log.Println("POST => /move/")

// 	var m map[string]interface{}
// 	c.Bind(&m)
// 	log.Println(m)
// }

// // POST => /upload/
// func POST_Upload_Handler(c *gin.Context) {
// 	log.Println("POST => /upload/")

// 	form, _ := c.MultipartForm()
// 	files := form.File["files"]
// 	for _, file := range files {
// 		log.Println(file.Filename)
// 		//c.SaveUploadedFile(file, "./files/"+file.Filename)
// 	}
// }

func MapRouter(server *gin.Engine) {
	// static assets folder
	server.Static("/assets/", webFolder+"assets/")
	server.StaticFile("/favicon.ico", webFolder+"favicon.ico")

	// middlewares
	server.Use(Cors)

	//GET:json
	server.GET("/api/folder/", GET_RootFolder_Handler)
	server.GET("/api/folder/:folderid", GET_Folder_Handler)

	// //GET:file
	server.GET("/api/file/:fileid", GET_File_Handler)

	// //GET:poster
	// server.GET("/poster/:filename", GET_Poster_Handler)

	// //GET:exit
	// server.GET("/exit/:password", GET_Exit_Handler)

	// //POST:create
	// server.POST("/createfolder/", POST_CreateFolder_Handler)
	// server.POST("/createfile/", POST_CreateFile_Handler)

	// //POST:delete
	// server.POST("/deletefolder/", POST_DeleteFolder_Handler)
	// server.POST("/deletefile/", POST_DeleteFile_Handler)

	// //POST:rename
	// server.POST("/renamefolder/", POST_RenameFolder_Handler)
	// server.POST("/renamefile/", POST_RenameFile_Handler)

	// //POST:move
	// server.POST("/move/", POST_Move_Handler)

	// //POST:upload
	// server.POST("/upload/", POST_Upload_Handler)

	// web Middleware
	server.Use(spa.Middleware("/", webFolder))
}
