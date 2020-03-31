// Copyright 2020 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package xcfg_test

import (
	"fmt"
	"github.com/orivil/xcfg"
	"reflect"
)

var data = `
# mysql数据库配置
[mysql]
# 连接地址
host = "127.0.0.1"
# 连接端口
port= "3306"
# 用户名
user = "root"
# 密码
password = "123456"
# 数据库
db_name = "ginadmin"
# 连接参数
parameters = "charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true"

# postgres数据库配置
[postgres]
# 连接地址
host = "127.0.0.1"
# 连接端口
port= "5432"
# 用户名
user = "root"
# 密码
password = "123456"
# 数据库
db_name = "ginadmin"
`

type config struct {
	Mysql    *Mysql    `toml:"mysql"`
	Postgres *Postgres `toml:"postgres"`
}

type Mysql struct {
	Host       string `toml:"host"`
	Port       string `toml:"port"`
	User       string `toml:"user"`
	Password   string `toml:"password"`
	DBName     string `toml:"db_name"`
	Parameters string `toml:"parameters"`
}

type Postgres struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"db_name"`
}

func ExampleEnv_Unmarshal() {
	env, err := xcfg.Decode([]byte(data))
	if err != nil {
		panic(err)
	} else {
		// 读取所有配置
		cfg := &config{}
		err = env.Unmarshal(cfg)
		if err != nil {
			panic(err)
		}
		// 解析 mysql 配置
		mysql := &Mysql{}
		err = env.UnmarshalSub("mysql", mysql)
		if err != nil {
			panic(err)
		}
		// 解析 postgres 配置
		postgres := &Postgres{}
		err = env.UnmarshalSub("postgres", postgres)
		if err != nil {
			panic(err)
		}
		needMysql := Mysql{
			Host:       "127.0.0.1",
			Port:       "3306",
			User:       "root",
			Password:   "123456",
			DBName:     "ginadmin",
			Parameters: "charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true",
		}
		fmt.Println(reflect.DeepEqual(*cfg.Postgres, *postgres))
		fmt.Println(reflect.DeepEqual(*cfg.Mysql, *mysql))
		fmt.Println(reflect.DeepEqual(needMysql, *mysql))
	}

	// Output:
	// true
	// true
	// true
}
