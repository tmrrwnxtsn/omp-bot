package serial

func (s *DummySerialService) Remove(serialID uint64) (bool, error) {
	serial, ok := s.mapper[serialID]
	if !ok {
		return false, ErrSerialNotFound
	}
	serial.IsDeleted = true
	return true, nil
}
