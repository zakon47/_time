package _time

import (
	"strconv"
	"strings"
	"unicode/utf8"
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
	return goGet(metka, func(metka string) (string, string) {
		var pos = len(metka)
		for k := len(metka) - 1; k >= 0; k-- {
			n, _ := utf8.DecodeRuneInString(string(metka[k]))
			if n > 47 && n < 58 {
				pos--
			} else {
				break
			}
		}
		var symbol, number string
		if pos == len(metka) {
			symbol = metka[:pos]
			number = "0"
		} else {
			symbol = metka[:pos]
			number = metka[pos:]
		}
		return symbol, number
	})
}

/*
Преобразовать строковую метку в uin32 число!
Пример: 10m
*/
func Metka(metka string) *metka2 {
	return goGet(metka, func(metka string) (string, string) {
		var pos = 0
		for k := 0; k < len(metka); k++ {
			n, _ := utf8.DecodeRuneInString(string(metka[k]))
			if n > 47 && n < 58 {
				pos++
			} else {
				break
			}
		}
		var symbol, number string
		if pos == 0 {
			symbol = metka[pos:]
			number = "0"
		} else {
			symbol = metka[pos:]
			number = metka[:pos]
		}
		return symbol, number
	})
}

func goGet(metka string, fn func(metka string) (string, string)) *metka2 {
	metka = strings.Trim(metka, " ")
	if len(metka) == 0 {
		return &metka2{}
	}
	metka = strings.ToLower(metka)
	var pos = 0
	for k := 0; k < len(metka); k++ {
		n, _ := utf8.DecodeRuneInString(string(metka[k]))
		if n > 47 && n < 58 {
			pos++
		} else {
			break
		}
	}

	symbol, number := fn(metka)

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
	return &metka2{
		num:    num,
		itog:   uint32(num) * KK,
		name:   metka,
		symbol: symbol,
	}
}
