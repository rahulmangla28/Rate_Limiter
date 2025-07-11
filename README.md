# 🚦 Rate_Limiter

[![Go Version](https://img.shields.io/badge/go-1.24.4-blue.svg)](https://golang.org/dl/)

A simple, efficient, and customizable **Rate Limiter** implementation in Go designed to control the rate of actions or requests to prevent abuse or overuse of resources. ⚡️🐹


---


## 📚 Table of Contents

- [About](https://github.com/rahulmangla28/Rate_Limiter?tab=readme-ov-file#-about)  
- [Features](https://github.com/rahulmangla28/Rate_Limiter?tab=readme-ov-file#-features)  
- [Installation](https://github.com/rahulmangla28/Rate_Limiter?tab=readme-ov-file#-installation)  
- [Usage](https://github.com/rahulmangla28/Rate_Limiter?tab=readme-ov-file#-usage)  
- [Configuration](https://github.com/rahulmangla28/Rate_Limiter?tab=readme-ov-file#-configuration)  
- [Examples](https://github.com/rahulmangla28/Rate_Limiter?tab=readme-ov-file#-examples)  
- [Contributing](https://github.com/rahulmangla28/Rate_Limiter?tab=readme-ov-file#-contributing)  


---


## ❓ About

Rate limiting is a method for controlling the frequency of requests made to a network, server, or other resource. It is crucial for preventing abuse, ensuring fair access, and maintaining overall system stability and performance.

This project delivers a streamlined and effective Go implementation of a rate limiter, built to restrict the number of service calls or requests within a specified time interval. It's well-suited for APIs, web services, and backend systems that need robust throttling to protect against traffic spikes and ensure consistent performance.

By regulating request flow, rate limiting helps defend against denial-of-service (DoS) and distributed denial-of-service (DDoS) attacks, prevents individual clients from monopolizing system resources, and ensures a smooth and responsive user experience. Additionally, it can help reduce infrastructure costs by avoiding unnecessary scaling due to uncontrolled traffic bursts.

Whether you're developing a high-traffic API or safeguarding a backend service, this rate limiter provides a dependable foundation for managing request throughput effectively.

---


## ✨ Features

- ✅ Easy to integrate with any Go project  
- 🔄 Supports different rate limiting strategies (e.g., token bucket, leaky bucket, fixed window, sliding window)  
- 🧵 Thread-safe and efficient  
- ⚙️ Configurable limits and intervals  
- 📦 Minimal dependencies  


---


## 🛠 Installation

You can clone the repository and start using it immediately:

```bash
git clone https://github.com/rahulmangla28/Rate_Limiter.git
cd Rate_Limiter
```


---


## Usage

This example demonstrates how to use the `RateLimiter` package to limit the number of allowed calls within a time interval.

```go
package main

import (
    "fmt"
    "github.com/rahulmangla28/Rate_Limiter"
    "time"
)

func main() {
    // Create a RateLimiter that allows 5 calls per 10 seconds
    limiter := ratelimiter.NewRateLimiter(5, 10*time.Second)

    for i := 0; i < 10; i++ {
        if limiter.Allow() {
            fmt.Printf("Call %d allowed \n", i+1)
        } else {
            fmt.Printf("Call %d blocked \n", i+1)
        }
    }
}
```


---


## ⚙️ Configuration

| Parameter | Description                      | Default |
|-----------|--------------------------------|---------|
| maxCalls  | Maximum number of allowed calls | 5       |
| period    | Time period (in seconds) to reset the counter | 10      |

You can customize these parameters based on your requirements when creating a new RateLimiter instance, for example:

```go
limiter := ratelimiter.NewRateLimiter(maxCalls, time.Duration(period)*time.Second)
```


---


## 💡 Examples

- Limit API request rates ⏳  
- Control event handling frequency 🎯  
- Prevent abuse in web applications 🚫  


---


## 🚧 Future Enhancements

- **Distributed Rate Limiting:** Support rate limiting across multiple instances using Redis or other distributed stores.
- **Customizable Callbacks:** Allow users to specify callbacks or hooks when a request is blocked.
- **Metrics & Monitoring:** Integrate metrics to expose rate limiter statistics (e.g., allowed vs blocked calls).
- **Rate Limiting by Key:** Support different limits based on user API keys, IP addresses, or other identifiers.
- **Dynamic Configuration:** Enable changing rate limits at runtime without restarting services.


---


## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the issues page if you want to contribute.

To contribute:

1. Fork the repository 🍴  
2. Create your feature branch (`git checkout -b feature/my-feature`) 🌿  
3. Commit your changes (`git commit -m 'Add some feature'`) 💬  
4. Push to the branch (`git push origin feature/my-feature`) 📤  
5. Open a pull request 🔀  


---
