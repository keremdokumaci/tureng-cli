package app

type TurengProxy struct {
	Language string
}

func (p *TurengProxy) Query(word string) (string, error) {
	return "", nil
}
