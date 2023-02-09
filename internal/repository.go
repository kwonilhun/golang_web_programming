package internal

import (
	"errors"
)

var ErrNotFoundMembership = errors.New("not found membership")

type Repository struct {
	data map[string]Membership
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}

// 데이터 삽입, 삭제, 검색 부분은 repository에 넣는다.

func (r *Repository) Create(membership Membership) {
	r.data[membership.UserName] = membership
}

func (r *Repository) GetById(id string) (Membership, error) {
	for _, membership := range r.data {
		if membership.ID == id {
			return membership, nil
		}
	}
	return Membership{}, ErrNotFoundMembership
}

func (r *Repository) UpdateById(membership Membership) (Membership, error) {
	_, exist := r.data[membership.ID]
	if !exist {
		return Membership{}, errors.New("데이터가 없습니다.")
	}
	r.data[membership.ID] = membership
	return membership, nil
}

func (r *Repository) deleteById(id string) error {
	_, exist := r.data[id]
	if !exist {
		return errors.New("삭제 실패")
	}
	delete(r.data, id)
	return nil
}
