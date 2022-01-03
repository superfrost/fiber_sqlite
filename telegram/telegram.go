package telegram

import (
	"fiber-sqlite/database"
	"fiber-sqlite/models"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
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

	bot.Handle(telebot.OnPhoto, func(m *telebot.Message) {

		fileId := uuid.NewV4()
		fileNameFull := fmt.Sprintf("%v.jpg", fileId)

		locatonAndFileName := fmt.Sprintf("./public/img/%s.jpg", fileId)
		bot.Download(&m.Photo.File, locatonAndFileName)
		bot.Send(m.Sender, "Processing image... Wait a minute...")

		caption := m.Photo.Caption
		segmentSize, err := strconv.Atoi(caption)
		if err != nil {
			segmentSize = 0
		}

		if segmentSize < 8 {
			segmentSize = 8
		}

		// Send to user processed image
		// cmd := exec.Command("./python_scripts/venv/Scripts/python.exe", "./python_scripts/segmentation.py", fileNameFull, fmt.Sprint(segmentSize), fileId.String(), "1")

		// For heroku
		cmd := exec.Command("python", "./python_scripts/segmentation.py", fileNameFull, fmt.Sprint(segmentSize), fileId.String(), "1")

		out, err := cmd.Output()
		if err != nil {
			fmt.Println(out)
			fmt.Println(err)
			bot.Send(m.Sender, "Sever error. Try again later.")
			return
		}

		resultFileName := fmt.Sprint(string(out))
		bot.Send(m.Sender, resultFileName)

		p := &telebot.Photo{
			File:    telebot.FromDisk(fmt.Sprintf("./public/img/result/%v", resultFileName)),
			Caption: fmt.Sprintf("Segment size: %v", segmentSize),
		}
		bot.Send(m.Sender, p)
	})

	bot.Handle(telebot.OnDocument, func(m *telebot.Message) {
		bot.Send(m.Sender, "OnDocument")
		str := fmt.Sprintf("%+v", m.Document)
		bot.Send(m.Sender, str)
	})

	r := &telebot.ReplyMarkup{}
	btnHelp := r.Text("ℹ Help")
	btnSettings := r.Text("⚙ Settings")

	r.Reply(r.Row(btnHelp, btnSettings))

	s := &telebot.ReplyMarkup{}
	btnPrev := s.Data("⬅", "prev")
	btnNext := s.Data("➡", "next")

	s.Inline(s.Row(btnPrev, btnNext))

	// bot.Handle(&btnPrev, func(c *telebot.Callback) {
	// 	bot.Respond(c, &telebot.CallbackResponse{URL: "google.com"})
	// })

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

	bot.Handle("/users", func(m *telebot.Message) {
		var users []models.User
		database.DB.Find(&users)

		str := fmt.Sprintf("%v", users)
		bot.Send(m.Sender, str)
	})

	bot.Handle("/run", func(m *telebot.Message) {
		// cmd := exec.Command("./python_scripts/venv/Scripts/python.exe", "./python_scripts/scr.py")

		cmd := exec.Command("python", "./python_scripts/scr.py")
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			return
		}

		str := fmt.Sprint(string(out))
		if err != nil {
			fmt.Println(err)
			return
		}
		bot.Send(m.Sender, str)
	})

	bot.Handle("/add", func(m *telebot.Message) {
		someString := "one    two   three four "

		words := strings.Fields(someString)

		fmt.Println(words, len(words)) // [one two three four] 4
	})

	bot.Handle("/photo", func(m *telebot.Message) {
		// Send photo from https://thispersondoesnotexist.com/image

		url := "https://thispersondoesnotexist.com/image"
		res, err := http.Get(url)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		defer res.Body.Close()

		fileName := uuid.NewV4()
		file, err := os.Create(fmt.Sprintf("./public/img/person-%v.jpg", fileName))
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		defer file.Close()

		b, err := io.Copy(file, res.Body)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		fmt.Println("File size", b)

		p := &telebot.Photo{File: telebot.FromDisk(fmt.Sprintf("./public/img/person-%v.jpg", fileName)), Caption: "Random person"}
		bot.Send(m.Sender, p)
	})

	bot.Handle("/doc", func(m *telebot.Message) {
		url := "https://thispersondoesnotexist.com/image"
		res, err := http.Get(url)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		defer res.Body.Close()

		fileName := uuid.NewV4()
		file, err := os.Create(fmt.Sprintf("./public/img/person-%v.jpg", fileName))
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		defer file.Close()

		b, err := io.Copy(file, res.Body)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		fmt.Println("File size", b)
		p := &telebot.Document{File: telebot.FromDisk(fmt.Sprintf("./public/img/person-%v.jpg", fileName)), Caption: "Random person"}
		bot.Send(m.Sender, p)
	})

	bot.Handle(&btnHelp, func(m *telebot.Message) {
		bot.Send(m.Sender, "HELP!!! ")
		time.Sleep(time.Second * 1)
		bot.Send(m.Sender, "...")
		time.Sleep(time.Second * 2)
		bot.Send(m.Sender, "Nobody respond...")
	})

	bot.Handle(&btnSettings, func(m *telebot.Message) {
		bot.Send(m.Sender, "Settings")
	})

	bot.Handle(telebot.OnText, func(m *telebot.Message) {
		// all the text messages that weren't
		// captured by existing handlers
		bot.Send(m.Sender, "Don't know this command")
	})

	go bot.Start()
}

func createWorkingDirs() {

}
