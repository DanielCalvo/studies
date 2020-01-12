### 204: End to end tests
- How do you validate the set up of your cluster to make sure it works?

#### Manual tests
- You can start by checking if the nodes on the cluster are healthy with `kubectl get nodes` and `kubectl get nodes`
- To check the control plane components, you can also: `kubectl get pods -n kube-system`
- If the control plane components are deployed as services, you can also check their status with: `service kube-apiserver status` and same for kube-controller-manager, kube-scheduler and then kubelet
- You can also try deploying a sample app to see if it works by doing:
    - `kubectl run nginx`
    - `kubectl get pods`
    - `kubectl scale --replicas=3 deploy/nginx`
    - `kubectl get pods` <- See if they got scheduled properly on different nodes

- Try exposing applications as services and see if you can access them:
    - `kubectl expose deployment nginx --port=80 --type=NodePort`
    - `kubectl get service`
    - `curl http://worker:31850`

- You can run many more manual tests to see if other features are working in the cluster such as secrets, configmaps, volumes, encryptions, networking, etc


####End to end tests
- All of these tests together make the end-to-end tests!
- These tests already exists and are on the kubernetes/test-infra repository. There are about 1000 end-to-end tests there!

- For each test it usually builds a namespace and then it creates the necessary objects within that namespace
- It then records results and deletes that namespace

- The test infra has tests, but it can also help you build., deploy, test and cleanup any version of kubernetes. Helpful if you're developing.
- Any solution that claims to be a kubernetes based solution must at least pass the comformance tests (about 160)
- The full e2e tests take about 12 hours to complete :o
- Conformance tests take about 1.5 hours
- Sonobuoy is also a tool that is often used to test k8s clusters

### 205: End to end tests - Run and Analyze
- On a k8s master node, provided you have Goland installed, run `go get -u k8s.io/test-infra/kubetest`
- Then run kubetest with the relevant version of k8s you're running by going: `kubetest --extract=v1.11.3` (don't forget to change the version for the one you're using)
- `cd kubernetes`
- `kubetest --test --provider=skeleton > testout.txt`
- If using the skeleton provider (meaning cluster is hosted on baremetal) you must pass certain parameters by environment variables, such as:
- `export KUBE_MASTER_IP="192.168.26.10:6433"`
- `export KUBE_MASTER=kube-master`
- If you don't want to run the entire e2e tests, you can pass the following to run only certain tests:
- `kubetest --test --provider=skeleton --test-args="--ginkgo.focus=Secrets"`
- `kubetest --test --provider=skeleton --test-args="--ginkgo.focus=\[Conformance\]"`

### 206: Smoke test
- For this lecture, mumshad goes over the smoke tests as set up on k8s the hardway here: https://github.com/kelseyhightower/kubernetes-the-hard-way/blob/master/docs/13-smoke-test.md
- It creates a secret, a deployment, a service, checks logs and attempts to exec something on a node

### 207: End to end tests
- Mumshad went over the tests here: https://github.com/mmumshad/kubernetes-the-hard-way/blob/master/docs/16-e2e-tests.md