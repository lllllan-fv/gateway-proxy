package conf

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

func InitModule(path string, config interface{}) {
	paths := strings.Split(path, "/")
	if len(paths) < 2 {
		panic("illegal path")
	}

	v := viper.New()

	initDefaults(v, config)

	v.SetConfigName(paths[len(paths)-1])
	v.SetConfigType("toml")
	v.AddConfigPath(strings.Join(paths[:len(paths)-1], "/"))

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&config); err != nil {
		panic(err)
	}

	log.Println("[INFO] Successfully initialize the configuration module:", path)
}

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirect returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil).
func indirect(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

func initDefaults(v *viper.Viper, config interface{}) {
	initDefaultsRecursive(v, reflect.TypeOf(indirect(config)), "")
}

func initDefaultsRecursive(v *viper.Viper, t reflect.Type, prefix string) {
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		ft := f.Type

		if f.Anonymous {
			initDefaultsRecursive(v, ft, prefix)
		} else {
			name := strings.ToLower(f.Name)
			if value, ok := f.Tag.Lookup("mapstructure"); ok {
				name = value
			}
			if ft.Kind() == reflect.Struct {
				initDefaultsRecursive(v, ft, fmt.Sprintf("%s%s.", prefix, name))
			} else {
				if value, ok := f.Tag.Lookup("default"); ok {
					key := prefix + name
					v.SetDefault(key, value)
				}
			}
		}
	}
}
