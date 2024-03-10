# Geogame

Geogame it is a web application game that uses geographic coordinates to provide hints about a country and the player must guess which country it is.

The link it is available here: 
[Geogame](https://geogame-frontend.yellowstone-f1fc5476.eastus.azurecontainerapps.io/)

# How works:

- User enter
- A random country is provided
- The user can use hints (or not) to figure out which country it is
- If the user found the correct answer, must click at "send location"
- A label will appear displaying the result, and other informations like score, score bonus and distance.
  
It uses Go-Haversine package (created by myself) to calculate the distance between the guess coordinates and the country coordinates. Then, the difference distance is returned as Kilometers. Finally, the score will be calculated and returned for the user.

You can access the go-haversine package here: [go-haversine](https://github.com/andrefsilveira1/go-haversine)

# How I made it:

I used Golang, Fiber (Go framework), React, Ant-design, Terraform, Docker, Azure SQL Database, Azure Server and Azure Container Apps. In this repository, I hide infrastructure details for security. 

# How execute:

If you wanna execute in your local environment, you can do it following these steps:

- Install Go
- Install yarn
- Move to frontend folder and run: `yarn start`
- Move to backend folder and run: `go run main.go`

Notice that the database connection is based at Azure SQL Database. If you do not has any Azure connection Available, you can use any different database service and provide the connection.

Contribuitions: 

I made this alone, so, probably has some errors here. If you find any issues, please, do not hesitate to reach me out :)
