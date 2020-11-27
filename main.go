package main

// Import the fmt for formatting strings
// Import os so we can read environment variables from the system
import (
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, Kubernetes！I'm from Jenkins CI！")
	env_branch := os.Getenv("branch")
	if env_branch == "" {
		fmt.Fprintf(w, "env_branch is nil")
	}
	fmt.Fprintf(w, env_branch)
}

func main() {
	// fmt.Println("Hello, Kubernetes！I'm from Jenkins CI！")
	// fmt.Println("BRANCH_NAME:", os.Getenv("branch"))
	http.HandleFunc("/", hello)
	//ListenAndServe监听srv.Addr指定的TCP地址，并且会调用Serve方法接收到的连接。如果srv.Addr为空字符串，会使用":http"。
	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}
