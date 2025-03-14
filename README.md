## A simple TCP server which makes the responses with jokes by interval.

The jokes sample data has been taken from this [repo](https://github.com/kylecs/jokes-dataset-json).

### How to run it:
By default the server will be running on: ``localhost:8080``.

With "Makefile":
- Run the command ``make run``. If want to specify the port, run the following command ``make run ARGS="--port=8000"

Go build:
- Run the command ``go build ./cmd/jokes`` and the run the binary.

Open a new terminal window, and run the command using a ``ns``: ``ns localhost 8080``. After 2 secods, you going receive the jokes within interval.
