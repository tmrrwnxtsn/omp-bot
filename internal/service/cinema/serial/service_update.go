package serial

import "github.com/ozonmp/omp-bot/internal/model/cinema"

func (s *DummySerialService) Update(serialID uint64, serial cinema.Serial) error {
	updatingSerial, ok := s.mapper[serialID]
	if !ok {
		return ErrSerialNotFound
	}
	updatingSerial.Title = serial.Title
	updatingSerial.Genre = serial.Genre
	updatingSerial.SeasonsNum = serial.SeasonsNum
	return nil
}
