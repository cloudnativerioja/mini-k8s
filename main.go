package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	var k3d_binary string
	if checkK3dInstalled() {

		fmt.Println("k3d is installed on this system.")
		k3d_binary = "k3d"
	} else {
		fmt.Println("k3d is not installed on this system.")
		url := "https://github.com/k3d-io/k3d/releases/download/v5.4.6/k3d-linux-amd64"
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		out, err := os.Create("bin/k3d")
		if err != nil {
			panic(err)
		}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			panic(err)
		}

		fmt.Println("k3d binary downloaded and saved to /usr/local/bin/k3d")

		err = os.Chmod("bin/k3d", 0755)
		if err != nil {
			panic(err)
		}
		k3d_binary = "bin/k3d"
		fmt.Println("k3d binary set to executable")
	}
	//Create a flag to create/delete cluster
	action := flag.String("action", "", "'create' to create a k3d cluster, 'delete' to delete a k3d cluster")
	flag.Parse()

	if *action == "" || (*action != "create" && *action != "delete") {
		fmt.Println("Please provide a valid action flag: 'create' or 'delete'")
		os.Exit(1)
	}

	switch *action {
	case "create":
		// create k3d cluster
		fmt.Println("Creating k3d cluster...")
		cmd := exec.Command(k3d_binary, "cluster", "create", "--config", "config/cluster.yaml")
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error creating k3d cluster: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("K3d cluster created successfully\n")
	case "delete":
		// delete k3d cluster
		fmt.Println("Deleting k3d cluster...")
		cmd := exec.Command(k3d_binary, "cluster", "delete", "--config", "config/cluster.yaml")
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error deleting k3d cluster: %v\n", err)
			os.Exit(1)
		}
	}
}

func checkK3dInstalled() bool {
	_, err := exec.LookPath("k3d")
	if err != nil {
		return false
	}

	return true
}
