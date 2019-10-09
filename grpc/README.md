# gRPC

`Nuxt.js`のフロントで、バックエンドが`gRPC`の場合のデモアプリ。

```
$ go run main.go
```

```
$ protoc -I proto/ proto/post.proto --go_out=plugins=grpc:proto
```

```
$ npm install
$ npm run dev
```