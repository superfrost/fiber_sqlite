package telegram

import (
	"fiber-sqlite/database"
	"fiber-sqlite/models"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	tb "gopkg.in/tucnak/telebot.v2"
)

var Bot *tb.Bot

// Start Telegram bot with set of handlers as go routine
func Run() {
	bot, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TELEGRAM_BOT_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		panic("Can't connect to telegram bot")
	}
	Bot = bot

	// Inline btns
	r := &tb.ReplyMarkup{}
	btnHelp := r.Text("ℹ Help")
	btnSettings := r.Text("⚙ Settings")
	r.Reply(r.Row(btnHelp, btnSettings))

	bot.Handle("/id", func(m *tb.Message) {
		bot.Send(m.Sender, fmt.Sprintf("Your ID: %+v", m.Sender.ID))
	})

	bot.Handle(tb.OnPhoto, func(m *tb.Message) {

		fileId := uuid.NewV4()
		fileNameFull := fmt.Sprintf("%v.jpg", fileId)

		locatonAndFileName := fmt.Sprintf("./public/img/%s.jpg", fileId)
		bot.Download(&m.Photo.File, locatonAndFileName)
		bot.Send(m.Sender, "Processing image... Wait a minute...")
		bot.Send(m.Sender, fmt.Sprintf("%+v", m))

		caption := m.Caption
		segmentSize, err := strconv.Atoi(caption)
		bot.Send(m.Sender, fmt.Sprint(segmentSize))
		if err != nil {
			segmentSize = 0
		}
		if segmentSize < 8 {
			segmentSize = 8
		}

		cwd, _ := os.Getwd()
		scriptPath := path.Join(cwd, "python_scripts", "segmentation.py")
		pythonPath := os.Getenv("PYTHON_PATH")

		cmd := exec.Command(pythonPath, scriptPath, fileNameFull, fmt.Sprint(segmentSize), fileId.String(), "1")

		out, err := cmd.Output()
		if err != nil {
			fmt.Println(out)
			fmt.Println(err)
			bot.Send(m.Sender, "Sever error. Try again later.")
			return
		}

		resultFileName := fmt.Sprint(string(out))
		bot.Send(m.Sender, resultFileName)

		p := &tb.Photo{
			File:    tb.FromDisk(fmt.Sprintf("./public/img/result/%v", resultFileName)),
			Caption: fmt.Sprintf("Segment size: %v", segmentSize),
		}
		bot.Send(m.Sender, p)
	})

	bot.Handle(tb.OnDocument, func(m *tb.Message) {
		bot.Send(m.Sender, "OnDocument")
		str := fmt.Sprintf("%+v", m.Document)
		bot.Send(m.Sender, str)
	})

	bot.Handle("/user", func(m *tb.Message) {

		var user models.User
		database.DB.Where("id = ?", "1").Find(&user)
		if user.Id == 0 {
			bot.Send(m.Sender, "Can't find anything")
			return
		}
		str := fmt.Sprintf("User %+v", user)
		bot.Send(m.Sender, str)
	})

	bot.Handle("/users", func(m *tb.Message) {
		var users []models.User
		database.DB.Find(&users)

		str := fmt.Sprintf("%v", users)
		bot.Send(m.Sender, str)
	})

	bot.Handle("/run", func(m *tb.Message) {

		cwd, _ := os.Getwd()
		scriptPath := path.Join(cwd, "python_scripts", "scr.py")
		pythonPath := os.Getenv("PYTHON_PATH")

		cmd := exec.Command(pythonPath, scriptPath)

		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			bot.Send(m.Sender, "Sever error. Try again later.")
			return
		}

		str := string(out)
		bot.Send(m.Sender, str)
	})

	bot.Handle("/add", func(m *tb.Message) {

		someString := "one    two   three four "
		words := strings.Fields(someString)
		fmt.Println(words, len(words))
	})

	bot.Handle("/photo", func(m *tb.Message) {

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

		p := &tb.Photo{File: tb.FromDisk(fmt.Sprintf("./public/img/person-%v.jpg", fileName)), Caption: "Random person"}
		bot.Send(m.Sender, p)
	})

	bot.Handle("/doc", func(m *tb.Message) {
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
		p := &tb.Document{File: tb.FromDisk(fmt.Sprintf("./public/img/person-%v.jpg", fileName)), Caption: "Random person"}
		bot.Send(m.Sender, p)
	})

	bot.Handle("/version", func(m *tb.Message) {
		cmd := exec.Command(os.Getenv("PYTHON_PATH"), "--version")
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			bot.Send(m.Sender, "Sever error. Try again later.")
			return
		}

		result := string(out)
		bot.Send(m.Sender, result)
	})

	bot.Handle(&btnHelp, func(m *tb.Message) {
		bot.Send(m.Sender, "HELP!!! ")
		time.Sleep(time.Second * 1)
		bot.Send(m.Sender, "...")
		time.Sleep(time.Second * 1)
		bot.Send(m.Sender, "Nobody respond...")
	})

	bot.Handle(&btnSettings, func(m *tb.Message) {
		bot.Send(m.Sender, "Settings")
	})

	bot.Handle(tb.OnText, func(m *tb.Message) {
		// all the text messages that weren't captured by existing handlers
		bot.Send(m.Sender, "Don't know this command")
	})

	go bot.Start()
}
