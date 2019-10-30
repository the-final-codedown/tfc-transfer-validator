package cap_updater

import (
	tfc_cap_updater "github.com/the-final-codedown/tfc-cap-updater/proto/tfc/cap/updater"
	"google.golang.org/grpc"
)

var service *grpc.ClientConn

func GetCapStub(serverAdress string) (tfc_cap_updater.CapUpdaterServiceClient, error) {

	service, err := grpc.Dial(serverAdress, grpc.WithInsecure())
	// use the generated client stub
	cl := tfc_cap_updater.NewCapUpdaterServiceClient(service)
	return cl, err
}
func CloseConnection() {
	_ = service.Close()

}
