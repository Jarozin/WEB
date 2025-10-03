package dto

import (
	"app/internal/service/core/dto"

	tdbdto "app/internal/tests_for_testing/unitTests/testDataBuilder/dto"
)

type BookParamsDTOObjectMother struct {
}

func NewBookParamsDTOObjectMother() *BookParamsDTOObjectMother {
	return &BookParamsDTOObjectMother{}
}

func (bmom *BookParamsDTOObjectMother) DefaultBookParams() *dto.BookParamsDTO {
	return tdbdto.NewBookParamsDTOBuilder().Build()
}
