package main

import (
  "net/http"
	"text/template"
	"math/big"
)

type Page struct {
	Op string
	Left string
	Right string
	Result string
}

// func (p Page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
func calcHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{
		Left: r.FormValue("left"),
		Right: r.FormValue("right"),
		Op: r.FormValue("op"),
		Result: "",
	}
	// p.Left = r.FormValue("left")
  // p.Right = r.FormValue("right")
	// p.Op =  r.FormValue("op")

	leftInt := &big.Int{}
	rightInt := &big.Int{}

	_, leftOK := leftInt.SetString(p.Left, 10)
	_, rightOK := rightInt.SetString(p.Right, 10)

	if leftOK && rightOK {
		resultInt := &big.Int{}
		// 演算子ごとに分岐
		switch p.Op {
		case "add":
				resultInt.Add(leftInt, rightInt)
		case "sub":
				resultInt.Sub(leftInt, rightInt)
		case "multi":
				resultInt.Mul(leftInt, rightInt)
		case "div":
				resultInt.Div(leftInt, rightInt)
		}
		p.Result = resultInt.String()
	}


  tmpl, err := template.ParseFiles("form.html")
  if err != nil {
    panic(err)
  }
  err = tmpl.Execute(w, p)
  if err != nil {
    panic(err)
  }
}

func main() {
	http.HandleFunc("/", calcHandler)
	http.ListenAndServe(":3000", nil)

  // http.Handle("/", Page{})
  // http.ListenAndServe("localhost:3000", nil)
}