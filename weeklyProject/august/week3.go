package august // week3

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

type server struct {
	addr string
	mu   sync.Mutex   // protects guestbook slice
	book []string     // guestbook entries
}

// HTML template
var tpl = `
<!DOCTYPE html>
<html>
<head>
    <title>Go Guestbook</title>
</head>
<body>
    <h1>Welcome to the Guestbook</h1>
    <form method="POST" action="/sign">
        <input type="text" name="name" placeholder="Enter your name" required>
        <button type="submit">Sign</button>
    </form>

    <h2>Entries:</h2>
    <ul>
        {{range .}}
            <li>{{.}}</li>
        {{else}}
            <li>No entries yet.</li>
        {{end}}
    </ul>
</body>
</html>
`

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		// Show guestbook page
		t := template.Must(template.New("page").Parse(tpl))
		s.mu.Lock()
		defer s.mu.Unlock()
		t.Execute(w, s.book)

	case "/sign":
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		name := r.FormValue("name")
		if name != "" {
			s.mu.Lock()
			s.book = append(s.book, name)
			s.mu.Unlock()
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)

	default:
		http.NotFound(w, r)
	}
}

func GuestbookApp() {
	s := &server{addr: ":8080", book: []string{}}
	fmt.Println("Guestbook running at http://localhost:8080/")
	err := http.ListenAndServe(s.addr, s)
	if err != nil {
		panic(err)
	}
}
