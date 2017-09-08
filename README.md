# Terraform Provider

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)
- irc.freenode.net: #terraform

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

## Building The Provider

Clone repository to: `$GOPATH/src/github.com/elijahcaine/terraform-provider-provider-etcd`

```sh
$ git clone https://github.com/elijahcain/terraform-provider-provider-etcd $GOPATH/src/github.com/elijahcaine/terraform-provider-provider-etcd
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/elijahcaine/terraform-provider-provider-etcd
$ make build
```

## Using the provider

FILL THIS IN

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*).
You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`.
This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make bin
...
$ $GOPATH/bin/terraform-provider-provider-etcd
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

## Future work

- Become feature complete.
- Merge into [Terraform Providers](https://github.com/terraform-providers).
- Check `TODO.txt` for a status on feature implementations.

## Scope

This project was created out of necessity, so the features implemented are only the ones required by the author.

If you would like to add a feature please do so!
Please make an issue where other contributors can discuss with you how best to implement the feature.
When you are ready to make the contribution yourself, make a Pull Request.

If you do not feel technically capable to make a code contribution yourself don't be afraid!
Feel free to make the issue and it will be attended to at the community's best effort.

### Things to keep in mind when contributing

- Document changes.
- Write tests for your changes.
- Can't write code? Docs are always useful!
