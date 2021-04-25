package Aplication

type Operations interface {
	get(key string) (string, error)
	set(key string, value string) (bool, error)
	upsert(key string, value string) (bool, error)
	delete(key string) (bool, error)
}
