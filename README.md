# Cognises
 
 Gather information about cloud servers for various things. 

## Usage

### Generate config

`generate-config` writes the 'default' config.

example:

```
cognises generate-config > ~/.cognises.yaml
```

### SSH Config

`ssh-config` writes generated SSH config based on configured providers.

example:

```
cognises ssh-config > ~/.ssh/generated_config
```

It is recommended to set the generated output to a secondary SSH config file. For example, use an include in `.ssh/config`:

```
Include ~/.ssh/gererated_config
```

## Cloud Providers

Relies on cloud provider credentials to exist in a normal location, eg. if `aws s3 ls` fails, this will likely fail too. 

Currently only supports the following:

- AWS

### AWS


Require IAM policies:

```
ec2:DescribeRegions
ec2:DescribeInstances
```