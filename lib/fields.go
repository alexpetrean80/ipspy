package lib

type Fields struct {
	All       bool
	Geo       bool
	DNS       bool
	Mobile    bool
	ISP       bool
	InitialIP bool
	Proxy     bool
	Hosting   bool
}

const (
	Continent     = "continent"
	ContinentCode = "continentCode"
	Country       = "country"
	CountryCode   = "countryCode"
	Region        = "region"
	RegionName    = "regionName"
	City          = "city"
	District      = "district"
	ZipCode       = "zip"
	Latitude      = "lat"
	Longitude     = "lon"
	Timezone      = "timezone"
	Offset        = "offset"
	Currency      = "currency"
	ISP           = "isp"
	Organization  = "org"
	ASNumber      = "as"
	ASName        = "asname"
	ReverseDNS    = "reverse"
	IsMobile      = "mobile"
	IsProxy       = "proxy"
	IsHosting     = "hosting"
	IP            = "query"
)
