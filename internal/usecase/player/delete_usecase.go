package player

func (s *Service) DeletePlayer(id string) error {
	return s.PlayerRepository.DeleteById(id)
}
