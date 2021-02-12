### 12. Starting the demo system
- Kiali, Jaegger and Grafana!
- Oof, there's no way to add the proxy container to a running pod, you just have to restart all of your pods
- You don't need to apply istio by default, and you also don't have to use a buncha yamls
- Just applied the yamls

### 13. Kiali Deeper Dive
#### Telemetry requirements
- Simpler than you might think!
- You need an Envoy sidecar proxy running in each pod you want to monitor
- The control plane needs to be running (istiod, kiali, jaeger, grafana)
- You don't need any specific istio yaml configuration! (no need for virtual services, gateways, etc, those are only if you want those services, they are not required for telemetry)

#### Dashboard demo
- Interesting, if the line between two services is gray it means it was not used in the time selected for the graph
- The line can be removed is there is no traffic. It's not a static picture of your system arch, it only shows parts that are sending and receiving traffic in the window that you have specified

- oooOOooOh you can double click on an item!
    - Shows incoming and outgoing routes from/to that object
- There's a legend button on the bottom left, useful to understand better what's up
- A workload represents all of the pods for a given service
    - The workload graph has response times though. Neat!

- oooOOooOh you can right click an item and there's a pop up menu!
    - You can show details on a service! Oh my word!
    - Overview! Traffic! Inbound Metrics! I like it
- Workload has even more options!
    - Overview, traffic, logs, inbound metrics and outbound metrics
    - There are even drag downs to choose a given pod on a workload and a given container in a pod. Neat!
- There are traffic animations! I love this tool! ahah
- The Kiali docs are massive: https://kiali.io/documentation/latest/features/
- You can also do modifications to the system through Kiali, so that comes next

- FOR LATER
    - Deep dive into what the different graphs do and what they can show you
    - Read the Kiali docs!

### 14. Kiali Dynamic Traffic Routing
- Autho argues: Just the observability can be enough of a feature to use Istio
- Other things you can do: Canary releases, a/b testing, you can even deploy things in prod only for engineers (a dark release)
    - This could mean you may not even need to have staging
- Right click on service on the UI
    - Show details
    - On the right of the screen: Actions > Suspend traffic! :o
    - All this is done by istio yaml, but these buttons can generate the yaml and apply that automatically
- These now exist:
    - `kubectl get virtualservices`
    - `kubectl get destinationrules`
- Uses to modifying things on the UI
   - Emergencies
   - Troubleshooting
   - Later of course, edit your yamls!

### 15. Distributed tracing overview
- opentracing.io
- Jaeger!
- We don't need to use the tracing libraries at all, because Istio already has that built in into the proxy
- One catch: You need to propagate headers

### 16. Using Jaeger UI
- Oh you can actually see how long the proxy took to do it's work at the end of the graph. Neat!
- Oh you can zoom in!
- You can also click the request to see more details. Neat!

### 17. Why you need to propagate headers
- When a proxy first receives a request from the outside world, it will add a header (request-id) to it
- These headers have to be forwarded: https://istio.io/latest/faq/distributed-tracing/#how-to-support-tracing
- If you don't forward the request ids, the sidecars will just create new request ids every time the code passes through them

### 18. What happens if you don't propagate headers
- You get broken traces :(

### 19. Metrics with Grafana
- The service and workload dashboards are the most useful ones
- The Istio Mesh dashboard is also pretty interesting