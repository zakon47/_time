package _time

import (
	"github.com/zakon47/_strings"
	"strconv"
	"strings"
)

type metka2 struct {
	name   string
	symbol string
	num    int
	itog   uint32
}

func (m *metka2) String() string {
	return strconv.Itoa(m.num) + m.symbol
}
func (m *metka2) StringSwap() string {
	return m.symbol + strconv.Itoa(m.num)
}
func (m *metka2) Uint32() uint32 {
	return m.itog
}
func (m *metka2) Int64() int64 {
	return int64(m.itog)
}

/*
Преобразовать строковую метку в uin32 число!
Пример: m10
*/
func MetkaLast(metka string) *metka2 {
	metka = strings.Trim(metka, " ")
	if len(metka) == 0 {
		return &metka2{}
	}
	metka = strings.ToLower(metka)

	number, symbol := _strings.NumberRight(metka)
	num, KK := goGet(number, symbol)
	return &metka2{
		num:    num,
		itog:   uint32(num) * KK,
		name:   metka,
		symbol: symbol,
	}
}

/*
Преобразовать строковую метку в uin32 число!
Пример: 10m
*/
func Metka(metka string) *metka2 {
	metka = strings.Trim(metka, " ")
	if len(metka) == 0 {
		return &metka2{}
	}
	metka = strings.ToLower(metka)

	number, symbol := _strings.NumberLeft(metka)
	num, KK := goGet(number, symbol)
	return &metka2{
		num:    num,
		itog:   uint32(num) * KK,
		name:   metka,
		symbol: symbol,
	}
}

func goGet(number, symbol string) (int, uint32) {
	num, err := strconv.Atoi(number)
	if err != nil || num < 0 {
		num = 0
	}
	//Определем коэфицент умножения
	var KK uint32 = 0
	switch symbol {
	case "m":
		KK = 60
	case "h":
		KK = 60 * 60
	case "d":
		KK = 60 * 60 * 24
	case "mm":
		KK = 60 * 60 * 24 * 31
	case "y":
		KK = 60 * 60 * 24 * 31 * 12
	case "w":
		KK = 60 * 60 * 24 * 7
	}
	return num, KK
}
