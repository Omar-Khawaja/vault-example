Simple example that uses Vault's Go client library to interact with Vault and
read secrets. The secret key is "hello" and the value is "world" as show below
using the CLI.

```shell
$ vault kv get secret/foo
====== Metadata ======
Key              Value
---              -----
created_time     2018-10-29T01:16:21.500198665Z
deletion_time    n/a
destroyed        false
version          1

==== Data ====
Key      Value
---      -----
hello    world
```

```shell
$ vault kv get -field=hello secret/foo
world
```
