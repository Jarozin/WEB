package models

import (
	"app/internal/service/core/models"

	tdbmodels "app/internal/tests_for_testing/unitTests/testDataBuilder/models"
)

type ReaderModelObjectMother struct {
}

func NewReaderModelObjectMother() *ReaderModelObjectMother {
	return &ReaderModelObjectMother{}
}

func (rmom *ReaderModelObjectMother) DefaultReader() *models.ReaderModel {
	return tdbmodels.NewReaderModelBuilder().Build()
}
