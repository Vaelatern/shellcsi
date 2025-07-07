package main

import (
	"log"
	"net"
	"os"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc"

	driver "github.com/Vaelatern/shellcsi/internal/csi"
)

func main() {
	sock := "/tmp/csi.sock"
	listener, err := net.Listen("unix", sock)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer os.Remove(sock)
	defer listener.Close()

	err = os.Chmod(sock, 0600)
	if err != nil {
		log.Fatalf("Failed to set permissions on %s to 0600", sock)
	}

	s := grpc.NewServer()
	d := driver.NewDriver("shell.csi", "0.0.1", driver.Config{})

	csi.RegisterIdentityServer(s, d)
	csi.RegisterControllerServer(s, d)
	csi.RegisterNodeServer(s, d)

	log.Printf("Starting CSI driver on %s", sock)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
