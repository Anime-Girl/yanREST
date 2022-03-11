package domain

type PromoMeta struct {
	Name        string
	Description string
}

type PromoMetaWithID struct {
	ID          int64
	Name        string
	Description string
}

type Promo struct {
	ID           int64
	Name         string
	Description  string
	Prizes       []Prize
	Participants []Participant
}
