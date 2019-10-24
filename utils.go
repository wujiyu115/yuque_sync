package main

import (
	"fmt"
	"os"
	"reflect"
	"errors"
	"bytes"
	"regexp"
)

//GenNameSpace gen namespace
func GenNameSpace(config SyncConfig) string {
	return fmt.Sprintf("%s/%s", config.Login, config.Repo)
}

//CreateFile create file
func CreateFile(path string) error {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}

// ReflectStructMethod if an interface is either a struct or a pointer to a struct
func ReflectStructMethod(Iface interface{}, MethodName string) error {
	ValueIface := reflect.ValueOf(Iface)

	if ValueIface.Type().Kind() != reflect.Ptr {
		ValueIface = reflect.New(reflect.TypeOf(Iface))
	}

	Method := ValueIface.MethodByName(MethodName)
	if !Method.IsValid() {
		return fmt.Errorf("Couldn't find method `%s` in interface `%s`, is it Exported?", MethodName, ValueIface.Type())
	}
	return nil
}

// ReflectStructField if an interface is either a struct or a pointer to a struct
func ReflectStructField(Iface interface{}, FieldName string) error {
	ValueIface := reflect.ValueOf(Iface)
	if ValueIface.Type().Kind() != reflect.Ptr {
		ValueIface = reflect.New(reflect.TypeOf(Iface))
	}
	Field := ValueIface.Elem().FieldByName(FieldName)
	if !Field.IsValid() {
		return fmt.Errorf("Interface `%s` does not have the field `%s`", ValueIface.Type(), FieldName)
	}
	return nil
}

//ReflectStrVal get str val
func ReflectStrVal(Iface interface{}, FieldName string) (string, error) {
	if err := ReflectStructField(Iface, FieldName); err != nil {
		return "", err
	}
	ValueIface := reflect.ValueOf(Iface)
	return ValueIface.FieldByName(FieldName).String(), nil
}
//Cal func
func Call(m interface{}, params ...interface{}) ([]reflect.Value, error) {
	f := reflect.ValueOf(m)
	if f.Kind() != reflect.Func {
		return nil, errors.New("funcInter is not func")
	}
	if len(params) != f.Type().NumIn() {
		return nil, errors.New("the number of input params not match!")
	}
	in := make([]reflect.Value, len(params))
	for k, v := range params {
		in[k] = reflect.ValueOf(v)
	}
	return f.Call(in), nil
}

func FormatTags(tags []string) string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	tagLen := len(tags)
	for idx, st := range tags {
		buffer.WriteString(st)
		if idx == (tagLen - 1) {
			continue
		}
		buffer.WriteString(",")
	}
	buffer.WriteString("]")
	return buffer.String()
}

func FormatRaw(body string) string{
	multiBr, _ := regexp.Compile("<div style=\"display:none\">[\\s\\S]*?<\\/div>");
	hiddenContent, _ := regexp.Compile("(<br>){2}")
	return hiddenContent.ReplaceAllString(multiBr.ReplaceAllString(body, ""), "<br/>")
}