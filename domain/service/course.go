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

func (srv *Course) Transit(src string, inCur string, val float64, outCur string) (float64, error) {
	data, ok := srv.data.Load(src)
	if !ok {
		return 0.0, model.ErrNotFound
	}

	crsList, ok := data.(model.ExchangeList)
	if !ok {
		return 0.0, model.ErrSomethingStrange
	}

	for _, v := range crsList.Exchanges {
		if v.First == inCur && v.Second == outCur {
			return val / v.Rate, nil
		}

		if v.First == outCur && v.Second == inCur {
			return v.Rate / val, nil
		}
	}

	return 0, model.ErrNotFound
}
