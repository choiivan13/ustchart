package db

type DbHandler interface {
	Update() (bool, error)
}
