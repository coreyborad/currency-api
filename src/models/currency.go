package models

type InputCurrency struct {
	InCurrency  string  `form:"in_currency" json:"in_currency" validate:"required,oneof=TWD JPY USD"`
	OutCurrency string  `form:"out_currency" json:"out_currency" validate:"required,oneof=TWD JPY USD"`
	Price       float64 `form:"price" json:"price" validate:"required,min=0"`
}
