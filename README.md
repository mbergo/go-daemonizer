# Go Daemonizer

```
Run the daemon

Run the daemon with the following command:

go run daemon.go sleep 10

The command will run for 10 seconds and then exit.

Check the process

Check the process with the following command:

ps -ef | grep sleep

You should see the following output:

501  8717  8716   0  8:50PM ttys005    0:00.00 sleep 10
501  8727  8441   0  8:50PM ttys006    0:00.00 grep sleep

The process is running as a daemon.

Kill the process

Kill the process with the following command:

kill -9 8717

Check the process

Check the process with the following command:

ps -ef | grep sleep

You should see the following output:

501  8727  8441   0  8:50PM ttys
```