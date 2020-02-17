### 8: Compute Options
- What is the simplest way to host a website on a cloud platform like GCP?
- Maybe your site is just 2 html pages
- Google App Engine: PaaS option: Serverless and ops free
- Compute Engine: IaaS option: Fully controllable down to OS
- Google Container Engine / GKE: Clusters of machines running k8s and hosting containers

### Use case for static website hosting:
- Use Google Cloud Storage. Create a bucket and upload data there.
- Or use static generators (like Hugo?)
- You can store your data on github and use a webhook to run a update script
- You can also use Jenkins and specify a post-build step to push artefacts to GCP

- For more advanced static site storage with SSL, CDN and other features
- Firebase hosting + Google cloud storage
- Also offers atomic deployment and 1 click rollback
- More like Google Cloud Storage++

### 9: Google Compute Engine (GCP)
- If you'd like to control load balancing of your web app, would you use SaaS, PaaS or IaaS?
- Use IaaS! With compute instances and load balancing/scaling rules defined by yourself
- There's a thing called Google Cloud Launcher that can help you out with GCP (like to launch Wordpress)
- You have loads of storage options!
    - Cloud Storage Buckets
    - Standard persistent disks
    - SSDs
    - Local SSDs
- You can also access from your VM
    - Cloud SQL
    - No SQL
- Load Balancing Options:
    - Network load balancing
    - HTTP load balancing
- Rule of thumb: The more sophisticated load balancing you wanna do, the higher on the OS stack you need to go
- You can also have Internal load balancing with compute instances using HAProxy and Consul

#### DevOps things you can with GCP:
- Compute engine Management with Puppet, Chef, Salt and Ansible
- Automated image builds with Jenkins, Packer ans K8s (?)
- Distributed load TEsting with K8s (?)
- CI with Travis CI
- Managing deployments with Spinnaker
- There are a lot of GCP equivalents to all of these