package time_interval

import (
	"fmt"
	"github.com/zakon47/find_value"
	"strconv"
	"strings"
)

//Интерфейс для маркеров
type TimeMarkers interface {
	GetSymbol() string
	GetSeconds() uint32
	GetName1D() string
	GetNameD1() string
	fmt.Stringer
}
type timeMarker struct {
	name    string
	symbol  string
	num     int
	seconds uint32
}

func (m *timeMarker) GetName1D() string {
	return strconv.Itoa(m.num) + m.symbol
}
func (m *timeMarker) GetNameD1() string {
	return m.symbol + strconv.Itoa(m.num)
}
func (m *timeMarker) GetSymbol() string {
	return m.symbol
}
func (m *timeMarker) GetSeconds() uint32 {
	return m.seconds
}
func (m *timeMarker) String() string {
	return fmt.Sprintf("%d", m.seconds)
}

/*
Преобразовать строковую метку в uin32 число!
Пример: m10
*/
func NewMarkerD1(metka string) TimeMarkers {
	metka = strings.Trim(metka, " ")
	if len(metka) == 0 {
		return &timeMarker{}
	}
	metka = strings.ToLower(metka)

	number, Symbol := find_value.NumberReverse(metka)
	num, KK := goGet(number, Symbol)
	return &timeMarker{
		num:     num,
		seconds: uint32(num) * KK,
		name:    metka,
		symbol:  Symbol,
	}
}

/*
Преобразовать строковую метку в uin32 число!
Пример: 10m => 600
*/
func NewMarker1D(metka string) TimeMarkers {
	metka = strings.Trim(metka, " ")
	if len(metka) == 0 {
		return &timeMarker{}
	}
	metka = strings.ToLower(metka)

	number, Symbol := find_value.Number(metka)
	num, KK := goGet(number, Symbol)
	return &timeMarker{
		num:     num,
		seconds: uint32(num) * KK,
		name:    metka,
		symbol:  Symbol,
	}
}

//получить из символа - соответствующее число
func goGet(number, Symbol string) (int, uint32) {
	num, err := strconv.Atoi(number)
	if err != nil || num < 0 {
		num = 0
	}
	//Определем коэфицент умножения
	var KK uint32 = 0
	switch Symbol {
	case "s":
		KK = 1
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
