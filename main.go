package main

import (
	"os"
	"text/template"
)

func main() {
	jobConfig := `
	apiVersion: batch/v1
	kind: Job
	metadata:
	name: blender-model-proccess-1831
	spec:
	template:
		spec:
		containers:
			- name: blender
			image: gocsuntianye/blender-test:2206151637
			command:
				[
				"/app/blender/blender",
				"-b",
				]
		restartPolicy: Never
	`
	t1 := template.New("t1")
	t1, err := t1.Parse(jobConfig)
	if err != nil {
		panic(err)
	}

	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}
