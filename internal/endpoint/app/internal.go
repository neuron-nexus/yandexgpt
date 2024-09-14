package App

import (
	"bytes"
	"encoding/json"
	model "github.com/neuron-nexus/yandexgpt/internal/models/sync"
	"io"
	"net/http"
)

const URL = "https://llm.api.cloud.yandex.net/foundationModels/v1/completion"

func (a *App) sendRequestToYandexGPT(request *model.Request) (model.Response, error) {
	byteData, err := json.Marshal(request)
	if err != nil {
		return model.Response{}, err
	}

	httpRequest, err := http.NewRequest("POST", URL, bytes.NewReader(byteData))
	if err != nil {
		return model.Response{}, err
	}

	httpRequest.Header.Set("Authorization", a.Credential.KeyType+" "+a.Credential.Key)

	client := &http.Client{}

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return model.Response{}, err
	}

	defer httpResponse.Body.Close()

	var response model.Response

	byteResponse, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return model.Response{}, err
	}

	err = json.Unmarshal(byteResponse, &response)

	return response, err
}
