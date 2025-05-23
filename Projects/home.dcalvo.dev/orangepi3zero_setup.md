# Things done from the initial image so you can have your own starting point

```shell
#whatever you got from dhcp here
ssh orangepi@192.168.1.46 #defalt password is orangepi
sudo su
apt-get update && apt-get upgrade -y && apt-get install git vim curl wget sshfs htop build-essential make apt-transport-https ca-certificates gnupg-agent software-properties-common -y
adduser daniel
sudo adduser daniel sudo
#login as yourself, add ssh key

#lol, as root:
rm /lib/systemd/system/getty@.service.d/override.conf
rm /lib/systemd/system/serial-getty@.service.d/override.conf
init 6
deluser orangepi --remove-all-files

#Change for the machine number
echo 'opi-X' > /etc/hostname
```

To edit the image:

# Set a static address -- this part wasn't automated
- vim /etc/netplan/orangepi-default.yaml

Do note that the interface name changes, it might not always be eth0
```yaml
network:
  version: 2
  renderer: networkd
  ethernets:
    end0: 
      dhcp4: false
      addresses:
        - 192.168.1.200/24
      routes:
        - to: default
          via: 192.168.1.1
      nameservers:
          addresses: [192.168.1.1]
```

The default is:
```yaml
network:
  version: 2
  renderer: NetworkManager
```