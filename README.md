# golang-concurrency
Golang concurrency practise dump

Adding golang concurreny patterns and observations which will help understand concurrency in golang better.


Random observations.
**Read on a closed channel does not block or throw an error.**

**Who creates the channel , closes the channel. This is because the sender never knows when the channel is closed , but the receiver does. We can check with an OK flag returned when reading from the channel. It is true when channel is active , false when closed.**

**Buffered channel only blocks for sender when the buffer is full.**
