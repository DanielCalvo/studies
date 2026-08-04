## network set up 

```
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
```

# copy over your public ssh key
ssh-copy-id daniel@192.168.1.211
ssh-copy-id daniel@192.168.1.212

## changed root passwd but it just wont let me copy the keys -- figure this out later
ssh-copy-id root@192.168.1.211
ssh-copy-id root@192.168.1.212

