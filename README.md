# Neuron Nexus Yandex GPT

Neuron Nexus Yandex GPT is a Go framework for interacting with the Yandex GPT API. Read More in [Wiki](https://github.com/neuron-nexus/yandexgpt/wiki).

## Installation



```bash
go get -u github.com/neuron-nexus/yandexgpt
```

## Usage

### Single Message Mode

This mode is for sending a system pompt and a single message. It is NOT intended to create a dialog with Yandex GPT.

```go
package main

import (
	"fmt"
	"github.com/neuron-nexus/yandexgpt/pkg/models/gpt"
	"github.com/neuron-nexus/yandexgpt/pkg/models/key"
	single "github.com/neuron-nexus/yandexgpt/pkg/sync/yandexGPTSyncApp/message"
)

const (
	GPT_API_KEY = "AQVN***************"
	STORAGE_ID  = "b1*****************"
)

func main() {

    // Initializing the single-mode model
	singleMode := single.New(GPT_API_KEY, key.API_KEY, STORAGE_ID)
	singleMode.SetModel(gpt.PRO) // Also posible gpt.LITE
	singleMode.SetTemperature(0.5) // Any float64 nubmer from 0 to 1
	singleMode.SetSystemPrompt("You are a professional Go programmer") // Any system prompt
    // Request from user
	singleMode.SetSingleMessage("Write a program in Go that outputs \"Hello World\" to the console") 

	resp, _ := singleMode.SendRequest()

	fmt.Println("\nAssistant:\n", resp.Result.Alternatives[0].Message.Text)
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
	"github.com/neuron-nexus/yandexgpt/pkg/models/gpt"
	"github.com/neuron-nexus/yandexgpt/pkg/models/key"
	"github.com/neuron-nexus/yandexgpt/pkg/models/role"
	"github.com/neuron-nexus/yandexgpt/pkg/sync/yandexGPTSyncApp/dialog"
	"os"
	"strings"
)

const (
	GPT_API_KEY = "AQVN***************"
	STORAGE_ID  = "b1*****************"
)

func main() {
	dialogMode := dialog.New(GPT_API_KEY, key.API_KEY, STORAGE_ID)
	dialogMode.SetModel(gpt.PRO)
	dialogMode.SetTemperature(0.5)
	dialogMode.SetSystemPrompt("You are a professional Go programmer")

	for {
		fmt.Print("You: ")
		message, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		_ = dialogMode.AddMessage(role.User, strings.TrimSpace(message))
		resp, _ := dialogMode.SendRequest()
		fmt.Println("Assistant: ", resp.Result.Alternatives[0].Message.Text)
	}
}
```
This model automatically adds the response from Yandex GPT to the message list, so you will only need to add messages from users.

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.
