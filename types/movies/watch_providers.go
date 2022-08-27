package movies

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
	At Ar `json:"AT"`
	Au Ar `json:"AU"`
	Be Ar `json:"BE"`
	Br Ar `json:"BR"`
	CA Ar `json:"CA"`
	Ch Ar `json:"CH"`
	Cl Ar `json:"CL"`
	Co Ar `json:"CO"`
	Cz Ar `json:"CZ"`
	De Ar `json:"DE"`
	Dk Ar `json:"DK"`
	Ec Ar `json:"EC"`
	Ee Ar `json:"EE"`
	Es Ar `json:"ES"`
	Fi Ar `json:"FI"`
	Fr Ar `json:"FR"`
	GB Ar `json:"GB"`
	Gr Ar `json:"GR"`
	Hu Ar `json:"HU"`
	ID Ar `json:"ID"`
	Ie Ar `json:"IE"`
	In Ar `json:"IN"`
	It Ar `json:"IT"`
	Jp Ar `json:"JP"`
	Kr Ar `json:"KR"`
	Lt Ar `json:"LT"`
	LV Ar `json:"LV"`
	MX Ar `json:"MX"`
	My Ar `json:"MY"`
	Nl Ar `json:"NL"`
	No Ar `json:"NO"`
	Nz Ar `json:"NZ"`
	PE Ar `json:"PE"`
	Ph Ar `json:"PH"`
	Pl Ar `json:"PL"`
	Pt Ar `json:"PT"`
	Ro Ro `json:"RO"`
	Ru Ar `json:"RU"`
	SE Ar `json:"SE"`
	Sg Ar `json:"SG"`
	Th Ar `json:"TH"`
	Tr Ar `json:"TR"`
	Us Ar `json:"US"`
	Ve Ar `json:"VE"`
	Za Ar `json:"ZA"`
}

type Ar struct {
	Link     string `json:"link"`
	Flatrate []Buy  `json:"flatrate,omitempty"`
	Rent     []Buy  `json:"rent,omitempty"`
	Buy      []Buy  `json:"buy"`
}

type Buy struct {
	DisplayPriority int    `json:"display_priority"`
	LogoPath        string `json:"logo_path"`
	ProviderID      int    `json:"provider_id"`
	ProviderName    string `json:"provider_name"`
}

type Ro struct {
	Link     string `json:"link"`
	Flatrate []Buy  `json:"flatrate"`
}
