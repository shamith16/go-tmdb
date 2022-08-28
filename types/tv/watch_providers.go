package tv

import "encoding/json"

func UnmarshalWatchProviders(data []byte) (WatchProviders, error) {
	var r WatchProviders
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WatchProviders) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type WatchProviders struct {
	ID      int     `json:"id"`
	Results Results `json:"results"`
}

type Results struct {
	Ar Ar `json:"AR"`
	At At `json:"AT"`
	Au At `json:"AU"`
	Be Ar `json:"BE"`
	Br Ar `json:"BR"`
	CA At `json:"CA"`
	Ch At `json:"CH"`
	Cl Ar `json:"CL"`
	Co Ar `json:"CO"`
	Cz Ar `json:"CZ"`
	De At `json:"DE"`
	Dk At `json:"DK"`
	Ec Ar `json:"EC"`
	Es Ar `json:"ES"`
	Fi At `json:"FI"`
	Fr At `json:"FR"`
	GB At `json:"GB"`
	Hu Ar `json:"HU"`
	Ie At `json:"IE"`
	In Ar `json:"IN"`
	It At `json:"IT"`
	Jp At `json:"JP"`
	MX At `json:"MX"`
	Nl Ar `json:"NL"`
	No At `json:"NO"`
	Nz Ar `json:"NZ"`
	PE Ar `json:"PE"`
	Pl At `json:"PL"`
	Pt Ar `json:"PT"`
	Ro Ar `json:"RO"`
	Ru Ru `json:"RU"`
	SE At `json:"SE"`
	Tr Ar `json:"TR"`
	Us At `json:"US"`
	Ve Ar `json:"VE"`
	Za Ar `json:"ZA"`
}

type Ar struct {
	Link     string     `json:"link"`
	Flatrate []Flatrate `json:"flatrate"`
}

type Flatrate struct {
	DisplayPriority int    `json:"display_priority"`
	LogoPath        string `json:"logo_path"`
	ProviderID      int    `json:"provider_id"`
	ProviderName    string `json:"provider_name"`
}

type At struct {
	Link     string     `json:"link"`
	Buy      []Flatrate `json:"buy,omitempty"`
	Flatrate []Flatrate `json:"flatrate"`
	Ads      []Flatrate `json:"ads,omitempty"`
	Rent     []Flatrate `json:"rent,omitempty"`
}

type Ru struct {
	Link     string     `json:"link"`
	Flatrate []Flatrate `json:"flatrate"`
	Free     []Flatrate `json:"free"`
}
