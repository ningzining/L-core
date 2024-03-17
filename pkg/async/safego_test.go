package async

import (
	log "github.com/ningzining/L-log"
	"testing"
	"time"
)

func TestSaveGo(t *testing.T) {
	SaveGo(Log)
	time.Sleep(time.Second)
}

func TestSaveErrorGroup(t *testing.T) {
	if err := SaveErrorGroup(
		func() error {
			_, err := Return1()
			return err
		}, func() error {
			_, err := Return2()
			return err
		},
		func() error {
			Panic1()
			return nil
		},
	); err != nil {
		t.Error(err)
		return
	}
	time.Sleep(time.Second)
}

func Log() {
	log.Error("111")
	panic("error")
}

func Return1() (int, error) {
	return 1, nil
}

func Return2() (string, error) {
	return "str", nil
}

func Panic1() {
	panic("panic1")
}
