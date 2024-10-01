package sync

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	model "github.com/neuron-nexus/yandexgpt/internal/models"
)

const URL = "https://llm.api.cloud.yandex.net/foundationModels/v1/completion"

func (a *App) sendRequestToYandexGPT(request *model.Request) (model.Response, error) {
	byteData, err := json.Marshal(&request)
	if err != nil {
		return model.Response{}, err
	}

	httpRequest, err := http.NewRequest("POST", URL, bytes.NewReader(byteData))
	if err != nil {
		return model.Response{}, err
	}

	httpRequest.Header.Set("Authorization", a.Credential.KeyType+" "+a.Credential.Key)
	httpRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return model.Response{}, err
	}

	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != 200 {
		return model.Response{}, errors.New("yandex: " + httpResponse.Status)
	}

	var response model.Response

	byteResponse, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return model.Response{}, err
	}

	err = json.Unmarshal(byteResponse, &response)
	return response, err
}
