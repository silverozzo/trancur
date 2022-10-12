package service

import (
	"sync"

	"trancur/domain/model"
)

type Course struct {
	data sync.Map
}

func NewCourse() *Course {
	srv := &Course{}

	return srv
}

func (srv *Course) GetCourseListBySource(src string) (*model.ExchangeList, error) {
	// data := &model.ExchangeList{
	// 	Exchanges: []model.Exchange{
	// 		{
	// 			First:  "RUB",
	// 			Second: "USD",
	// 			Rate:   10.0,
	// 		},
	// 		{
	// 			First:  "RUB",
	// 			Second: "EUR",
	// 			Rate:   9.0,
	// 		},
	// 	},
	// 	Updated: time.Now(),
	// }

	data, ok := srv.data.Load(src)
	if !ok {
		return nil, model.ErrNotFound
	}

	crsList, ok := data.(model.ExchangeList)
	if !ok {
		return nil, model.ErrSomethingStrange
	}

	return &crsList, nil
}

func (srv *Course) SaveCourseListBySource(src string, data *model.ExchangeList) error {
	if data == nil {
		return model.ErrEmptyData
	}

	srv.data.Store(src, *data)

	return nil
}
