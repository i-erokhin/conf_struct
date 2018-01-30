package env

import (
	"fmt"
	"os"

	lib "github.com/i-erokhin/conf_struct"
)

type Source struct {
	Prefix string
}

func (e Source) get(name string) lib.Var {
	fullName := e.Prefix + name
	value, found := os.LookupEnv(fullName)
	return lib.Var{
		Value: value,
		Found: found,
		Name:  fullName,
	}
}

func (e Source) Required(name string) (v lib.Var, err error) {
	v = e.get(name)
	if !v.Found {
		err = fmt.Errorf("required ENV variable is not set: %q", v.Name)
	} else if v.Value == "" {
		err = fmt.Errorf("required ENV variable is empty: %q", v.Name)
	}
	return
}

func (e Source) Optional(name string) (lib.Var, error) {
	return e.get(name), nil
}

func (e Source) Default(dft string) lib.VarGetter {
	return func(name string) (lib.Var, error) {
		v := e.get(name)
		if !v.Found {
			v.Found = true
			v.Value = dft
		}
		return v, nil
	}
}
