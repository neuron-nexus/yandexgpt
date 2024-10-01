# Neuron Nexus Yandex GPT

Neuron Nexus Yandex GPT is a Go framework for interacting with the Yandex GPT API. Read More in [Wiki](https://github.com/neuron-nexus/yandexgpt/wiki).

## Installation



```bash
go get -u github.com/neuron-nexus/yandexgpt/v2
```

## Usage (v2.0.0)

### Single Message Mode

This mode is for sending a system pompt and a single message. It is NOT intended to create a dialog with Yandex GPT.

```go
package main

import (
	"fmt"
	"github.com/neuron-nexus/yandexgpt/v2"
)

const (
	GPT_API_KEY = "AQVN***************"
	STORAGE_ID  = "b1*****************"
)

func main() {

    app := yandexgpt.NewYandexGPTSyncApp(
		KEY,
		yandexgpt.API_KEY,
		STORAGE_ID,
		yandexgpt.GPTModelPRO,
	)

	configs := []yandexgpt.GPTParameter{
		{
			Name:  yandexgpt.ParameterPrompt, // Important!
			Value: "You are professional Go programmer",
		},
		{
			Name:  yandexgpt.ParameterTemperature, // Default: 0.3
			Value: "0.7",
		},
		{
			Name:  yandexgpt.ParameterMaxTokens, // Default: 2000
			Value: "1000",
		},
	}

	app.Configure(configs...)
	message := yandexgpt.GPTMessage{
		Role: yandexgpt.RoleUser,
		Text: "Write a programm that prints 'Hello World'",
	}

	app.AddMessage(message)
	res, err := app.SendRequest()
	println(res.Text)
}
```

Result:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```
### Dialog Mode
A simple example of creating a dialog with Yandex GPT. It will allow you to communicate with the model inside the console.
```go
package main

import (
	"bufio"
	"fmt"
	"github.com/neuron-nexus/yandexgpt/v2"
	"os"
	"strings"
)

const (
	GPT_API_KEY = "AQVN***************"
	STORAGE_ID  = "b1*****************"
)

func main() {
	app := yandexgpt.NewYandexGPTSyncApp(
		KEY,
		yandexgpt.API_KEY,
		STORAGE_ID,
		yandexgpt.GPTModelPRO,
	)

	configs := []yandexgpt.GPTParameter{
		{
			Name:  yandexgpt.ParameterPrompt, // Important!
			Value: "You are professional assistant",
		},
	}

	err := app.Configure(configs...)
	if err != nil {
		panic(err)
	}

	for {
		//read text from console
		fmt.Print("You: ")
		text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		if text == "exit" {
			return
		}

		app.AddMessage(yandexgpt.GPTMessage{
			Role: yandexgpt.RoleUser,
			Text: strings.TrimSpace(text),
		})
		res, _ := app.SendRequest()
		fmt.Println("Assistant: ", res.Text)
		app.AddRawMessage(res.Result.Alternatives[0].Message)
	}
}
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.
