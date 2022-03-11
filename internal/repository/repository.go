package repository

import (
	"errors"
	"sync"
	"testo/internal/domain"
)

var ErrNotFound = errors.New("not found")

type promoRepo struct {
	Promo []domain.Promo
	mu    sync.Mutex
}

func NewRepository() *promoRepo {
	return &promoRepo{}
}

func (r *promoRepo) InsertPromo(promo domain.PromoMeta) (int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.insertPromo(promo)

}

func (r *promoRepo) insertPromo(promo domain.PromoMeta) (int64, error) {
	if len(r.Promo) == 0 {
		r.Promo = append(r.Promo, domain.Promo{
			ID:          1,
			Name:        promo.Name,
			Description: promo.Description,
		})

		return 1, nil
	}

	newID := r.Promo[len(r.Promo)-1].ID + 1

	r.Promo = append(r.Promo, domain.Promo{
		ID:          newID,
		Name:        promo.Name,
		Description: promo.Description,
	})

	return newID, nil
}

func (r *promoRepo) SelectAllPromos() ([]domain.PromoMetaWithID, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.selectAllPromos()
}

func (r *promoRepo) selectAllPromos() ([]domain.PromoMetaWithID, error) {
	result := make([]domain.PromoMetaWithID, 0, len(r.Promo))

	for _, promo := range r.Promo {
		result = append(result, domain.PromoMetaWithID{
			ID:          promo.ID,
			Name:        promo.Name,
			Description: promo.Description,
		})
	}

	return result, nil
}

func (r *promoRepo) SelectPromo(id int64) (domain.Promo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.selectPromo(id)
}

func (r *promoRepo) selectPromo(id int64) (domain.Promo, error) {
	for _, promo := range r.Promo {
		if promo.ID == id {
			return promo, nil
		}
	}

	return domain.Promo{}, ErrNotFound
}

func (r *promoRepo) UpdatePromo(id int64, promo domain.PromoMeta) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.updatePromo(id, promo)
}

func (r *promoRepo) updatePromo(id int64, promo domain.PromoMeta) error {
	for i, val := range r.Promo {
		if val.ID == id {
			r.Promo[i] = domain.Promo{
				ID:           id,
				Name:         promo.Name,
				Description:  promo.Description,
				Prizes:       val.Prizes,
				Participants: val.Participants,
			}

			return nil
		}
	}

	return ErrNotFound
}

func (r *promoRepo) DeletePromo(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.deletePromo(id)
}

func (r *promoRepo) deletePromo(id int64) error {
	for i, val := range r.Promo {
		if val.ID == id {
			r.Promo = append(r.Promo[:i], r.Promo[i+1:]...)
			return nil
		}
	}

	return ErrNotFound
}

func (r *promoRepo) InsertParticipant(promoID int64, participant domain.ParticipantMeta) (int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.insertParticipant(promoID, participant)
}

func (r *promoRepo) insertParticipant(promoID int64, participant domain.ParticipantMeta) (int64, error) {
	index := -1
	for i, val := range r.Promo {
		if val.ID == promoID {
			index = i
		}
	}

	if index == -1 {
		return 0, ErrNotFound
	}

	if len(r.Promo[index].Participants) == 0 {
		r.Promo[index].Participants = append(r.Promo[index].Participants, domain.Participant{
			ID:   1,
			Name: participant.Name,
		})

		return 1, nil
	}

	newID := r.Promo[index].Participants[len(r.Promo[index].Participants)-1].ID + 1

	r.Promo[index].Participants = append(r.Promo[index].Participants, domain.Participant{
		ID:   newID,
		Name: participant.Name,
	})

	return newID, nil
}

func (r *promoRepo) DeleteParticipant(promoID, participantID int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.deleteParticipant(promoID, participantID)
}

func (r *promoRepo) deleteParticipant(promoID, participantID int64) error {
	index := -1
	for i, val := range r.Promo {
		if val.ID == promoID {
			index = i
		}
	}

	if index == -1 {
		return ErrNotFound
	}

	for i, val := range r.Promo[index].Participants {
		if val.ID == promoID {
			r.Promo[index].Participants = append(r.Promo[index].Participants[:i],
				r.Promo[index].Participants[i+1:]...)
			return nil
		}
	}

	return ErrNotFound
}

func (r *promoRepo) InsertPrize(promoID int64, prize domain.PrizeMeta) (int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.insertPrize(promoID, prize)
}

func (r *promoRepo) insertPrize(promoID int64, prize domain.PrizeMeta) (int64, error) {
	index := -1
	for i, val := range r.Promo {
		if val.ID == promoID {
			index = i
		}
	}

	if index == -1 {
		return 0, ErrNotFound
	}

	if len(r.Promo[index].Prizes) == 0 {
		r.Promo[index].Prizes = append(r.Promo[index].Prizes, domain.Prize{
			ID:          1,
			Description: prize.Description,
		})

		return 1, nil
	}

	newID := r.Promo[index].Participants[len(r.Promo[index].Participants)-1].ID + 1

	r.Promo[index].Prizes = append(r.Promo[index].Prizes, domain.Prize{
		ID:          newID,
		Description: prize.Description,
	})

	return newID, nil
}

func (r *promoRepo) DeletePrize(promoID, prizeID int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.deletePrize(promoID, prizeID)
}

func (r *promoRepo) deletePrize(promoID, prizeID int64) error {
	index := -1
	for i, val := range r.Promo {
		if val.ID == promoID {
			index = i
		}
	}

	if index == -1 {
		return ErrNotFound
	}

	for i, val := range r.Promo[index].Prizes {
		if val.ID == promoID {
			r.Promo[index].Prizes = append(r.Promo[index].Prizes[:i],
				r.Promo[index].Prizes[i+1:]...)
			return nil
		}
	}

	return ErrNotFound
}
