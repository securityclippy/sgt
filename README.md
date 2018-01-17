## Welcome to Sgt!
![](docs/images/11036-200.png)

SGT is an osquery management server written in Golang and built in aws.  Sgt (Simple Go TLS)
is backed entirely by AWS services, making its infrastructure requirements extremely
simple, robust and scalable.

SGT is managed entirely through terraform

### Getting started.

Getting started with sgt is designed to be very simple with minimal setup required.  To get started, however, you will need a FEW things first.

1. An aws account with admin access to DynamoDB, EC2, ES (ElastisearchService), Kinesis/Firehose and IAM. (note, this must be
programatic access, so you can have an access key and secret to use)
2. [Golang 1.8.2+](https://golang.org/doc/install)
3. [Terraform](https://www.terraform.io/intro/getting-started/install.html)
4. A domain with dns [managed via Route53](http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/MigratingDNS.html)
5. An SSL cert with public and private keypair.  This will be used to terminate TLS connections to our server
see [Obtaining a free ssl cert for SGT with Letsencrypt](docs/letsencrypt_cert_instructions.md) for one method of aquiring a certificate


Once you have these things installed and ready to go, it's time to get started with the real fun.

Create a new aws profile in your ~/.aws/credentials file.  This can be done either manually or by [configuring the AWS CLI](http://docs.aws.amazon.com/cli/latest/userguide/cli-config-files.html)

The name of this profile can be whatever you want it to be, but we recommend choosing something that matches with
the environment you're working with.  For example, if you're going to be spinning up SGT in a testing environment, naming your profile ['testing'] is
quite appropriate.

clone the this repo
```commandline
git clone git@git.repo
```

##### Building SGT.
cd into the sgt directory

```commandline
cd sgt
```

get dependencies and build the binary

```commandline
go get
go build
```

### configuring certificates

In order to terminate our TLS connection, we need the TLS certificate and associated private key.

SGT expects both the certificate and private key in PEM format, so make sure both files are in the proper format.
If you're using letencrypt or certbot, this is the default format.  (For instructions on getting a letsencrypt cert,
see [this link](docs/letsencrypt_cert_instructions.md)

Once you have both the cert and the private key files, place them in the `sgt/certs` directory and name them appropriately.
While you can configure them to be named whatever you like, we highly recommend naming them in this format:
```
subdomain.domain.fullchain.pem
```
and
```
subdomain.domain.privkey.pem
```

This naming scheme will allow you to easily identify which certs belong to which environment if you ever end up with more than
one (While the truly brave test in prod, usually we recommend at least a dev/prod setup :) )

#### Deployment Wizard

Once you've installed Go and Terraform, and built your SGT binary, its time to run your deployment!

The wizard will walk you through everything you need to configure a new environment,
create the proper directory structure and the environment specific configuration
files and stand up the environment if you choose to do so

```commandline
./sgt -wizard
```


