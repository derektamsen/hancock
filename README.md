# Golang S3 URL Signer for Amazon AWS
This service provides authorized access to S3 assets via url signing. It is designed in an optimized way so that locality of S3 buckets or cloudfront endpoints can still be leveraged.

The service receives http requests, signs the url using your AWS secret key, and finally redirects the client to S3 or cloudfront with the signed url. This service is to be used in conjunction with an authorization such as PKI client cert authentication via nginx.

This service runs as a http server to receive requests, sign urls, and redirect users. It is recommended that this be placed behind a nginx proxy or an Amazon ELB.

## Design

```
user -> nginx (auth + proxy) -> aws-s3-url-signer (url signing + redirect) -> user (redirected) -> S3/cloudfront
```

1. `user`
    1. User make a request to `<your domain>/some/s3/asset`.
2. `nginx`
    1. Authenticates the user with client PKI certificates or another form of authentication.
    2. Proxies the traffic to this service.
3. `aws-s3-url-signer`
    1. Takes the requested url from the user and signs it using the aws sdk
    2. Redirects user to the S3 or cloudfront endpoint with a signed url
4. `user`
    1. User issues a direct request to S3 or cloudfront for the requested resource.

## Building
```
make
```

You can also use make to clean your built artifact with `make clean`
