require 'aws-sdk-s3' # Requires the aws-sdk-s3 gem

# Set up S3 client with default configuration
s3 = Aws::S3::Client.new

begin
  # List all S3 buckets
  response = s3.list_buckets

  puts "Your S3 Buckets:"
  response.buckets.each do |bucket|
    puts "- #{bucket.name}"
  end
rescue Aws::S3::Errors::ServiceError => e
  puts "Error listing buckets: #{e.message}"
end
