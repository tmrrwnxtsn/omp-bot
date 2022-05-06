package cinema

import "fmt"

type Serial struct {
	ID         uint64
	Title      string
	Genre      string
	SeasonsNum int
	IsDeleted  bool
}

func (s *Serial) String() string {
	return fmt.Sprintf(
		"Serial{ID=%d,Title=%s,Genre=%s,SeasonsNum=%d,IsDeleted=%t}",
		s.ID, s.Title, s.Genre, s.SeasonsNum, s.IsDeleted,
	)
}
