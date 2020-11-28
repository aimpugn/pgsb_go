package main

import (
	"fmt"
	"net/http"
	"os"

	/*
		go.mod 상단에 선언된 module 이름을 기준으로 `import path`가 결정된다
		# `module pgsb`인 경우
		import "pgsb/api"

		# `module github.com/aimpugn/pgsb`인 경우
		import "github.com/aimpugn/pgsb/api"
	*/
	"github.com/aimpugn/pgsb/api"
)

// If can not access this go server via browser, before move forward, check firewalld rather than iptables
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", api.EchoHandleFunc)

	// retrieve all books or save books
	http.HandleFunc("/api/books", api.BooksHandleFunc)
	http.HandleFunc("/api/books/", api.BookHandleFunc)

	err := http.ListenAndServe(port(), nil)
	fmt.Println(err)
}

func port() string {
	port := os.Getenv("PORT")
	fmt.Println(port)
	if len(port) == 0 {
		port = "8089"
	}

	return ":" + port
}

/*
https://stackoverflow.com/a/13255928/8562273
# `http.ResponseWriter`
- 인터페이스로 보이지 않는다(not visible)
- 이 인터페이스를 구현하는 기존 타입이 포인터이므로, 인터페이스에 대한 포인터를 사용할 필요가 없다
- `http.ResponseWriter` 인터페이스는 실제로는 `http.response` 구조체가 백업한다
- https://www.airs.com/blog/archives/281
  - 인터페이스는 컴파일러가 컴파일 시 보는 정적 타입이며, 다른 타입들은 동적 타입으로, 런타임 시 보이게(visible at runtime) 된다
  - 할당 또는 함수 호출 등으로 `인터페이스 값을 복사할 때, 동적 타입의 값을 복사하게 되며, 이는 일반적으로 대부분의 타입이 작동하는 방식이다.
  - 하지만, 인터페이스를 사용하는 가장 일반적인 케이스는 동적 타입이 포인터 타입이 되는 것
  - 인터페이스를 복사할 때, 포인터를 복사하는 것이지, 포인터가 가리키는 값을 복사하는 것이 아니다

# `*http.Request`
- 인터페이스 아닌 구조체
- ```Except for reading the body, handlers should not modify the provided Request.```
- `Handler`는 요청 구조체를 수정하지 말고, 수정된 복사본을 넘겨야 한다(https://stackoverflow.com/a/56875204/8562273)
*/
func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Cloud Native Go.")
}
