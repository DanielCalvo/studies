## Metanotes
I went over section 1 of the book with voice to text and an AI pass to fix typos. A quick pass on sections 2 and 3 is still pending. That and maybe organizing thes notes a bit better...
---

Chapter 1 the reliability stack

I like that. On page two it says the question "is my service available?" is analogous to "is my service doing what its users need it to do?"

This implies that availability needs to be measured from a user perspective instead of creating a model to track reliability and then chasing the model instead of the user

If users don't think you're being reliable, then you're not

On page 3, it says a good SLI measures your service from the perspective of your users. Interesting!

An SLI is most useful if it can result in a binary answer, good or bad. So for instance, you might have an SLI for the page to load under 2 seconds. If the page does load under 2 seconds then the users are happy. If it takes longer than 2 seconds then we can say that's bad

So you take the good visits and divide them by all visitor visits and then you have your SLI percentage

Your SLO then is a target for what that percentage should be.

An error budget is a way of measuring how your SLI has performed against your SLO over a period of time. It defines how unreliable your service is permitted to be. If you go over your error budget, then you need to take corrective action. Like for example, dedicate a single engineer for reliability improvements.

An SLA is a service level agreement that you have with the customer usually over a contract. SO FOR INSTANCE SLO IS USUALLY SOMETHING YOU HAVE INTERNALLY WITHIN YOUR ORG FOR AN OBJECTIVE THAT YOU'RE AIMING AT, while an SLA is something that you have agreed upon with a customer with a contract and if you break that contract, there might be some penalty or you might have to pay something back or you might not get paid or depends on the contract.

The book then says SLIs are the most important part of the reliability stack and might be the most important part of the book! You might never get to the point of having a reasonable SLO target or a calculated error budget, but taking a step back and thinking about the service from your user's perspective can be very useful.
On totality

So the book argues that of the whole word soup, the SLI is the most important one, probably because figuring out what SLI you create makes you think about what you're going to measure from the user perspective.

An SLI is at its most basic, a metric that tells you how your service is operating from the perspective of your users.

So even though an API availability or an error rate might be a valid starting point, a better SLI would be like. How long does it take for a user to authenticate against your service and retrieve some data or maybe to log in or maybe to create a new VM from a VM dashboard. Depends on the service!

A good SLI can be expressed in a sentence that all stakeholders can understand.

SLOs are targets for how often you can fail or otherwise not operate properly and still ensure that your users aren't upset. So for instance, if a visitor tries to visit a website and it loads very slowly, they might refresh and it might work and it's okay. But if it fails every time they visit then they'll probably abandon that website then go do stuff elsewhere. Remember that an SLO is an objective. It's not a contractual agreement!

Error budgets are the most sophisticated part of the stack. Maybe because you need SLI and SLO first. The author says that calculating an error budget can be more complicated than you expect.

You can have event-based and time-based error budgets. With event-based, you think of good or bad events in the service or product. With time-based, you measure bad minutes. It says like on these 30 days we had 43 bad minutes and so on and so forth. So which one you pick seems to depend on your preference

Author then says surplus error budget, ship more features

Exceeding error budgets. Do engineering reliability work

What is a service? Well a service is anything that has users. These can be internal, external or even other services or machines

You can also have an SLI and an SLO for a data processing pipeline. You can also measure batch jobs


The book then gives some example of services like a website, a video streaming service, or a web-based email provider and so on


Then the book goes on. The book says that commonly when talking about SLOs, we'll hear about request and response from APIs mostly for two reasons because they're incredibly common and because they're incredibly easy to measure in comparison with some other services. So establishing a meaningful SLI for a request and response API is much easier than a data processing pipeline or some other service. Measuring API responses is easy!


Be careful not to have SLOs only because it is a trend and making it mean whatever you wanted to mean.


SLOs are just data and hopefully it is data that can help you have better discussions and make better decisions. It's something that makes you think about your service from a new perspective and not strict ideology. An SLO is supposed to guide you but not demand anything from you


The book encourages iterating. So for instance, first you pick some SLIs and then once you observe them, you figure out you can pick an SLO target and then once you do those, you realize that maybe you can pick an error budget. But then you realize that your SLI wasn't as good as you thought because you need to represent the users' view better. So you change the SLI but then you need to change the SLO and then you need to communicate this. And this is all fine. It's a journey and not the destination so iterating is fine


The author then says it's all about humans because SLOs should lead to happier engineers, product people, business and users. That should be the goal, not appending extra nines at the end of your SLO target

chapter 2 How to think about reliability?

