# DB Migration & Seeding Flow

GO sample that showcases the DB workflow with migrations and seeders from traditional frameworks in other programming languages. It uses the <https://github.com/golang-migrate/migrate> for the migrations side and a custom development for the seeding part using <https://github.com/brianvoe/gofakeit> and based in the <https://ieftimov.com/post/simple-golang-database-seeding-abstraction-gorm/> article. It also uses the official MongoDB driver for GO.

You should install MongoDB and start the service with `systemctl start mongod` in order to work with this sample. Alternatively, you can have Docker installed and execute the following command `docker run -d -p 27017:27017 --name m1 mongo` in order to start a MongoDB container to play with.

## More Resources

* <https://play.golang.org/p/rDodxp22Jq>
