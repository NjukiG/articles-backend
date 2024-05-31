# Articles API

This is a minimal backend for a blog-like application.
It has ability to add users, authentication using jwt, add articles and add comments to articles.
It also ensures diferent users have different permissions.
e.g Only users are allowed to add articles, patch and delete only their articles;
users allowed to add comments to other users' articles and only commenters can patch and delete their own comments unless entire rticle is deleted by original user.

## Routes

User Routes:
POST "/public/signup"
POST "/public/login"
GET "/protected/validate"

Article Routes:
GET "/protected/articles"
GET "/protected/articles"
GET "/protected/articles/:id"
POST "/protected/articles"
PUT "/protected/articles/:id"
DELETE "/protected/articles/:id"

Comments Routes:
POST "/articles"
GET "/articles/:id/comments"
GET "/comments/:id"
POST "/articles/:id/comments"
PUT "/comments/:id"
DELETE "/comments/:id"

You can clone the app, cd into it and run: go run main.go and it will run locally on your machine. I t will also download all the modules used to build.
