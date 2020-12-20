package prisk

type Area struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	Key        string `json:"key"`
	ChildAreas []Area `json:"items"`
}
