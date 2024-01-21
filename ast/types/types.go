package types

type Type interface {
	Equal(Type) bool
	Assign(Type) bool
	String() string
}
