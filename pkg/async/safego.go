package async

import (
	"fmt"
	log "github.com/ningzining/L-log"
	"golang.org/x/sync/errgroup"
)

type function func()

func SaveGo(f function) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf(fmt.Sprintf("%v", err))
			}
		}()

		f()
	}()
}

type errFunction func() error

func SaveErrorGroup(funcList ...errFunction) error {
	if len(funcList) == 0 {
		return nil
	}

	var eg errgroup.Group

	// 并发执行多个函数
	for _, f := range funcList {
		saveF := func() error {
			defer func() {
				if err := recover(); err != nil {
					log.Errorf(fmt.Sprintf("%v", err))
				}
			}()

			return f()
		}

		eg.Go(saveF)
	}

	return eg.Wait()
}
