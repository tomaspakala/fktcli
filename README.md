# fktcli

fktcli generates random data mainly by using [gofakeit](https://github.com/brianvoe/gofakeit) project

## Usage

### Misc

- Get guid:

```
fktcli guid [flags]
> 7abbe496-4d8f-4d06-ae3b-3ee84d36acd1
```

- Get date time now:

```
fktcli dateTime [flags]
> 2023-06-04T15:35:17.6049409Z
```

### Person

- Get name:

```
fktcli person name [flags]
> Leopold Effertz
```

- Get email:

```
fktcli person email [flags]
> fredrickchamplin@kirlin.name
```

### Auth

- Get user name:

```
fktcli auth userName [flags]
> Bergstrom3953
```

- Get password:

```
fktcli auth password [flags]
> Kv#hP8ifs42glZ
```

More details can be found in [docs](./docs/fktcli.md)
