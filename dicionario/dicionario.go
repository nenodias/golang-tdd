package dicionario

type Dicionario map[string]string

const (
	ErroPalavraNaoExistente = ErrDicionario("palavra não encontrada")
	ErroPalavraExistente    = ErrDicionario("a palavra já está no dicionário: ")
)

type ErrDicionario string

func (e ErrDicionario) Error() string {
	return string(e)
}

func (d Dicionario) Busca(key string) (string, error) {
	value, existe := d[key]
	if !existe {
		return "", ErroPalavraNaoExistente
	}
	return value, nil
}

func (d Dicionario) Adciona(key, value string) error {
	_, err := d.Busca(key)
	switch err {
	case ErroPalavraNaoExistente:
		d[key] = value
	case nil:
		return ErroPalavraExistente
	default:
		return err
	}

	return nil
}
