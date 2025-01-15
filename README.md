# Terraform Provider NTOPNG

## Auth ##

This provider allows env var auth as well as provider{} auth inline

```
export NTOPNG_USERNAME="your-username"
export NTOPNG_TOKEN="your-token"
export NTOPNG_HOST=https://test.foo.com
```

Inline:

```
provider "ntopng" {
  host     = "http://1.2.3.4:3000"
  username = "your-username"
  token    = "your-token"
}


```

## Build provider

Run the following command to build the provider

```shell
$ go build -o terraform-provider-ntopng
```

## Test sample configuration

First, bump the version so that its unique and won't pull from the registry:

```shell
$ vim Makefile
```

Next, build and install the provider.

```shell
$ make install
```

Then, navigate to the `examples` directory.

```shell
$ cd _examples
```

To run this locally you'll need to add a `~/.terraformrc` file with:

```
provider_installation {
  filesystem_mirror {
    path    = "/Users/marcus.young/.terraform.d/plugins"
    include = [
      "github.com/myoung34/ntopng",
      "registry.terraform.io/myoung34/ntopng",
    ]

  }
  direct {
    exclude = ["myoung34/ntopng"]
  }
}
```

Run the following command to initialize the workspace and apply the sample configuration.

```shell
$ rm -rf .terraform .terraform.lock.hcl
$ terraform init
$ export NTOPNG_API_KEY="your-api-key"
$ export NTOPNG_API_SECRET="your-api-secret"
$ terraform plan
```
