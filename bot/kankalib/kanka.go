package kankalib

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type KankaClient struct {
	Token string
	Campain string	
}

func(kanka KankaClient) CharacterAttributes(entity string) (Attributes, error){
	url := fmt.Sprintf("https://api.kanka.io/1.0/campaigns/%s/entities/%s/attributes", kanka.Campain, entity)
	log.Println(url)
	req, err := http.NewRequest("GET", 
		url, nil,
	)
	if err != nil {
		return Attributes{}, fmt.Errorf("error while creating Request : %v", err)
	}

	req.Header.Add("Authorization", "Bearer "+kanka.Token)
	req.Header.Add("Content-type", "application/json")


	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Attributes{}, fmt.Errorf("error while playing Request : %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Attributes{}, fmt.Errorf("request failed with status : %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Attributes{}, fmt.Errorf("error while reading response : %v", err)
	}

	var attributes Attributes
	err = json.Unmarshal(body, &attributes)
	if err != nil {
		return Attributes{}, fmt.Errorf("error while unmarshalling the response: %v", err)
	}

	return attributes, nil
}
