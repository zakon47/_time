package _time

import (
	"github.com/zakon47/_strings"
	"strconv"
	"strings"
)

type Marker struct {
	Name    string
	Symbol  string
	Num     int
	Seconds uint32
}

func (m *Marker) Name1D() string {
	return strconv.Itoa(m.Num) + m.Symbol
}
func (m *Marker) NameD1() string {
	return m.Symbol + strconv.Itoa(m.Num)
}

/*
Преобразовать строковую метку в uin32 число!
Пример: m10
*/
func NewMarkerD1(metka string) *Marker {
	metka = strings.Trim(metka, " ")
	if len(metka) == 0 {
		return &Marker{}
	}
	metka = strings.ToLower(metka)

	number, Symbol := _strings.NumberRight(metka)
	num, KK := goGet(number, Symbol)
	return &Marker{
		Num:     num,
		Seconds: uint32(num) * KK,
		Name:    metka,
		Symbol:  Symbol,
	}
}

/*
Преобразовать строковую метку в uin32 число!
Пример: 10m
*/
func NewMarker1D(metka string) *Marker {
	metka = strings.Trim(metka, " ")
	if len(metka) == 0 {
		return &Marker{}
	}
	metka = strings.ToLower(metka)

	number, Symbol := _strings.NumberLeft(metka)
	num, KK := goGet(number, Symbol)
	return &Marker{
		Num:     num,
		Seconds: uint32(num) * KK,
		Name:    metka,
		Symbol:  Symbol,
	}
}

func goGet(number, Symbol string) (int, uint32) {
	num, err := strconv.Atoi(number)
	if err != nil || num < 0 {
		num = 0
	}
	//Определем коэфицент умножения
	var KK uint32 = 0
	switch Symbol {
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
