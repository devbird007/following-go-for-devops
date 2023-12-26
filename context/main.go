package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://www.golang.org", nil)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	// Attach it to our request.
	req = req.WithContext(ctx)

	// Get our resp.
	cancel()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	// Print the page to stdout
	io.Copy(os.Stdout, resp.Body)
}
