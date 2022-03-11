package controller

type (
	getPromoResponse struct {
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	getPromoByIDResponse struct {
		ID           int64         `json:"id"`
		Name         string        `json:"name"`
		Description  string        `json:"description"`
		Prizes       []prize       `json:"prizes"`
		Participants []participant `json:"participants"`
	}

	prize struct {
		ID          int64  `json:"id"`
		Description string `json:"description"`
	}

	participant struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	postPromoRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	postPromoResponse struct {
		ID int64 `json:"id"`
	}

	postParticipantByPromoIDRequest struct {
		Name string `json:"name"`
	}

	postParticipantByPromoIDResponse struct {
		ID int64 `json:"id"`
	}

	postPrizeByPromoIDRequest struct {
		Description string `json:"description"`
	}

	postPrizeByPromoIDResponse struct {
		ID int64 `json:"id"`
	}

	postRaffleResponse struct {
		Winner participant `json:"winner"`
		Prize  prize       `json:"prize"`
	}

	putPromoByIDRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
)
