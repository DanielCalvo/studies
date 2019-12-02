Excellent blog post: https://blog.golang.org/pipelines

### Highlights

- Multiple functions can read from the same channel until that channel is closed; this is called fan-out. This provides a way to distribute work amongst a group of workers to parallelize CPU use and I/O. 
- A function can read from multiple inputs and proceed until all are closed by multiplexing the input channels onto a single channel that's closed when all the inputs are closed. This is called fan-in.

##### There is a pattern to our pipeline functions
- Stages close their outbound channels when all the send operations are done.
- stages keep receiving values from inbound channels until those channels are closed.

This pattern allows each receiving stage to be written as a range loop and ensures that all goroutines exit once all values have been successfully sent downstream.


  
