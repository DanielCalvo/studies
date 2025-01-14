require 'aws-sdk-s3'

s3 = Aws::S3::Client.new
resp = s3.list_buckets
puts resp.buckets.map(&:name)
puts s3.inspect
puts s3.class

resp = s3.list_objects_v2({
  bucket: "dcalvo-dev-bucket",
  max_keys: 2
})

puts resp.to_h
puts

#https://docs.aws.amazon.com/sdk-for-ruby/v3/api/Aws/S3/Client.html