The book says past performance matters. Your users probably expect the availability of your service to remain the same. So if it was reliable in the past, they expect it to be reliable in the future

Users already implicitly expect things from you

So for instance, if you have a video streaming service, some buffering at the beginning is okay, but some buffering in the middle of the film might be disruptive and users might not like that

The book then comments that 100% is unnecessary and impossible. So for instance some cars don't start the first time but as long as they start either on the second or third time almost all the time it's fine. You're not going to buy a new car because it struggles to start occasionally

Another example of this is a watch that falls behind a minute or two per month, but as long as the watch is unreliable in more or less the same way all the time, the owner might not be too bothered by this. They'll simply add a minute or two to whatever the watch says and that result is good enough and then they'll just have to adjust the watch once a month or so


Or maybe you ordered a pizza from this place you like? And every tenth order or so takes an hour to arrive which is longer than usual. But since you like the pizza, you're okay so you know the pizza arrives quickly 90% of the time and that's good enough

Also, bear in mind that the more reliable you want your system to be the more expensive it gets. So for instance, you might need instances of your service in multiple locations or you might need more machines and then more rigorous testing

Also, if you want four nines of reliability over a month means your service can only be down for 4 and 1/2 minutes during 30 days, which means that anyone on call needs to have a response time in the order of seconds which is very difficult

This means you might need to have engineers on call at all times. Very glued to a laptop

So that means that the closer you want to get to 100%, the more you'll be chasing diminishing returns with increasing cost

When thinking of reliability, what really helps is to take a step back and put yourself in the shoes of your users. What are the things they need from you? How often do they need those things? And how performant do you need to be in the first place and then you can probably figure out the level of reliability that you can have?

This way you're thinking of your users first and they'll be happier and so will your engineering and operations teams because you won't be asking them to do the impossible

Chapter 3 developing meaningful service level indicators

SLIs are the most important part of the whole entire process

However, human happiness is the ultimate goal of using SLO based approaches

SLIs are the foundation of the stack. You can't have meaningful SLOs or error budgets if your SLIs aren't meaningful.

Remember that an SLI is defined as a metric that tells you how your service is operating from the perspective of your users


Furthermore, your service is not reliable if your users think it's not reliable. User opinion is what matters the most

The author encourages shifting the focus from what your service needs to do to what the users expect the service to do. By doing so, you can build better telemetry about the user needs


The author then says that having user focused SLIs is often more complicated than just relying on traditional metrics which means that you might need to set up non-trivial amounts of code like new services only to help determine what this telemetry would look like

The author then argues that it doesn't matter if you have errors in your logs or higher database latency or container crashes as long as your service is performing its job as determined by the comprehensive SLIs then these all can wait until normal work hours. No one needs to be paged


Then there's a section about a request and response service and the author imagines that you have a simple request and response service

So the first thing you need to care about is if the service is actually up so you probably should have a check that asks. Is the service actually up?

Then he asks is the service available and by available he means is it probably available to the end user through the network?

Then the third thing the author asks is now that the service is up and available. Is it responding correctly? So is it responding in a timely fashion? Or is it taking say 10 seconds to respond to a simple request?

Now it could still be up, available and responding but it could be returning errors. So one thing to add would be is your service returning an acceptable number of good responses?

Now the author then argues that maybe your service is responding correctly, but it should maybe be responding with some JSON with some fields or what have you and it might be responding with empty JSON to a valid request. So are the responses in the correct format?

But then none of that matters. If you're returning, let's say yesterday's data. So you also need to ask is the correct data being returned?

So an SLI could be the 95th percentile of requests to our service will be responded to with the correct data within 400 milliseconds

However, most services are more complicated than just an API request. How do you measure that?

The author does encourage measuring user journeys as SLIs

