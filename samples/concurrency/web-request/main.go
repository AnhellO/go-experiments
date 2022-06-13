package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const URL = "https://excuser.herokuapp.com/v1/excuse/office/"

type controller struct {
	service service
	stopCh  chan struct{}
}

type service struct {
	url      string
	interval int64
}

func NewController(url string) *controller {
	service := service{
		url:      url,
		interval: int64(5),
	}

	return &controller{
		service: service,
		stopCh:  make(chan struct{}),
	}
}

func main() {
	controller := NewController(URL)
	ticker := time.NewTicker(time.Duration(controller.service.interval) * time.Second)

	go func() {
		fmt.Println("entra go-routine")
		for {
			select {
			case <-controller.stopCh:
				return
			case t := <-ticker.C:
				log.Println("Request at", t)
				resp, err := http.Get(controller.service.url)
				if err != nil {
					log.Fatalln(err)
				}

				defer resp.Body.Close()
				log.Println("Request done, getting response body...", t)

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatalln(err)
				}

				log.Print("Response body:\n", string(body))
			}
		}
	}()

	time.Sleep(20 * time.Second)
	ticker.Stop()
	controller.stopCh <- struct{}{}
	fmt.Println("Done")
}
