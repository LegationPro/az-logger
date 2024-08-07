<h1>Az Logger</h1>

A simple, color-coded, file-based logging package for Go applications.


<h2>Overview</h2>

This package provides a logger with various log levels (`Debug`, `Info`, `Warn`, `Error`, `Panic`, `Fatal`). It logs messages to files named by their log levels and supports color-coded output to the terminal for easier reading.

<h2>Installation</h2>
To use this logger in your project, import the package:

```go
import "github.com/LegationPro/go-logger/logger"
```

<h1>Usage</h1>

<h2>Initialization</h2>

Before logging any messages, you need to initialize the logger. This can be done by calling NewLogger() with the path where the log files will be stored.

```go
logger.NewLogger("logs")
```

<h2>Logging Messages</h2>
You can log messages at different levels using the following methods:

  - `Debug(message string)`

  - `Info(message string)`

  - `Warn(message string)`

  - `Error(message string)`

  - `Panic(message string)`

  - `Fatal(message string)`

<h3>Example:</h3>

```go
logger.GetLogger().Debug("This is a debug message")
logger.GetLogger().Info("This is an info message")
```

<h2>Cleaning Logs</h2>

To clean (truncate) all log files, you can use the `Clean()` method:

```go
logger.GetLogger().Clean()
```

<h2>Example</h2>
Here's a complete example of using the logger:

```go
package main

import "github.com/LegationPro/go-logger/logger"

func main() {
    // Initialize the logger
    logger.NewLogger("logs")

    // Log messages of different levels
    logger.GetLogger().Debug("Debug message")
    logger.GetLogger().Info("Info message")
    logger.GetLogger().Warn("Warning message")
    logger.GetLogger().Error("Error message")
    logger.GetLogger().Panic("Panic message")
    logger.GetLogger().Fatal("Fatal message")

    // Clean up log files
    logger.GetLogger().Clean()
}
```
<h2>Implementation Details</h2>

<h3>Log Levels</h3>

The package defines the following log levels:

- Debug
- Info
- Warn
- Error
- Panic
- Fatal

Each level corresponds to a specific color output in the terminal and a dedicated log file.

<h2>Log Files</h2>

Log files are created in the directory specified when initializing the logger. Each file is named according to its log level (e.g., `debug.log`, `info.log`).

Color Output
Messages are color-coded based on their log level using the `fatih/color` package:

- Debug: Cyan
- Info: Green
- Warn: Yellow
- Error: Red
- Panic: Bright Red
- Fatal: Magenta

<h2>File Permissions</h2>

The log directory and files are created with permissions **0755**.

<h2>License</h2>

MIT License

Copyright (c) 2024 Anže Peternel

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

<h2>Contributions</h2>
Contributions are welcome! Please submit pull requests or open issues for any suggestions or improvements.

Feel free to modify or expand upon this README to fit your needs!