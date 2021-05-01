
- This became obsolete. Re-evaluate & redo, or use it for something else, or archive

## ACTION PLAN
### Part 1
- Create an EKS cluster, using
    - Terraform: The entire cluster should be deployed with "terraform up"
    - Ingress: An AWS ALB (deployed through terraform)
    - DNS entry: A wildcard SSL entry pointing to the load balancer
    - Then: Deploy a hello world web app with an ingress

#### Notes on Part 1:
- Terraform destroy isn't working right -- why?
    - Maybe write that the helm stuff depends on the cluster so that it gets deleted first when you destroy it?
- Or does it work after 2 tries?
- Can you quickly code something to list all aws resources?

Idea 1: Hmm, use helm to template your ingress?
Idea 2: Get the ARN of your ingress using a data source in terraform and generate an ACM certificate from there?
- What is your source of truth for your deployed branches? Hmm
- I think you can do the above separately
Idea 3: You can use helm to template your ingress resource to work with an external "source of truth" regarding which branches you have deployet to manage that on an ingress

### Part 2
1. Deploy all the infra with Terraform
2. Deploy a helm chart with a deployment, service and ingress (you can also deploy it manually for now)
3. Get the ARN of your ingress using a data source in terraform  <- Next step is here
4. Point some DNS to your ALB so you can access your app with mylab.dcalvo.dev (wildcard dns later perhaps?)
5. TLS setup: Generate a certificate for that domain, add the validations steps, and google how to terminate tls on an ALB with k8s

### Misc notes
- See if it's possible to use the ALB chart here instead of the helm incubator: https://aws.github.io/eks-charts/

### Reference notes:
- https://docs.aws.amazon.com/eks/latest/userguide/alb-ingress.html
- https://github.com/kubernetes-sigs/aws-load-balancer-controller
- Docs
    - https://kubernetes-sigs.github.io/aws-load-balancer-controller/latest/

### Off topic interesting ideas
- https://aws.amazon.com/blogs/containers/api-gateway-as-an-ingress-controller-for-eks/
