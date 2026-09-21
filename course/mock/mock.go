package mock

import (
	"errors"

	"github.com/dasilro/gocourse_domain/domain"
)

type CourseSdkMock struct {
	GetMock func(id string) (*domain.Course, error)
}

func (m *CourseSdkMock) Get(id string) (*domain.Course, error) {
	if m.GetMock == nil {
		return nil, errors.New("GetMock is not set")
	}
	return m.GetMock(id)
}
