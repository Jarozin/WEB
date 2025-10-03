package models

import (
	"app/internal/service/core/models"

	tdbmodels "app/internal/tests_for_testing/unitTests/testDataBuilder/models"
)

type RatingModelObjectMother struct {
}

func NewRatingModelObjectMother() *RatingModelObjectMother {
	return &RatingModelObjectMother{}
}

func (rmom *RatingModelObjectMother) DefaultRating() *models.RatingModel {
	return tdbmodels.NewRatingModelBuilder().Build()
}
