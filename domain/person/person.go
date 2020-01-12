package person

type Person struct {
	id   string
	name string
}

func New(id, name string) Person {
	return Person{
		id:   id,
		name: name,
	}
}

func (p Person) ID() string {
	return p.id
}

func (p Person) Name() string {
	return p.name
}
