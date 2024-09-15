## Link to initial blog post 

https://alistair.cockburn.us/hexagonal-architecture/

<img width="691" alt="Screenshot 2024-09-15 at 12 41 06" src="https://github.com/user-attachments/assets/ece83695-76ba-4534-9834-932513474277">


## Definition

Hexagonal architecture, proposed by Alistair Cockburn in 2005, is an architectural pattern that aims to build loosely coupled application components that can be connected via ports and adapters.
In this pattern, the consumer opens the application at a port via an adapter, and the output is sent through a port to an adapter. 

## Application

An application is a technology-agnostic component that contains the business logic that orchestrates functionalites or use cases. A hexagon represents the application that receives
write and read queries from the ports and sends them to external actors, such as database and third-party services, via ports. A hexagon visually represents multiple port/adapter combinations for
an application and shows the difference between the left side (or driving side) and right side (or driven side).

## Actors

In hexagonal architecture, actors are categorized into primary and secondary actors based on their role in interacting with the system:

<b>Primary Actors:</b> These are the actors that initiate interactions with your application. They represent the entities that use your application's capabilities. Primary actors typically interact with the system via driving adapters (e.g., controllers, API endpoints).

Example: A user submitting a form through a web interface or an external system making an API call.

<b>Secondary Actors:</b> These are external systems or services that your application interacts with to perform certain operations. Secondary actors are accessed through driven adapters (e.g., database connections, external APIs, message queues).

Example: A database where the application stores data, a payment gateway the app uses to process payments, or RabbitMQ for sending/receiving messages.

## Ports

In Hexagonal Architecture (also known as Ports and Adapters), the term "ports" refers to the interfaces or abstractions that define the core business functionality of an application. Ports are essentially the points through which external systems (e.g., user interfaces, databases, APIs) interact with the core domain logic of the application.

There are two types of ports:

<p>Inbound Ports: These allow external systems to interact with the application (e.g., service interfaces, controller methods). They represent how the core application is "driven" by outside forces.</p>
<p>Outbound Ports: These allow the application to interact with external systems (e.g., repositories, third-party services). They represent how the core application "drives" external systems.</p>

## Adapters

Adapters deal primarily with transforming a request from an actor to an application, and vice versa. Data transformation helps the application undestand the requests that come from actors. For example, a specific driver adapter can transform a techonogly-specific request into a call to an application service. In the same way, a driven adapter can covert a technology-agnostic request
from the application into a technology-specific request on the driven port.

## Dependencies

In hexagonal architecture, outer layers depend on inner layers, which makes it easier to implement the application core first and then implement outer layers to depend on them.

## Lower level vs higher level

Low level considered as inner levels which is mainly domain layer. And high level is outer level adapters like database or message queue implementation details. (3rd party)
