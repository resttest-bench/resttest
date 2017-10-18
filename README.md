# resttest
REST test (sort of)

Decided to go a bit off-book. Since you mentioned GRPC, I figured I'd do this with GRPC rather than REST.
No unit tests, but instead you can run the client against the server which returns a bunch of mocked data, both generated from the protobuf definition in `transactions.proto`.
I went with a cursor rather than pages, assuming this is real-time data (though the server doesn't actually do anything meaningful with the cursor.)

To run:
If you don't have Go set up, there are some instructions here: https://golang.org/doc/install

`make dependencies` to brew install protobuf and go get a few libs
`make proto` to build the client and server libs
`make run-server` to run the server
`make run-client` to run the client