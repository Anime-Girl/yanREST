package service

import (
	"errors"
	"math/rand"
	"testo/internal/domain"
	"time"
)

var ErrRaffle = errors.New("raffle error")

type (
	Repository interface {
		Promo
		Participant
		Prize
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
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) InsertPromo(promo domain.PromoMeta) (int64, error) {
	return s.repo.InsertPromo(promo)
}

func (s *service) SelectAllPromos() ([]domain.PromoMetaWithID, error) {
	return s.repo.SelectAllPromos()
}

func (s *service) SelectPromo(id int64) (domain.Promo, error) {
	return s.repo.SelectPromo(id)
}

func (s *service) UpdatePromo(id int64, promo domain.PromoMeta) error {
	if promo.Name == "" {
		return errors.New("cant delete name")
	}
	return s.repo.UpdatePromo(id, promo)
}

func (s *service) DeletePromo(id int64) error {
	return s.repo.DeletePromo(id)
}

func (s *service) InsertParticipant(promoID int64, participant domain.ParticipantMeta) (int64, error) {
	return s.repo.InsertParticipant(promoID, participant)
}

func (s *service) DeleteParticipant(promoID, participantID int64) error {
	return s.repo.DeleteParticipant(promoID, participantID)
}

func (s *service) InsertPrize(promoID int64, prize domain.PrizeMeta) (int64, error) {
	return s.repo.InsertPrize(promoID, prize)
}

func (s *service) DeletePrize(promoID, prizeID int64) error {
	return s.repo.DeletePrize(promoID, prizeID)
}

func (s *service) SelectWinners(promoID int64) ([]domain.Result, error) {
	promo, err := s.repo.SelectPromo(promoID)
	if err != nil {
		return nil, err
	}

	if len(promo.Prizes) == len(promo.Participants) {
		return nil, ErrRaffle
	}

	indexes := make([]int, len(promo.Prizes))
	for i := range promo.Prizes {
		indexes = append(indexes, i)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(indexes), func(i, j int) {
		indexes[i], indexes[j] = indexes[j], indexes[i]
	})

	result := make([]domain.Result, len(promo.Prizes))

	for i := range indexes {
		result = append(result, domain.Result{
			Winner: promo.Participants[i],
			Prize:  promo.Prizes[i],
		})
	}

	return result, nil
}
