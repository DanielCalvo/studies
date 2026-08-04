## Preliminary SLO
- 99.9% of valid requests return the correct resized image within 5 seconds over a 30 day rolling window

## SLIs
- Percentage of valid requests returned correctly
- Percentages of valid requests returned below 5s

## Observations
- An SLO just puts a target on one or more SLIs (in our case the target is a percentage)
- Only valid requests should count towards the 99.9% of correct operations.  so if a user uploads a png,  that shouldn-t count. We need to ensure that metrics track that properly
- there is a risk that there might be a bug in which the application incorrectly rejects a valid j peg. a synthetic test with known valid inputs would help detecting that.

## Other topics
- Which metrics are important again?
- RED metrics?