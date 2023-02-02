package internal

import (
	"errors"
	"github.com/google/uuid"
	"strings"
)

const (
	NAVER = "naver"
	TOSS  = "toss"
	PAYCO = "payco"
)

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

func checkName(name string) error {
	if len(strings.TrimSpace(name)) == 0 {
		return errors.New("이름을 입력하지 않았습니다.")
	}
	return nil
}

func checkType(memberType string) error {
	if len(strings.TrimSpace(memberType)) == 0 {
		return errors.New("타입을 입력하지 않았습니다.")
	}
	if memberType != NAVER && memberType != TOSS && memberType != PAYCO {
		return errors.New("타입이 맞지 않습니다.")
	}
	return nil
}

func checkRequest(memberType string, name string) error {
	nameError := checkName(name)
	typeError := checkType(memberType)

	switch {
	case nameError != nil && typeError != nil:
		return errors.New("이름과 타입을 다시 확인해주세요")
	case nameError != nil && typeError == nil:
		return nameError
	case nameError == nil && typeError != nil:
		return typeError
	default:
		return nil
	}
}

func checkMembership(memberType string, name string, memberId string) error {
	checkRequestWithOutId := checkRequest(memberType, name)
	switch {
	case checkRequestWithOutId != nil:
		return checkRequestWithOutId
	case len(strings.TrimSpace(memberId)) == 0:
		return errors.New("id를 입력해주세요")
	default:
		return nil
	}
}

func (app *Application) Create(request CreateRequest) (CreateResponse, error) {

	requestName := request.UserName
	requestType := request.MembershipType

	requestError := checkRequest(requestType, requestName)

	if requestError != nil {
		return CreateResponse{}, requestError
	}
	for _, value := range app.repository.data {
		if value.UserName == requestName {
			return CreateResponse{}, errors.New("이미 등록되었습니다.")
		}
	}
	id, _ := uuid.NewUUID()
	memberInfo := Membership{ID: id.String(), UserName: requestName, MembershipType: requestType}
	app.repository.data[id.String()] = memberInfo
	return CreateResponse{id.String(), request.MembershipType}, nil
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	requestName := request.UserName
	requestType := request.MembershipType
	requestId := request.ID
	requestError := checkMembership(requestType, requestName, requestId)

	if requestError != nil {
		return UpdateResponse{}, requestError
	}
	_, exist := app.repository.data[requestId]
	if !exist {
		return UpdateResponse{}, errors.New("변경할 데이터가 없습니다.")
	} else {
		for _, value := range app.repository.data {
			if value.UserName == requestName && value.ID != requestId {
				return UpdateResponse{}, errors.New("변경할 이름이 이미 등록되어 있습니다")
			}
		}
		m := app.repository.data[request.ID]
		m.UserName = requestName
		m.MembershipType = requestType
		app.repository.data[request.ID] = m
		return UpdateResponse{requestId, requestName, requestType}, nil
	}
}

func (app *Application) Delete(id string) error {
	_, exist := app.repository.data[id]
	if !exist {
		return errors.New("id가 유효하지 않습니다.")
	}
	delete(app.repository.data, id)
	return nil
}
