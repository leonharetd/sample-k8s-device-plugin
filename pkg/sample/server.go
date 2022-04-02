package sample

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	utils "sample-k8s-device-plugin/utils"
	"time"

	"google.golang.org/grpc"
	devicepluginv1beta1 "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

type DevicePluginServer struct {
	sugerNum int
	devicepluginv1beta1.UnimplementedDevicePluginServer
}

var _ devicepluginv1beta1.DevicePluginServer = &DevicePluginServer{}

func (d *DevicePluginServer) GetDevicePluginOptions(ctx context.Context, req *devicepluginv1beta1.Empty) (*devicepluginv1beta1.DevicePluginOptions, error) {
	log.Println("Get DevicePlugin Options")
	return &devicepluginv1beta1.DevicePluginOptions{}, nil
}

func (d *DevicePluginServer) ListAndWatch(req *devicepluginv1beta1.Empty, stream devicepluginv1beta1.DevicePlugin_ListAndWatchServer) error {
	log.Println("List and Watch")
	for {
		resp := &devicepluginv1beta1.ListAndWatchResponse{
			Devices: []*devicepluginv1beta1.Device{
				{ID: "b", Health: devicepluginv1beta1.Healthy},
				{ID: "a", Health: devicepluginv1beta1.Healthy},
				{ID: "c", Health: devicepluginv1beta1.Healthy},
			},
		}
		if err := stream.Send(resp); err != nil {
            return fmt.Errorf("failed to send %s", err)
		}
		time.Sleep(time.Second*2)
	}
}

func (d *DevicePluginServer) Allocate(ctx context.Context, req *devicepluginv1beta1.AllocateRequest) (*devicepluginv1beta1.AllocateResponse, error) {
	log.Printf("Allocate")
	log.Printf("allocate deviceids:")
	for _, containerReq := range req.ContainerRequests {
		log.Printf(fmt.Sprint(containerReq.DevicesIDs))
	}
	var containerReqs []*devicepluginv1beta1.ContainerAllocateResponse
	for _, containerReq := range req.ContainerRequests {
		containerReps := devicepluginv1beta1.ContainerAllocateResponse{
			Envs: map[string]string{
				utils.SamplePodEnvKey: strings.Join(containerReq.DevicesIDs, ","),
			},
		}
        containerReqs = append(containerReqs, &containerReps)
	}
	return &devicepluginv1beta1.AllocateResponse{
		ContainerResponses: containerReqs,
	}, nil
}

func Serve(sugerNum int) error {
	socket := filepath.Join(utils.DevicePluginsDir, utils.SampleSocket)
	_ = os.Remove(socket)
	listener, err := net.Listen("unix", socket)
	if err != nil {
		return fmt.Errorf("failed to listen %s: %s", socket, err)
	}
	defer listener.Close()

	server := grpc.NewServer()
	devicepluginv1beta1.RegisterDevicePluginServer(server, &DevicePluginServer{sugerNum: sugerNum})
	return server.Serve(listener)
}

