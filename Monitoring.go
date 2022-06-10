package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/bwmarrin/discordgo"
	"log"
	"net/http"
	"strings"
	"time"
)

type config struct {
	Token string
	Url   string
	ch    chan string
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	DarkRustDm := config{
		Token: "OTg0NTU0NjUzNDEwODE2MDYw.GBKmCj.GXfvWERGC8QB4NUvY1vnzSb3cE_Nr4mDbcJ6Vg",
		Url:   "https://www.gs4u.net/ru/s/228660.html",
		ch:    ch1,
	}
	DarkRustMain := config{
		Token: "OTg0NTU0NzYzMzUxODkyMDIy.GkCJyf.lvozNzD7vpQkF2b-oJuqz0rcrjsdvUFb1nlEpk",
		Url:   "https://www.gs4u.net/ru/s/229032.html",
		ch:    ch2,
	}
	go DarkRustMain.getOnline()
	go DarkRustMain.discordStatus()
	go DarkRustDm.getOnline()
	go DarkRustDm.discordStatus()

	time.Sleep(6256 * time.Hour)

}

func (cfg *config) getOnline() {
	for {
		time.Sleep(time.Second * 5)
		response, err := http.Get(cfg.Url)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()
		if response.StatusCode > 400 {
			fmt.Println("Error:", response.StatusCode)
		}
		doc, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			log.Fatalln(err)
		}
		a := doc.Find("div.text").Text()
		a = strings.ReplaceAll(a, "\n", "")
		a = strings.ReplaceAll(a, "\t", "")
		cfg.ch <- a

		continue
	}
}
func (cfg *config) discordStatus() {
	gd, err := discordgo.New(cfg.Token)
	if err != nil {
		log.Fatalln(err)
	}
	err = gd.Open()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Bot is up")
	fmt.Println("Use <ctrl+c> for exit")
	for {
		gd.UpdateGameStatus(0, <-cfg.ch)
		time.Sleep(time.Second * 10)
	}
}
