package ntopng

type JSONUnpackError struct{}

func (m *JSONUnpackError) Error() string {
	return "key not found"
}
