# dc-to-k8s

## debug

this project uses delve and vscode, because that's just where i am rn. to debug:

```sh
$ yarn
...
$ yarn dev-dlv
[nodemon] 2.0.15
[nodemon] to restart at any time, enter `rs`
[nodemon] watching path(s): *.*
[nodemon] watching extensions: go
[nodemon] starting `yarn go-dlv ./`
API server listening at: [::]:12345
2022-03-27T20:00:59-07:00 warning layer=rpc Listening for remote connections (connections are not authenticated nor encrypted)
```

now u go the the run tab and click on `Debug go` then the play icon.

### what

so my understanding is that we build an output via gdb with proper source mapping. i am starting to think that the chrome debugger protocol for node is some fukken standard i should know...
