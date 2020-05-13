### 138. Introduction to Messaging
- When you start deploy applications, they will need to communicate with one another.
- There are two patterns for application communication:
    - Synchronous communication
    - Asynchronous communication

- Synchronous:
    - Buying service > Shipping service

- On asynchronous communication, there's a queue or middleware that connects both services
    - Buying service > Queue > Shipping service

- Synchronous traffic between applications can be problematic if there are sudden spikes of traffic
- What if you suddenly need to enconde 1000 videos instead of the usual 10?

- You can decouple your applications with a decoupling layer
    - Using SQS: Queue model
    - Using SNS: Pub/sub model
    - Kinesis: real time streaming model (big data related)
- These services can scale independently from your application!

### 139. AWS SQS 
- A Queue service that receives messages from producers that can be polled by consumers

#### Standard Queue
- SQS is old! About 10 years old
- Fully Managed
- Scales from 1 message per second to 10.000 to seconds
- Default message retention: 4 days, maximum of 14 days
- No limit to how many messages can be in the queue
- Low latency service (<10ms)
- You can have scaling terms of consumers and producers (or just consumers?)
- Can have duplicate messages (rare, with high throughput)
- Can have out of order messages (best effort ordering)
- Limit of 256kb per message sent

#### Delay Queue
- 


### 140. 
### 141. 
### 142. 
### 143. 
### 144. 
### 145. 
### 146. 
### 147. 
### 148.
### 149.
### 150. 
 
 
