package foundation

import (
	"flag"
	"github.com/labstack/echo"
	"path/filepath"
	"os"
)

const (
	ApplicationHost = ""
	ApplicationPort = "8080"
	ApplicationStorageLog = "/storage/logs/app.log"
	ApplicationHalJSON = "application/hal+json"
)

type ApplicationConfig struct {
	Host    *string
	Port    *string
}

type ApplicationFileHandler struct {
	FileHandler *os.File
}

func NewApplicationConfig() *ApplicationConfig {
	host := flag.String("host", ApplicationHost, "host name option")
	port := flag.String("port", ApplicationPort, "port number option")
	flag.Parse()
	return &ApplicationConfig{Host:host, Port:port}
}

func RegisterFileLogger(e *echo.Echo, handler *ApplicationFileHandler) {
	e.SetLogOutput(handler.FileHandler)
}

var sharedFileHandler *ApplicationFileHandler = newLogFileHandler()

func ApplicationLogFile() (*ApplicationFileHandler) {
	return sharedFileHandler
}
// shared log file handler
func newLogFileHandler() (*ApplicationFileHandler) {
	current, _ := filepath.Abs(".")
	file, _ := os.Create(current + ApplicationStorageLog)
	return &ApplicationFileHandler{FileHandler:file}
}
