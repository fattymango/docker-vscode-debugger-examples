# VS Code Remote Debugger With Docker
This is a simple example of how to use the attach VS Code debugger to a running program with hot reload in a Docker container using different programming languages.

## Setup
1. Run the docker compose file to start the containers
    ```bash
    docker compose up -d
    ```
    or use Makefile
    ```bash
    make up
    ```

2. Open any example folder in VS Code and start the debugger
3. Make a change in the code and see the hot reload in action
4. Set a breakpoint and see the debugger in action


## Documentation
You can find the documentation for each programming language and how it connects to the debugger in the following links:
- [Node.js](./node/README.md)
- [Go](./go/README.md)