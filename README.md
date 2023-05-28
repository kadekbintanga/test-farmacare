# How to running the project

## Project Information
- This project is created with Go and Gin Framework
- This project use PostgreSQL for the database

## How to run the project
### Requirement :
1. Make sure you have installed Go programming language in your device.
2. Make sure you have installed PostgreSQL database in your device.

### Step - step to run the project :
1. Create a database with PostgreSQL in your device
2. Clone this project to your device
3. Open project with your code editor.
4. Change the name of ".env.example" file in this project with ".env"
5. Open .env and adjust the database configuration with your database configuration that have you created before. Example:
```bash
DB_USER=postgres
DB_PASS=your password
DB_NAME=your database name
DB_HOST=localhost
DB_PORT=5432
```
6. Open the terminal and run this command :
```bash
go mod download
```
7. After that you can run the project with this command :
```bash
go run main.go
```
8. The project will run in port 8000
9. Now you can hit API service provided in this project

## API service in this project
There are 6 API service in this project :
### 1. Get Pokemon
This service for get list data of pokemons that is provided from POKEAPI service. To run this, you can hit the project API to :
```bash
http://127.0.0.1:8000/api/v1/pokemon?page=10&limit=20
```
In this service, the request method is GET and you have to set the page and limit parameters for get the data. If you do not set the parameters. The service will process the data with default page and limit. Explanation of parameters.
- page = the page of list data. The default is 1.
- limit = the limit of list data. The default is 10

### 2. Post Battle Manual
This service for post data battle of the pokemon with the position. To run this, you can hit the project API to :
```bash
http://127.0.0.1:8000/api/v1/battle/manual
```
In this service, the request method is POST and you have to send data JSON in this API. Example :
```bash
{
    "battle_name":"Battle Poke Manual 1",
    "position":["pikachu", "fearow", "ekans", "ditto","nidoqueen"]
}
```
Explanation :
- battle_name = the name of battle. type data is string
- position = the names of pokemon in the battle. You have to sort the positions from 1 to 5. In this example "pikachu" is position 1, "fearow" is position 2, "ekans" is position 3, etc. You have to insert the correct name of pokemon. if you insert wrong name, the service will response with error bad request status. Type data is array.

### 3. Post Battle Auto
This service for post data battle of the pokemon like service number 2. Post Battle Manual, but in this service the position will be generated automatically in this service. To run this, you can hit the project API to :
```bash
http://127.0.0.1:8000/api/v1/battle/auto
```
In this service, the request method is POST and you have to send data JSON in this API. Example :
```bash
{
    "battle_name":"Battle Poke Campion",
    "pokemons":["spearow", "fearow", "ekans", "pikachu","nidoqueen"]
}
```
Explanation :
- battle_name = the name of battle. type data is string
- pokemons = the names of pokemon in the battle. You do not need to sort the positions of pokemon. In this service the position will generated automatically. You have to insert the correct name of pokemon. if you insert wrong name, the service will response with error bad request status. Type data is array.

### 4. Get Battle List
This service for get history of battle that was generated before. To run this, you can hit the project API to :
```bash
http://127.0.0.1:8000/api/v1/battle?start_date=2023-05-28 14:46:00&end_date=2023-05-28 14:48:25&page=1&limit=10
```
In this service, the request method is GET and you have to set the start_date, end_date, page and limit parameters for get the data. If you do not set the parameter. The service will process the data with default value of parameter. Explanation of parameters :
- start_date = start of range battle date. it is recommended to use "yyyy-mm-dd hh:mm:ss" format with UTC timezone.
- end_date = end of range battle date. it is recommended to use "yyyy-mm-dd hh:mm:ss" format with UTC timezone
- page = the page of list data. The default is 1.
- limit = the limit of list data. The default is 10

### 5. Get Total Score
This service for get score of all pokemon that have got battle before. To run this, you can hit the project API to :
```bash
http://127.0.0.1:8000/api/v1/battlepokemon/score
```
In this service, the request method is GET.

### 6. Update Annulled Position
This service for annul position of pokemon in a battle before. To run this, you can hit the project API to :
```bash
http://127.0.0.1:8000/api/v1/battlepokemon/annulled?uuid_pokemon=d3926d65-b06b-4aff-ac24-2bc6349c10db
```
In this service, the request method is PUT and you have to set the uuid parameter. If you do not set the parameter or the uuid is not registered in database, the status response service is bad request. Explanation of parameters:
- uuid = is the uuid of pokemon that you want to annul.

## Postman Example
For support this project. I attach postman collection with name "TEST-FARMACARE.postman_collection.json".


### THANK YOU
created by : Kadek Bintang Anjasmara @2023