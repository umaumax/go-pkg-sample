package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type Member struct {
	Id   int
	Name string
	Tech string
}

func main() {
	fmt.Println("----固定文字列----")
	{
		const templateText = `
	Name: {{.Name}}
	ID  : {{.Id}}
	Tech: {{.Tech}}`
		tpl := template.Must(template.New("mytemplate").Parse(templateText))
		member := Member{1, "なのは", "RH"}
		if err := tpl.Execute(os.Stdout, member); err != nil {
			log.Fatalln(err)
		}
	}
	fmt.Println()
	fmt.Println("----tplファイル----")
	{
		tpl := template.Must(template.ParseFiles("TemplateSample.tpl"))
		member := Member{7083, "フェイト", "BD"}
		if err := tpl.Execute(os.Stdout, member); err != nil {
			log.Fatalln(err)
		}
	}
}
