### 47. Beanstalk intro 

```text

```

### 48. Beanstalk overview

```text
Most web apps have the same architecture (LB + ASG)
This is were elasticbeanstalk comes in, a developer centric view to deploy apps in AWS.
Uses: EC2, ASG, ELB, RDS and all the most popular things.
Easy to set up, with a lot of control on how to fine tune things

elasticbeanstalk is a managed service.
Instance config / os handled by beanstalk
Deployment strategy is configuyrable but performed by beanstalk

Just the application code is the responsability of the developer

elasticbeanstalk has 3 components:
Application
Application version
Environment name.

You deploy applications to environments and can promote application versions around
There's a rollback feature! 

Elasticbeanstalk supports A LOT of platforms!
```

### 49. Beanstalk Hands on

```text
Launched the sample node.js app
You can modify the whole thing (instance size, load balancer, subnets and what not)

```

### 50. Beanstalk Deployment Modes

```text
Single instance, great for dev. 
High availability with load balancer, great for prod.

Deployment update modes:
All at once (fastest, but instances aren't available to serve traffic)
Rolling: Updates a few instances at a time
Rolling with additional batches: Like rolling, but spins up new instances to move the batch (so old app remains available)
Immutable: Spins up new instances in a new ASG, deploys to these instances and then swaps all the instances when everything is ready
```

### 51. Beanstalk Deployment Modes Hands-on

```text
```

### 52. Beanstalk for SysOps

```text
Beanstalk questions are very basic in the sysop exam compared to the developer exam

```

### 53. Beanstalk clean up
```text

```

### Quiz!
