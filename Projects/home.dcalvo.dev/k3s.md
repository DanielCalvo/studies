How do I set up k3s on an orangepi with ubuntu again?

# Login as root, change your password, delete the default account, create one for yourself, change hostname

```shell
#whatever you got from dhcp here
ssh orangepi@192.168.1.46
sudo su
adduser daniel
sudo adduser daniel sudo

#lol
rm /lib/systemd/system/getty@.service.d/override.conf
rm /lib/systemd/system/serial-getty@.service.d/override.conf
init 6
deluser orangepi --remove-all-files

echo 'opi1' > /etc/hostname
```

# Set a static
- `vim /etc/netplan/orangepi-default.yaml`
```yaml
network:
  version: 2
  renderer: networkd
  ethernets:
    eth0:
      dhcp4: false
      addresses:
        - 192.168.1.201/24
      routes:
        - to: default
          via: 192.168.1.1
      nameservers:
          addresses: [192.168.1.1]
```

# Single host, basic k3s install
- https://docs.k3s.io/quick-start
`curl -sfL https://get.k3s.io | sh -`

# Accessing the cluster from your desktop
On the cluster node: `chmod +r /etc/rancher/k3s/k3s.yaml`
On your local:
```shell
scp daniel@192.168.1.201:/etc/rancher/k3s/k3s.yaml ~/.kube/config
sed -i 's/127.0.0.1/192.168.1.201/g' ~/.kube/config
```



# App with ingress:
See [k3s_app](./k3s_app)
