# kafkawireformat
Go files that implement kafkas wire format.

This is an implementation of kafkas byte wire format (kafka >= 0.11) as it is described in the [kafka protocol guide](https://kafka.apache.org/protocol).
This is explicitly not an implementation of a kafka server or client. This project can be used to do exactly that though. It merely contains go files that help you render out the bytes to communicate with a kafka server or client.