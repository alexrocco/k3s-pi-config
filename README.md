# k3s Pi Config tool

k3s-pi-config is a configure tool to deploy and manage Kubernetes cluster, using k3s (https://k3s.io), on Raspberry Pi devices.

## Architecture

The application is a Golang CLI, so it uses [Cobra](https://github.com/spf13/cobra) to simplify its development. All the commands runnuing on remote hosts will use Golang native ssh implementation, [golang.org/x/crypto/ssh](https://pkg.go.dev/golang.org/x/crypto/ssh?tab=doc).


## Backlog

1. Configure Raspberry Pi ports on server node. The port 6443 needs to be accessible by the agent nodes.
2. Install and deploy k3s as a server (master) using default installation guide on k3s docs (https://rancher.com/docs/k3s/latest/en/installation/install-options/)
3. Configure Raspberry Pi ports on agent nodes. The nodes need to be able to reach other nodes over UDP port 8472 (Flannel VXLAN).
4. Install and deploy k3s as a node. The k3s server URL must be provided as a flag and the token will be recovered from the server node.

Notes:
- For now, as a PoC, all the ssh connection will be using default Raspberry Pi user and password.
- This is not Production ready! 