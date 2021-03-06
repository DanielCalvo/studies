### Pre-exam checklist
- You should try ddoing some exercises that combined everything
- Create your own questions! Cluster set up & maintenance seems difficult.
- You gotta become familiar with the docs
- Recheck lecture 8 with "Resources for this lecture". There might be something useful there, the taints & tolerations pdf appeared interesting!
- Try doing all the practive sections without checking your course resources, just checking the kubernetes reference as allowed per exam
- Don't forget to run the browser compatibility tool on whatever system you're going to use before doing the exam!
- Use the code - DEVOPS15 - while registering for the CKA or CKAD exams at Linux Foundation to get a 15% discount.
- Exam Curriculum (Topics): https://github.com/cncf/curriculum
- Exam Tips: http://training.linuxfoundation.org/go//Important-Tips-CKA-CKAD
When done with the course, check the handbook to see if there's anything you missed:
- Candidate Handbook: https://www.cncf.io/certification/candidate-handbook
- To test GateOne: `docker run -p 443:443 dcwangmit01/gateone`

### To review
- Exam curriculum and resources: https://www.cncf.io/certification/cka/
- Maybe take the most difficult questions from each quiz and practice doing them only checking the official doc?

### Section 2
- Memorize how to start a pod on the command line
- Get good with kubectl! Remember to remember all the flags and options and output settings!
- (Follow up from Dani: kubectl create returns more useful error messages than kubectl apply)
- (Follow up from Dani: Can you write a replicaset definition from memory?)
- (Follow up from Dani: Learn how to use the edit command to get running resources and make a copy from them)
- Bookmark for the exam: https://kubernetes.io/docs/reference/kubectl/conventions/ and make sure to familiarize yourself with the generator commands!
- (Follow up from Dani: Carefully study the output of kubectl get services to know exactly what everything means. Give that a go on kubectl describe services too)
- See imperactive commands with kubectl on lecture 35


### Section 3
- A pod failed to start because the scheduler wasn't running. What components are always running on vanilla k8s? (corends, etc, proxy, what else? ) (Follow up from Dani: Which pods from Kubernetes should be available on a default install of Kubernetes?)
3:
The cert: https://www.cncf.io/certification/cka/
Candidate handbook: https://www.cncf.io/certification/candidate-handbook
Exam tips: https://training.linuxfoundation.org/wp-content/uploads/2019/05/Important-Tips-CKA-CKAD-May.pdf
- Is there some vi config for auto ident you can use?
- (Follow up from Dani: How can you see "events" in Kubernetes? Is it by describing the pod?)
- Remember to how to use args in case they're needed on pod creation: Create a pod definition file in the manifests folder. Use command kubectl run --restart=Never --image=busybox static-busybox --dry-run -o yaml --command -- sleep 1000 > /etc/kubernetes/manifests/static-busybox.yaml
- Will the exam ask you to ssh into a worker node and look for a static pod?
- (Follow up from Dani: Is there a definitive way to say if pods are static pods, other than their pod names?)
- Can you create a pod with certain arguments from the command line? like busybox sleep? by memory?
- I didnt do "running your own scheduler", that was really boring
- How to manually generate a config to connect to a cluster?
- I think I'm missing lecture 52 or something
- Recommended for exam: https://kubernetes.io/docs/reference/kubectl/conventions/
- Memorize which k8s components should always be running

### Section 5:
- Can you investigate all kubectl subcommands?
- (Follow up from Dani: Creating a pod loading the configs from a configmap was really tough. You need to work on this!
- What objects can you find on the kube-system namespace that you can copy and paste around?
- Give pod design patterns a google! 
- Give liveness and readiness probes a look, even though these are only part of the CKAD course

### Section 6:
- ETCD backup is pretty insane to memorize and... the official doc seems lacking?
- Improving your knowledge of kubeadm could help
- Hey it would be neat to backup etcd and see if you can copy & paste your cluster around with the etcd binary data

### Section 7:
- (Follow up from Dani: There's definitely some learning to be done as far as troubleshooting a broken cluster (as in when you can't use kubectl)
- The TLS troubleshooting at lecture was very boring, reapproch at the future if you feel like
- (Follow up from Dani: Can you investigate all API endpoints and interact with them later?)
- (Follow up from Dani: How to create other users in other namespaces?)
- Redo the lecture 128 practice test if you feel like it, you forgot to differentiate between roles and roles bindings and cluster roles and cluster role bindings 
- (Follow up from Dani: It got really boring copy and pasting roles and rolebindings around. Do some exercises on these subjects later if you feel like it)
- Redo the practice test for network policies if you're feeling like it (the hands on part requires a lot of copying and pasting)

### Section 8:
- Not really exam related, but it would be interesting to see how the Container Storage Interface talks to EBS on AWS.
- (Follow up from Dani: The explanation here was nice, but very generic. I think you should try implementing a storage solution on cloud and/or on premise and seeing how this works)

### Section 9:
- Try traefik as an ingress!
- (Follow up from Dani: ingress configuration is very large, perhaps you should have a file with multiple examples?)
- (Follow up from Dani: Run commands from vim, also learn tmux to split your terminal)
- (Follow up from Dani: Remember! The ingress needs to be on the same namespace as the service!)

### Section 10:
- Skim Kelsey's k8s the hardway one last time just to make sure you remember more or less how it goes
- Maybe watch all of Mumshad's videos on this section, they could be helpful

### Section 13:
- It would be helpful to explore all the flags for all the kubectl commands (`kubectl logs -f --previous` seemed helpful)
- Explore and map out the documentation pages!
- Give the docs a thorough look! Map it out!

Check the FAQ here to mentally prepare a list of commands you'll use:
https://www.cncf.io/certification/cka/

### Section 14:
- I did not manage to get json path queries working, they were confusing! Revisit later
- Give custom output and columns and loops a revisit, they're difficult to work with.

Kodekloud and Katakoda seem worth a look!

It would be nice to have a sample config of a pod with two containers somewhere.
Also it would be cool to build a bunch of stuff at the end of the course with several objects (deployments, services, secrets and so on)

Madman ideas: Automate k8s the hard way with some golang
Absolute madman idea: Connect to the kubernetes api manually and write a simple client!

Catch up on the "ip" command, as ifconfig is now deprecated.

I think you should also investigate all the --options when running k8s the hard way (what do they mean?)

Follow up on some experimenting on k8s on AWS, is kops still the only way forward?

Post course:
CKA search on reddit
Go through all the docs: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/

Put on your inbox later:
- RAFT protocol, CAP theorem