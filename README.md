# LogWatcher

LogWatcher is a powerful log monitoring tool written in Go. It enables users to efficiently watch, filter, and analyze logs in real-time. With its robust features, LogWatcher helps developers and system administrators gain insights into their applications and services, ensuring smooth operations and quick troubleshooting.

## Features
- **Real-time Log Monitoring**: Instantly see new log entries as they come in.
- **Filtering and Searching**: Easily filter logs by keywords, severity, or date.
- **Custom Alerts**: Set alerts for specific log patterns or thresholds.

## How to use

### Using go run 
You can use by typing your terminal `go run main.go` but it won't work if you are not in the project directory.

### Using go build
```
git clone https://github.com/BedirMirac/logWatcher.git
cd logWatcher
go build .
```

After these steps you can start the program by typing `logWatcher` in your terminal if you are in the project directory.

### Using go install (RECOMMENDED)
```
git clone https://github.com/BedirMirac/logWatcher.git
cd logWatcher
go install
```
It looks like the go build method but the difference is now you can start the program  by typing `logWatcher` in your terminal without being in the project directory.

## Contributing
We welcome contributions! Please read our contributing guidelines before submitting a pull request.

## License
LogWatcher is open-source software licensed under the MIT License.
