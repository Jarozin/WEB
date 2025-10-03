package models

import (
	"app/internal/service/core/models"

	tdbmodels "app/internal/tests_for_testing/unitTests/testDataBuilder/models"
)

type BookModelObjectMother struct {
}

func NewBookModelObjectMother() *BookModelObjectMother {
	return &BookModelObjectMother{}
}

func (bmom *BookModelObjectMother) DefaultBook() *models.BookModel {
	return tdbmodels.NewBookModelBuilder().Build()
}
