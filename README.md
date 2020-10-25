# Routes application
## Calculate distance and driving time between source coordinates and each of given destination, ordered by driving time and distance (if driving time is equal)

**How to run on your machine** <br/>
Application is made as a Go module with the vendor directory to make it easier to run by eliminating the necessity of dependency download. 
- Clone repository
- Open termianal in root of the project
- Type "go run main.go"
- Make sure your 8080 port is free, because application is listening on this port

**Application is also available online on the following address, deployed with Google App Engine** <br/>
https://routes-293609.ey.r.appspot.com/api/routes?src=34.123456,53.123456&dst=34.654321,53.654321 <br/><br/>

**How to use** <br/>
Example request should looks like this:<br/>
Local machine: <br/>
localhost:8080/api/routes?src=13.388860,52.517037&dst=13.397634,53.529407 <br/>
Online: <br/>
https://routes-293609.ey.r.appspot.com/api/routes?src=34.123456,53.123456&dst=34.654321,53.654321 <br/>

Proper request must contain **exactly one src** query parameter and **at least one dst**, so in order to get distances between multiple destinations from source, you can call: <br/>
https://routes-293609.ey.r.appspot.com/api/routes?src=34.123456,53.123456&dst=34.654321,53.654321&dst=34.222222,53.22222&&dst=34.333333,53.333333 <br/><br/>

**Example response**<br/>
Routes from src to each dst are sorted by duration and distance (if duration is equal)
```json
{
  "Source": "34.123456,53.123456",
  "Routes": [
    {
      "Destination": "34.220000,53.222220",
      "Duration": 967.2,
      "Distance": 19654.9
    },
    {
      "Destination": "34.330000,53.333333",
      "Duration": 2111.3,
      "Distance": 32564
    },
    {
      "Destination": "34.654321,53.654321",
      "Duration": 5599.2,
      "Distance": 99773
    }
  ]
}
```

**External service usage**<br/>
Application is using osrm service to get distance and duration between given points using example endpoint: <br/>
http://router.project-osrm.org/route/v1/driving/13.388860,52.517037;13.397634,52.529407;13.428555,52.523219?overview=false
