package main

import internalhttp "github.com/vsPEach/Framework/internal/server/http"

func main() {
	s := internalhttp.NewServer()
	_ = s.Start()
}
