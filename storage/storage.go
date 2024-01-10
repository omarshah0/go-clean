package storage

type Storage interface {
	GetAllUsers() (string, error)
}
