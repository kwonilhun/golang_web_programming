package practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//github.com/google/uuid

func TestTestify(t *testing.T) {
	t.Run("일치 여부 확인", func(t *testing.T) {
		num1 := 1
		num2 := 1
		assert.Equal(t, num1, num2)
	})

	t.Run("empty 확인", func(t *testing.T) {
		assert.NotEmpty(t, "uuid")
		assert.NotEmpty(t, []int{1, 2, 3})
	})

	t.Run("숫자 값 일치 확인", func(t *testing.T) {
		num1 := 1
		num2 := int64(1)

		//assert.Equal(t, num1, num2) // FAIL
		assert.NotEqual(t, num1, num2)
		assert.EqualValues(t, num1, num2)
	})

	t.Run("숫자 검증", func(t *testing.T) {
		assert.Zero(t, 0)
		assert.Positive(t, 1)
	})

	t.Run("element 확인", func(t *testing.T) {
		type Member struct {
			Name string
			Age  int
		}
		var members1, members2 []Member

		members1 = []Member{
			{Name: "Amy", Age: 10},
			{Name: "John", Age: 20},
		}
		members2 = []Member{
			{Name: "Amy", Age: 10},
			{Name: "John", Age: 20},
		}
		assert.Equal(t, members1, members2)

		members1 = []Member{
			{Name: "Amy", Age: 10},
			{Name: "John", Age: 20},
		}
		members2 = []Member{
			{Name: "John", Age: 20},
			{Name: "Amy", Age: 10},
		}
		//assert.Equal(t, members1, members2) // FAIL
		assert.ElementsMatch(t, members1, members2)
	})

	t.Run("len 확인", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		assert.Len(t, numbers, 3)
		assert.Equal(t, 3, len(numbers))
	})
}
