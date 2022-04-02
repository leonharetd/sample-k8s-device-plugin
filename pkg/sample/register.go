package sample

import (
	"context"
	"fmt"
	utils "sample-k8s-device-plugin/utils"

	"google.golang.org/grpc"
	devicepluginv1beta1 "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

func Register() error {
	conn, err := grpc.Dial("unix://"+devicepluginv1beta1.KubeletSocket, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return fmt.Errorf("failed to dial %s: %s", devicepluginv1beta1.KubeletSocket, err)
	}
	defer conn.Close()

	client := devicepluginv1beta1.NewRegistrationClient(conn)
	req := &devicepluginv1beta1.RegisterRequest{
		Version:      "v1beta1",
		Endpoint:     utils.SampleSocket, // endpoint 就是文件名字
		ResourceName: utils.SampleResourceName,
	}
	_, err = client.Register(context.Background(), req)
	return err
}
