Gepeto says: prometheus has also developed substantially better OpenTelemetry ingestion support

It seems Prometheus 3.0 came out shortly after the second edition of the book. so the second edition is not that current either --  reading the documentation on what is new on version 3.0  is likely your best interest if you wanna dive deeper on that

Or alternatively you can just focus on the fundamentals plus the Kubernetes operator

Not really Prometheus operator specific but remember: CRDs are cluster-wide. 

I would be interested to see how prometheus ingests and handles traces!

## Questions to answer
- Very basic one: How do you ingest and dashboard cluster metrics?
- What app do you want to instrument to ingest the metrics?
- How is push gateway represented by the prometheus operator?
- Can you have a list of CRDs and what they do? How do they compare to a plain text prometheus set up?

## Other notes