package practice

import (
	"fmt"
	"testing"
)

func TestGoroutine(t *testing.T) {
	t.Run("goroutine으로 값 출력하기", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			go func() {
				fmt.Println(i)
			}()
		}
	})

	t.Run("goroutine 끝날때까지 기다리기", func(t *testing.T) {
		go func() {
			for i := 0; i < 10; i++ {
				fmt.Println(i)
			}
		}()
	})
}
