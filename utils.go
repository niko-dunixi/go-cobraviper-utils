// Code generated with `make generate`; DO NOT EDIT.
//go:build !generating

package cobraviperutils

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type FlagType interface {
	int | int32 | int64 | string | uint | uint32 | uint64
}

func BindFlag[T FlagType](flagset *pflag.FlagSet, name string, value T, usage string) {
	switch any(value).(type) {
	case int:
		flagset.Int(name, any(value).(int), usage)
	case int32:
		flagset.Int32(name, any(value).(int32), usage)
	case int64:
		flagset.Int64(name, any(value).(int64), usage)
	case string:
		flagset.String(name, any(value).(string), usage)
	case uint:
		flagset.Uint(name, any(value).(uint), usage)
	case uint32:
		flagset.Uint32(name, any(value).(uint32), usage)
	case uint64:
		flagset.Uint64(name, any(value).(uint64), usage)
	}
	viper.BindPFlag(name, flagset.Lookup(name))
	viper.BindEnv(name, strings.ToUpper(strings.ReplaceAll(name, "-", "_")))
}

func GetFlag[T FlagType](name string) T {
	asTPointer := func(value any) *T {
		castValue, ok := value.(T)
		if !ok {
			placeholder := new(T)
			panic(fmt.Sprintf("could not cast value (%v) to type %T", value, &placeholder))
		}
		return &castValue
	}

	placeholder := new(T)
	switch any(*placeholder).(type) {
	case int:
		placeholder = asTPointer(viper.GetInt(name))
	case int32:
		placeholder = asTPointer(viper.GetInt32(name))
	case int64:
		placeholder = asTPointer(viper.GetInt64(name))
	case string:
		placeholder = asTPointer(viper.GetString(name))
	case uint:
		placeholder = asTPointer(viper.GetUint(name))
	case uint32:
		placeholder = asTPointer(viper.GetUint32(name))
	case uint64:
		placeholder = asTPointer(viper.GetUint64(name))
	}
	return *placeholder
}
