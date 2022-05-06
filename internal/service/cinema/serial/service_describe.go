package serial

import (
	"errors"
	"github.com/ozonmp/omp-bot/internal/model/cinema"
)

var ErrSerialNotFound = errors.New("serial not found")

func (s *DummySerialService) Describe(serialID uint64) (*cinema.Serial, error) {
	serial, ok := s.mapper[serialID]
	if !ok {
		return nil, ErrSerialNotFound
	}
	return serial, nil
}
