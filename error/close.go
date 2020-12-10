package ge

type Closer interface {
	Close() error
}
func Close(c Closer) error {
	if c != nil {
		return c.Close()
	}
	return nil
}