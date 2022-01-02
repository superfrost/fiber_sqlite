package telegram

import (
	"fiber-sqlite/database"
	"fiber-sqlite/models"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"gopkg.in/tucnak/telebot.v2"
)

var Bot *telebot.Bot

func Run() {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  os.Getenv("TELEGRAM_BOT_TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		panic("Can't connect to telegram bot")
	}

	Bot = bot

	r := &telebot.ReplyMarkup{}
	btnHelp := r.Text("ℹ Help")
	btnSettings := r.Text("⚙ Settings")
	r.Contact("Send phone number")
	btnLocation := r.Location("Send location")
	r.Data("Show help", "help")
	r.URL("Visit", "https://google.com")

	s := &telebot.ReplyMarkup{}
	btnPrev := s.Data("⬅", "prev")
	btnNext := s.Data("➡", "next")

	s.Inline(s.Row(btnPrev, btnNext))

	bot.Handle(&btnPrev, func(c *telebot.Callback) {
		bot.Respond(c, &telebot.CallbackResponse{URL: "google.com"})
	})

	bot.Handle("/user", func(m *telebot.Message) {

		var user models.User
		database.DB.Where("id = ?", "1").Find(&user)
		if user.Id == 0 {
			bot.Send(m.Sender, "Can't find anything")
			return
		}
		str := fmt.Sprintf("User %+v", user)
		bot.Send(m.Sender, str, r)

	})

	bot.Handle("/run", func(m *telebot.Message) {
		cmd := exec.Command("python", "./python_scripts/scr.py")
		out, err := cmd.Output()

		str := fmt.Sprint(string(out))
		if err != nil {
			log.Fatal(err)
		}
		bot.Send(m.Sender, str)
	})

	r.Reply(r.Row(btnHelp, btnSettings))

	bot.Handle(&btnHelp, func(m *telebot.Message) {
		bot.Send(m.Sender, "HELP!!! ")
		time.Sleep(time.Second * 2)
		bot.Send(m.Sender, "...")
		time.Sleep(time.Second * 2)
		bot.Send(m.Sender, "Nobody respond...")
	})

	bot.Handle(&btnSettings, func(m *telebot.Message) {
		bot.Send(m.Sender, "Settings")
	})

	bot.Handle(&btnLocation, func(m *telebot.Message) {
		bot.Send(m.Sender, "Location")
	})

	go bot.Start()
}
