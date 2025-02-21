## 55 MVC app structure 
- Model: Resources (objects?) used in your app: user, post, article, stock, etc. Most likely requires a db for persistence.
- Views: Front end of the application. Embedded ruby, that can template html apparently
- Controller: Contains the code that handles your model? not that well explained here

## 56 General flow of an mvc app, sorta
- user requests something at the browser
- request received at the router of the rails app
- request routed to the appropriate action in a controller
- controller#action either renders a view or communicates with the model
- model communicates with the db
- model sends back info to controller
- controller renders view

# 58 and onwards, creating your first rails app
- add a route to routes.rb pointing to a controller
- write your controller under app/controller
- have an appropriately named action
- and then have... something on the pages folder and something on the home action
- `rails generate controller pages`
- router > controller > view (in this case we don't have model, so no db connection)

## 59 root route, controller and move mvc
1. config/routes.rb -> root 'pages#home'
2. `rails generate controller pages`
3. app/controllers/page_controller.rb -> def home
4. /app/views/pages/home.html.erb (put hello world in there!)

## 60 structure of a rails application
- Goes over all of the folders and what they do. I think I'll consult this one on a "need to remember" basis

## 73 The backend: CRUD scaffold and wrap up for section 3
- rails generate scaffold article title:string description: text
- rails db:migrate
- rails routes
- rails routes --expanded

 