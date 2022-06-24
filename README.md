# micro-service-example

Go kit (https://github.com/go-kit/kit) を使用したマイクロサービスのサンプルコード。  
異なる2つのサーバアプリケーションがgRPCを使いリクエストを送り合うようなシナリオを想定しています。  
このレポは確認用なので、一方がサーバ、もう一方がクライアントとして機能します。

## How to start

```bash
# 証明書生成
make cert

# サーバの起動
make start-server

# クライアントの起動 (サーバへリクエストの送信)
make start-client
```

## コード解説

`pkg1`以下にサーバ用コード, `pkg2`以下にクライアント用コード, `shared`に共通のコードを配置しています。  
サーバは`Server`, `Service`, `Repository`の3層構造になっており、既存のREST API構造と似たような構成を取っています。

## gRPC通信
サーバ・クライアント間通信にはgRPCを使用しています。  
protobufスキーマは`pkg1/pb`以下の`.proto`ファイルで定義しています。  
同一ディレクトリの`.pb.go`は自動生成ファイルで、`make protoc`で再生成できます。 

## 参考

- https://grpc.io/docs/guides/auth/
- https://dev.to/techschoolguru/how-to-secure-grpc-connection-with-ssl-tls-in-go-4ph
- https://itnext.io/practical-guide-to-securing-grpc-connections-with-go-and-tls-part-1-f63058e9d6d1
