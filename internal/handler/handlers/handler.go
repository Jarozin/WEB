package handlers

import (
	"app/internal/service/intf"
	"app/internal/service/pkg/auth"
	"app/internal/service/pkg/hash"
	"io"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	bookService        intf.IBookService
	libCardService     intf.ILibCardService
	readerService      intf.IReaderService
	reservationService intf.IReservationService
	ratingService      intf.IRatingService
	tokenManager       auth.ITokenManager
	hasher             hash.IPasswordHasher
	accessTokenTTL     time.Duration
	refreshTokenTTL    time.Duration
}

func NewHandler(
	bookService intf.IBookService,
	libCardService intf.ILibCardService,
	readerService intf.IReaderService,
	reservationService intf.IReservationService,
	ratingService intf.IRatingService,
	tokenManager auth.ITokenManager,
	accessTokenTTL time.Duration,
	refreshTokenTTL time.Duration,
) *Handler {
	return &Handler{
		bookService:        bookService,
		libCardService:     libCardService,
		readerService:      readerService,
		reservationService: reservationService,
		ratingService:      ratingService,
		tokenManager:       tokenManager,
		accessTokenTTL:     accessTokenTTL,
		refreshTokenTTL:    refreshTokenTTL,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	router := gin.Default()

	router.Use(h.corsSettings())

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/auth/sign-up", h.signUp)
			v1.POST("/auth/sign-in", h.signIn)
			v1.POST("/auth/refresh", h.refresh)

			v1.GET("/books", h.getPageBooks)
			v1.GET("/books/:id", h.getBookByID)

			v1.GET("/books/:id/ratings/avg", h.getAvgRatingByBookID)
			v1.GET("/books/:id/ratings", h.getRatingsByBookID)

			registered := v1.Group("/", h.readerIdentity)
			{
				registered.POST("/books/:id/ratings", h.addNewRating)

				registered.GET("/readers/:id", h.getReaderByID)
				registered.POST("/readers/:id/favorite_books", h.addToFavorites)

				registered.GET("/readers/:id/lib_cards", h.getLibCardByReaderID)
				registered.PUT("/readers/:id/lib_cards", h.updateLibCard)
				registered.POST("/readers/:id/lib_cards", h.createLibCard)

				registered.POST("/readers/:id/reservations", h.reserveBook)
				registered.GET("/readers/:id/reservations", h.getReservationsByReaderID)
				registered.GET("/readers/:id/reservations/:reservation_id", h.getReservationByID)
				registered.PATCH("/readers/:id/reservations/:reservation_id", h.updateReservation)
			}
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

func (h *Handler) corsSettings() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowOrigins: []string{
			"*",
		},
		AllowCredentials: true,
		AllowHeaders: []string{
			"Authorization",
			"Content-Type",
		},
		ExposeHeaders: []string{
			"Content-Type",
		},
	})
}

// убрал
//authenticate := router.Group("/auth")
//{
//authenticate.POST("/sign-up", h.signUp)
//authenticate.POST("/sign-in", h.signIn)
//authenticate.POST("/admin/sign-in", h.signInAsAdmin)
//authenticate.POST("/refresh", h.refresh)
//}
// _ "github.com/nikitalystsev/BookSmart/docs_swagger"
