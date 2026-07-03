package engine

type Storage interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

type Engine struct {
	data map[string]string
}

func New() *Engine {
	return &Engine{
		data: make(map[string]string),
	}
}

func (e *Engine) Set(key, value string) error {
	e.data[key] = value
	return nil
}

func (e *Engine) Get(key string) (string, error) {
	return e.data[key], nil
}

func (e *Engine) Delete(key string) error {
	delete(e.data, key)
	return nil
}
