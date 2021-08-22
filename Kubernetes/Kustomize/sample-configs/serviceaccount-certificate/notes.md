This seems interesting: 
- https://github.com/kubernetes-sigs/kustomize/blob/master/examples/patchMultipleObjects.md

The Kustomize docs aren't the most intuitive ones...

Bleh: https://github.com/kubernetes-sigs/kustomize/issues/1113

## So you can't replace arbitrary fields as command line arguments apparently
But you can set:
- image
- label
- namespace
- nameprefix
- namesuffix
- replicas

## Interesting: 
- https://learnk8s.io/templating-yaml-with-code
- https://kubectl.docs.kubernetes.io/faq/kustomize/eschewedfeatures/

Kustomize is opinionated in the way that it asks you to keep all the config in git/code


