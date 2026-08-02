The chapter goes over how chaos engineering is having un hypothesis and then introducing a bit of chaos to test your hypothesis

 but this is nice because I was unfamiliar with chaos engineering so the whole hypothesis thing makes a lot of sense

 like you don-t just break things we no right or reason you're like oh I bet introducing a one hundred millisecond latency to our systems will not break anything but then maybe you discover you have misconfigured time out somewhere or something like that

The book then describes using a tool called the traffic control that can simulate packet loss, increase latency, and  throw put limits

The book then describes using the stress utility to create memory pressure on a system-
- This is to test the hypothesis that the grocery store processes will serve fewer requests as they cannot obtain the resources to process requests, in this case memory.
- it also had Poth sizes that latent civil increase
 - and that metrics should allow you to quickly identify this problem

The other experiments de book performs is unexpected shutdowns. this goes to show that services should and are expected to manage shutdowns gracefully. to

So this chapter is like the author showing how open telemetry can help you diagnose these issues when applications run into trouble or or resource starved or there are network issues and so on and so forth