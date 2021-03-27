### 37. Cascading Failures 
- Circuit breaking aims to avoid cascading failures: https://en.wikipedia.org/wiki/Cascading_failure
- A circuit breaker monitors the success rate of the requests that a given service is making
- If too many requests fail, the circuit breaker will stop relaying requests
- Envoy has circuit breakers built in!
- Not everyone needs circuit breakers, but if it's easy for you to do, consider implementing them!

### 38. Configuring outlier detection
- Author sets service with 2 replicas: one and a bad one
- Author configures a DestinationRule with a trafficPolicy and outlierDetection

### 39. Testing Circuit Breakers
- The circuit breakers are applied on a pod by pod basis!
- Consecutive failures: A number of consecutive failures that a pod responds with
- Noice! In the example in the video after a pod responder with too many 502s, istio then routes the traffic away from that pod. Neat!
- The badly behaved pod was ejected for 30 seconds based on that bad configuration! That's really cool to see
- Author argues that circuit breakers can be complicated, go into them with your eyes open and paying attention!