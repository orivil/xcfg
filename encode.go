// Copyright 2019 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package xcfg

import (
	"github.com/BurntSushi/toml"
	"io"
	"io/ioutil"
	"os"
)

func Decode(data []byte) (env Env, err error) {
	env = make(Env)
	err = toml.Unmarshal(data, &env)
	return
}

func DecodeFile(filename string) (env Env, err error) {
	var data []byte
	data, err = ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	} else {
		return Decode(data)
	}
}

func AppendConfig(w io.Writer, data []byte) error {
	_, err := w.Write(data)
	return err
}

func AppendFile(filename, data string, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, perm)
	if err != nil {
		return err
	}
	err = AppendConfig(f, []byte(data))
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}
