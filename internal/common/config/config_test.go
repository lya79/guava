package config

import (
	"testing"
)

func Test_GetConfig(t *testing.T) {
	p, err := getPathOfConfig()
	if err != nil {
		t.Error("getPathOfConfig, err:", err)
	}
	t.Log("getPathOfConfig:", p)
	config := GetConfig()
	if config.Chatroom.Host == "" {
		t.Error("config.Chatroom.Host empty")
	} else if config.Chatroom.Port == 0 {
		t.Error("config.Chatroom.Port 0")
	}
}

func Test_IsLocalEnv(t *testing.T) {
	if !IsLocalEnv() {
		t.Error("IsLocalEnv false")
	}
}

func Test_IsDebugMode(t *testing.T) {
	if !IsDebugMode() {
		t.Error("IsDebugMode false")
	}
}

func Test_GetAppSite(t *testing.T) {
	if GetAppSite() != "chatroom" {
		t.Error("GetAppSite != chatroom")
	}
}
