package conf_struct

type Var struct {
	Value string
	Found bool
	Name  string
}

type VarGetter func(string) (Var, error)

type Source interface {
	Required(string) (Var, error)
	Optional(string) (Var, error)
	// Default must always set Var.Found to true, otherwise
	// methods like StringPointer will be incorrect.
	Default(string) VarGetter
}
