# argoproj.github.io

## Argo Workflows
- For orchestrating workflows on k8s, think data processing. Allows you to configure a sequence of tasks. Neat!
- Conceptually it reminds me of step functions, but without the serverless part 

### Use cases
- ML pipelines
- Data and batch processing (neat)
- Infra automation and ci/cd

## Argo Events
- Allows you to react to events from certain sources (s3, kafka, http requests)

## Argo Rollouts
- Advanced k8s deployment strategies like canary, bluegreen and the like
- Also gives you: traffic shifting, automated rollbacks, and ingress integration. Neat! 

## ArgoCD
- This one you know, the gitops tool that reconciles configuration from git into a cluster

## Other
- Neat: https://github.com/akuity/awesome-argo
    - Argocd autopilot seems interesting

---

# Fluxcd

## Flux
- The gitops tool that reconciles configuration from git into a cluster. Slightly different than argo, but same deal

## Flagger
- Pogressive delivery for k8s, seems to be equivalent to argocd rollouts
- Has: Progressive releases with canary, ab testing and bluegreen

It seems there's not much else to fluxcd at this time, looks like the community around it isn't as large as argo.