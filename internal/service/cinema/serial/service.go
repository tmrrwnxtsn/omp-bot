package serial

import (
	"github.com/ozonmp/omp-bot/internal/model/cinema"
)

type SerialService interface {
	Describe(serialID uint64) (*cinema.Serial, error)
	List(cursor uint64, limit uint64) ([]cinema.Serial, error)
	Create(serial cinema.Serial) (uint64, error)
	Update(serialID uint64, serial cinema.Serial) error
	Remove(serialID uint64) (bool, error)
}

type DummySerialService struct {
	storage []*cinema.Serial
	mapper  map[uint64]*cinema.Serial
	serial  uint64
}

func NewDummySerialService() *DummySerialService {
	service := &DummySerialService{
		storage: make([]*cinema.Serial, 0),
		mapper:  make(map[uint64]*cinema.Serial),
		serial:  0,
	}
	fillStorageBySampleData(service)
	return service
}
