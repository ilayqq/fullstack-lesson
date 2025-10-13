package car

type Repository interface {
	FindAll() []string
}

type RepositoryImpl struct {
}

func (r *RepositoryImpl) FindAll() []string {
	return []string{"TOYOTA", "BMW", "MCLAREN"}
}
