package sample

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sample-k8s-device-plugin/utils"

	"github.com/fsnotify/fsnotify"
)

// 监听目录里自己的socket文件
func WatchKubeletRestart(stop chan struct{}) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("create kubelet watch error %v", err)
	}
	defer watcher.Close()
	go func() {
		for {
			select {
			case event, ok := <- watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Remove == 0 {
                    continue
				}
				if fileName, err := filepath.Rel(utils.DevicePluginsDir, event.Name); err == nil && fileName == utils.SampleSocket {
					log.Printf("socket has been removed, exiting")
					os.Exit(0)
				}
			case err, ok := <- watcher.Errors:
				if !ok {
					return
				}
				log.Printf("watcher error: %s", err)
			}
		}
	}()

	if err := watcher.Add(utils.DevicePluginsDir); err != nil {
		return fmt.Errorf("failed to start watching %s: %s", utils.DevicePluginsDir, err)
	}

	<-stop
    return nil
}
