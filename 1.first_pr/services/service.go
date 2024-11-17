package services

import (
	"app/services/convertor"
)

type Service struct {
	Convertor convertor.Convertor
}

func NewService() Service {

	return Service{

		Convertor: convertor.NewConvertor(),
	}
}
