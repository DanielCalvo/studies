### Pre-exam checklist
- Don't forget to run the browser compatibility tool on whatever system you're going to use before doing the exam!
- Use the code - DEVOPS15 - while registering for the CKA or CKAD exams at Linux Foundation to get a 15% discount.
- Exam Curriculum (Topics): https://github.com/cncf/curriculum
- Exam Tips: http://training.linuxfoundation.org/go//Important-Tips-CKA-CKAD

When done with the course, check the handbook to see if there's anything you missed:
- Candidate Handbook: https://www.cncf.io/certification/candidate-handbook

### FOLLOW UPS FROM LATER:
#### Section 2
- Memorize how to start a pod on the command line
- Get good with kubectl! Remember to remember all the flags and options and output settings!
- (Follow up from Dani: kubectl create returns more useful error messages than kubectl apply)
- (Follow up from Dani: Can you write a replicaset definition from memory?)
- (Follow up from Dani: Learn how to use the edit command to get running resources and make a copy from them)
- Bookmark for the exam: https://kubernetes.io/docs/reference/kubectl/conventions/ and make sure to familiarize yourself with the generator commands!
- (Follow up from Dani: Carefully study the output of kubectl get services to know exactly what everything means. Give that a go on kubectl describe services too)
- See imperactive commands with kubectl on lecture 35

#### Section 3
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

#### Section 5:
- Can you investigate all kubectl subcommands?


Check the FAQ here to mentally prepare a list of commands you'll use:
https://www.cncf.io/certification/cka/

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

