### 27. Session affinity (stickiness) 
- Is it possible to use the weighted destination rules to make a single "stick" to a canary?
- No, not using the weighting functionality of a virtual service
- For DestinationRule, there's a spec.trafficPolicy.loadBalancer field!
- You can hash based on cookie, http header, source ip and... ring size?
- This is a bit convoluted. The weighting rules are applied before the load balancer rules
- You can also use headers, but apparently you need to forward them?
- Your frontend can generate a header for your particular user
- Session affinity does not work simultaneously with weighted routings?
- It's okay to use this for performance enhancements (like some sort of caching) but don't use this to establish a session
- FOR LATER: This lecture is a bit difficult to grasp. Rewatch it again laterto better grasp concepts 

### 28. What is consistent Hashing useful for?
- 