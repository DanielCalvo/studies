 I want to change a tag a k:v pair on an object, while keeping other tags. For example, assuming we have a bucket with these tags:
 
```yaml
Name: Mybucket
Service: MyEmailService
Owner: Daniel
```

I want to have:
```yaml
Name: Mybucket
Service: MyEmailService
Person: SomeoneElse
```
