# Single host, basic k3s install
- https://docs.k3s.io/quick-start
`curl -sfL https://get.k3s.io | sh -`

# Accessing the cluster from your desktop
On the cluster node: `chmod +r /etc/rancher/k3s/k3s.yaml`
On your local:
```shell
scp root@192.168.1.201:/etc/rancher/k3s/k3s.yaml ~/.kube/config
sed -i 's/127.0.0.1/192.168.1.201/g' ~/.kube/config
```

# App with ingress:
See [k3s_app](./k3s_app)

## Installing metallb

## Reinstalling everything
(on the k8s host)

```
/usr/local/bin/k3s-uninstall.sh
(then i usually restarh the machine)
/usr/local/bin/k3s-uninstall.sh
curl -sfL https://get.k3s.io | sh -
```

Then copy the config again to your local machine
scp root@192.168.1.201:/etc/rancher/k3s/k3s.yaml ~/.kube/config