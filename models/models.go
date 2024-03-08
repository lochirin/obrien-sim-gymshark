package models

type OrderSpecification struct {
	Sizes    []int `json:"sizes"`
	Capacity int   `json:"capacity"`
}
