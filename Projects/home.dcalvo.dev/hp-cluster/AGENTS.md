This is a folder for the config and setup of a Kubernetes cluster in my home lab.

# Cluster info
The cluster is composed of two HP Prodesk machines with 16bg of ram, 250gb of ssd and a Core i7-4790S each. It runs RKE2 version v1.35.6+rke2r1

You can ssh into the two nodes with non root credentials if necessary to assist the user. Node details are as follows:

# Node information
## hp1
- Access: ssh daniel@192.168.1.211
- This cluster hosts the Kubernetes control panel, but also runs non-control plane pods, functioning additionally as a worker node.

## hp2
- Access: ssh daniel@192.168.1.212
- This is a Kubernetes worker node only.
