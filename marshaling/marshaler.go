package marshaling

type Marshaler interface {
	Marshal(interface{}) error
}

type Unmarshaler interface {
	Unmarshal(interface{}) error
}

type MarshalUnmarshaler interface {
	Marshaler
	Unmarshaler
}
