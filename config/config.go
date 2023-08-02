package config

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/valyala/fastjson"
)

type Config struct {
	WhiteList string
	ScipList  string
	Project   string
}

func NewConfig(fpath string) (c Config, err error) { //nolint:all
	// filename is the JSON file to read
	config, err := ioutil.ReadFile(fpath)
	if err != nil {
		return
	}

	v, err := fastjson.ParseBytes(config)
	if err != nil {
		return
	}
	if !v.Exists("WhiteList") {
		err = fmt.Errorf("not init WhiteList in %s", fpath)
		return
	}
	c.WhiteList = string(v.Get("WhiteList").GetStringBytes())

	if !v.Exists("ScipList") {
		err = fmt.Errorf("not init ScipList in %s", fpath)
		return
	}
	c.ScipList = string(v.Get("ScipList").GetStringBytes())

	if !v.Exists("Project") {
		err = fmt.Errorf("not init Project in %s", fpath)
		return
	}
	c.Project = string(v.Get("Project").GetStringBytes())

	return
}

func GetList(cpath string) []string {
	file, err := os.Open(cpath)
	if err != nil {
		return nil
	}
	defer file.Close()

	list := make([]string, 0, 10)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 2 {
			list = append(list, scanner.Text())
		}
	}

	return list
}
