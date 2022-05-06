package serial

import (
	"github.com/ozonmp/omp-bot/internal/model/cinema"
)

func (s *DummySerialService) Create(serial cinema.Serial) (uint64, error) {
	s.serial++

	serial.ID = s.serial
	serial.IsDeleted = false

	s.storage = append(s.storage, &serial)
	s.mapper[serial.ID] = &serial
	return serial.ID, nil
}
