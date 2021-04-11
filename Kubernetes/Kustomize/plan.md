- Rename this file before pushing it to github!
    - Or at least chage the above line!

- Compile the kubectl docs!

### Kustomize plan!
1. Review that kustomization reference
    - https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/

- Create a deployment that using kustomize has two different environmnents with two different deployment images
    - one for staging!
    - one for prod!

- Use a patch for some sort of example you still have to define!
    - It would be particular interesting to focus on the difference between path and patcheStrat
      egicMerge


So in the end the plan is to create 3 sample configs
- deployment with different images
- patches
- patchStrategicMerge

- 