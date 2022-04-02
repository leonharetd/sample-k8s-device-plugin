package utils
import (
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

// common var statement
const DevicePluginsDir = pluginapi.DevicePluginPath

// sample device plugin const
const (
	SampleSocket = "sample-k8s-device-plugin.sock"
	SampleResourceName = "hardware-vendor.example/suger"
	SamplePodEnvKey = "suger"
)