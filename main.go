package main

import (
	"backend-food/internal/api/router"
)

func main() {
	// go func() {
	// 	for i := 0; i < 1000; i++ {
	// 		fmt.Println(i)
	// 		time.Sleep(1000 * time.Millisecond)
	// 	}

	// }()
	r := router.NewRouter()
	r.Engine.Run()

}
