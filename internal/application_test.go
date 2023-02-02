package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMembership(t *testing.T) {
	t.Run("멤버십을 생성한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, req.MembershipType, res.MembershipType)
	})

	t.Run("이미 등록된 사용자 이름이 존재할 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		req2 := CreateRequest{"jenny", "naver"}
		app.Create(req)
		_, err2 := app.Create(req2)

		assert.Errorf(t, err2, "이미 등록된 사용자 이름이 존재할 경우 실패한다.")
	})

	t.Run("사용자 이름을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"", "naver"}
		_, err := app.Create(req)
		assert.Errorf(t, err, "사용자 이름을 입력하지 않은 경우 실패한다.")
	})

	t.Run("멤버십 타입을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", ""}
		_, err := app.Create(req)
		assert.Errorf(t, err, "멤버십 타입 입력하지 않아 에러 발생")
	})

	t.Run("naver/toss/payco 이외의 타입을 입력한 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "google"}
		_, err := app.Create(req)
		assert.Errorf(t, err, "naver/toss/payco 이외의 타입을 입력한 경우 실패한다.")
	})
}

func TestUpdate(t *testing.T) {
	t.Run("멤버십 정보를 갱신한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, _ := app.Create(req)
		req2 := UpdateRequest{res.ID, "jenny", "toss"}
		res2, err2 := app.Update(req2)
		assert.Equal(t, res.ID, res2.ID)
		assert.Equal(t, req2.MembershipType, "toss")
		assert.Nil(t, err2)
	})

	t.Run("수정하려는 사용자의 이름이 이미 존재하는 사용자 이름이라면 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req1 := CreateRequest{"jenny", "naver"}
		req2 := CreateRequest{"elsa", "naver"}
		res1, _ := app.Create(req1)
		app.Create(req2)

		req3 := UpdateRequest{res1.ID, "elsa", "naver"}
		_, err2 := app.Update(req3)
		assert.Errorf(t, err2, "sss")
	})

	t.Run("멤버십 아이디를 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req1 := CreateRequest{"jenny", "naver"}
		req2 := CreateRequest{"elsa", "naver"}
		app.Create(req1)
		app.Create(req2)

		req3 := UpdateRequest{"", "elsa", "naver"}
		_, err2 := app.Update(req3)
		assert.Errorf(t, err2, "멤버십 아이디를 입력하지 않은 경우, 예외 처리한다.")
	})

	t.Run("사용자 이름을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req1 := CreateRequest{"jenny", "naver"}
		req2 := CreateRequest{"elsa", "naver"}
		res1, _ := app.Create(req1)
		app.Create(req2)

		req3 := UpdateRequest{res1.ID, "", "naver"}
		_, err2 := app.Update(req3)
		assert.Errorf(t, err2, "사용자 이름을 입력하지 않은 경우, 예외 처리한다.")
	})

	t.Run("멤버쉽 타입을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req1 := CreateRequest{"jenny", "naver"}
		req2 := CreateRequest{"elsa", "naver"}
		res1, _ := app.Create(req1)
		app.Create(req2)

		req3 := UpdateRequest{res1.ID, "peter", ""}
		_, err2 := app.Update(req3)
		assert.Errorf(t, err2, "멤버쉽 타입을 입력하지 않은 경우, 예외 처리한다.")
	})

	t.Run("주어진 멤버쉽 타입이 아닌 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req1 := CreateRequest{"jenny", "naver"}
		req2 := CreateRequest{"elsa", "naver"}
		res1, _ := app.Create(req1)
		app.Create(req2)

		req3 := UpdateRequest{res1.ID, "peter", "google"}
		_, err2 := app.Update(req3)
		assert.Errorf(t, err2, "멤버쉽 타입을 입력하지 않은 경우, 예외 처리한다.")
	})
}

func TestDelete(t *testing.T) {
	t.Run("멤버십을 삭제한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req1 := CreateRequest{"jenny", "naver"}
		res1, _ := app.Create(req1)

		err := app.Delete(res1.ID)
		assert.Nil(t, err)
	})

	t.Run("id를 입력하지 않았을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req1 := CreateRequest{"jenny", "naver"}
		app.Create(req1)

		err := app.Delete("")
		assert.Errorf(t, err, "id를 입력하지 않았을 때 예외 처리한다.")
	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req1 := CreateRequest{"jenny", "naver"}
		app.Create(req1)

		err := app.Delete("qwwer")
		assert.Errorf(t, err, "입력한 id가 존재하지 않을 때 예외 처리한다.")
	})
}
