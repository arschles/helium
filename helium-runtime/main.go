package main

import (
	"os"
	"strconv"

	"github.com/gobuffalo/envy"
	"github.com/kubehelium/helium/pkg/runtime"
)

func main() {
	portStr := envy.Get("PORT", "8080")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		logger.Printf("Error: invalid port (%s)", portStr)
		os.Exit(1)
	}
	srv := runtime.NewServer(port)
	logger.Printf("Running on port %d", port)
	if err := srv.Run(); err != nil {
		logger.Printf("Error: server failed (%s)", err)
		os.Exit(1)
	}
}
