package heartbeat

import (
	"context"
	"encoding/xml"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/encoding/charmap"

	"trancur/domain/model"
)

const (
	mainCurrency = "RUB"
	sourceName   = "RUS"
)

type CourseService interface {
	SaveCourseListBySource(string, *model.ExchangeList) error
}

type Rus struct {
	srv     CourseService
	infoLog *log.Logger
	errLog  *log.Logger
}

type ResponseValCurs struct {
	XMLName xml.Name         `xml:"ValCurs"`
	Valutes []ResponseValute `xml:"Valute"`
}

type ResponseValute struct {
	XMLName  xml.Name `xml:"Valute"`
	CharCode string   `xml:"CharCode"`
	Nominal  int      `xml:"Nominal"`
	Value    string   `xml:"Value"`
}

var buff [1024 * 1024]byte

func NewRus(srv CourseService, infoLog, errLog *log.Logger) *Rus {
	hb := &Rus{
		srv:     srv,
		infoLog: infoLog,
		errLog:  errLog,
	}

	return hb
}

func (hb *Rus) StartBeat(ctx context.Context) {
	hb.infoLog.Println("запускем сердцебиение по ЦБ РФ")

	hb.tick()

	tcr := time.NewTicker(time.Hour * 24)

	for {
		select {
		case <-ctx.Done():
			tcr.Stop()
			hb.infoLog.Println("сердцебиение от ЦБ РФ остановили")
			return

		case <-tcr.C:
			hb.tick()
		}
	}
}

func (hb *Rus) tick() {
	hb.infoLog.Println("тик")

	rsp, err := http.Get("http://www.cbr.ru/scripts/XML_daily.asp")
	defer rsp.Body.Close()

	if err != nil {
		hb.errLog.Println("не смогли отработать запрос в ЦБ РФ:", err)

		return
	}

	var res ResponseValCurs

	d := xml.NewDecoder(rsp.Body)
	d.CharsetReader = charset
	err = d.Decode(&res)
	if err != nil {
		hb.errLog.Println("не смогли распарсить xml от ЦБ РФ:", err)

		return
	}

	hb.infoLog.Println("от ЦБ РФ курсы:", len(res.Valutes))

	var data model.ExchangeList

	data.Exchanges = make([]model.Exchange, len(res.Valutes))
	for i, v := range res.Valutes {
		rateFl, err := strconv.ParseFloat(strings.Replace(v.Value, ",", ".", 1), 64)
		if err != nil {
			hb.errLog.Println("не смогли распарсить строку:", v.Value)

			continue
		}

		data.Exchanges[i] = model.Exchange{
			First:  mainCurrency,
			Second: v.CharCode,
			Rate:   rateFl,
		}
	}

	hb.srv.SaveCourseListBySource(sourceName, &data)
}

func charset(chh string, input io.Reader) (io.Reader, error) {
	switch chh {
	case "windows-1251":
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	default:
		return nil, errors.New("unknown charset:" + chh)
	}
}
