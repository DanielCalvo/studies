require 'github_api'

github = Github.new auto_pagination: true

github.repos.list org: 'Kubernetes' do |repo|
  # puts "Repo list returns something with a class of #{repo.class}"
  # puts "Ancestors: #{repo.class.ancestors}" #Ah, so its a Hashie::Hash, mostly!
  #what was the thing? ancestors?
  puts repo.clone_url
  # repo.keys.each do |key|
  #   puts key
  #   end
  #break
end


#p 5.class.superclass.superclass.superclass.superclass #nil -- basicobject has no superclass!
#p 5.class.ancestors #returns all the superclasses and some modules apparently
#p 3.14.class.ancestors