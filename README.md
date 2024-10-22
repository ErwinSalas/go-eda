# Project Idea: Real-Time Food Ordering Platform

## Project Overview
You can develop a food ordering platform where events are triggered by different phases of the order (creation, payment, preparation, shipping, delivery). This architecture allows for the modeling of several microservices, independent databases for each service, and an event-driven system for communication between them. Additionally, you can utilize queues and topics for asynchronous messaging and emulate AWS services with LocalStack.

## Project Description
The platform allows users to place orders, restaurants to prepare those orders, and delivery personnel to take the food to its destination. The entire flow will be event-based, enabling the various components of the system to communicate asynchronously.

## Services in the Project

### Order Service
- **Responsibility**: Manages the creation of orders.
- **Functionality**: When a user creates an order, an event is generated and consumed by other services. Each order has an ID, customer details, and a list of requested dishes.

### Payment Service
- **Responsibility**: Processes payments for orders.
- **Functionality**: When a new order is created, an event is emitted to this service to process the payment. Once the payment is completed, the order status changes, and a "payment completed" event is sent.

### Food Preparation Service
- **Responsibility**: Manages food preparation.
- **Functionality**: Receives a "payment completed" event and handles the food preparation. Once the order is ready, it emits a "order ready" event.

### Delivery Service
- **Responsibility**: Handles the delivery of orders.
- **Functionality**: Listens for "order ready" events and assigns a delivery person to take the food to the customer. Once the order is delivered, it emits an "order delivered" event.

## Tools to Use
- **Go**: Each microservice can be a Go application that communicates with other services via events.
- **MySQL**: Each service has its own independent MySQL database (as in your initial Docker Compose setup).
- **SQS**: For messaging, each event is a task that goes through a queue (e.g., "order created," "payment completed").
- **SNS**: For notifying multiple services about events. This is used to alert various microservices when something relevant occurs (e.g., notifying both the payment and food services when a new order is created).
- **LocalStack**: To emulate SQS and SNS, avoiding the need to use actual AWS services.

## Event Flow

### Order Created:
1. The Order Service receives an order from a user and saves it in its database.
2. The Order Service emits an event to SQS to notify that an order has been created.

### Payment Completed:
1. The Payment Service listens for the "order created" event via SNS and processes the payment.
2. If the payment is successful, the Payment Service sends a "payment completed" event to SQS.

### Order in Preparation:
1. The Food Preparation Service listens for the "payment completed" event.
2. When the food is ready, the service sends an "order ready" event.

### Delivery in Progress:
1. The Delivery Service listens for the "order ready" event, assigns a delivery person, and sends the order to the customer.
2. Once delivered, it sends an "order delivered" event.
