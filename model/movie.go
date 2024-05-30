package model

type Companies struct {
	CompanyCd string `json:"companyCd"`
	CompanyNm string `json:"companyNm"`
}

type Companys struct {
	CompanyCd string `json:"companyCd"`
	CompanyNm string `json:"companyNm"`
}

type Directors struct {
	PeopleNm string `json:"peopleNm"`
}

type Movie struct {
	Companies   []Companies `json:"companies"`
	Companys    []Companys  `json:"companys"`
	Directors   []Directors `json:"directors"`
	GenreAlt    interface{} `json:"genreAlt"`
	MovieCd     string      `json:"movieCd"`
	MovieNm     string      `json:"movieNm"`
	MovieNmEn   string      `json:"movieNmEn"`
	NationAlt   string      `json:"nationAlt"`
	OpenDt      string      `json:"openDt"`
	PrdtStatNm  string      `json:"prdtStatNm"`
	PrdtYear    string      `json:"prdtYear"`
	RepGenreNm  string      `json:"repGenreNm"`
	RepNationNm string      `json:"repNationNm"`
	TypeNm      string      `json:"typeNm"`
}
