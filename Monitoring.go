package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/bwmarrin/discordgo"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		gd, err := discordgo.New("OTg0NTU0NzYzMzUxODkyMDIy.Gt4xFm.qEdhxTkGc5-9MWr_isvaSk4FFYAEpFjLWZEM_Y")

		if err != nil {
			log.Println(err)
		}
		err = gd.Open()
		if err != nil {
			log.Println(err)
		}
		log.Println("bot dark main is up")
		for {
			data := darkMain()
			err = gd.UpdateGameStatus(0, data)

			if err != nil {
				log.Println(err)
			}
			time.Sleep(10 * time.Second)
			debug.FreeOSMemory()
		}
	}()

	go func() {
		gd, err := discordgo.New("OTg0NTU0NjUzNDEwODE2MDYw.GBVsca.2apCqCe6Mf2o7pbQXmFR7khypT_zdwAjsIhGJ0")
		if err != nil {
			log.Println(err)
		}

		err = gd.Open()
		if err != nil {
			log.Println(err)
		}

		log.Println("bot dark dm is up")

		for {
			data := darkDM()
			err = gd.UpdateGameStatus(0, data)

			if err != nil {
				log.Println(err)
			}
			time.Sleep(10 * time.Second)
			debug.FreeOSMemory()
		}

	}()
	wg.Wait()

}

func darkMain() string {

	response, err := http.Get("https://www.gs4u.net/ru/s/229032.html")

	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()
	if response.StatusCode > 400 {
		log.Println(err)
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		log.Println(err)
	}

	a := doc.Find("div.text").Text()
	a = strings.ReplaceAll(a, "\n", "")
	a = strings.ReplaceAll(a, "\t", "")
	a = strings.ReplaceAll(a, " ", "")
	a = strings.ReplaceAll(a, "из", "/")

	return a

}
func darkDM() string {

	response, err := http.Get("https://www.gs4u.net/ru/s/228660.html")

	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()
	if response.StatusCode > 400 {
		log.Println(err)
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		log.Println(err)
	}

	a := doc.Find("div.text").Text()
	a = strings.ReplaceAll(a, "\n", "")
	a = strings.ReplaceAll(a, "\t", "")
	a = strings.ReplaceAll(a, " ", "")
	a = strings.ReplaceAll(a, "из", "/")

	return a
}
