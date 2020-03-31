// Copyright 2019 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package xcfg

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"strconv"
)

type Env map[string]interface{}

func (e Env) Unmarshal(schema interface{}) error {
	buf := new(bytes.Buffer)
	err := toml.NewEncoder(buf).Encode(e)
	if err != nil {
		return err
	}
	return toml.Unmarshal(buf.Bytes(), schema)
}

func (e Env) GetSub(namespace string) (env Env, ok bool) {
	mp, o := e[namespace].(map[string]interface{})
	if !o {
		return nil, false
	} else {
		return mp, true

	}
}

func (e Env) Len() int {
	return len(e)
}

func (e Env) GetStr(name string) string {
	return e[name].(string)
}

func (e Env) GetInt(name string) int {
	return e[name].(int)
}

func (e Env) GetFloat(name string) float64 {
	return e[name].(float64)
}

func (e Env) GetBool(name string) bool {
	return e[name].(bool)
}

func (e Env) GetSliceStr(name string) []string {
	return e[name].([]string)
}

func (e Env) GetSliceInt(name string) []int {
	return e[name].([]int)
}

func (e Env) GetSliceFloat(name string) []float64 {
	return e[name].([]float64)
}

func (e Env) GetSliceBool(name string) []bool {
	return e[name].([]bool)
}

// LoadOSEnv for loading the OS environment values
func (e Env) LoadOSEnv() (err error) {
	for key, value := range e {
		osValue := os.Getenv(key)
		if osValue != "" {
			switch value.(type) {
			case string:
				e[key] = osValue
			case int:
				e[key], err = strconv.Atoi(osValue)
				if err != nil {
					return fmt.Errorf("OS env value [%s]: %s", key, err)
				}
			case bool:
				switch osValue {
				case "y", "Y", "yes", "YES", "Yes", "1", "t", "T", "true", "TRUE", "True":
					e[key] = true
				case "n", "N", "no", "NO", "No", "0", "f", "F", "false", "FALSE", "False":
					e[key] = false
				default:
					return fmt.Errorf("OS env value [%s]: need boolean", key)
				}
			case float64:
				e[key], err = strconv.ParseFloat(osValue, 64)
				if err != nil {
					return fmt.Errorf("OS env value [%s]: %s", key, err)
				}
			default:
				return errors.New("os config value only support 'string', 'int', 'float64' or 'bool'")
			}
		}
	}
	return nil
}
