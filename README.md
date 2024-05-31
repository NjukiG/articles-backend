# Articles API

This is a minimal backend for a blog-like application.
It has ability to add users, authentication using jwt, add articles and add comments to articles.
It also ensures diferent users have different permissions.

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
