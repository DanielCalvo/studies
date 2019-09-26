### 182. Section introduction 
- Yay we'll check a bunch of stuff! DNS is cool. There are graphs!
- Also 3rd party domains

### 183. Route 53 Overview
- Managed DNS
- DNS is a collection of records and rules that help clients understand how to reach a server through a URL
- Most common records:
    - A
    - AAAA
    - CNAME: URL to URL
    - ALIAS: URL to AWS resource
- Route 53 can use public domain names that own (or buy)
- Private domain names that can be resolved only by your instances in your vpc (ex  myapp.company.internal)

#### Advanced features:
- Load Balancing
- Health checks
- Routing policies
- You pay 0.50 USD per hosted zone

### 184. Route 53 Hands on
- Global service
- Added an A record through the UI, very easy

### 185. Route 53 - EC2 setup
- Created an instance in Ireland and one in N. Virginia and one in Tokyo
- Created an application load balancer in Ireland pointing to one of those EC2 instances

### 186. Route 53 - TTL
- TTL: Time a client caches a DNS response as to not overwhelm a DNS server
- High TLL: About 24h. Less traffic on DNs, but possible chance of outdated records
- Low TTL: About 60 seconds. Lots of traffic on DNS, but records will be outdated for less time, also easy to change records 
- The `dig` command line tool show's the TTL. Neat, I didn't know that.

### 187. CNAME vs Alias
- Popular question at the exam!
- AWS resources (load balancer, cloud front, etc) expose an AWS URL (ex: lb-123.us-east-2.elb.amazonaws.com) but you want that to be myapp.mydomain.com
- CNAME Points a URL to another URL (app.mydomain.com > blabla.anything.com)
- CNAMES are only for non root domains! (ex: Needs to be something.mydomain.com, it can not be mydomain.com)
- You can't cant CNAME from myrootdomain.com into a load balancer, for instance
- ALIAS points a URL to an AWS resource (app.mydomain.com > blabla.anything.com)
- It works both for root domains and nonroot domains
- Alias is free of charge and has a native health check
- Possible exam question: "We want to point mydomain.com to our load balancer, do you use CNAME or Alias?" You use alias! CNAMEs don't support root domains!

### 188. Routing policy - Simple

### 189. Routing policy - weighted

### 190. Routing policy - latency

### 191. Route 53 health checks

### 192. Routing policy - failover

### 193. Routing policy - geolocation

### 194. Routing policy - multi value

### 195. 3rd party domains & Route 53

### 196. Section clean up

### Quiz