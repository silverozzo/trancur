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
	rusMainCurrency = "RUB"
	sourceName      = "RUS"
)

type RusHearbeat struct {
	srv     CourseService
	infoLog *log.Logger
	errLog  *log.Logger
}

type responseValCurs struct {
	XMLName xml.Name         `xml:"ValCurs"`
	Valutes []responseValute `xml:"Valute"`
	Date    string           `xml:"Date,attr"`
}

type responseValute struct {
	XMLName  xml.Name `xml:"Valute"`
	CharCode string   `xml:"CharCode"`
	Nominal  float64  `xml:"Nominal"`
	Value    string   `xml:"Value"`
}

var buff [1024 * 1024]byte

func NewRusHeartbeat(srv CourseService, infoLog, errLog *log.Logger) *RusHearbeat {
	hb := &RusHearbeat{
		srv:     srv,
		infoLog: infoLog,
		errLog:  errLog,
	}

	return hb
}

func (hb *RusHearbeat) StartBeat(ctx context.Context) {
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

func (hb *RusHearbeat) tick() {
	hb.infoLog.Println("тик по ЦБ РФ")

	rsp, err := http.Get("http://www.cbr.ru/scripts/XML_daily.asp")
	defer rsp.Body.Close()

	if err != nil {
		hb.errLog.Println("не смогли отработать запрос в ЦБ РФ:", err)

		return
	}

	var res responseValCurs

	d := xml.NewDecoder(rsp.Body)
	d.CharsetReader = winCharset
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
			First:  rusMainCurrency,
			Second: v.CharCode,
			Rate:   rateFl / v.Nominal,
		}
	}
	data.Updated, err = shitDateParse(res.Date)
	if err != nil {
		hb.errLog.Println("не смогли распарсить строку:", res.Date, err)

		return
	}

	hb.srv.SaveCourseListBySource(sourceName, &data)
}

func winCharset(chh string, input io.Reader) (io.Reader, error) {
	switch chh {
	case "windows-1251":
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	default:
		return nil, errors.New("unknown charset:" + chh)
	}
}

func shitDateParse(input string) (time.Time, error) {
	dateSl := strings.SplitN(input, ".", 3)

	if len(dateSl) != 3 {
		return time.Time{}, errors.New("not enough parts in date")
	}

	//	todo проблема с указанием часового пояса
	date, err := time.Parse("2006-01-02", dateSl[2]+"-"+dateSl[1]+"-"+dateSl[0])
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}
