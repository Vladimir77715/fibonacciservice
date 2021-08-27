package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type config struct {
	ServerInfo Server   `json:"server"`
	Methods    []Method `json:"methods"`
}

type configParser struct {
	config *config
}

type Server struct {
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type Method struct {
	Method string                 `json:"method"`
	Args   map[string]interface{} `json:"args"`
}

func ParseConfig(path string) (*configParser, error) {
	jsonFile, err := os.Open(path)
	defer jsonFile.Close()
	if err != nil {
		return nil, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var conf config
	if e := json.Unmarshal(byteValue, &conf); e != nil {
		return nil, e
	}
	var cp = &configParser{config: &conf}
	return cp, nil
}

func (cp *configParser) GetCallUrls() ([]string, error) {
	if len(cp.config.Methods) == 0 {
		return nil, errors.New("В config.json отсутствуют методы для вызова")
	}
	si := cp.config.ServerInfo

	var hostUrl = fmt.Sprintf("%v://%v:%v", si.Protocol, si.Host, si.Port)
	var callUrls = make([]string, len(cp.config.Methods))
	for i, m := range cp.config.Methods {
		args := ""
		for k, v := range m.Args {
			if args == "" {
				args += "?"
			} else {
				args += "&"
			}
			args += fmt.Sprintf("%v=%v", k, v)
		}
		callUrls[i] = fmt.Sprintf("%v/%v%v", hostUrl, m.Method, args)
	}
	return callUrls, nil
}
