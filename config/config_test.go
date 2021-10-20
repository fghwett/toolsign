package config

import "testing"

func TestInit(t *testing.T) {
	conf, err := Init("../config.yml")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(conf.Config.Cookie)
}
