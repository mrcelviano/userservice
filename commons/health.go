package commons

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
	"sync/atomic"
)

var stopped = int64(0)

func StartHttp(e *echo.Echo, port int) {
	log.Println("START HTTP SERVER ON PORT ", port)
	go func() {
		<-stopChan
		atomic.StoreInt64(&stopped, 1)
	}()
	go func() {
		err := e.Start(":" + strconv.Itoa(port))
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err, " Failed to run echo!")
		}
	}()

}
