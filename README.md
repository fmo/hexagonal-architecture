# Definition

Hexagonal architecture, also known as the Ports and Adapters architecture, was popularized by Alistair Cockburn, a prominent software engineer, and author. 
Alistair Cockburn introduced the concept in the early 2000s to address issues related to maintaining and scaling software systems. 
His ideas on this architectural style were aimed at making software more modular, easier to test, and adaptable to changes.

The term "hexagonal architecture" came from the visual representation of the architecture, where the central business logic is surrounded by "ports" and "adapters," 
forming a hexagonal shape. This design pattern emphasizes the separation of concerns, ensuring that the core logic of the application is independent of external factors such as user interfaces, 
databases, and other external systems.

# Primary (Driving) Actors

* GUI
* Console

# Secondary (Driven) Actors

* File System
* DB
* Message Queue

## Link to initial blog post 

https://alistair.cockburn.us/hexagonal-architecture/
