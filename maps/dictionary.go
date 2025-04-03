package maps

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because doesn't exist")
)

type DictionaryErr string

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {

	value, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

// You can modify maps without passing as an address
// You are copying a map, but just the pointer part, not the underlying data
// structure that contains the data

// NEVER init a nil map variable
func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value

	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, newValue string) error {

	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist

	case nil:
		d[key] = newValue

	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) error {

	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, key)
	default:
		return err
	}
	return nil
}

func Search(dict map[string]string, key string) (string, error) {
	return dict[key], nil
}

func (e DictionaryErr) Error() string {
	return string(e)
}
