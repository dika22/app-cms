package structs

type ResponseGetAllTags struct {
	DetailTag []DetailTag `json:"tags"`
}

func (t Tag) NewResponseGetAllTags() ResponseGetAllTags {
	return ResponseGetAllTags{}
}