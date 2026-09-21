package mock

import (
	"errors"

	"github.com/dasilro/gocourse_domain/domain"
)

type UserSdkMock struct {
	GetMock func(id string) (*domain.User, error)
}

func (m *UserSdkMock) Get(id string) (*domain.User, error) {
	if m.GetMock == nil {
		return nil, errors.New("GetMock is not set")
	}
	return m.GetMock(id)
}
