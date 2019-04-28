package structs

type Configs struct {
	Micro *Micro `json:"micro,omitempty"`
	Other *Other `json:"other,omitempty"`
}

type Micro struct {
	Demo
}

type Other struct {
	Demo
}

type Demo struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	Hi      string `json:"hi,omitempty"`
	Age     int    `json:"age,omitempty"`
}
