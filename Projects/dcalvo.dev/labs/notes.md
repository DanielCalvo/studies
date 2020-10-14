#### ALB setup
- https://docs.aws.amazon.com/eks/latest/userguide/alb-ingress.html

0.
`aws eks --region eu-west-1 update-kubeconfig --name training-eks-cTakQmD9`

2.
```shell script
eksctl utils associate-iam-oidc-provider \
       --region eu-west-1 \
       --cluster training-eks-cTakQmD9 \
       --approve
```
3.
```shell script
curl -o iam-policy.json https://raw.githubusercontent.com/kubernetes-sigs/aws-alb-ingress-controller/v1.1.8/docs/examples/iam-policy.json
```

4. Will tell you it already exists:
```shell script
aws iam create-policy \
    --policy-name ALBIngressControllerIAMPolicy \
    --policy-document file://iam-policy.json
```

5.
```shell script
kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/aws-alb-ingress-controller/v1.1.8/docs/examples/rbac-role.yaml
```

6.
- Create an IAM role named eks-alb-ingress-controller
- Attach the ALBIngressControllerIAMPolicy IAM policy that you created in a previous step to it

```shell script
AWS_ACCOUNT_ID=$(aws sts get-caller-identity --query "Account" --output text)
OIDC_PROVIDER=$(aws eks describe-cluster --name training-eks-cTakQmD9 --query "cluster.identity.oidc.issuer" --output text | sed -e "s/^https:\/\///")
```

```shell script
read -r -d '' TRUST_RELATIONSHIP <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Federated": "arn:aws:iam::${AWS_ACCOUNT_ID}:oidc-provider/${OIDC_PROVIDER}"
      },
      "Action": "sts:AssumeRoleWithWebIdentity",
      "Condition": {
        "StringEquals": {
          "${OIDC_PROVIDER}:sub": "system:serviceaccount:kube-system:alb-ingress-controller"
        }
      }
    }
  ]
}
EOF
echo "${TRUST_RELATIONSHIP}" > trust.json
```
`aws iam create-role --role-name alb-role --assume-role-policy-document file://trust.json --description "alb-role"`

`aws iam attach-role-policy --role-name alb-role --policy-arn=arn:aws:iam::452788477576:policy/ALBIngressControllerIAMPolicy`

Hmm is this right? Not sure if the role name is right
```shell script
kubectl annotate serviceaccount -n kube-system alb-ingress-controller \
eks.amazonaws.com/role-arn=arn:aws:iam::452788477576:role/alb-role #Or is it alb-role?
```

kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/aws-alb-ingress-controller/v1.1.8/docs/examples/alb-ingress-controller.yaml


#### K8s questions:
- How to upgrade a cluster?
- How to handle user access?
- Trying out a Chaos tool (it breaks things!)
- How to make 2 clusters talk to each other?
    - Internal load balancer?
- See if you can add HTTPs termination with the cert managed by amazon on an ALB

### Further ideas:
- Maybe rename the dns_zone to general_infra and put the VPC config in there? 

### Reading to get started
- https://learn.hashicorp.com/tutorials/terraform/eks
- https://aws.amazon.com/blogs/startups/from-zero-to-eks-with-terraform-and-helm/

- Do a K8s one with backups of everything for when the cluster goes down and restores for when it comes back up
- Check the sysadmin book and the AWS certification course for infrastrcuture blueprints and implement them!
