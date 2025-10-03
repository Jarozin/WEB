package models

import (
	"app/internal/service/core/models"
	"time"

	tdbmodels "app/internal/tests_for_testing/unitTests/testDataBuilder/models"
)

type ReservationModelObjectMother struct {
}

func NewReservationModelObjectMother() *ReaderModelObjectMother {
	return &ReaderModelObjectMother{}
}

func (rmom *ReaderModelObjectMother) DefaultReservation() *models.ReservationModel {
	return tdbmodels.NewReservationModelBuilder().Build()
}

func (rmom *ReaderModelObjectMother) ExpiredReservation() *models.ReservationModel {
	return tdbmodels.NewReservationModelBuilder().
		WithReturnDate(time.Now().AddDate(0, -1, 0)).
		Build()
}
