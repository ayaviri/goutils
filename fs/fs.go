package fs

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func InitialiseServer(port int64, servingDirectoryEnvvar string) {
	var SERVING_DIRECTORY string = os.Getenv("SERVING_DIRECTORY")

	if SERVING_DIRECTORY == "" {
		log.Fatalf("Read empty serving directory name")
	}

	loggingHandler := newLoggingHandler(os.Stdout)
	http.Handle("/", loggingHandler(http.FileServer(http.Dir(SERVING_DIRECTORY))))
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func newLoggingHandler(destination io.Writer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return handlers.LoggingHandler(destination, next)
	}
}
