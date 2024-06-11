## Binance Price Fetcher

This project is a Go application that fetches cryptocurrency prices from the Binance API. It supports fetching prices for multiple symbols concurrently using worker goroutines and monitors the number of requests being made to the Binance API.

## Features
• Fetches prices for multiple cryptocurrency symbols from the Binance API.
• Uses worker goroutines to handle concurrent requests.
• Monitors and logs the number of requests made.
• Graceful shutdown handling using context and signals.

## Prerequisites
• Internet connection to access the Binance API.

## Installation
1.Clone the repository:

git clone https://github.com/usmanov0/ticker-parser-task.git
Ensure you have Go installed and set up correctly.

2.Run the application:

go run cmd/main.go

## Usage

1. Start the application:
   Run the main application using the command mentioned in the installation step. The application will start fetching prices for the symbols specified in config.yaml.

2. Graceful Shutdown:
The application can be stopped gracefully using a signal (e.g., Ctrl+C). It will finish the ongoing requests and then stop the workers.

3. Monitor Requests:
The application logs the total number of requests made to the Binance API. This can help in monitoring the request rate and adjusting the number of workers accordingly.

## Code Overview

cmd/main.go

The entry point of the application. It initializes the configuration, creates services, starts workers, and handles graceful shutdown.

internal/adapter/adapter.go

Contains the BinanceRepository which is responsible for fetching prices from the Binance API.

internal/config/config.go

Handles loading and parsing the configuration file (config.yaml).

internal/input/input.go

Waits for user input (like Ctrl+C) to stop the application gracefully.

internal/monitor/monitor.go

Monitors and logs the number of requests made to the Binance API.

internal/service/service.go

Defines the service that interacts with the BinanceRepository to fetch prices.

internal/worker/worker.go

Contains the logic for running worker goroutines that fetch prices concurrently.