package main

import (
	"html/template"
	"net/http"
)

// Define the template
var tpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>I'm Sorry ❤️</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            text-align: center;
            background-color: #ffcccb;
            padding: 50px;
        }
        .card {
            background: white;
            padding: 20px;
            border-radius: 10px;
            box-shadow: 0px 0px 10px gray;
            display: inline-block;
        }
        button {
            background-color: #ff4d4d;
            color: white;
            padding: 10px 20px;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            font-size: 18px;
        }
        button:hover {
            background-color: #cc0000;
        }
    </style>
</head>
<body>
    <div class="card">
        <h1>I'm so sorry ❤️</h1>
        <p>I didn't mean to upset you. Please forgive me?</p>
        <form action="/forgive" method="post">
            <button type="submit">Forgive Me? 😊</button>
        </form>
    </div>
</body>
</html>
`

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.New("webpage").Parse(tpl)
	tmpl.Execute(w, nil)
}

func forgiveHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Yay! Thank you for forgiving me! ❤️  now call me ammadi !"))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/forgive", forgiveHandler)
	http.ListenAndServe(":8080", nil)
}
