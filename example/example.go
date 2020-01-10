package main

import (
	"fmt"
	"github.com/fs714/kmutex"
	"time"
)

func main() {
	key1 := "key1"
	key2 := "key2"

	km := kmutex.New()

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println("--- Task 1.1 Start ---")
			km.Lock(key1)

			time.Sleep(1 * time.Second)

			km.Unlock(key1)

			fmt.Println("--- Task 1.1 Finished ---")
		}()

		go func() {
			fmt.Println("--- Task 1.2 Start ---")
			km.Lock(key1)

			time.Sleep(2 * time.Second)

			km.Unlock(key1)

			fmt.Println("--- Task 1.2 Finished ---")
		}()

		go func() {
			fmt.Println("--- Task 1.3 Start ---")
			km.Lock(key1)

			time.Sleep(3 * time.Second)

			km.Unlock(key1)

			fmt.Println("--- Task 1.3 Finished ---")
		}()

		go func() {
			fmt.Println("--- Task 2.1 Start ---")
			km.Lock(key2)

			time.Sleep(1 * time.Second)

			km.Unlock(key2)

			fmt.Println("--- Task 2.1 Finished ---")
		}()

		go func() {
			fmt.Println("--- Task 2.2 Start ---")
			km.Lock(key2)

			time.Sleep(2 * time.Second)

			km.Unlock(key2)

			fmt.Println("--- Task 2.2 Finished ---")
		}()

		go func() {
			fmt.Println("--- Task 2.3 Start ---")
			km.Lock(key2)

			time.Sleep(3 * time.Second)

			km.Unlock(key2)

			fmt.Println("--- Task 2.3 Finished ---")
		}()

		time.Sleep(10 * time.Second)
		fmt.Println("==================")
	}
}
