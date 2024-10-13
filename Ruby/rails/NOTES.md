Create a test app:
- https://guides.rubyonrails.org/getting_started.html#creating-the-blog-application

rails new helloworld

To start it:
bin/rails server

Add to config/routes.rb
```ruby
  get "/articles", to: "articles#index"
```

The route above declares that GET /articles requests are mapped to the index action of ArticlesController.

To create ArticlesController and its index action:
`bin/rails generate controller Articles index --skip-routes`

Lets generate a model:
bin/rails generate model Article title:string body:text


## Other notes
- oh wow there's a routes command, I wonder what other commands are there!
- so many helpers! (like link_to, I wonder if there's a list of all of them somewhere?)