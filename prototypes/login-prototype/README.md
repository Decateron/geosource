# login-prototype

Demoing Google+ login. Implementation is based off of the [gplus-quickstart-go](https://github.com/googleplus/gplus-quickstart-go) project, adapted to Polymer. 

This prototype also supports username creation, and stores the data into a simple Postgresql database which records a user's email and username.

### Setup

Follow Step 1 of the [gplus-quickstart-go tutorial](https://developers.google.com/+/web/samples/go) to set up the Google+ API.

Then, set up a Postgresql database. It should be accessible via the command line by running a command such as:

```bash
psql -h host -d database -U user
```

The database should be password protected for security reasons. Once that is in place, you should initialize the database with the `dbinit.sql` file by running the following command in the database:

```bash
\copy dbinit.sql
```

You'll also need to create a `config.gcfg` file, based off of the given `example.gcfg` file. This is necessary so as not to expose critical security information such as the database's password on GitHub.

After that, it should be as simple as running:

```bash
go run *.go
```

within the `server/` folder. The site will then be accessible via `http://localhost:8000`.

To later destroy the database, use the `dbdrop.sql` file, and run the following command in the database:

```bash
\copy dbdrop.sql
```
