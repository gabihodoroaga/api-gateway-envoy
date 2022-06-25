package main

/*
====NOTICE====
This code was copied form somewhere but I don't remember where.
If anybody consider this code his, ping me and I will put the
author name and to copyright notice.
*/

import (
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

func readTemplate(instream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(instream)
	return buf.String()
}

func readEnvVars(rawEnv []string) (environ map[string]string) {
	environ = make(map[string]string)
	for _, item := range rawEnv {
		parts := strings.SplitN(item, "=", 2)
		environ[parts[0]] = parts[1]
	}
	return
}

func writeTemplateToStream(tplSource string, environ map[string]string, outStream io.Writer) {
	tpl := template.New("_root_")
	tpl.Funcs(template.FuncMap{
		"split":  tplSplitStr,
		"exists": tplCheckExists,
	})
	tpl.Option("missingkey=error")
	_, err := tpl.Parse(tplSource)
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.Execute(outStream, environ)
	if err != nil {
		log.Fatal(err)
	}
}

func tplSplitStr(args ...interface{}) ([]string, error) {
	rawValue := args[0].(string)
	sep := args[1].(string)
	limit := -1
	if len(args) > 2 {
		parsedLimit, ok := args[2].(int)
		if !ok {
			err := errors.New("Limit arg (3rd) to `split` is not integer")
			return nil, err
		}
		limit = parsedLimit
	}
	return strings.SplitN(rawValue, sep, limit), nil
}

func tplCheckExists(args ...interface{}) (bool, error) {
	datamap, ok := args[0].(map[string]string)
	if !ok {
		return false, errors.New("data-map arg (1st) to `exists` should be a map[string]string, did you mean '.' or '$'?")
	}
	key, ok := args[1].(string)
	if !ok {
		return false, errors.New("lookup-key arg (2nd) to `exists` should be a string")
	}
	_, exists := datamap[key]
	return exists, nil
}

func main() {
	writeTemplateToStream(readTemplate(os.Stdin), readEnvVars(os.Environ()), os.Stdout)
}
