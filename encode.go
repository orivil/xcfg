// Copyright 2019 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package xcfg

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

const (
	jsonExt = ".json"
	ymlExt  = ".yml"
	yamlExt = ".yaml"
)

// Unmarshal parses the JSON/YAML data and stores the result
// in the value pointed to by schema. If v is nil or not a pointer,
// Unmarshal returns an error.
//
// the config value comes form 3 ways:
// 1. form config file if config file exist
// 2. form the content value if config file not exist
// 3. form os environment, and os environment value will cover the config value
//
// file is the file name under the Dir directory
// content is the default file value, if the file not exist, create one by the
// content value.
func Unmarshal(file string, schema interface{}) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	ext := filepath.Ext(file)

	// for collecting config keys
	cfgs := make(Data, 5)
	switch ext {
	case jsonExt:
		err = json.Unmarshal(data, &cfgs)
		if err != nil {
			return err
		}
		// get config value form os environment
		err = cfgs.LoadOSEnv()
		if err != nil {
			return err
		}
		data, err = json.Marshal(cfgs)
		if err != nil {
			return err
		}

	case yamlExt, ymlExt:
		err = yaml.Unmarshal(data, &cfgs)
		if err != nil {
			return err
		}
		// get config value form os environment
		err = cfgs.LoadOSEnv()
		if err != nil {
			return err
		}
		data, err = yaml.Marshal(cfgs)
		if err != nil {
			return err
		}
	default:
		return err
	}

	switch ext {
	case jsonExt:
		err = json.Unmarshal(data, schema)
		if err != nil {
			return err
		}

	case yamlExt, ymlExt:
		err = yaml.Unmarshal(data, schema)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("config file type [%s] not supported", ext)
	}
	return nil
}

// UnmarshalData is a shorthand function for loading the config data to map
func UnmarshalData(file string) (data Data, err error) {
	data = make(Data, 5)
	err = Unmarshal(file, &data)
	if err != nil {
		return nil, err
	} else {
		return data, nil
	}
}
