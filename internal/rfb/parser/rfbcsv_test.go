package parser

import (
	"os"
	"t-bonatti/gopj/internal/rfb"
	"testing"

	"github.com/stretchr/testify/assert"
)

func loadFile(t *testing.T, fileName string) *os.File {
	f, err := os.Open("../../../test/data/rfb_csv/" + fileName)
	assert.Nil(t, err)
	return f
}

func TestCnaeParseCsvFile(t *testing.T) {

	expect := []rfb.Cnae{
		{Code: "0111301", Description: "Cultivo de arroz"},
		{Code: "0111302", Description: "Cultivo de milho"},
		{Code: "0111303", Description: "Cultivo de trigo"},
		// FIXME: {Code: "0111399", Description: "Cultivo de outros cereais não especificados anteriormente"},
	}

	f := loadFile(t, "cnaes")
	defer f.Close()

	data := GetCnaesFromCsv(f)
	assert.ElementsMatch(t, expect, data)
}

func TestCityParseCsvFile(t *testing.T) {
	expect := []rfb.City{
		{Code: "0001", Description: "GUAJARA-MIRIM"},
		{Code: "0002", Description: "ALTO ALEGRE DOS PARECIS"},
		{Code: "0003", Description: "PORTO VELHO"},
	}

	f := loadFile(t, "cities")
	defer f.Close()

	data := GetCitiesFromCsv(f)
	assert.ElementsMatch(t, expect, data[0:3])
}

func TestLegalNatureParseCsvFile(t *testing.T) {
	expect := []rfb.LegalNature{
		{Code: "0000", Description: "NATUREZA JURIDICA NAO INFORMADA"},
		{Code: "1104", Description: "Autarquia Federal"},
		{Code: "1112", Description: "Autarquia Estadual ou do Distrito Federal"},
	}

	f := loadFile(t, "legalnatures")
	defer f.Close()

	data := GetLegalNaturesFromCsv(f)
	assert.ElementsMatch(t, expect, data[0:3])
}

func TestCountryParseCsvFile(t *testing.T) {
	expect := []rfb.Country{
		{Code: "000", Description: "COLIS POSTAUX"},
		{Code: "013", Description: "AFEGANISTAO"},
		{Code: "017", Description: "ALBANIA"},
	}

	f := loadFile(t, "countries")
	defer f.Close()

	data := GetCountriesFromCsv(f)
	assert.ElementsMatch(t, expect, data[0:3])
}

func TestPartnerQualificationParseCsvFile(t *testing.T) {
	expect := []rfb.PartnerQualification{
		{Code: "00", Description: "NAO INFORMADA"},
		{Code: "05", Description: "Administrador"},
		{Code: "09", Description: "Curador"},
	}

	f := loadFile(t, "partnerqualifications")
	defer f.Close()

	data := GetPartnerQualificationsFromCsv(f)
	assert.ElementsMatch(t, expect, data[0:3])
}

func TestCompaniesParseCsvFile(t *testing.T) {
	expect := []rfb.Company{
		{CnpjBase: "36627984", Name: "RAFAELA FRANCISCO FIGUEIREDO 46701565867", LegalNature: "2135", ResponsibleQualification: "50", ShareCapital: 300.00, Size: 1, ResponsibleFederativeEntity: ""},
		{CnpjBase: "36627985", Name: "JULIANA RAMOS DE MELO MARQUES 36687602845", LegalNature: "2135", ResponsibleQualification: "50", ShareCapital: 500.00, Size: 1, ResponsibleFederativeEntity: ""},
		{CnpjBase: "36627986", Name: "FABRICIO DA SILVA SEIXAS 01555083420", LegalNature: "2135", ResponsibleQualification: "50", ShareCapital: 500.00, Size: 1, ResponsibleFederativeEntity: ""},
	}

	f := loadFile(t, "companies")
	defer f.Close()

	data := GetCompaniesFromCsv(f)
	assert.ElementsMatch(t, expect, data[0:3])
}

func TestEstablishmentsParseCsvFile(t *testing.T) {
	expect := []rfb.Establishment{
		{
			CnpjBase:           "",
			CnpjOrder:          "",
			CnpjDv:             "",
			HeadquartersBranch: "",
			FantasyName:        "",
			Status:             0,
			StatusDate:         0,
			StatusReason:       "",
			ExtCity:            "",
			Country:            "",
			StartDate:          0,
			MainCnae:           "",
			OtherCnae:          "",
			TypeAddress:        "",
			Number:             "",
			Complement:         "",
			District:           "",
			Cep:                "",
			State:              "",
			City:               "",
			Ddd1:               "",
			Phone1:             "",
			Ddd2:               "",
			Phone2:             "",
			FaxDdd:             "",
			Fax:                "",
			Email:              "",
			SpecialStatus:      "",
			SpecialStatusDate:  0,
		},
	}

	f := loadFile(t, "companies")
	defer f.Close()

	data := GetEstablishmentsFromCsv(f)
	assert.ElementsMatch(t, expect, data[0:3])
}
