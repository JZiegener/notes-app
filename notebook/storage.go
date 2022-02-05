package notebook

type Storage interface {
	StoreFile(path string) error
}

// NoOpStorage does nothing
type NoOpStorage struct {
}

func (s NoOpStorage) StoreFile(path string) error {
	return nil
}
