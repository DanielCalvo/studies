### The goal
- Create a release pipeline for dcalvo.dev

### How I think this works
- Create a packer image with your dcalvo.dev index.html and the mdscanner stuff
    - Hey does mdscanner have a cronjob for the scans?
- Update an image setting on an autoscaling group behind and ALB or something using an Elastic IP in TF?
- Update DNS or Elastic IP attachment? Or nothing else if using autoscaling?
- Release complete?

### mdscanner idea/update
- Host the results in S3
- Have the workers run in... EC2? Lambda mayhaps?