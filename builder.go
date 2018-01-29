package go_conf_struct

import (
	"fmt"
	"strconv"
	"strings"
)

type Builder struct {
	Errors []error
}

func (b *Builder) StringArray(name string, getter VarGetter) []string {
	v, err := getter(name)
	if err != nil {
		b.Errors = append(b.Errors, err)
		return []string{}
	}
	result := strings.Split(v.Value, ",")
	for i, value := range result {
		result[i] = strings.TrimSpace(value)
	}
	return result
}

func (b *Builder) String(name string, getter VarGetter) string {
	v, err := getter(name)
	if err != nil {
		b.Errors = append(b.Errors, err)
		return ""
	}
	return v.Value
}

func (b *Builder) Int(name string, getter VarGetter) int {
	v, err := getter(name)
	if err != nil {
		b.Errors = append(b.Errors, err)
		return 0
	}
	result, err := strconv.Atoi(v.Value)
	if err != nil {
		b.Errors = append(b.Errors, fmt.Errorf("to int convertation error for %s=%q: %v", v.Name, v.Value, err))
		return 0
	}
	return result
}

func (b *Builder) StringPointer(name string, getter VarGetter) *string {
	v, err := getter(name)
	if err != nil {
		b.Errors = append(b.Errors, err)
		return nil
	}
	if !v.Found {
		return nil
	} else {
		return &v.Value
	}
}

func (b *Builder) Bool(name string, getter VarGetter) bool {
	v, err := getter(name)
	if err != nil {
		b.Errors = append(b.Errors, err)
		return false
	}
	return v.Value != ""
}
