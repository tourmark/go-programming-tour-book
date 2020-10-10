package configs

import (
	_ "github.com/go-programming-tour-book/blog-service/pkg/setting"
)

type DatabaseSettingS struct {
	DBType   string
	UserName string
	Pwd      string
	Host     string
	DBName   string
	Charset  string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	return nill
}
