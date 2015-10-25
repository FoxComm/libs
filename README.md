# FoxCommerce Shared Libraries

A set of helper libraries that make using FoxCommerce services easy.

## Set it up!

- Clone it

    ```
    git clone https://github.com/FoxComm/libs
    ```

- Configure the repository

    ```
    make configure
    ```

- Build the services

    ```
    make build
    ```

- Install the services

    ```
    make install
    ```

## Included Libraries

#### Announcer (deprecated)

`announcer` is a library that every service may implement in order to tell the 
FoxCommerce Router that the service is online and running. As the process for
service discovery has changed, new services should not use this library.

#### Configs

`configs` is a set of helpers designed to allow services to find configuration
values from a variety of configuration storage mechanisms.

#### DB

`db` is a set of helpers for interacting with PostgreSQL or Mongo instances.

#### Endpoints

`endpoints` is a mapping to other FoxCommerce services. This is only guaranteed
to be valid when running inside of core FoxCommerce infrastructure.

#### etcd_client (deprecated)

`etcd_client` is a wrapper for interacting with etcd.

#### Logger

`logger` is a wrapper around Logrus that puts log messages in the format that
we like to use.

#### Spree

`spree` is a helper that allows services to all Spree API endpoints.

#### Utils

`utils` is a miscellaneous set of individual utilities.

## License

The contents of this repository are the property to FoxCommerce, Inc. All rights reserved.
