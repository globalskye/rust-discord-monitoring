package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/bwmarrin/discordgo"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime/debug"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		gd, err := discordgo.New("MTAwMDM0OTMwNjg2MTEyNTY3NA.GXJhSS.CEBhJhQtVuyeAYiTsxNwi850sEwoWhOKaQySgY")
		if err != nil {
			log.Println(err)
		}

		err = gd.Open()
		if err != nil {
			log.Println(err)
		}

		log.Println("bot hunt rust is up")

		for {
			data := huntMain()
			fmt.Printf(data)
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

func huntMain() string {

	response, err := http.Get("https://huntrust.shop/#/app/store/")

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

	a := doc.Find("div.").Text()

	return a

}
