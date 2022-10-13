package controller

import (
	"trancur/domain/model"
)

type CourseService interface {
	GetCourseListBySource(string) (*model.ExchangeList, error)
	Transit(string, string, float64, string) (float64, error)
}

type Config interface {
	GetDefaultCourseSource() string
}
