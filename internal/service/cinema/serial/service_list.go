package serial

import (
	"github.com/ozonmp/omp-bot/internal/model/cinema"
)

func (s *DummySerialService) List(cursor uint64, limit uint64) ([]cinema.Serial, error) {
	var from uint64
	if cursor == 0 {
		from = 0
	} else {
		from = cursor - 1
	}

	until := from + limit
	if until >= uint64(len(s.storage)) {
		return []cinema.Serial{}, nil
	}

	result := make([]cinema.Serial, 0, limit)
	for i := from; i <= until && i <= uint64(len(s.storage)); i++ {
		serial := *s.storage[i]
		if serial.IsDeleted {
			if until+1 < uint64(len(s.storage)) {
				until++
			}
			continue
		}
		result = append(result, serial)
	}
	return result, nil
}
