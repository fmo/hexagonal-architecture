# Definition

Hexagonal architecture, proposed by Alistair Cockburn in 2005, is an architectural pattern that aims to build loosely coupled application components that can be connected via ports and adapters.
In this pattern, the consumer opens the application at a port via an adapter, and the output is sent through a port to an adapter. 

# Primary (Driving) Actors

* GUI
* Console

# Secondary (Driven) Actors

* File System
* DB
* Message Queue

## Link to initial blog post 

https://alistair.cockburn.us/hexagonal-architecture/
