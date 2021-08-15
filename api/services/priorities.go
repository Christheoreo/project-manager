package services

import "github.com/Christheoreo/project-manager/interfaces"

type PrioritiesService struct {
	PrioritiesRepository interfaces.IPrioritiesRepository
}

func (s *PrioritiesService) Exists(ID int) (bool, error) {
	return s.PrioritiesRepository.Exists(ID)
}
