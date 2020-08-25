package alerts

import (
	"time"
	client "github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
)

type Firing struct {
	FiringID     string      `json:"firingId"`
	DefinitionID string      `json:"definitionId"`
	Name         string      `json:"name"`
	StartTime    time.Time   `json:"startTime"`
	EndTime      time.Time   `json:"endTime"`
	Service      string      `json:"service"`
	FieldMap     map[string]interface{}    `json:"fieldMap"`
}


type FiringsList struct {
	Firings []Firing `json:"data"`
}

func ListActiveFirings() (*FiringsList, error) {
	req, err := client.NewRequest(
		Config,
		"GET",
		"/alerts/v2/alert-firings/active",
		nil,
	)

	if err != nil {
		return nil, err
	}

	res, err := client.Do(Config, req)

	if err != nil {
		return nil, err
	}

	if client.IsError(res) {
		return nil, client.NewAPIError(res)
	}

	var response FiringsList
	if err = client.BodyJSON(res, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
