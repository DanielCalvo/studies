## Installing rails from source
```shell
mkdir -p /home/daniel/bin/rubybin #if not there already
cd /tmp
wget https://cache.ruby-lang.org/pub/ruby/3.3/ruby-3.3.5.tar.gz
tar -xvzf ruby-3.3.5.tar.gz 
cd ruby-3.3.5
./configure --prefix=/home/daniel/bin/rubybin
make -j8
make install
echo 'export PATH=$PATH:/home/daniel/bin/rubybin/bin' >> ~/.bashrc
```

And then:
```
gem install rails
```
