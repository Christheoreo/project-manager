package interfaces

type IPrioritiesRepository interface {
	Exists(ID int) (bool, error)
}

type IPrioritiesService interface {
	Exists(ID int) (bool, error)
}
