# GoMonitor - Lightweight Monitoring System

GoMonitor is a lightweight monitoring system built in Go that collects metrics from services and exposes them via a simple REST API. It allows you to scrape and query metrics in real-time, inspired by Prometheus but without the complexity of a full-fledged time series database.

## Features
- **Concurrent Metric Collection**: Scrapes metrics from various services concurrently.
- **Real-time Metrics**: Provides up-to-date metrics through a simple REST API.
- **Custom Metrics**: You can instrument your own services to expose metrics in the desired format.
- **API for Querying**: Retrieve metrics through a query API.

## Future Improvements
Persistence: Implement support for a time series database to store and query metrics over time.
Alerting: Add alerting functionality for when certain metric thresholds are crossed.
Dashboard: Create a web-based dashboard to visualize collected metrics.
Service Discovery: Implement automatic service discovery to scrape metrics from services dynamically.

### Why 
I built this monitoring system as a personal project to learn Go. It was a fun way to explore Go’s concurrency model with goroutines and channels, as well as its ability to create lightweight services. Through this project, I gained experience in writing concurrent systems, handling HTTP servers, and designing simple APIs.
