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
		"Serial{\n"+
			"\tID=%d,\n"+
			"\tTitle=%s,\n"+
			"\tGenre=%s,\n"+
			"\tSeasonsNum=%d,\n"+
			"\tIsDeleted=%t\n"+
			"}",
		s.ID,
		s.Title,
		s.Genre,
		s.SeasonsNum,
		s.IsDeleted,
	)
}
