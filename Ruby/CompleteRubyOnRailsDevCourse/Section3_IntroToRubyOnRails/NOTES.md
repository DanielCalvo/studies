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
