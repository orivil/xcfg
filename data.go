// Copyright 2019 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package xcfg

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"strconv"
)

type Data map[string]interface{}

func (d Data) Unmarshal(schema interface{}) error {
	data, err := yaml.Marshal(d)
	if err != nil {
		return err
	} else {
		return yaml.Unmarshal(data, schema)
	}
}

func (d Data) Copy() Data {
	dst := make(Data, len(d))
	for key, value := range d {
		dst[key] = value
	}
	return dst
}

func (d Data) Set(name string, value interface{}) {
	d[name] = value
}

func (d Data) GetStr(name string) string {
	return d[name].(string)
}

func (d Data) GetInt(name string) int {
	return d[name].(int)
}

func (d Data) GetFloat(name string) float64 {
	return d[name].(float64)
}

func (d Data) GetBool(name string) bool {
	return d[name].(bool)
}

func (d Data) GetSliceStr(name string) []string {
	return d[name].([]string)
}

func (d Data) GetSliceInt(name string) []int {
	return d[name].([]int)
}

func (d Data) GetSliceFloat(name string) []float64 {
	return d[name].([]float64)
}

func (d Data) GetSliceBool(name string) []bool {
	return d[name].([]bool)
}

// LoadOSEnv for loading the OS environment values
func (d Data) LoadOSEnv() (err error) {
	for key, value := range d {
		osValue := os.Getenv(key)
		if osValue != "" {
			switch value.(type) {
			case string:
				d[key] = osValue
			case int:
				d[key], err = strconv.Atoi(osValue)
				if err != nil {
					return fmt.Errorf("OS env value [%s]: %s", key, err)
				}
			case bool:
				switch osValue {
				case "y", "Y", "yes", "YES", "Yes", "1", "t", "T", "true", "TRUE", "True":
					d[key] = true
				case "n", "N", "no", "NO", "No", "0", "f", "F", "false", "FALSE", "False":
					d[key] = false
				default:
					return fmt.Errorf("OS env value [%s]: need boolean", key)
				}
			case float64:
				d[key], err = strconv.ParseFloat(osValue, 64)
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
