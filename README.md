# Project Saffron Orchestrator
###  Notics I just generate README.md from ChatGPT lol, but don't worry things it mentions here is correct. and i checked
Welcome to the Project Saffron Orchestrator repository. This project is designed to provide a lightweight orchestrator framework that simplifies workload management and distribution across a computing cluster. While it draws inspiration from robust orchestrators like Kubernetes, Saffron Orchestrator is specifically tailored for use cases that require simpler orchestration mechanisms without the overhead of managing distributed computing complexities.

## Overview
Saffron Orchestrator is an orchestrator framework that focuses on scheduling and managing tasks across a cluster of machines. It is intentionally designed without features like distributed computing, service discovery, high availability, load balancing, and advanced security mechanisms. This design choice makes Saffron Orchestrator ideal for educational purposes, development environments, and small-scale production setups where the aforementioned features are not critical.

## Features
Saffron Orchestrator comprises several key components that work together to manage the lifecycle of tasks within the system:

- **The Task**: Represents a unit of work that needs to be executed. A task could be anything from a simple script to a complex application.
- **The Job**: A collection of tasks that together accomplish a specific objective. Jobs define how tasks relate to each other and their execution order.
- **The Scheduler**: Responsible for scheduling jobs and tasks on different workers based on the current workload and the available resources.
- **The Manager**: Oversees the operation of the cluster, including task distribution, health monitoring, and resource allocation.
- **The Worker**: A node within the cluster that executes the tasks assigned by the scheduler. Workers report their status and resource usage back to the manager.
- **The Cluster**: The collection of manager and worker nodes that form the operational backbone of Orchestrator-Lite.
- **The Command-Line Interface (CLI)**: Provides users with the tools to interact with Orchestrator-Lite, including deploying jobs, monitoring their status, and managing cluster resources.

## Getting Started
### Prerequisites
Before you begin, ensure you have the following installed on your system:

- Git
- Docker
- Go

### Installation
1. Clone the repository:


```bash
git clone https://github.com/yourusername/orchestrator-lite.git
```

License
This project is licensed under the Apache License 2.0 - see the LICENSE file for details.