The author also says that you can introduce a service tracing solution to measure incoming user request time instead of generating artificial ones (I wonder if they're referring to OpenTelemetry here? Could I use that? This would be an interesting laboratory thing to try!)

The book then briefly talks about user alignment and SLI and for instance an SLI for adding something to a shopping cart is actually a user journey from a product management standpoint and that's actually correct. The conversation that you can have with the product person is for instance, if the SLI for adding a thing to a shopping cart starts to degrade and that affects an SLO, then perhaps the conversation you can have is are we okay with having a user taking 5 seconds to put something on a shopping cart? Isn't that too much and reflects in a poor user experience? Then that SLI would help you have a conversation with that product manager to then shift the focus to work and reliability from his point of view. So it's like a single page that we can all be on which is quite nice!

The author then emphasizes that service level indicators are the most important part of an SLO based approach. You can have SLIs even if you don't have SLOs.

Chapter 4 choosing good service level objectives

SLOs are targets and good SLOs usually have two things in common. If you're exceeding your SLO target, your users are happy. If you are missing your SLO target, your users are unhappy.

The book then goes into the problem of being too reliable which means then you don't have the freedom of doing what you want. Like perform some chaos engineering, ship features quicker or just introduce downtime to see how the dependencies react. Additionally, there's the risk of operational underload because people learn to fix things by doing so in complex systems, you can learn a lot from failures. If things never fail, you'll miss out on all of that

The book then mentions that the difference between 99.9 and 99.99 is actually much greater than you might realize at first, and you should be looking at the numbers in between as well. So if your service has to be unavailable 2 hours per month for platform maintenance or something then 99.7 would be a correct starting point

The book then advises caution with the fact that you can have too many SLOs. What the book says is you want to capture the most important features of your system and you can often accomplish this by measuring only a subset of those features

If you have too many SLOs, it might be difficult to make decisions using your data, so having too much data can make decisions hard to make

So for instance, if you have a storage service with a cache layer, you don't need an SLO for cache hit and cache miss, you can just have an SLO for general read latency

The other problem is that it becomes more complicated to report on the availability of your service. If you have to sort through dozens of SLOs with different targets, target percentages and histories and what not, it's difficult to communicate the availability of whatever you're trying to measure

Then there's also the multiple comparison problem which is that if you have many different measurements of the same system, there are greater chances of incorrect measurements taking place

And if you're looking at too many things, you'll always find something that looks just slightly off, which can waste your time by sending it down an endless needless rabbit hole

Your service can have soft dependencies and hard dependencies and a hard dependency is something that has to be reliable for your service to be reliable. Like a database that it has to read from in order to function. A soft dependency is something that your service needs but it can still reliably do without and the book then says that one of the best things you can do is convert your hard dependencies into soft ones to make the service more reliable

So it's obvious to say that your service cannot be more reliable than the hard dependencies it has. If it depends on the database, it can only be as reliable as the database. To figure out how reliable that database is, you can measure how many requests to this database complete without a failure. Or you can check the database SLO if there is one

Soft dependencies can be a bit trickier to define. So for instance, if you have some mapping software on your phone and the traffic features stop working. This impacts the reliability of the map application, but it doesn't make it wholly, unreliable or unusable

One of the best things you can do for your app is to turn your hard dependencies into soft ones. So for instance, going back to the database example, you could introduce a caching layer if much of the data is similar or doesn't have to be updated to the second or in the case of your mapping software. If the traffic feature doesn't work, then don't crash the entire app and same thing if you have some third party that you count on only for a certain part of your service, which is cool. It's a nice idea

The book then says you can measure hard drives by getting data from the internet like hard disk failures from Backblaze

When choosing a target, the book reminds you that you shouldn't try to make your target too high and remember that SLOs are not SLAs and you should keep in mind that you can change your SLO if the situation warrants it. There's no shame in picking a target and then changing it if you learn that it is wrong

However, the best way to figure out how your service might operate in the future is studying how it has operated in the past

The book then says that as SLOs are informed by SLIs, oftentimes when picking your SLI you might need to pick new metrics and you might need to collect this data for at least a month to have an idea of what your SLO could be like

The book then goes over ranges, max, min, mean, average and median. The median value is kind of cool because I already had read about it but shamefully it had never clicked into my head. The median value is the value where half are below that value and half are above that value which is different from the average!

Then there are percentiles in which you can commonly formulate SLOs around like oh, I want the 90th percentile of my requests to be below this time because the other 10% might be outliers that take very long, for whatever reason

Then there's a problem with low resolution metrics which you don't have if you log metrics for every single request. But if you only have one data point per minute, then you need to measure how many bad minutes you have. And if your SLO is like 99.9, then if the metric is bad like a few minutes in a day with an outlier or two, then you've already reached your threshold. So you might need to count like two or three bad metrics in a row to assume that they're bad

Another problem relates to quantity. So even if you collect data about your service at one second intervals, maybe your service only has activity far less frequently than that. Maybe like a batch scheduler job or a data pipeline with the lengthy processing periods or even a request and response that only gets called a few times a day. You might not have enough data points

For instance, if you have a pipeline that runs every hour and one of them fails that's already 95% availability, which could be fine if the pipeline is allowed to fail

For other things like an API that has low traffic during the night and or it's only used during office hours, then maybe you can say anything between 11:00 p.m. and 6:00 a.m. doesn't count and you don't count it

Also, maybe your data is noisy or not good quality so perhaps you need to measure things in a way that you need to be below a certain threshold for 5 minutes for you to consider a bad event

So you can use percentiles to calculate an SLO with like three rules and take into account the long tail and making sure that it stays normal and the data is distributed as intended

So for instance, you could have the p95 of all requests complete within 2 seconds

The p98 within 2 and 1/2 seconds

And a p99 within 4 seconds

All of that at 99.9% of the time! Neat!

If you only look at the 95th percentile, you will miss changes in the long tail and won't be able to address problems in the longest 5% of your responses! Imagine your slowest 5% get twice as slow for some reason! That's something you probably want to know about!

If you're setting up an SLO for a service that doesn't have a history yet well then try to take an educated guess!

Remember that SLOs are objectives, not a formal agreement. Maybe you can pick data from other places like similar services or data from staging requests

Chapter 5 how to use error budgets?

Error budgets are the final part of the reliability stack and it takes a lot of effort and resources to use them properly and not every team or company always gets to that part

Which is interesting cuz I very rarely see practical examples or people talking about practical application of doing something with the error budget itself. Like I read theory out there but it's rare to see concrete examples of this I feel

Remember that reliability isn't just uptime. It's doing what your users need you to do

The author then says that the pinnacle of having SLOs that work and are observed is alerting only on error budget burn rate which is interesting

Error budgets are budgets, the author says, and rarely should you actually stop releasing things because you can freeze releases when the error budget is burned, but then you might have a difficult-to-resolve issue because all changes queued up during the freeze are released at once. But maybe don't release any new big features when your error budget is low and maybe instead do smaller features that are less risky or something like that

The author then argues that an error budget can help with the eternal debate between development and operation teams where the former wants to ship many features and the latter wants to ship fewer features to keep things stable while the error budget helps somewhat objectively show which direction do we go towards?

Also, measuring error budgets over time can give you a good insight on what risk factors impact your service. So by knowing what kind of events and failures are bad enough to burn your error budget, you can discover what factors caused you the most problems over time. This is a very interesting insight

You can use your error budget maybe to try increasing your cache retention from 1 hour to 2 hours because you don't know what's going to happen or you can do load testing or maybe you can do a risky Kubernetes upgrade I don't know

Or for instance you could burn some error budget to rely entirely on some automated tests that you've been thinking of putting into practice and not using manual reviews for a while

You can also do load tests against production. They're a great time to find out if you're calculating your burn rate correctly and if your SLI measurements are correct!

You can also do a black hole exercise if you have multiple regions or zones serving traffic from a service. In a black hole exercise, you just turn off one region like simulating a region failure to see if your service actually remains available

There's a famous and true story about a service named Chubby on Google that users were warned that the global version should not be depended upon, but eventually a bunch of new services ended up having Chubby as a dependency and what they did at Google was when they had error budgets left over. They would just turn off Chubby for a couple of minutes. And once the fires were put out, this actually led to positive outcomes like teams learning not to rely on Chubby too much and maybe who knows, converted from a hard to a soft dependency

Then the book goes into a bunch of details and calculating error budgets, choosing time windows like services that don't have to be available in the middle of the night and so on and so forth. I think this is all very useful reference, but perhaps it's best kept as reference as error budgets are very rarely reached in a capacity in which you have some surplus of it and that you can propose risky stuff yourself instead of it coming from various other sources. But still I think this is useful to keep in mind

Then the book goes and discusses error budget policies.  These could document what you do if you burn a certain part of your error budget or all of it.
For instance, you could have one of your engineers working on reliability after you've burnt a percentage of your budget.

Error budgets might also have recommendations about communication because if you're out of error budget and services that depend on yours and they have their error budgets severely impacted the teams that are responsible for their services will want to know about this or maybe other departments. Yeah, I can see how this would happen so perhaps it is useful to document what to do

I think the book emphasizes that you probably should be flexible. For instance, what if your error budget policy says you must halt future work, but your CEO says you must ship a certain feature. The important thing to remember is that SLOs and error budgets give you data so that you can perhaps push back or have a conversation about those requests so they provide a starting point from which you can disagree. But I think the important part here is the conversation

So in summary, error budgets are the most difficult part of the reliability stack to get to, but first you need to think about SLIs. Then you need to think about your SLO second and then you need to calculate how you've performed against your target third and then using your error budget is actually the fourth and final step. They give you ways to make decisions about your service. They also give you indicators that you can use to decide when to ship features or experiment or what your biggest risks are which is nice
