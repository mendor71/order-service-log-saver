package domain

type Position struct {
	Id            int64  `json:"id"`
	PosCategoryId int64  `json:"posCategoryId"`
	PosName       string `json:"posName"`
	PosStatusId   int64  `json:"posStatusId"`
}

type SalePoint struct {
	Id         int    `json:"id"`
	SpStatusId int    `json:"spStatusId"`
	SpName     string `json:"spName"`
}

type Status struct {
	Id           int64  `json:"id"`
	StEntityType string `json:"stEntityType"`
	StName       string `json:"stName"`
}

type TransferOrder struct {
	SalePoint SalePoint  `json:"salePoint"`
	Status    Status     `json:"status"`
	Positions []Position `json:"positions"`
}
