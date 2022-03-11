package controller

import (
	"net/http"
	"strconv"
	"testo/internal/domain"
	"testo/internal/service"

	"github.com/gin-gonic/gin"
)

func (co *Controller) setPromo(c *gin.Context) {
	var entity postPromoRequest

	if err := c.BindJSON(&entity); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, err := co.service.InsertPromo(domain.PromoMeta{
		Name:        entity.Name,
		Description: entity.Description,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, postPromoResponse{
		ID: id,
	})
}

func (co *Controller) setParticipant(c *gin.Context) {
	var entity postParticipantByPromoIDRequest

	if err := c.BindJSON(&entity); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	requestPromoID, ok := c.Get("id")
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	promoIDstr, ok := requestPromoID.(string)
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	promoID, err := strconv.ParseInt(promoIDstr, 10, 0)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, err := co.service.InsertParticipant(promoID, domain.ParticipantMeta{
		Name: entity.Name,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, postPromoResponse{
		ID: id,
	})
}

func (co *Controller) setPrize(c *gin.Context) {
	var entity postPrizeByPromoIDRequest

	if err := c.BindJSON(&entity); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	requestPromoID, ok := c.Get("id")
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	promoIDstr, ok := requestPromoID.(string)
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	promoID, err := strconv.ParseInt(promoIDstr, 10, 0)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, err := co.service.InsertPrize(promoID, domain.PrizeMeta{
		Description: entity.Description,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, postPromoResponse{
		ID: id,
	})
}

func (co *Controller) raffle(c *gin.Context) {
	requestPromoID, ok := c.Get("id")
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	promoIDstr, ok := requestPromoID.(string)
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	promoID, err := strconv.ParseInt(promoIDstr, 10, 0)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	winners, err := co.service.SelectWinners(promoID)
	if err == service.ErrRaffle {
		c.AbortWithStatus(http.StatusConflict)
		return
	}
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	result := make([]postRaffleResponse, 0, len(winners))
	for _, val := range winners {
		result = append(result, postRaffleResponse{
			Winner: participant(val.Winner),
			Prize:  prize(val.Prize),
		})
	}

	c.JSON(http.StatusOK, result)
}

func (co *Controller) getPromo(c *gin.Context) {
	requestPromoID, ok := c.Get("id")
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	promoIDstr, ok := requestPromoID.(string)
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	promoID, err := strconv.ParseInt(promoIDstr, 10, 0)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	promo, err := co.service.SelectPromo(promoID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	prizes := make([]prize, 0, len(promo.Prizes))
	participants := make([]participant, 0, len(promo.Participants))

	for _, val := range promo.Prizes {
		prizes = append(prizes, prize{
			ID:          val.ID,
			Description: val.Description,
		})
	}

	for _, val := range promo.Participants {
		participants = append(participants, participant{
			ID:   val.ID,
			Name: val.Name,
		})
	}

	c.JSON(http.StatusOK, getPromoByIDResponse{
		ID:           promo.ID,
		Name:         promo.Name,
		Description:  promo.Description,
		Prizes:       prizes,
		Participants: participants,
	})
}

func (co *Controller) getPromos(c *gin.Context) {
	promos, err := co.service.SelectAllPromos()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	result := make([]getPromoResponse, 0, len(promos))
	for _, promo := range promos {
		result = append(result, getPromoResponse{
			ID:          promo.ID,
			Name:        promo.Name,
			Description: promo.Description,
		})
	}

	c.JSON(http.StatusOK, result)
}

func (co *Controller) putPromo(c *gin.Context) {
	var entity putPromoByIDRequest

	if err := c.BindJSON(&entity); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	requestPromoID, ok := c.Get("id")
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	promoIDstr, ok := requestPromoID.(string)
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	promoID, err := strconv.ParseInt(promoIDstr, 10, 0)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = co.service.UpdatePromo(promoID, domain.PromoMeta{
		Name:        entity.Name,
		Description: entity.Description,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

func (co *Controller) deletePromo(c *gin.Context) {
	requestPromoID, ok := c.Get("id")
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	promoIDstr, ok := requestPromoID.(string)
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	promoID, err := strconv.ParseInt(promoIDstr, 10, 0)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = co.service.DeletePromo(promoID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

func (co *Controller) deleteParticipant(c *gin.Context) {
	requestPromoID, ok := c.Get("promoId")
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	promoIDstr, ok := requestPromoID.(string)
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	promoID, err := strconv.ParseInt(promoIDstr, 10, 0)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	requestParticipantID, ok := c.Get("participantId")
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	participantIDstr, ok := requestParticipantID.(string)
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	participantID, err := strconv.ParseInt(participantIDstr, 10, 0)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = co.service.DeleteParticipant(promoID, participantID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

func (co *Controller) deletePrize(c *gin.Context) {
	requestPromoID, ok := c.Get("promoId")
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	promoIDstr, ok := requestPromoID.(string)
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	promoID, err := strconv.ParseInt(promoIDstr, 10, 0)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	requestPrizeID, ok := c.Get("prizeId")
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	prizeIDstr, ok := requestPrizeID.(string)
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	prizeID, err := strconv.ParseInt(prizeIDstr, 10, 0)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = co.service.DeletePrize(promoID, prizeID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
