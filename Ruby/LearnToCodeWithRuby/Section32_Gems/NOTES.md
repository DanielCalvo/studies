# ruby on rails is a gem, neat!
- https://rubygems.org/
- https://rubygems.org/gems/rails
- https://rubygems.org/gems/csv
- https://rubygems.org/gems/sidekiq <- Ha, I heard about this one!
- https://rubygems.org/gems/puma
- https://rubygems.org/gems/faker

# on the versioning:
- 3.2.1
- 3 is the major version
- 2 is the minor version
- 1 is the patch version

Major version increments usually have breaking changes. Assume it to be incompatible with the previous versions.
Minor version increments usually have new features. Assume it to be compatible with the previous versions.
Patch version increments usually have bug fixes or small improvements. Assume it to be compatible with the previous versions.

Do note that some gems can depend on other gems

```shell
gem update --system #will auto update the gem system
gem install faker #installs the latest version of faker
gem uninstall faker #but thats not what we want for now!
```

Lets create a gemfile! (see the gemfile)
And now run: `bundle install`

This installs the bundle and also generates Germfile.lock.