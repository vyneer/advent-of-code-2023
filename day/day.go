package day

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

const inputFolder string = "./inputs"

var re = regexp.MustCompile("[0-9]+")

type Solver interface {
	Solve() ([]string, error)
	GetDayString() string
}

type Opts struct {
	Name        string
	Path        string
	packageName string
}

type OptFunc func(*Opts)

func DefaultOpts(d Solver) *Opts {
	o := Opts{}

	t := reflect.TypeOf(d)
	o.packageName = strings.ToLower(t.Elem().Name())
	o.Name = fmt.Sprintf("Day %s", o.getDayNumber())
	o.Path = fmt.Sprintf("%s/%s.txt", inputFolder, o.packageName)

	return &o
}

func WithTest(fileNames ...string) OptFunc {
	return func(o *Opts) {
		o.Path = fmt.Sprintf("%s/tests/%s.txt", inputFolder, o.packageName)
		if len(fileNames) == 1 {
			o.Path = fmt.Sprintf("%s/tests/%s", inputFolder, fileNames[0])
		}
	}
}

func (o *Opts) getDayNumber() string {
	return re.FindString(o.packageName)
}
