do you really need large instaence for karpenter or for your small use case can you have something smaller?
also try karpenter with just 1 node


## brief checks
aws eks list-clusters

aws eks update-kubeconfig --name k8s-gitops


## on
kubectl apply -f karpenter.yaml
ec2nodeclass.karpenter.k8s.aws/default created
nodepool.karpenter.sh/default created

## troubleshootig
How do you get all the logs of a kubernetes deployment that has more than one pod again?
use a label: kubectl logs -l app=my-app -f

how you got to the bottom of this: read the karpenter logs

Ask AI how could you improve this setup and see what he comes up with
 ask AI to see if it can double check that there are any straggling resources related to this set up

 also: when deleting the clusters some instances persist.  is there anyway you can insure they are deleted? do you need to script something or can you handle this within terraform?

## mistakes made
- You renamed everything and forgot to add the auto discovery tags on the vpc subnets
- you also forget to change the role name on the carpenter yama definitions since you renamed the cluster
- Additionally it would seem the carpenter and inflate yams were not applied by default --  I wonder if this was a mistake I made or something that was ona read me somewhere and I didn't read it

## Observations
 why is the inflate pod running on a note that apparently is for carpenter only?