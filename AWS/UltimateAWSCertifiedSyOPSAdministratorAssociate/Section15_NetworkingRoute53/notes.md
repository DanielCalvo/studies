- Chapter status: Took notes attentively, got the quiz a bit wrong
- Follow up on better understanding the difference between CNAMES and ALIAS, Stephan's explanation was a bit ambigious
- Rewatch lectures 195 to better understand domain delegaton and lecture 187 to better understand CNAME vs ALIAS

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
- Maps a domain name to a URL
- Use it when you need to redirect to a single resource
- You can't attach health checks to a simple routing policy
- If multiple values are returned, a random one is chosen by the client

### 189. Routing policy - weighted
- Controls the % of requests that will go to a specific endpoint
- You can assign different IP addresses with different weights to a name entry
- Sum doesn't have to be 100, AWS will sum and work with the average
- Helpful to test 1% of traffic on a new app version for example 
- Helpful to split traffic between two regions
- Can be associated with health checks, so if one instance isn't working properly you can stop sending traffic to it

#### Hands on
- New address entry, routing policy > weighted
- Add multiple records to the same name with different weights. Neat!
- When you use dig, you just get one address as a response, you're not aware of the weighted mechanism on the client side

### 190. Routing policy - latency
- Will redirect the user to the server that has the least latency to him
- Helpful when latency of users is a priority 

#### Hands on
- New address entry, routing policy > latency
- Depending on your geographical location, you will get a different DNS response, which will be the one with the lowest latency from you

### 191. Route 53 health checks
- An instance/endpoint is considered unhealthy if it fails 3 health checks in a row 
- Default health check interval is 30 seconds. It can be changed, but this incurs higher costs.
- In the background on Amazon's side, about 15 health chekers will check the endpoint health
- You can have HTTP, TCP and HTTPS health checks (no SSL verification)
- You can integrate these healthchecks with cloudwatch if you wanted to
- Health checks can be linked to Route 53 queries, so the behaviour of route 53 can change depending on what the health checks result

#### Hands on
- Created 3 health checks on route 53

### 192. Routing policy - failover
- You have a main EC2 instance and a failover one
- Route 53 has a health check pointing to the primary instance
- If the route check fails, it'll point you to the failover instance

#### Hands on
- New address entry, routing policy > failover
- When creating the main A rentry, set the failover type to primary and associate it with a healthcheck
- Create the second entry with the same name on the name entry and the routing policy of failover, type secondary. The secondary failover doens't need a health check

### 193. Routing policy - Geolocation
- Different from latency based!
- Routing based on user location 
- Ex: Traffic that originates from the UK should go to this specific IP
- You should create a default policy in case there's no match on location!

#### Hands on
- New address entry, routing policy > Geolocation
- You can then choose a location! Continents or countries
- Remember to create multiple A entries to the same name with different geolocation settings so different people from different places get different results from your DNS 
- Don't forget to create a default routing policy at the end!

### 194. Routing policy - Multi Value
- Use when you want to route traffic to multiple resources
- You can associate a route 53 health check with errors
- Returns up to 8 healthy records for each multi value query
- Multi value is not a good replacement for having an ELB!
- Diagram: Multiple values associated with one A entry. Each entry is associated with a health check. If the health check is unhealthy, route 53 no longer returns that when someone queries that name.

#### Hands on
- New address entry, routing policy > Multivalue answer, associate it with a healthcheck
- Create a bunch of records with the same name and multivalue routing policy!

### 195. 3rd party domains & Route 53
- Route 63 is also a registrar!
- A domain name registrar as an organization that manages the reservation of internet domain names
- Remember: Domain Registrar != DNS
- If you buy your domain on a 3rd party website, you can still use Route 53!
    - Create a hosted zone in route 53
    - Update NS records on the 3rd party website to use route 53 name servers

#### Hands on
- Go on google domains and on the settings, set google domains to use a custom name server, and in there set the AWS name servers
- Possible exam question: "How do you integrate a third party domain with route 53?" Create your domain elsewhere but have the name servers point to your AWS route 53 hosted zone

### 196. Section clean up

### Quiz