Sample of using protobufs in Go
===============================

The protobuf is declared in ``person/person.proto``. To re-build the ``.pb.go``
file, run from this directory:

.. sourcecode:: text

	PATH=$PATH:$GOPATH/bin protoc --go_out=person -Iperson person/person.proto

This assumes ``protoc`` is installed, along with its Go plugin. It creates the
file ``person/person.pb.go``, which defines the Go package ``person``.
