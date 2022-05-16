// main.go

package main

func main() {
	a := App{}
	a.Initialize(
		"postgres",
		"CDAS",
		"cdas_ex02")

	a.Run(":5432")
}
