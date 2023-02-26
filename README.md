# go-logstash
Logstash Logger written in Go that pushes logs directly to [Logstash](https://www.elastic.co/products/logstash). It supports writing outputs to Logstash in JSON format as well as String format.

### **Features**

- Supports TCP and UDP protocols for sending logs to Logstash.
- Supports writing logs in both JSON and string formats.
- Provides customisable options for configuring the Logstash connection.

### **Getting Started**

#### 1. Install the package using:

```
go get github.com/KaranJagtiani/go-logstash
```

#### 2. Import the package in your Go application:

```go
import "github.com/KaranJagtiani/go-logstash"
```
#### 3. Initialize the **`logstash_logger`** instance with your Logstash server details and connection type:

```go
logger := logstash_logger.Init("<host>", <port>, "protocol", <timeout>)
```

The `logstash_logger.Init()` has the following configuration options:

- Host: Logstash server hostname.
- Port: Logstash server port number.
- Protocol: Logstash connection protocol, either "tcp" or "udp".
- Timeout: Connection timeout in seconds.

Example:
```go
logger := logstash_logger.Init("logstash", 5228, "udp", 5)
```

#### 4. Use one of the different methods provided for sending logs to Logstash:

```go
payload := map[string]interface{}{
		"message": "test message",
		"error":   false,
}

// Generic - For logging the payload as it is
logger.Log(payload)

// Adds a attribute called "severity": "INFO" to the payload
logger.Info(payload)

// Adds a attribute called "severity": "DEBUG" to the payload
logger.Debug(payload)

// Adds a attribute called "severity": "WARN" to the payload
logger.Warn(payload)

// Adds a attribute called "severity": "ERROR" to the payload
logger.Error(payload)

// For sending a string message to Logstash
logger.LogString("String Message")
```

### **Contribution**
Contributions to go-logstash are welcome! If you find a bug or want to add a new feature, please create an issue or submit a pull request here on GitHub.

#### Submit a Pull Request
1. Fork this repository.
2. Create your feature branch (git checkout -b my-new-feature).
3. Commit your changes (git commit -am 'Add some feature').
4. Push to the branch (git push origin my-new-feature).
5. Create a new Pull Request.

