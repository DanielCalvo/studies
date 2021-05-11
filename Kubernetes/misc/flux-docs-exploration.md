#### TODO / Questions to answer
- How can flux help with cd/cd though?
- The "toolkit components" section seems nice if you wanna nerd out a bit more later too

### What flux does, in a nutshell

> Flux is a tool for keeping Kubernetes clusters in sync with sources of configuration (like Git repositories), and automating updates to configuration when there is new code to deploy.

In other words
1. You push configs to git
2. Flux reads configs from git
3. Flux applies those configs to a cluster
4. Flux keeps repeating steps 2 and 3 so that whatever configs/instructions you push to a repository are reflected in a cluster. This is called the reconciliation loop.


### Things important to understand before getting started with flux
 - Listed here: http://127.0.0.1:8000/get-started/
- Gitops, sources, kustomization

### Other notes
- The get started guide is OK http://127.0.0.1:8000/get-started/
- How to install flux in a cluster using terraform? 
        - But you need to write how to perform that install with terraform, that would be more better!

### Getting started according to the docs
- flux bootstrap
- use the command line to create a GitRepository yaml (pod info source)
- another command is also used to create the kustomization 

- flux create source|kustomization
- flux get is useful too
- explore all the arguments that the flux command line tool takes!

### Guides

#### Installation
- `flux bootstrap` can be used to install flux from yamls in a repo!
- You can also generate the install yamls with `flux install --export`
- You can extract the... uh I forgot
- OooOooOooh there is a dev install!
#### Managing helm releases
- Guide just goes over how you can have a HelmRepository and a HelmRelease
- You can also pass variables to the configmap
#### Notifications
- There's a component for you to spam yourself through slack when a given reconciliation fails. There's also an alert component!
#### Push receivers
- flux is pull based but you can also push notifications to it if you have to to notify it of something
#### Monitoring
- You can use prometheus to monitor the flux control plane! Set up seems easy enough
#### Sealed secrets
- You can instal kubeseal on the cluster to manage secrets
#### Mozilla SOPS
- You can also use SOPS to upload encrypted secrets to a repo and configure flux to decrypt those internally. Neat!
#### Automate image updates to Git
- You can use Flux's image toolkit to detect changes to certain docker images on container registries and apply them to a cluster, really neat!
- The components are ImageRepository (a repo) ImagePolicy (which tags you want to scan) ImageUpdateAutomation (updates data about those on your git repo so you do gitops)
- You can also add an image policy as a comment (?) to any k8s resource, like HelmReleases, Kustomizations and so on
- You can also use a PR workflow to review image changes before they are applied to your cluster, noice!
- You can also set up a webhook receiver to trigger a deployment as soon as an image is pushed to your container registry, you can configure dockerhub to send a webhook to flux for example (noice)
- You an also suspend these automations during an incident (from the command line or adding `spec.suspend` to an ImageUpdateAutomation
- ImageRepository still does not have native auth to container repos like ECR, so you have to implement a workaround for that which is hacky :/
- The image toolkit is still in Alpha!


### tmp notes
- Following dev install: https://toolkit.fluxcd.io/guides/installation/#dev-install
- You can install a kustomize manifest and a helm release... kinda no big deal. Maybe useful to validate things

#### Terraforming it though
- This: 
- Lead me here: https://github.com/fluxcd/terraform-provider-flux












