#btracing

-You can select a sampling point in the sampler for open telemetry in your instrument mentation
-  the trace provider receives a sampler as your configuration parameter
-  once a trace is created a decision is made on whether to sample the trace. This decision is stored in these pan context
-  once span has ended the span processor applies the sampling decision

# metrics
- Sampling in the case of metrics mas really alter the data rend and breath rendering it useless
-  for instance imagine if you're gonna record data for each request like how you're gonna sample a counter you're gonna alter de find value
-  so the book isn't conclusive here but it tell you that simply mattress isn't a good idea however
-  the book says that reducing the amount of data produced by metrics focuses on aggregating the data 

#logs
- The book is a bit out of date here as it said that there was no way to sample logs
-  checking with codex there is a way to sample logs
- You can use this percent jeff whatever you want (ex: 10)

# sampling strategies
You can make the decision on what to sample in various different points

##head ssampling
- The quickest way to the side what so sample is at the very beginning of the trace journey this is known as head son sampling
- The application that creates the first span of the trace that is the root span decides whether to sample the trace or not and propagates that decision by the context to every subsequent service called
-  head sampling reduces the overhead for the entire system
- However this doesn-t always work as it's possible for applications to configure sampling differently and one another

##tail sampling
- Tale sampling is a common strategy that plates until a trace is complete before making a sampling decision
- this allows the sampler to perform some analysis on the trace to detect potentially anomalous or interesting occurrences
-  with tail sampling all the applications must produce and transmit the telemetry to a destination that the sides to simple the data or not like
- Depending on where the tail sampling is performed this option could cause significant amounts of data to be transmitted over the network and then dropped, so you're sending that around to do nothing with it
- Additionally the sampler must buffer in memory or store the that effort entirety of the trace until it is ready to decide if he wants to keep it
-  this inevitably leads to an increase in memory and storage consumed -- To mitigate for memory concerns you can configure a maximum trace duration but this can also lead to gaps for traces that never finished or things like that


## probability sampling
- So this ensures that that I selected randomly
-  this appearculated by applying configurable ratio to the hash of the trace id and since the id is propagated across holes system all components configured with the same ratio and they trace id ratio based sampler Apply the same logic at decision time

# samplers available
- So you can have always owned and always off blank
- you can then choose to sample by trace id, or parent based
- but again this is from 20>21 so this information might not be the most accurate one

# Rest of the chapter
-The book then goes into a python code example for sampling at the application level with the SDK
-Then it shows you how to use the open telemetric elector to sample dataUsing the tail sampling processor
  - It-s cool to see that details sampling processor can make decisions for you for what that sorry what traces to sample like
  -  you can sample an overall trace duration
  -  or certain span attribute values
  -  or the status code of a span
- Then it shows you how to capture only 10% of all traces but always capture traces that take longer than one second, neat
- Perhaps as an exercise for later it would be cool like at the collector level to configure it to always send failed traces but otherwise if a trace is successful maybe only send 20% of successful traces