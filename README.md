# fsjl

**F**ormat **S**tructured **J**SON **L**ogs

Orthogonal, opinionated JSON log transformer.

For https://docs.kofile.systems/dev-docs/microservices/good-citizenship-for-services/

### Usage

[Pipe in JSON log output](https://getpino.io/#/) to `fsjl`:

```sh
node run-server.js | fsjl
```

... will transform the server logs from stdout ...

```sh
{"pid":94473,"hostname":"MacBook-Pro-3.home","level":30,"msg":"hello world","time":1459529098958,"v":1}
{"pid":94473,"hostname":"MacBook-Pro-3.home","level":50,"msg":"this is at error level","time":1459529098959,"v":1}
{"pid":94473,"hostname":"MacBook-Pro-3.home","level":30,"msg":"the answer is 42","time":1459529098960,"v":1}
{"pid":94473,"hostname":"MacBook-Pro-3.home","level":30,"msg":"hello world","time":1459529098960,"obj":42,"v":1}
{"pid":94473,"hostname":"MacBook-Pro-3.home","level":30,"msg":"hello world","time":1459529098960,"obj":42,"b":2,"v":1}
{"pid":94473,"hostname":"MacBook-Pro-3.home","level":30,"msg":"another","time":1459529098960,"obj":{"aa":"bbb"},"v":1}
{"pid":94473,"hostname":"MacBook-Pro-3.home","level":50,"msg":"an error","time":1459529098961,"type":"Error","stack":"Error: an error\n    at Object.<anonymous> (/Users/davidclements/z/nearForm/pino/example.js:14:12)\n    at Module._compile (module.js:435:26)\n    at Object.Module._extensions..js (module.js:442:10)\n    at Module.load (module.js:356:32)\n    at Function.Module._load (module.js:311:12)\n    at Function.Module.runMain (module.js:467:10)\n    at startup (node.js:136:18)\n    at node.js:963:3","v":1}
{"pid":94473,"hostname":"MacBook-Pro-3.home","level":30,"msg":"hello child!","time":1459529098962,"a":"property","v":1}
{"pid":94473,"hostname":"MacBook-Pro-3.home","level":30,"msg":"hello baby..","time":1459529098962,"another":"property","a":"property","v":1}
{"pid":94473,"hostname":"MacBook-Pro-3.home","level":30,"msg":"after setImmediate","time":1459529098963,"v":1}
```

... into ...

<img width="1382" alt="screen shot 2018-05-07 at 4 11 09 pm" src="https://user-images.githubusercontent.com/29997/39729563-6c45fd08-5211-11e8-9d40-8b9166365faf.png">

### Options

| flag               | default | description                                                                                                      |
| ------------------ | ------- | ---------------------------------------------------------------------------------------------------------------- |
| `-no-color`        | false   | terminal color support should be detected automatically; if not, specify this flag to turn off ANSI color output |
| `-fall-through`    | false   | lines that miss the timestamp and level detection will get swallowed silently unless this is true                |
| `-ignore-all-meta` | false   | if true, does not print any metadata, eg. `{key}={value}` pairs                                                  |

##### Example of `-no-color`

```sh
$ node run-server.js | fsjl -no-color
```

<img width="1387" alt="screen shot 2018-05-07 at 4 13 27 pm" src="https://user-images.githubusercontent.com/29997/39729586-9b250308-5211-11e8-8511-52c92adfd64a.png">

##### Example of `-fall-through`

```sh
$ node run-server.js | fsjl -fall-through
```

<img width="1380" alt="screen shot 2018-05-07 at 4 15 30 pm" src="https://user-images.githubusercontent.com/29997/39729654-e61fd3b0-5211-11e8-80a1-1cf9a5bf7b9c.png">

##### Example of `-ignore-all-meta`

```sh
$ node run-server.js | fsjl -ignore-all-meta
```

<img width="1053" alt="screen shot 2018-05-07 at 4 16 05 pm" src="https://user-images.githubusercontent.com/29997/39729668-f7c477f6-5211-11e8-9164-164248d48219.png">

## Build

`make build`

## Run tests

TODO: write tests
