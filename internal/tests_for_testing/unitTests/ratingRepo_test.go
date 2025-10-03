package unitTests

import (
	"app/internal/service/core/models"
	"app/internal/service/impl"
	"app/internal/service/intfRepo"
	"context"
	"errors"
	"testing"

	implRepo "app/internal/repo/impl"
	ommodels "app/internal/tests_for_testing/unitTests/objectMother/models"

	"app/pkg/logging"

	"github.com/jmoiron/sqlx"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

type RatingRepoTestsSuite struct {
	suite.Suite
}

func (rrts *RatingRepoTestsSuite) BeforeEach(t provider.T) {
	t.Tags("BookSmart", "rating repo", "suite", "steps")
}

func (rrts *RatingRepoTestsSuite) Test_Create_Success(t provider.T) {
	var (
		ratingRepo intfRepo.IRatingRepo
		rating     *models.RatingModel
		mock       sqlxmock.Sqlmock
		db         *sqlx.DB
		err        error
	)

	t.Title("Test Create Rating Success")
	t.Description("The new rating was successfully created")

	t.WithNewStep("Arrange", func(sCtx provider.StepCtx) {
		if db, mock, err = sqlxmock.Newx(); err != nil {
			t.Fatal(err)
		}

		rating = ommodels.NewRatingModelObjectMother().DefaultRating()
		ratingRepo = implRepo.NewRatingRepo(db, logging.GetLoggerForTests())

		mock.ExpectExec(`insert into bs.rating values`).
			WithArgs(rating.ID, rating.ReaderID, rating.BookID,
				rating.Review, rating.Rating).
			WillReturnResult(sqlxmock.NewResult(1, 1))
	})

	t.WithNewStep("Act", func(sCtx provider.StepCtx) {
		err = ratingRepo.Create(context.Background(), rating)
	})

	t.WithNewStep("Assert", func(sCtx provider.StepCtx) {
		t.Assert().Nil(err)
	})
}

func (rrts *RatingRepoTestsSuite) Test_Create_ErrorExecutionQuery(t provider.T) {
	var (
		ratingRepo intfRepo.IRatingRepo
		rating     *models.RatingModel
		mock       sqlxmock.Sqlmock
		db         *sqlx.DB
		err        error
	)

	t.Title("Test Create Rating Error: execution query error")
	t.Description("Execution query error")

	t.WithNewStep("Arrange", func(sCtx provider.StepCtx) {
		if db, mock, err = sqlxmock.Newx(); err != nil {
			t.Fatal(err)
		}

		rating = ommodels.NewRatingModelObjectMother().DefaultRating()
		ratingRepo = implRepo.NewRatingRepo(db, logging.GetLoggerForTests())

		mock.ExpectExec(`insert into bs.rating values`).
			WithArgs(rating.ID, rating.ReaderID, rating.BookID,
				rating.Review, rating.Rating).
			WillReturnError(errors.New("insert error"))
	})

	t.WithNewStep("Act", func(sCtx provider.StepCtx) {
		err = ratingRepo.Create(context.Background(), rating)
	})

	t.WithNewStep("Assert", func(sCtx provider.StepCtx) {
		t.Assert().NotNil(err)
		t.Assert().Equal(errors.New("insert error"), err)
		err = mock.ExpectationsWereMet()
		t.Assert().NoError(err)
	})
}

func (rrts *RatingRepoTestsSuite) Test_GetByReaderAndBook_Success(t provider.T) {
	var (
		ratingRepo intfRepo.IRatingRepo
		rating     *models.RatingModel
		findRating *models.RatingModel
		mock       sqlxmock.Sqlmock
		db         *sqlx.DB
		err        error
	)

	t.Title("Test Get Rating By Reader And Book Success")
	t.Description("Rating was successfully retrieved by reader and book")

	t.WithNewStep("Arrange", func(sCtx provider.StepCtx) {
		if db, mock, err = sqlxmock.Newx(); err != nil {
			t.Fatal(err)
		}

		rating = ommodels.NewRatingModelObjectMother().DefaultRating()
		ratingRepo = implRepo.NewRatingRepo(db, logging.GetLoggerForTests())

		rows := sqlxmock.NewRows([]string{"id", "reader_id", "book_id", "review", "rating"}).
			AddRow(rating.ID, rating.ReaderID, rating.BookID, rating.Review, rating.Rating)

		mock.ExpectQuery(`select (.+) from bs.rating where (.+)`).
			WithArgs(rating.ReaderID, rating.BookID).WillReturnRows(rows)
	})

	t.WithNewStep("Act", func(sCtx provider.StepCtx) {
		findRating, err = ratingRepo.GetByReaderAndBook(context.Background(), rating.ReaderID, rating.BookID)
	})

	t.WithNewStep("Assert", func(sCtx provider.StepCtx) {
		t.Assert().Nil(err)
		t.Assert().Equal(rating, findRating)
	})
}

func (rrts *RatingRepoTestsSuite) Test_GetByReaderAndBook_ErrorExecutionQuery(t provider.T) {
	var (
		ratingRepo intfRepo.IRatingRepo
		rating     *models.RatingModel
		mock       sqlxmock.Sqlmock
		findRating *models.RatingModel
		db         *sqlx.DB
		err        error
	)

	t.Title("Test Get Rating By Reader And Book Error: execution query error")
	t.Description("Error executing query")

	t.WithNewStep("Arrange", func(sCtx provider.StepCtx) {
		if db, mock, err = sqlxmock.Newx(); err != nil {
			t.Fatal(err)
		}

		rating = ommodels.NewRatingModelObjectMother().DefaultRating()
		ratingRepo = implRepo.NewRatingRepo(db, logging.GetLoggerForTests())

		mock.ExpectQuery(`select (.+) from bs.rating where (.+)`).
			WithArgs(rating.ReaderID, rating.BookID).WillReturnError(errors.New("select error"))
	})

	t.WithNewStep("Act", func(sCtx provider.StepCtx) {
		findRating, err = ratingRepo.GetByReaderAndBook(context.Background(), rating.ReaderID, rating.BookID)
	})

	t.WithNewStep("Assert", func(sCtx provider.StepCtx) {
		t.Assert().NotNil(err)
		t.Assert().Equal(errors.New("select error"), err)
		t.Assert().Nil(findRating)
		err = mock.ExpectationsWereMet()
		t.Assert().NoError(err)
	})
}

func (rrts *RatingRepoTestsSuite) Test_GetByBookID_Success(t provider.T) {
	var (
		ratingRepo  intfRepo.IRatingRepo
		rating      *models.RatingModel
		findRatings []*models.RatingModel
		mock        sqlxmock.Sqlmock
		db          *sqlx.DB
		err         error
	)

	t.Title("Test Get Rating By Book ID Success")
	t.Description("Rating was successfully retrieved by book ID")

	t.WithNewStep("Arrange", func(sCtx provider.StepCtx) {
		if db, mock, err = sqlxmock.Newx(); err != nil {
			t.Fatal(err)
		}

		rating = ommodels.NewRatingModelObjectMother().DefaultRating()
		ratingRepo = implRepo.NewRatingRepo(db, logging.GetLoggerForTests())

		rows := sqlxmock.NewRows([]string{"id", "reader_id", "book_id", "review", "rating"}).
			AddRow(rating.ID, rating.ReaderID, rating.BookID, rating.Review, rating.Rating)

		mock.ExpectQuery(`select (.+) from bs.rating where (.+)`).
			WithArgs(rating.BookID, impl.ReviewsPageLimit, 0).WillReturnRows(rows)
	})

	t.WithNewStep("Act", func(sCtx provider.StepCtx) {
		findRatings, err = ratingRepo.GetByBookID(context.Background(), rating.BookID, impl.ReviewsPageLimit, 0)
	})

	t.WithNewStep("Assert", func(sCtx provider.StepCtx) {
		t.Assert().Nil(err)
		t.Assert().Equal([]*models.RatingModel{rating}, findRatings)
	})
}

func (rrts *RatingRepoTestsSuite) Test_GetByBookID_ErrorExecutionQuery(t provider.T) {
	var (
		ratingRepo  intfRepo.IRatingRepo
		rating      *models.RatingModel
		mock        sqlxmock.Sqlmock
		findRatings []*models.RatingModel
		db          *sqlx.DB
		err         error
	)

	t.Title("Test Get Rating By Reader And Book Error: execution query error")
	t.Description("Error executing query")

	t.WithNewStep("Arrange", func(sCtx provider.StepCtx) {
		if db, mock, err = sqlxmock.Newx(); err != nil {
			t.Fatal(err)
		}

		rating = ommodels.NewRatingModelObjectMother().DefaultRating()
		ratingRepo = implRepo.NewRatingRepo(db, logging.GetLoggerForTests())

		mock.ExpectQuery(`select (.+) from bs.rating where (.+)`).
			WithArgs(rating.BookID, impl.ReviewsPageLimit, 0).WillReturnError(errors.New("select error"))
	})

	t.WithNewStep("Act", func(sCtx provider.StepCtx) {
		findRatings, err = ratingRepo.GetByBookID(context.Background(), rating.BookID, impl.ReviewsPageLimit, 0)
	})

	t.WithNewStep("Assert", func(sCtx provider.StepCtx) {
		t.Assert().NotNil(err)
		t.Assert().Equal(errors.New("select error"), err)
		t.Assert().Nil(findRatings)
		err = mock.ExpectationsWereMet()
		t.Assert().NoError(err)
	})
}

func TestRatingRepoTestsSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(RatingRepoTestsSuite))
}
