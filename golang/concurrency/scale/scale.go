package main

import (
	"errors"
	"image"
	"path"
	"path/filepath"
	"runtime"
	"sync"
)

func scaleDir(srcDir, dstDir string, size image.Rectangle) error {
	var mu sync.Mutex
	var errs error
	pool := make(chan bool, runtime.GOMAXPROCS(0))
	wg := sync.WaitGroup{}

	srcFiles, err := filepath.Glob(filepath.Join(srcDir, "*.png"))

	if err != nil {
		return err
	}

	wg.Add(len(srcFiles))

	for _, src := range srcFiles {

		go func(srcFile string) {
			defer wg.Done()
			pool <- true
			defer func() { <-pool }()

			destFile := path.Join(dstDir, filepath.Base(srcFile))

			if err := scaleImage(srcFile, destFile, size); err != nil {
				mu.Lock()
				errs = errors.Join(errs, err)
				mu.Unlock()
			}
		}(src)

	}
	wg.Wait()

	return errs
}
