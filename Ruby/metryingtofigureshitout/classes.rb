#what does self do again?
# what does super do again?
# so arguments don't need to be in parenthesis right?
# whats with ::?
# what does < do in classes again?
# what about structs? https://ruby-doc.org/3.3.5/Struct.html
# whats with attr?
# what does underscore as in _variable do?

require 'github_api'

github = Github.new

#Here you can probe
github.repos.list org: 'Kubernetes' do |repo|
  puts "Repo list returns something with a class of #{repo.class}"
  puts "Ancestors: #{repo.class.ancestors}" #Ah, so its a Hashie::Hash, mostly!
  #what was the thing? ancestors?
  puts repo.clone_url
  # repo.keys.each do |key|
  #   puts key
  #   end
  break
end