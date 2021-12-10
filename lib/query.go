package lib

import (
	"fmt"
	"net"
	"strings"
)

const apiAddr = "http://ip-api.com/json"

type Query interface {
	Create(ip net.IP) string
}

type query string

func NewQuery(f Fields) query {
	fields := []string{}
	if f.All {
		fields = append(fields,
			Continent, ContinentCode,
			Country, CountryCode, Region, RegionName,
			City, District, ZipCode, Latitude,
			Longitude, Timezone, Offset, Currency,
			ISP, Organization, ASNumber, ASName,
			ReverseDNS, IsMobile, IsProxy, IsHosting, IP,
		)
	}

	if f.Geo {
		fields = append(fields,
			Continent, ContinentCode, Country, CountryCode,
			Region, RegionName, City, District,
			ZipCode, Latitude, Longitude, Timezone, Offset,
		)
	}

	if f.DNS {
		fields = append(fields, ReverseDNS)
	}

	if f.Mobile {
		fields = append(fields, IsMobile)
	}

	if f.Proxy {
		fields = append(fields, IsProxy)
	}

	if f.Hosting {
		fields = append(fields, IsHosting)
	}

	if f.InitialIP {
		fields = append(fields, IP)
	}
	if f.ISP {
		fields = append(fields, ISP, ASName, ASNumber, Organization)
	}

	fmt.Println(fields)
	return query(strings.Join(fields, ","))
}

func (q query) Create(ip net.IP) string {
	return fmt.Sprintf("%s/%s?fields=%s", apiAddr, ip, q)
}
