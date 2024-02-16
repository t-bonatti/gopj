package parser

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"
	"t-bonatti/gopj/internal/rfb"
)

func readFromCsv(f io.Reader, c chan<- []string) {
	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'

	for {
		csvLine, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		c <- csvLine
	}
	close(c)
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func toFloat(s string) float64 {
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, ",", ".")
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}

func GetCnaesFromCsv(f io.Reader) (cnaes []rfb.Cnae) {
	c := make(chan []string, 100)
	go readFromCsv(f, c)
	for csvLine := range c {
		var cnae rfb.Cnae
		cnae.Code = csvLine[0]
		cnae.Description = csvLine[1]
		cnaes = append(cnaes, cnae)
	}
	return
}

func GetCitiesFromCsv(f io.Reader) (cities []rfb.City) {
	c := make(chan []string, 100)
	go readFromCsv(f, c)
	for csvLine := range c {
		var city rfb.City
		city.Code = csvLine[0]
		city.Description = csvLine[1]
		cities = append(cities, city)
	}
	return
}

func GetLegalNaturesFromCsv(f io.Reader) (legalNatures []rfb.LegalNature) {
	c := make(chan []string, 100)
	go readFromCsv(f, c)
	for csvLine := range c {
		var legalNature rfb.LegalNature
		legalNature.Code = csvLine[0]
		legalNature.Description = csvLine[1]
		legalNatures = append(legalNatures, legalNature)
	}
	return
}

func GetCountriesFromCsv(f io.Reader) (coutries []rfb.Country) {
	c := make(chan []string, 100)
	go readFromCsv(f, c)
	for csvLine := range c {
		var country rfb.Country
		country.Code = csvLine[0]
		country.Description = csvLine[1]
		coutries = append(coutries, country)
	}
	return
}

func GetPartnerQualificationsFromCsv(f io.Reader) (partnerQualifications []rfb.PartnerQualification) {
	c := make(chan []string, 100)
	go readFromCsv(f, c)
	for csvLine := range c {
		var partnerQualification rfb.PartnerQualification
		partnerQualification.Code = csvLine[0]
		partnerQualification.Description = csvLine[1]
		partnerQualifications = append(partnerQualifications, partnerQualification)
	}
	return
}

func GetCompaniesFromCsv(f io.Reader) (companies []rfb.Company) {
	c := make(chan []string, 100)
	go readFromCsv(f, c)
	for csvLine := range c {
		var company rfb.Company
		company.CnpjBase = csvLine[0]
		company.Name = csvLine[1]
		company.LegalNature = csvLine[2]
		company.ResponsibleQualification = csvLine[3]
		company.ShareCapital = toFloat(csvLine[4])
		company.Size = toInt(csvLine[5])
		company.ResponsibleFederativeEntity = csvLine[6]
		companies = append(companies, company)
	}
	return
}

func GetEstablishmentsFromCsv(f io.Reader) (establishments []rfb.Establishment) {
	c := make(chan []string, 100)
	go readFromCsv(f, c)
	for csvLine := range c {
		var establishment rfb.Establishment
		establishment.CnpjBase = csvLine[0]
		establishment.CnpjOrder = csvLine[1]
		establishment.CnpjDv = csvLine[2]
		establishment.HeadquartersBranch = csvLine[3]
		establishment.FantasyName = csvLine[4]
		establishment.Status = toInt(csvLine[5])
		establishment.StatusDate = toInt(csvLine[6])
		establishment.StatusReason = csvLine[7]
		establishment.ExtCity = csvLine[8]
		establishment.Country = csvLine[9]
		establishment.StartDate = toInt(csvLine[10])
		establishment.MainCnae = csvLine[11]
		establishment.OtherCnae = csvLine[12]
		establishment.TypeAddress = csvLine[13]
		establishment.Number = csvLine[14]
		establishment.Complement = csvLine[15]
		establishment.District = csvLine[16]
		establishment.Cep = csvLine[17]
		establishment.State = csvLine[18]
		establishment.City = csvLine[19]
		establishment.Ddd1 = csvLine[20]
		establishment.Phone1 = csvLine[21]
		establishment.Ddd2 = csvLine[22]
		establishment.Phone2 = csvLine[23]
		establishment.FaxDdd = csvLine[24]
		establishment.Fax = csvLine[25]
		establishment.Email = csvLine[26]
		establishment.SpecialStatus = csvLine[27]
		establishment.SpecialStatusDate = toInt(csvLine[28])
		establishments = append(establishments, establishment)
	}
	return
}

func GetSimpleCompaniesFromCsv(f io.Reader) (simpleCompanies []rfb.SimpleCompany) {
	c := make(chan []string, 100)
	go readFromCsv(f, c)
	for csvLine := range c {
		var simpleCompany rfb.SimpleCompany
		simpleCompany.CnpjBase = csvLine[0]
		simpleCompany.ChoseSimple = csvLine[1]
		simpleCompany.ChoseSimpleDate = toInt(csvLine[2])
		simpleCompany.ChoseSimpleExclusionDate = toInt(csvLine[3])
		simpleCompany.ChoseMei = csvLine[4]
		simpleCompany.ChoseMeiDate = toInt(csvLine[5])
		simpleCompany.ChoseMeiExclusionDate = toInt(csvLine[6])
		simpleCompanies = append(simpleCompanies, simpleCompany)
	}
	return
}

func GetPartnersFromCsv(f io.Reader) (partners []rfb.Partner) {
	c := make(chan []string, 100)
	go readFromCsv(f, c)
	for csvLine := range c {
		var partner rfb.Partner
		partner.CnpjBase = csvLine[0]
		partner.PartnerIdentification = toInt(csvLine[1])
		partner.Name = csvLine[2]
		partner.Document = csvLine[3]
		partner.PartnerQualification = csvLine[4]
		partner.StartDate = toInt(csvLine[5])
		partner.CountryId = csvLine[6]
		partner.LegalRepresentative = csvLine[6]
		partner.LegalRepresentativeName = csvLine[6]
		partner.LegalRepresentativeQualification = csvLine[6]
		partner.AgeGroup = toInt(csvLine[6])
		partners = append(partners, partner)
	}
	return
}
