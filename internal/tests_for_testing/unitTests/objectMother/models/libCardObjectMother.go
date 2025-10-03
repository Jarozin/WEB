package models

import (
	"app/internal/service/core/models"
	"time"

	tdbmodels "app/internal/tests_for_testing/unitTests/testDataBuilder/models"
)

type LibCardModelObjectMother struct {
}

func NewLibCardModelObjectMother() *LibCardModelObjectMother {
	return &LibCardModelObjectMother{}
}

func (lcmom *LibCardModelObjectMother) DefaultLibCard() *models.LibCardModel {
	return tdbmodels.NewLibCardModelBuilder().Build()
}

func (lcmom *LibCardModelObjectMother) ExpiredLibCard() *models.LibCardModel {
	return tdbmodels.NewLibCardModelBuilder().
		WithIssueDate(time.Now().AddDate(0, 0, -370)).
		WithActionStatus(false).
		Build()
}
