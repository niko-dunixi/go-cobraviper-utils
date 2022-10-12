//go:build generating

package main

import (
	"embed"
	_ "embed"
	"encoding/json"
	"os"
	"sort"
	"strings"
	"text/template"
)

type TypeItem struct {
	Default  any       `json:"default"`
	Value    any       `json:"value"`
	Boundary *Boudnary `json:"boundary"`
}

type Boudnary struct {
	Left  string `json:"left"`
	Right string `json:"right"`
}

//go:embed templates
var embeddedTemplateFS embed.FS

func main() {
	file, err := os.Open("generate-permutations.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	permutations := make(map[string]TypeItem)
	if err := decoder.Decode(&permutations); err != nil {
		panic(err)
	}

	templates, err := template.New("").Funcs(template.FuncMap{
		"listInterfaceTypes": listInterfaceTypes,
		"lower":              strings.ToLower,
		"shouldBound":        shouldBound,
	}).ParseFS(embeddedTemplateFS, "templates/*.gotxt")
	if err != nil {
		panic(err)
	}

	if err := executeTemplate(templates, "utils.gotxt", permutations); err != nil {
		panic(err)
	}
	if err := executeTemplate(templates, "utils_test.gotxt", permutations); err != nil {
		panic(err)
	}
}

func executeTemplate(t *template.Template, filename string, items map[string]TypeItem) error {
	viperutilsFile, err := os.OpenFile(strings.TrimSuffix(filename, ".gotxt")+".go", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer viperutilsFile.Close()
	if err := t.ExecuteTemplate(viperutilsFile, filename, items); err != nil {
		return err
	}
	return nil
}

func listInterfaceTypes(items map[string]TypeItem) string {
	allTypes := make([]string, 0, len(items))
	for key := range items {
		allTypes = append(allTypes, strings.ToLower(key))
	}
	sort.Strings(allTypes)
	return strings.Join(allTypes, " | ")
}

func shouldBound(b *Boudnary) bool {
	return b != nil
}
