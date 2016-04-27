# Setup

### Install Git

```
sudo apt-get install git
```

### Install Go

[Reference](https://golang.org/doc/install)

```
sudo apt-get install golang
```
You'll also need to set GOPATH inside your `~/.bashrc` file. I chose to set mine to `$HOME/go`

### Install ImageMagick

[Reference](https://github.com/gographics/imagick)

1. `sudo apt-get install imagemagick`
2. `sudo apt-get install libmagickwand-dev`
3. `pkg-config --cflags --libs MagickWand`

### Retrieve the Repository

Get the repo and install all of the server's dependencies: 

```
go get github.com/joshheinrichs/geosource/server
```

And then add a symbolic link to the folder for convenience:
```
ln -s $GOPATH/src/github.com/joshheinrichs/geosource ~/geosource
```

### Install PostgreSQL

[Reference - PostgreSQL](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-postgresql-on-ubuntu-14-04)

[Reference - PostGIS](http://postgis.net/install/)

From inside the `geosource/database` folder, install PostgreSQL and PostGIS, create a database, and then populate it with tables via the `dbinit.sql` script:

1. `sudo apt-get install postgresql`
2. `sudo apt-get install postgis`
3. `sudo -i -u postgres`
4. Run `createuser --interactive` and create a user called "geosource"
5. `createdb geosource`
6. `sudo -i -u geosource`
4. `psql -d geosource -U geosource`
5. `create extension postgis;`
6. `\i dbinit.sql`

### Install Website

[Reference](http://bower.io/)

From inside the `geosource/app` folder, install bower and then download the website's dependencies:

1. `sudo apt-get install npm`
2. `ln -s /usr/bin/nodejs /usr/bin/node`
3. `sudo npm install -g bower`
4. `bower install`

### Set up HTTPS keys

[Reference](https://letsencrypt.org/getting-started/)

If you are doing local development, use [OpenSSL](https://www.openssl.org/). Otherwise, use LetsEncrypt to set up the encryption keys.

1. `sudo apt-get install letsnecrypt`
2. `letsencrypt certonly --standalone -d geosource.usask.ca`

The `fullchain.pem` and `privkey.pem` files should be located inside `/etc/letsencrypt/live/geosource.usask.ca/`

### Starting the Server

Inside the `geosource/server` folder:

1. set up a `config.gcfg` file
2. Run the server via `sudo -E go run main.go`
