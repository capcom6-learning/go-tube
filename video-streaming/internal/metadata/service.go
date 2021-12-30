package metadata

import (
	"encoding/json"
	"net/http"
)

type MetadataService struct {
	baseUrl string
}

func NewMetadataService(baseUrl string) *MetadataService {
	return &MetadataService{
		baseUrl: baseUrl,
	}
}

func (m *MetadataService) GetMetadataById(id string) (*Metadata, error) {
	req, err := http.NewRequest("GET", m.baseUrl+"/video", nil)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Set("id", id)
	req.URL.RawQuery = query.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, nil
	}

	var metadata Metadata
	if err := json.NewDecoder(res.Body).Decode(&metadata); err != nil {
		return nil, err
	}

	return &metadata, nil
}
