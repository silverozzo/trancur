package heartbeat

import (
	"time"

	"trancur/domain/model"
)

type CourseService interface {
	SaveCourseListBySource(string, *model.ExchangeList) error
}

type Config interface {
	GetHeartbeatDuration() time.Duration
}
