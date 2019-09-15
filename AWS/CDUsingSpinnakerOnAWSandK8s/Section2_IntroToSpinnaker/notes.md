
### 4. What is Spinnaker and it's history

- Open source multi-cloud CD platform
- Initially developed by Netflix
- Built for releasing software changes with high velocity and confidence
- Designed with pluggability in mind
- Suppors all major cloud providers

#### Spinnaker history
- Development started in 2014 by Netflix
- Tested by hundreds of teams with millions of deployments
- Spinnaker is the replacement of Netflix's Asgard
- Open sourced in 2015
- New but mature


### 5. Why you should use Spinnaker
- Spinnaker's goal is to deploy reliably
- Spinnaker offers deployment pipelines and allows for easy rollbacks
- Built for high velocity, to deploy fast and often
- Flexible pipeline management system
- Big active community
- Integrates with all the major cloud providers and Kubernetes

#### It's feature rich yo
- CI Integrations
- CLI for setup and admin
- VM bakery
- RBAC
- White-listed execution Windows
- Monitoring Integrations
- Deployment strategies
- Notifications
- Manual judgements
- Chaos Monkey Integration :D

#### More on the why
- The Kubernetes deployment API is not as advanced as Spinnaker
- Spinnaker complements Kubernetes by using powerful deployment pipelines allowing for multiple deployment strategies
- Kubernetes support in Spinnaker is feature complete
- You have direct access to the Spinnaker API without being forced to use the UI
- Documentation for Spinnaker is fully available
- Deployments are built-in and no custom scripting is required! This sounds interesting


#### Spinnaker workflow
1. Code check-in
2. CI
3. Bake
4. Deployment
5. Prod-deploy + test
6. Hotfix
7. Canary
8. Live

It's meant to be a solution to replace a lot of custom scripts/tools

### 6. Installing Spinnaker
- Halyard is a tool to configure, update and install Spinnaker
- Author recommends using S3 storage for Spinnaker data on AWS
- Minio: S3 compatible storage

Install chapters are not very note worthy (pun intended)

### 9. Spinnaker concepts
- CD Platform. Manages machines and deployments.
- Machines will be grouped into clusters
- Releases will be managed by deployments  

#### Spinnaker has 2 core sets of features to achieve this:

##### Cluster management:
- View and manage your resources in the cloud (EC2 instances for instance, or a k8s cluster, or GCP app engine)
- A cluster can be either at infrastructure level (Virtual machines) or at Orchestrator/PaaS level (Kubernetes, Google App engine)  

##### Deployment management: 
- Create and manage continous delivery workflows
- You're going to use workflows to describe the way applications should be deployed on your cluster
- These workflows you'll create are called the delivery pipeline in Spinnaker


### 10. Spinnaker Terminology
- Account: Contains a generic account name and a credential that is able to authenticate against something, such as: 
    - AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY
    - Or in K8S: Credentials to the docker registries, credentials to the Kubernetes api-server
- Instance: Can be a VM instance in AWS or a Kubernetes pod
- Server group: A collection of instances, running the deployed software 
- Cluster: A spinnaker deployment
    - A logical grouping of server groups
    - Spinnaker uses its own orchestration and can be different from built-in mechanisms

### 11. Providers
- In Spinnaker, the (cloud) provider is the set of virtual resources that Spinnaker has control over
    - ex: When using IaaS, it's the AWS/Azure/GCP resources
- The provider is where you'll deploy your resources, like server groups.
- A Server group is what Spinnaker calls the "source of deployable artifacts"
- You deploy applications on server groups, which are tied to a cloud provider 
- To get access to a provider, you need to configure an account.
    - You can have multiple accounts per provider, to separate environments for example

#### Supported providers
- Google App Engine
- AWS
- Azure
- DC/OS
- Google Compute Engine
- K8s (legacy) and K8s v2 (manifest based)


### 12. Deployment strategies
- Red/black (blue/green)
    - Starts the same amount of instances in a new server group
    - When the new group becomes healthy, the LB sends traffic to the new group
- Rolling red/black (blue/green)
    - Gradually starts instances in a new group and removes instances from the old group until all instances are running in the new deployment
- Canary analysis
    - Replaces a small amount of instances
    - A certain percentage of traffic is sent to the new version (ex: 1-5%)
    - This is to test whether the new version doesn't have any issues
    - If no issues arise within a certain timespan, the deployment can continue

### 13. Pipelines
- Part of the deployment management feature of Spinnaker
- You manage how your application gets deployed on your cluster using a deployment pipeline
- Multiple pipeline stages form the Spinnaker pipeline
- Parameters can be passed between stages
- Pipelines can be triggered in the UI by the user, or automatically by other software, such as Jenkins