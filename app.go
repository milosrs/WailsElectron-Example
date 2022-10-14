package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	go func() {
		var i = 0
		for ; ; i++ {
			runtime.EventsEmit(ctx, string(BlacklistUpdated), []string{fmt.Sprintf("blacklist_url: %d", i)})
			time.Sleep(time.Second * 10)
		}
	}()
}

const loginUrl = "https://3eyes.app/Login.ashx"

type LoginResp struct {
	Token string `json:"token"`
}

// Login
func (a *App) Login(username string, password string) (string, error) {
	reqBody := map[string]string{
		"username": username,
		"password": password,
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		return "", errors.New("Error sending request. Try again later.")
	}

	fmt.Printf("%v\n", data)

	resp, err := http.Post(loginUrl, "application/json", bytes.NewBuffer(data))
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}

	respData := LoginResp{}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println(string(respBytes))

	err = json.Unmarshal(respBytes, &respData)
	if err != nil {
		return "", err
	}

	fmt.Println(resp.Body)
	fmt.Println(respData)

	if respData.Token == "" {
		return "", errors.New("Wrong username or password.")
	}

	return respData.Token, nil
}
