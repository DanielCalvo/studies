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