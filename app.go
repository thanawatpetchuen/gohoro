package main

import (
	"log"
	"sync"
	"time"

	"github.com/tebeka/selenium"
)

type App struct {
	srv *selenium.Service
	cfg *Config
}

func (a *App) Start(worker int, wg *sync.WaitGroup) {
	wd, err := NewWebDriver(a.cfg.Selenium, a.cfg.ChromeDriver)
	errHandler(err)
	defer func() {
		wd.Quit()
		if wg != nil {
			wg.Done()
		}
	}()

	errHandler(wd.Get(a.cfg.SoundCloud.UserProfileURL))

	time.Sleep(5 * time.Second)

	cookieSelector := "button#onetrust-accept-btn-handler"
	cookieButton, err := wd.FindElement(selenium.ByCSSSelector, cookieSelector)
	errHandler(err)

	errHandler(cookieButton.Click())

	selector := ".soundTitle__titleContainer"
	// Get a reference to the text box containing code.
	elems, err := wd.FindElements(selenium.ByCSSSelector, selector)
	errHandler(err)

	log.Println(elems)
	for _, e := range elems {
		songSelector := ".soundTitle__usernameTitleContainer > a"

		songElem, err := e.FindElement(selenium.ByCSSSelector, songSelector)
		errHandler(err)

		href, err := songElem.GetAttribute("href")
		errHandler(err)

		if href == a.cfg.SoundCloud.SongName {
			log.Println("Found123")
			buttonSelector := ".soundTitle__playButton > a"
			buttonElem, err := e.FindElement(selenium.ByCSSSelector, buttonSelector)
			errHandler(err)

			errHandler(buttonElem.Click())
			break
		}

	}

	time.Sleep(5 * time.Second)
	log.Printf("Worker %d done.", worker)
}

func (a *App) Stop() {
	log.Println("Stopping Selenium service")
	errHandler(a.srv.Stop())
}

func (a *App) GetWorkers() int {
	return a.cfg.Workers
}

func NewApp() *App {
	cfg, err := NewConfig()
	errHandler(err)
	// Start a Selenium WebDriver server instance (if one is not already
	// running).
	service, err := NewSelenium(cfg.Selenium, cfg.ChromeDriver)
	errHandler(err)

	return &App{srv: service, cfg: cfg}

}
