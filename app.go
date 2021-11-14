package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/tebeka/selenium"
	"github.com/thanawatpetchuen/selego/util"
)

type App interface {
	Start(worker int, wg *sync.WaitGroup, data interface{})
	Stop()
	GetWorkers() int
}

type app struct {
	srv *selenium.Service
	cfg *Config
}

func (a *app) Start(worker int, wg *sync.WaitGroup, data interface{}) {
	var selectedDay, selectedMonth, selectedYear string
	wd, err := NewWebDriver(a.cfg.Selenium, a.cfg.ChromeDriver)
	errHandler(err)
	defer func() {
		// wd.Quit()
		if wg != nil {
			wg.Done()
		}
	}()

	errHandler(wd.Get(a.cfg.HoroSite.URL))

	time.Sleep(5 * time.Second)

	log.Println("Selecting Year")

	// Select Year
	yearSelector := "#container > div.content-main-fullwidth > div > div:nth-child(10) > div.col-xl-5.col-lg-5.col-md-12.col-sm-12.col-xs-12.pt-20 > div:nth-child(4)"
	yearContainer, err := wd.FindElement(selenium.ByCSSSelector, yearSelector)
	errHandler(err)

	errHandler(yearContainer.Click())

	yearsSelector := "#select2-dd_year-results li"
	yearsElems, err := wd.FindElements(selenium.ByCSSSelector, yearsSelector)
	errHandler(err)

	for i, y := range yearsElems {
		txt, err := y.Text()
		errHandler(err)

		cy := strings.Split(txt, " ")
		yx := cy[len(cy)-1]
		if yx == "2019" {
			log.Println("index:", i)
			selectedYear = yx
			time.Sleep(2 * time.Second)
			errHandler(y.Click())
			break
		}
	}

	time.Sleep(5 * time.Second)

	log.Println("Selecting Month")

	// Select Month
	clickMonth := func() {
		monthSelector := "#container > div.content-main-fullwidth > div > div:nth-child(10) > div.col-xl-5.col-lg-5.col-md-12.col-sm-12.col-xs-12.pt-20 > div:nth-child(3)"
		monthContainer, err := wd.FindElement(selenium.ByCSSSelector, monthSelector)
		errHandler(err)

		errHandler(monthContainer.Click())
	}

	clickMonth()

	monthLength := 0

	getMonthElems := func() []selenium.WebElement {
		monthsSelector := "#select2-dd_month-results li"
		monthsElems, err := wd.FindElements(selenium.ByCSSSelector, monthsSelector)
		errHandler(err)
		return monthsElems
	}

	monthsElems := getMonthElems()

	monthLength = len(monthsElems)

	// monthsElems = monthsElems[:1]

	for mm := 3; mm < monthLength; mm++ {
		time.Sleep(1 * time.Second)
		monthsElems := getMonthElems()
		if mm >= len(monthsElems) {
			log.Println("Month exceed")
			continue
		}

		d := monthsElems[mm]

		m, err := d.Text()
		selectedMonth = util.MonthParser(m)
		errHandler(err)

		errHandler(d.Click())

		log.Println("Selecting Day")

		// Select Day
		clickDay := func() {
			daySelector := `//*[@id="container"]/div[3]/div/div[8]/div[1]/div[2]`
			dayContainer, err := wd.FindElement(selenium.ByXPATH, daySelector)
			errHandler(err)

			errHandler(dayContainer.Click())
		}
		clickDay()

		time.Sleep(1 * time.Second)

		dayLength := 0

		getDayElems := func() []selenium.WebElement {
			daysSelector := "#select2-dd_day-results li"
			daysElems, err := wd.FindElements(selenium.ByCSSSelector, daysSelector)
			errHandler(err)
			return daysElems
		}
		daysElems := getDayElems()

		dayLength = len(daysElems)

		// daysElems = daysElems[:1]

		// for loop here
		for i := 0; i < dayLength; i++ {
			time.Sleep(1 * time.Second)
			daysElems := getDayElems()

			if i >= len(daysElems) {
				log.Println("Day exceed")
				break
			}

			dd := daysElems[i]

			time.Sleep(1 * time.Second)
			selectedDay, err = dd.Text()
			errHandler(err)

			errHandler(dd.Click())

			submitSelector := "input#img_submit"
			submitContainer, err := wd.FindElement(selenium.ByCSSSelector, submitSelector)
			errHandler(err)

			errHandler(submitContainer.Click())

			time.Sleep(1 * time.Second)

			resultSelector := "#panel1 > div > div.dtinf-txt.mt-5 > div.f123x.my-5"
			resultElem, err := util.FindCSSSelector(wd, resultSelector)

			if err == nil {
				resultText, err := resultElem.Text()
				errHandler(err)

				result := util.ThaiDateParser(resultText)
				log.Println("Result:", result)
				Write(fmt.Sprintf("result/%s.csv", selectedYear), fmt.Sprintf("%s/%s/%s,%s,%v,%d,%s", selectedDay, selectedMonth, selectedYear, result.Day, result.IsWaxing, result.MoonValue, strconv.Itoa(result.Month)))
			} else {
				log.Println("Error:", err)
				break
			}

			clickDay()
		}

		clickMonth()

	}

	time.Sleep(5 * time.Second)
	log.Printf("Worker %d done.", worker)
}

func (a *app) Stop() {
	log.Println("Stopping Selenium service")
	errHandler(a.srv.Stop())
}

func (a *app) GetWorkers() int {
	return a.cfg.Workers
}

func NewApp() App {
	cfg, err := NewConfig()
	errHandler(err)
	// Start a Selenium WebDriver server instance (if one is not already
	// running).
	service, err := NewSelenium(cfg.Selenium, cfg.ChromeDriver)
	errHandler(err)

	return &app{srv: service, cfg: cfg}

}
