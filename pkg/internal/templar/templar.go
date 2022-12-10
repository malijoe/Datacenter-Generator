package templar

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/go-task/slim-sprig"
)

func TemplateString(ts string, vars any, funcMap ...map[string]any) (string, error) {
	writer := bytes.NewBuffer([]byte{})

	funcs := sprig.GenericFuncMap()
	if len(funcMap) > 0 {
		for k, v := range funcMap[0] {
			funcs[k] = v
		}
	}

	templar, err := template.New("template").Funcs(funcs).Parse(ts)
	if err != nil {
		return "", err
	}

	if err = templar.Execute(writer, vars); err != nil {
		return "", err
	}

	str := writer.String()
	if str == "" && ts != "" {
		return "", fmt.Errorf("templating resulted in an empty string, template: {%s}, vars: {%+v}", ts, vars)
	}
	return str, nil
}
