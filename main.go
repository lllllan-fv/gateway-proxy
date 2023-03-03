package main

import (
	"fmt"

	"github.com/lllllan-fv/gateway-proxy/internal/dao"
)

func main() {
	for _, gsi := range dao.ListService(0) {
		fmt.Printf("gsi: %v\n", gsi)
	}
}
