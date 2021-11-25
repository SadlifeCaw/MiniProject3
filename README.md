# MiniProject3

## Notes
After an auction has finished, you must restart the servers if you want to hold another auction

The auction works if you only start 1 server, but there will obviously be no fault tolerance
## Run servers

Open 2 terminals and enter the Server folder:

```
cd .\Server\
```

Start each server by typing in the terminal:
```
go run .
```

**The system is designed for both servers to be up and running before any client connects!**

## Run client

Open as many terminals as you want and do 

```
cd .\Client\
```

Start each client by typing in the terminal:
```
go run .
```