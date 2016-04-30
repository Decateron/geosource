# geosource [![Build Status](https://travis-ci.org/joshheinrichs/geosource.svg?branch=master)](https://travis-ci.org/joshheinrichs/geosource) [![codecov.io](https://codecov.io/github/joshheinrichs/geosource/coverage.svg?branch=master)](https://codecov.io/github/joshheinrichs/geosource?branch=master) [![GoDoc](https://godoc.org/github.com/joshheinrichs/geosource/server?status.svg)](https://godoc.org/github.com/joshheinrichs/geosource/server)

GeoSource is web application for posting and browsing spatio-temporal data.

### Setup

Requires:
 * [Go 1.5+](https://golang.org/)
 * [PostgreSQL 9.4+](http://www.postgresql.org/)
 * [PostGIS 2.1+](http://postgis.net/)
 * [ImageMagick](http://www.imagemagick.org/script/index.php)
 * [Bower](http://bower.io/)
 * A TLS key and cert for HTTPS
  * I recommend using [OpenSSL](https://www.openssl.org/) for local development, and [Let's Encrypt](https://letsencrypt.org/) for deployment

[A more detailed setup procedure can be found here.](docs/setup.md)
