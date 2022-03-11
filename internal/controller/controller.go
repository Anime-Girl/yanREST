package controller

import (
	"testo/internal/domain"

	"github.com/gin-gonic/gin"
)

type (
	Service interface {
		Promo
		Participant
		Prize
		Raffle
	}

	Promo interface {
		InsertPromo(promo domain.PromoMeta) (int64, error)

		SelectAllPromos() ([]domain.PromoMetaWithID, error)
		SelectPromo(id int64) (domain.Promo, error)

		UpdatePromo(id int64, promo domain.PromoMeta) error

		DeletePromo(id int64) error
	}

	Participant interface {
		InsertParticipant(promoID int64, participant domain.ParticipantMeta) (int64, error)

		DeleteParticipant(promoID, participantID int64) error
	}

	Prize interface {
		InsertPrize(promoID int64, prize domain.PrizeMeta) (int64, error)

		DeletePrize(promoID, prizeID int64) error
	}

	Raffle interface {
		SelectWinners(promoID int64) ([]domain.Result, error)
	}
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) InitRoutes() *gin.Engine {
	router := gin.Default()

	promoGroup := router.Group("/promo")
	{
		promoGroup.POST("", c.setPromo)
		promoGroup.POST("/:id/participant", c.setParticipant)
		promoGroup.POST("/:id/prize", c.setPrize)
		promoGroup.POST("/:id/raffle", c.raffle)

		promoGroup.GET("/:id", c.getPromo)
		promoGroup.GET("", c.getPromos)

		promoGroup.PUT("/:id", c.putPromo)

		promoGroup.DELETE("/:promoId/prize/:prizeId", c.deletePrize)
		promoGroup.DELETE("/:promoId/participant/:participantId", c.deleteParticipant)
		//promoGroup.DELETE("/:id", c.deletePromo)
	}

	return router
}
