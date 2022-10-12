package heartbeat

import (
	"trancur/domain/model"
)

type CourseService interface {
	SaveCourseListBySource(string, *model.ExchangeList) error
}
