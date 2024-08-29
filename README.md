## Link to initial blog post 

https://alistair.cockburn.us/hexagonal-architecture/

<img width="709" alt="Screenshot 2024-08-28 at 10 41 09" src="https://github.com/user-attachments/assets/2fa6ff50-f076-42cb-8ba1-86e2fa3f823f">

# Definition

Hexagonal architecture, proposed by Alistair Cockburn in 2005, is an architectural pattern that aims to build loosely coupled application components that can be connected via ports and adapters.
In this pattern, the consumer opens the application at a port via an adapter, and the output is sent through a port to an adapter. 

## Application

An application is a technology-agnostic component that contains the business logic that orchestrates functionalites or use cases. A hexagon represents the application that receives
write and read queries from the ports and sends them to external actors, such as database and third-party services, via ports. A hexagon visually represents multiple port/adapter combinations for
an application and shows the difference between the left side (or driving side) and right side (or driven side).

# Primary (Driving) Actors

* GUI
* Console

# Secondary (Driven) Actors

* File System
* DB
* Message Queue
