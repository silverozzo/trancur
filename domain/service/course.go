package service

import (
	"time"

	"trancur/domain/model"
)

type Course struct{}

func NewCourse() *Course {
	srv := &Course{}

	return srv
}

func (srv *Course) GetCourseListBySource(src string) *model.ExchangeList {
	data := &model.ExchangeList{
		Exchanges: []model.Exchange{
			{
				First:  "RUB",
				Second: "USD",
				Rate:   10.0,
			},
			{
				First:  "RUB",
				Second: "EUR",
				Rate:   9.0,
			},
		},
		Updated: time.Now(),
	}

	return data
}
