package rfb

type DefaultEntity struct {
	Code        string
	Description string
}

type Cnae struct {
	Code        string
	Description string
}

type City struct {
	Code        string
	Description string
}

type LegalNature struct {
	Code        string
	Description string
}

type Country struct {
	Code        string
	Description string
}

type PartnerQualification struct {
	Code        string
	Description string
}

type Company struct {
	CnpjBase                    string
	Name                        string
	LegalNature                 string
	ResponsibleQualification    string
	ShareCapital                float64
	Size                        int
	ResponsibleFederativeEntity string
}

type Establishment struct {
	CnpjBase           string
	CnpjOrder          string
	CnpjDv             string
	HeadquartersBranch string
	FantasyName        string
	Status             int
	StatusDate         int
	StatusReason       string
	ExtCity            string
	Country            string
	StartDate          int
	MainCnae           string
	OtherCnae          string
	TypeAddress        string
	Number             string
	Complement         string
	District           string
	Cep                string
	State              string
	City               string
	Ddd1               string
	Phone1             string
	Ddd2               string
	Phone2             string
	FaxDdd             string
	Fax                string
	Email              string
	SpecialStatus      string
	SpecialStatusDate  int
}

type SimpleCompany struct {
	CnpjBase                 string
	ChoseSimple              string
	ChoseSimpleDate          int
	ChoseSimpleExclusionDate int
	ChoseMei                 string
	ChoseMeiDate             int
	ChoseMeiExclusionDate    int
}

type Partner struct {
	CnpjBase                         string
	PartnerIdentification            int
	Name                             string
	Document                         string
	PartnerQualification             string
	StartDate                        int
	CountryId                        string
	LegalRepresentative              string
	LegalRepresentativeName          string
	LegalRepresentativeQualification string
	AgeGroup                         int
}
