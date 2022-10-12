package http

type Config interface {
	GetDefaultCourseSource() string
	GetSelfHttpPort() string
}
