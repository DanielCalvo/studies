network:
  version: 2
  ethernets:
    eno1:
      dhcp4: false
      addresses:
        - 192.168.1.211/24
      routes:
        - to: default
          via: 192.168.1.1
      nameservers:
          addresses: [192.168.1.1]
