- The each and select ways of looping through an array seem really interesting, you should practice with them a bit!

## Stopped on lecture 36 on the area code exercise, seems like a fun one to do!

The author encourages:
- Reading this, for style: https://github.com/rubocop/ruby-style-guide
- Also for practive, https://try.ruby-lang.org/ is encouraged!


## To research
- google getters and setters and see what the internet recommends.
- hey try out rubymine

# Section 3 - rails
Guides, neat: https://guides.rubyonrails.org/
- The getting started (https://edgeguides.rubyonrails.org/getting_started.html) is very useful to 
- There are also other useful things on the website, like APIs! https://rubyonrails.org/
- Puma is the webserver distributed with rails by default

When creating a rails application, which steps do you have to perform to add a route again?

## 60. Structure of a rails application
Structure of a rails application, folders:
App / assets is where we store static assets like images and style sheets
App /Channels contains ruby code related to real time functionality, like chat related
Controllers: Where all your rails controllers will be housed. It would seem all controllers inherit from an application controller class.
App /Helpers: Helper functions or methods to be used in views templates
App /Models: Where you will store all your models. Looks like all models extend what is on the application record file. Similar to how all the application controllers would create will inherit from the application controller file
Views: Has your views. Of particular interest seems to be application.html.erb. According to the author all views get served through this file.

bin: n/a

config: Has the application config , neat. Off interest right now is routes.rb, which you use to define the HTTP paths and Associated routing for your app.

db: In development, actually contains your database. Which is SQLite. I believe the database migration files is also go here.

files:
Gemfile: Contains your gem config. You would modify this to include new Gems or to upgrade them
Gemfile.lock: Created by rails, locks versions of gems. You don't work directly with this file, it is auto generated

package.json: Lists packages and dependencies added to the application using yarn

README.md: Just a readme!

## 66. Front-end: Learn and practice HTML and CSS
Author encourages going through: 
- https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/HTML_basics
- https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/CSS_basics

Alternatively, you can go through this, which is more thorough
- https://learn.shayhowe.com/html-css/building-your-first-web-page/