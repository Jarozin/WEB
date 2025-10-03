package dto

import (
	"app/internal/service/core/dto"

	tdbdto "app/internal/tests_for_testing/unitTests/testDataBuilder/dto"
)

type ReaderSignInDTOObjectMother struct {
}

func NewReaderSignInDTOObjectMother() *ReaderSignInDTOObjectMother {
	return &ReaderSignInDTOObjectMother{}
}

func (rdom *ReaderSignInDTOObjectMother) DefaultReaderSignInDTO() *dto.SignInInputDTO {
	return tdbdto.NewReaderSignInDTOBuilder().Build()
}
