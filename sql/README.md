# PostgreSQL

PostgreSQL is a production-ready, open-source database. It's a great choice database for many web applications.

## How does PostgreSQL work?

Postgres, like most other database technologies, is itself a server. It listens for requests on a port (Postgres' default is :5432), and responds to those requests. To interact with Postgres, first you will install the server and start it. Then, you can connect to it using a client like psql or PGAdmin.

## Installation

### Linux

Here are the docs from Microsoft. The basic steps are:

```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```

## Ensure the installation worked
The psql command-line utility is the default client for Postgres. Use it to make sure you're on version 14+ of Postgres:

```bash
psql --version
```

## Start the Postgres server in the background

Linux
```bash
sudo service postgresql start
```

## Connect to the server using a client

Linux
On your command line run:

bash
```
sudo passwd postgres
```

Enter a new password and remember it. Then restart your shell session.

Host: localhost
Port: 5432
Username: postgres
Password: the password you created

## Create a database
A single Postgres server can host multiple databases. In the dropdown menu on the left, open the Localhost tab, then right click on "databases" 
and select "create database".

Name it whatever you like, but you'll need to know the name.

## Query the database
Right click on your database's name in the menu on the left, then select "query tool". You should see a new window open with a text editor. In the text editor, type the following query:

```sql
SELECT version();
```
