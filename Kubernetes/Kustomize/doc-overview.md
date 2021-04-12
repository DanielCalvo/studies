# Overview
- Just going through: https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/
- A quick overview of Kustomize's docs so I can have a rough idea of what functionalities the tool has!

### 1. bases
- Add resources from a kustomization dir
    - Uh, deprecated

### 2. commonAnnotations
- Add annotations to add all resources. If the anotation is already present on the resource, the value is overriden

### 3. commonLabels
Add labels and selectors to add all resources. Simple!

### 4. components
- This is a larger section -- there's an entire guide on it. Check later!

### 5. configMapGenerator
- Generatea ConfigMap resources
    - Large section. Check later!

### 6. crds
Adding CRD support

### 7. generatorOptions
Control behavior of ConfigMap and Secret generators.

### 8. images
Modify the name, tags and/or digest for images.

### 9. namePrefix
Prepends the value to the names of all resources and references.

### 10. namespace
Adds namespace to all resources.

### 11. nameSuffix
Appends the value to the names of all resources and references.

### 12. patches
- Patches can add or override fields on resources

### 13. patchesJson6902
Patch resources using the json 6902 standard

### 14. patchesStrategicMerge
Patch resources using the strategic merge patch standard.

### 15. replicas
- Change the number of replicas for a resource

### 16. resources
Resources to include.

### 17. secretGenerator
Generate Secret resources.

### 18. vars
Substitute name references.