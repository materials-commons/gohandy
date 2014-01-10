package marshalling

type Marshaler interface {
	Marshal(interface{}) error
}

type Unmarshaler interface {
	Unmarshal(interface{}) error
}

type MarshalerUnmarshaler interface {
	Marshaler
	Unmarshaler
}
