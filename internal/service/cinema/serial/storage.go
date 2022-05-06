package serial

import "github.com/ozonmp/omp-bot/internal/model/cinema"

var sampleData = []cinema.Serial{
	{
		Title:      "Рик и Морти",
		Genre:      "Мультфильм",
		SeasonsNum: 6,
		IsDeleted:  false,
	},
	{
		Title:      "Гравити Фолз",
		Genre:      "Мультфильм",
		SeasonsNum: 2,
		IsDeleted:  false,
	},
	{
		Title:      "Друзья",
		Genre:      "Комедия",
		SeasonsNum: 10,
		IsDeleted:  false,
	},
	{
		Title:      "Сопрано",
		Genre:      "Драма",
		SeasonsNum: 6,
		IsDeleted:  false,
	},
	{
		Title:      "Во все тяжкие",
		Genre:      "Криминал",
		SeasonsNum: 5,
		IsDeleted:  false,
	},
}

func fillStorageBySampleData(service SerialService) {
	for _, serial := range sampleData {
		_, _ = service.Create(serial)
	}
}
