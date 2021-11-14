package util_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thanawatpetchuen/selego/util"
)

func TestThaiNumberToArabic(t *testing.T) {
	text := "๒"

	result := util.ThaiNumberToArabic(text)
	assert.Equal(t, "2", result)
}

func TestThaiNumberToArabicLong(t *testing.T) {
	text := "๒๓๓"

	result := util.ThaiNumberToArabic(text)
	assert.Equal(t, "233", result)
}

func TestThaiDateParserDay(t *testing.T) {
	text := "ตรงกับวันอังคาร แรม ๑๐ ค่ำ เดือนอ้าย(๑) ปีจอ จ.ศ. ๑๓๘๐"

	result := util.ThaiDateParser(text)

	assert.Equal(t, "tuesday", result.Day)
}

func TestThaiDateParserWaxing(t *testing.T) {
	text := "ตรงกับวันอังคาร แรม ๑๐ ค่ำ เดือนอ้าย(๑) ปีจอ จ.ศ. ๑๓๘๐"

	result := util.ThaiDateParser(text)

	assert.Equal(t, false, result.IsWaxing)
}

func TestThaiDateParserMoonValue(t *testing.T) {
	text := "ตรงกับวันอังคาร แรม ๑๐ ค่ำ เดือนอ้าย(๑) ปีจอ จ.ศ. ๑๓๘๐"

	result := util.ThaiDateParser(text)

	assert.Equal(t, 10, result.MoonValue)
}

func TestThaiDateParserMonth(t *testing.T) {
	text := "ตรงกับวันอังคาร แรม ๑๐ ค่ำ เดือนอ้าย(๑) ปีจอ จ.ศ. ๑๓๘๐"

	result := util.ThaiDateParser(text)

	assert.Equal(t, 1, result.Month)
}
