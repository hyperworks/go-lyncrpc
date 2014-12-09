package lyncrpc

type RawJson struct{
	Json string
}

var NoParams = &RawJson{""}

func (raw *RawJson) MarshalJSON() ([]byte, error) {
	return []byte(raw.Json), nil
}
