### 131. CloudFront Overview
- CloudFornt is a content delivery network (CDN)
- It improves read performance, content is cached at the edge
- 216 points of presence globally (edge locations)
- Features: DDoS protection, integration with shield, AWS Web Application Firewall
- Can exporte external HTTPS and can talk to internal HTTPS backends

#### CloudsFront - Origins
- S3 Bucket
    - For distributing files globally and caching them at the edge
    - You also get enhanced security with CloudFront Origin Access Identity (OAI)
        - This allows your S3 bucket to only allow communication from CloudFront and nowhere else
    - CloudFront can also be used as an ingress to upload files to S3
- Custom Origin (HTTP)
    - Application Load Balancer
    - EC2 Instance
    - S3 website
    - Any HTTP backend you want (your own premise server, for example)

#### CloudFront at a high level 
- A bunch of edge locations all around the world, all connected to the origin (S3 bucket or something else, like your HTTP backend)
- Client sends an HTTP request to CloudFront
- Edge location forwards the request to the origin (including the query string and the request headers)
- Origin responds to the cache location, cache location then caches the response based on the cache settings that you have defined and return the response to the client
- The next time another client makes a similar request, the edge location will look into the cache before forwarding the request to the origin
    - (that's the whole purpose of having a CDN)

#### CloudFront - S3 as an Origin
- If you have some people in Los Angeles and an edge location in Los Angeles, CloudFront fetch the data from your S3 bucket and serve the data from there
- For the edge location of cloud front to access your S3 bucket, it uses an OAI (Origin Access Identity, it's an IAM Role) and using that role it's going to access your S3 bucket and the bucket policy will work with that role and send the file over
- This also works with other edge locations

#### CloudFront - ALB or EC2 as an Origin
- Your EC2 instances must be public (from an HTTP standpoint)
- Users access an edge location, and the edge location access your EC2 instance. It traverses (?) a security group, so this security group must allow the IP of CloudFront edge locations into the EC2 instances
- There is a public list of IP addresses for CloudFront that you can get
- If you use an ALB as an origin, it must also be public

#### CloudFront Geo Restriction
- You can restrict who can access your distribution
    - Whitelist: Allow your users to access your content only if they're in one of the countries in a list of approved countrs
    - Blacklist: Block users from certain countries
- "Country" is determined using a 3rd party Geo-IP database 
- Use case: Copyright Laws to control access to content

#### CloudFront vs S3 Cross Region Replication
- CloudFront:
    - Uses Global Edge Network
    - Files are cached for a given TTL (maybe a day)
    - Great for static content that must be available everywhere
- S3 Cross Region Replication
    - Must be set up for each region you want replication to happen
    - Files will be updated in near real time
    - Read only
    - Great for dynamic content that needs to be available at low-latency in a few regions
- CloudFront is for caching globally and S3 cross region replication for replication into selected regions

### 132. CloudFront with S3 - Hands on
- We'll create an S3 bucket
- We'll create a CLoudFront distribution
- We'll create an Origin Access Identity ("User" of CloudFront that will be accessing the S3 bucket)
- We'll limit the S3 bucket to be accessed only using this identity

#### Hands on
- Create bucket
- CloudFront > Create distribution > Web > Get Started
- Origin domain name: The bucket you just created
- Restrict Bucket Access: Yes
- Create new Identity
- Grant Read Permissions on Bucket: Yes pls update
- Force HTTPs
- Scroll all the way down and Create Distribution
- Distribution can take a few, maybe 10 minutes to get created
- You can see that you have an Origin access identity that was created
- If you go to your bucket > Permissions > Bucket Policy 
- You can see that a policy was automatically created for CloudFront. Neat!
- After the CloudFront is done spinning up, trying to open it redirects you to the bucket and you get an access denied message
    - This is a DNS issue and takes about 3 hours to resolve :(
- But you can make the file public in the S3 bucket to fix this temporarely
    - Hey awesome this works!
- Don't forget about the Origin Access Identity! It's important for the exam

### 133. [SAA-C02] CloudFront Signed URL / Cookies 
- Use case: You want to distribute paid shared content to premium users over the world
- For this, you can use CloudFront Signer URL / Cookie, We attach a policy with:
    - Includes URL expiration
    - Includes IP ranges to access the data from
    - Trusted Signers (which AWS accounts can create signed URLs)
- How long should the URL be valid for?
    - Shared content (movies or music): make it short, maybe a few minutes
    - Private content (private to user): you can make it last for years
- Signed URL = access to individual files (one signed URL per file)
- Signed Cookies = access to multiple files (one signed cookie for many files)

#### Diagram
- You have cloudfront with it's edge locations connecting to amazon S3 with it's OAI
- A client connects to an application (think video website)
- Application uses the AWS SDK to Generate a signed URL
- Application sends that signed URL to the client
- Client is able to use that signed url to get that object from cloudfront

#### CloudFront Signed URL vs S3 Pre-Signed URL
- CloudFrount Signed URL
    - Allows access to a path, no matter the origin (works with other HTTP backends, not just S3)
    - It uses an account wide key-pair, only the root user can manage it
    - Can filter by IP, path, date, expiration
    - Can leverage cache features
- S3 Pre-Signed URL
    - Issues a request as the person who pre-signed the URL
    - Has the same IAM rights as the person who granted the signed URL (Stephan: Uses the IAM key of the signing principal)
    - Limited lifetime

### 134. [SAA-C02] AWS Global Acelerator - Overview

#### Global users for our application, aka the problem Global Acelerator tries to solve
- You have deployed an application and have global users who want to access it directly
- But the app is only deployed in one region!
- They go over the public internet, which can add a lot of latency due to many hops
- We wish to go as fast as possible through the AWS network to minimize latency

#### Unicast IP vs Anycast IP
- Unicast IP: One server holds one IP address
- Anycast IP: All servers hold the same IP address and the client is routed to the nearest one :o

#### AWS Global Accelerator
- Leverages the AWS internal network to route to your application
- So what happens is that instead of a client in America trying to access an app in India going through the entire internet, this user is going to talk to the closes edge location and from this edge location it will go internally through Amazon's network to your app
- 2 Anycast IP are created for your application 
- The Anycast IP sends traffic directly to Edge applications
- The edge locations send the traffic to your application

#### AWS Global Accelerator part 2
- Works with Elastic IP, EC2 instances, ALB, NLB, public or private
- Consistent performance over AWS's internal network
    - Intelligent routing to lowest latency and fast regional failover
    - No issue with client cache (because the IP doesn't change)
    - Internal AWS network
- Health checks
    - Global Accelerator performs a health check of your appliations
    - Helps make your application global (failover less than 1 minute for unhealthy)
    - Great for disaster recovery (thanks to the health checks)
- Security
    - Only 2 external IPs need to be whitelisted
    - DDoS protection thanks to AWS Shield

#### AWS Global Accelerator vs CloudFront
- They both use the AWS global network and it's edge locations around the world
- Both services integrate with AWS Shield for DDoS protection
- CloudFront will improve the performance of:
    - Cacheable content (images and videos)
    - Dynamic content (such as API acceleration and dynamic site delivery)
    - Content is served at the edge 
- Global Accelerator
    - Improves the performance for a wide range of applications over TCP or UDP
    - Proxying packages at the edge to applications running in one or more AWS regions
    - Good for for non HTTP use cases such as gaming (UDP), IoT (MQTT) or VoIP
    . Good for HTTP use cases that require static IP addresses
    . Good for HTTP use cases that require static deterministic, fast regional failover

### 135. [SAA-C02] AWS Global Acelerator - Hands on
- Launched an EC2 instance in Ireland
- Launched an EC2 instance in Londo
- Go on Global accelerator
    - On first page: name it MyFirstAccelerator or something
    - Second page: Port 80, Proto TCP, No client affinity
    - Third page: Add eu-west-1 and eu-west-2 as endpoint groups
    - Add your two EC2 instances as endpoints
    - Create accelerator! It will be created in Oregon
- Neat it works!
    - You can see that your accelerator has 2 endpoitn groups
    - With endpoint ids and passing health checks!
    - If you make the health checks fail, the endpoint will be removed from the global accelerator
- Service seems a bit pricey with the data transfer fees

### Quiz!
- Pending