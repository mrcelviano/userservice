package commons

import (
	"context"
	"github.com/labstack/echo"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var StopTimeout = time.Minute * 2
var stopChan = make(chan bool)

// Обработчик сигналов операционной системы.
// Мягко (не принимает новые запросы и дожидается пока текущие обработаются) останавливает http сервер.
// Если в stopChan поступит false, приложение остановится через StopTimeout времени

func NewSignalHandler(e *echo.Echo) chan bool {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go handle(c)
	go func() {
		notStop := <-stopChan
		if e != nil {
			log.Println("STOP HTTP server")
			err := e.Shutdown(context.Background())
			if err != nil {
				log.Fatal("can't stop http server")
			}
		}
		if !notStop {
			time.Sleep(StopTimeout)
			os.Exit(99)
		}
	}()
	return stopChan
}

func handle(sigChan <-chan os.Signal) {
	stopped := false
	for sig := range sigChan {
		log.Println("STOP TRIGGER: signal")
		log.Println(sig.String() + " signal!!!")
		if stopped {
			time.Sleep(time.Second)
			os.Exit(561)
		}
		stopped = true
		go func() {
			for {
				stopChan <- true
			}
		}()
	}
}
