package user

type Store interface {
	GetAll() []Model
	Get(id int) Model
	Create(user Model) Model
	Delete(id int)
}

type StubStore struct {
	models map[int]Model
}

func NewDummyStore() *StubStore {
	return &StubStore{
		map[int]Model{
			1: {
				Name: "A",
				Age:  10,
			},
			2: {
				Name: "B",
				Age:  20,
			},
		},
	}
}

func (s *StubStore) GetAll() []Model {
	var models []Model
	for _, value := range s.models {
		models = append(models, value)
	}
	return models
}

func (s *StubStore) Get(id int) Model {
	model := s.models[id]
	return model
}

func (s *StubStore) Create(model Model) Model {
	s.models[10] = model // fixme
	return model
}

func (s *StubStore) Delete(id int) {
	delete(s.models, id)
}
