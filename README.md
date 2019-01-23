# Inventory availabilities rest API with GO

Simple app to demo building rest api using GO. The api serves inventory atp data

## Getting Started

Clone this repository into local GO workspace. Build the project and execute 'inventory_availabilities'. Building the project needs a
few GO libraries to be installed locally. Step by step instructions below will cover these details.
provided below.

### Prerequisites

1. GO installed https://golang.org/dl/
2. Postgresql server

```
Postgresql on mac can be installed using brew: $ brew install postgresql
For more details: https://wiki.postgresql.org/wiki/Homebrew
```

### Installing
After installing Postgresql locally or on a remote server, a dbuser, table and a few sample records need to be created. Sample commands
are below:

In mac following command can be used to create a Postgresql database
```
$ createdb atpdb
```
Rest of the commands can be executed from psql, which is Postgresql interactive terminal.
These commands will login to atpdb database, create a user called atpdbusr, reset its password, create a table called atp and finally insert
some sample records in the table.
```
$ psql atpdb
atpdb=# CREATE ROLE atpdbusr CREATEDB;
atpdb=# ALTER USER atpdbusr WITH PASSWORD 'new_password';
atpdb=# CREATE TABLE atp (item_id varchar(255), uom varchar(255), onhand int, demand int, atpqty int);
atpdb=# INSERT INTO atp (item_id, uom, onhand, demand, atpqty) values ('327689', 'EACH', 80, 20, 60);
atpdb=# INSERT INTO atp (item_id, uom, onhand, demand, atpqty) values ('988667', 'EACH', 80, 20, 60);
```
Update the application.properties file, with db credentials from above step.

Next set of commands will download and install necessary GO libraries for building this project.

```
$ go get -u "github.com/gorilla/mux"
$ go get -u "github.com/magiconair/properties"
$ go get -u "github.com/lib/pq"
```
Next, the project can be built and rest api brought up with the following:

From the root of the project folder -
```
$ go build
$ ./inventory_availabilities
```
The above will bring up the rest api on port 8000. Below is a sample URL and response. You can hit this URL from a browser, or from an
app like Postman, Insomnia etc

```
Request: http://localhost:8000/availabilities/327689

Response:
{
  "id": "327689",
  "uom": "EACH",
  "onhand": 80,
  "demand": 20,
  "atpqty": 60
}
```

## Acknowledgments

* Nice article about using mux to route http requests  https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo
* Connecting to a postgresql db with GO: https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
