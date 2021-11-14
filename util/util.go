package util

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/tebeka/selenium"
)

func Click(wd selenium.WebDriver, selector string) {
	elem, err := wd.FindElement(selenium.ByCSSSelector, selector)
	errHandler(err)

	errHandler(elem.Click())
}

func FindCSSSelector(wd selenium.WebDriver, selector string) (selenium.WebElement, error) {
	elem, err := wd.FindElement(selenium.ByCSSSelector, selector)
	// errHandler(err)
	return elem, err
}

func ThaiNumberToArabic(text string) string {
	temp := text
	for _, t := range temp {
		temp = strings.Replace(temp, string(t), ThaiNumbers[string(t)], -1)
	}
	return temp
}

//       0       1   2  3     4        5   6    7
// ตรงกับวันอังคาร แรม ๑๐ ค่ำ เดือนอ้าย(๑) ปีจอ จ.ศ. ๑๓๘๐
func ThaiDateParser(text string) Date {
	boom := strings.Split(text, " ")

	day := strings.Replace(boom[0], "ตรงกับวัน", "", 1)
	day = ThaiDay[day]

	isWaxing := boom[1] == "ขึ้น"

	moonValue, _ := strconv.Atoi(ThaiNumberToArabic(boom[2]))

	month := boom[4]
	e := `\((.*?)\)`
	r := regexp.MustCompile(e)
	result := r.FindAllStringSubmatch(month, -1)

	monthV, _ := strconv.Atoi(ThaiNumberToArabic(result[0][1]))

	log.Println(result)

	return Date{
		Day:       day,
		IsWaxing:  isWaxing,
		MoonValue: moonValue,
		Month:     monthV,
	}
}

func MonthParser(text string) string {
	boom := strings.Split(text, " ")
	e := `\((.*?)\)`
	r := regexp.MustCompile(e)
	result := r.FindAllStringSubmatch(boom[1], -1)
	m, err := strconv.Atoi(result[0][1])
	errHandler(err)

	return strconv.Itoa(m)
}
