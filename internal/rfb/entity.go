package rfb

type defaultEntity struct {
	Code        int
	Description string
}

type Cnae struct {
	defaultEntity
}

type City struct {
	defaultEntity
}

type LegalNature struct {
	defaultEntity
}

type Country struct {
	defaultEntity
}

type PartnerQualification struct {
	defaultEntity
}

type Company struct {
	CnpjBase                    string
	Name                        string
	LegalNature                 int
	ResponsibleQualification    int
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
	MainCnae           int
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
	email              string
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
	PartnerQualification             int
	StartDate                        int
	CountryId                        int
	LegalRepresentative              string
	LegalRepresentativeName          string
	LegalRepresentativeQualification int
	AgeGroup                         int
}
