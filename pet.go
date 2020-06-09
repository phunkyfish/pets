package pets

type Dog struct {
	name string
}

func NewDog(name string) *Dog {
	return &Dog{
		name: name,
	}
}

func (d *Dog) String() string {
	return d.name
}
