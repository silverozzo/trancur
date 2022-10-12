package heartbeat

import (
	"context"
	"encoding/xml"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html/charset"

	"trancur/domain/model"
)

const (
	thMainCurrency = "THB"
	thSourceName   = model.ThSource
	thTitleSuffix  = "Selling Rate"
)

type ThHeartbeat struct {
	srv     CourseService
	cfg     Config
	infoLog *log.Logger
	errLog  *log.Logger
}

type responseThBody struct {
	XMLName xml.Name         `xml:"RDF"`
	Items   []responseThItem `xml:"item"`
}

type responseThItem struct {
	XMLName        xml.Name `xml:"item"`
	Title          string   `xml:"title"`
	BaseCurrency   string   `xml:"baseCurrency"`
	TargetCurrency string   `xml:"targetCurrency"`
	Value          float64  `xml:"value"`
	Date           string   `xml:"date"`
}

func NewThHeartbeat(srv CourseService, cfg Config, infoLog, errLog *log.Logger) *ThHeartbeat {
	hb := &ThHeartbeat{
		srv:     srv,
		cfg:     cfg,
		infoLog: infoLog,
		errLog:  errLog,
	}

	return hb
}

func (hb *ThHeartbeat) StartBeat(ctx context.Context) {
	hb.infoLog.Println("запускем сердцебиение по ЦБ Тайланда")

	hb.tick()

	tcr := time.NewTicker(hb.cfg.GetHeartbeatDuration())

	for {
		select {
		case <-ctx.Done():
			tcr.Stop()
			hb.infoLog.Println("сердцебиение от ЦБ Тайланда остановили")
			return

		case <-tcr.C:
			hb.tick()
		}
	}
}

func (hb *ThHeartbeat) tick() {
	hb.infoLog.Println("тик по ЦБ Тайланда")

	rsp, err := http.Get("https://www.bot.or.th/App/RSS/fxrate-all.xml")
	defer rsp.Body.Close()

	if err != nil {
		hb.errLog.Println("не смогли отработать запрос в ЦБ Тайланда:", err)

		return
	}

	var res responseThBody

	d := xml.NewDecoder(rsp.Body)
	d.CharsetReader = charset.NewReaderLabel
	err = d.Decode(&res)
	if err != nil {
		hb.errLog.Println("не смогли распарсить xml от ЦБ Тайланда:", err)

		return
	}

	hb.infoLog.Println("от ЦБ Тайланда курсы:", len(res.Items))

	var data model.ExchangeList
	data.Exchanges = make([]model.Exchange, 0)
	for _, v := range res.Items {
		if !strings.HasSuffix(v.Title, thTitleSuffix) {
			continue
		}

		exchange := model.Exchange{
			First:  v.BaseCurrency,
			Second: v.TargetCurrency,
			Rate:   v.Value,
		}

		data.Exchanges = append(data.Exchanges, exchange)

		date, err := time.Parse("2006-01-02", v.Date)
		if err != nil {
			hb.infoLog.Println("не распознана строка:", v.Date)

			continue
		}

		data.Updated = date
	}

	hb.srv.SaveCourseListBySource(thSourceName, &data)
}
