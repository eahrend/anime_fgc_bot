docker run --name mongodb -d -p 27017:27017 -v ./mongo:/data/db mongo
env MONGODB_URI="mongodb://0.0.0.0:27017" go run main.go