### Overview

Bloxley Backend is being developed in microservice architecture. The power of microservice architecture is to deliver efficient and modular services, enabling seamless scalability and maintainability.

### Key Features

* **Microservices Architecture**: Independent and modular services for better scalability, maintainability, and fault isolation.
* **Cloud Infrastructure**: Deployed on Google Cloud Platform (GCP) to ensure high availability and optimal performance.
* **Container Orchestration**: Managed using Kubernetes for automated deployment, scaling, and management of containerized applications.
* **Scalability & Resilience**: Designed to handle high loads with automatic scaling and load balancing capabilities.
* **Continuous Integration/Continuous Deployment (CI/CD)**: Streamlined processes for rapid development and deployment.

### Architecture

The KISLL platform follows a microservices-based architecture where each service functions as an independent module. This architecture facilitates horizontal scaling and easy updates to individual components without affecting the entire system.

### Key Components

* **Microservices**: Individual units designed to handle specific business capabilities.
* **Kubernetes (K8s)**: Used for orchestrating containers and ensuring smooth operations across the platform.
* **Google Cloud Platform (GCP)**: Provides the infrastructure backbone, including Compute Engine, Cloud Storage, Cloud SQL, and other managed services.
* **Service Mesh**: Facilitates secure communication between services, traffic management, and observability.

### Deployment

Bloxley utilizes Kubernetes for deploying and managing the microservices across the GCP infrastructure. Helm charts are used to manage the configuration of Kubernetes applications, providing a seamless deployment experience.

### Prerequisites

* **Google Cloud Account**: Ensure you have a GCP project set up.
* **Kubernetes Cluster**: A GKE (Google Kubernetes Engine) cluster or self-managed Kubernetes setup.
* **Helm**: For managing Kubernetes manifests and deployments.
* **Docker**: To build and package microservices as container images.

### Deployment Steps

1. **Build Docker Images**:
```bash
docker build -t gcr.io/your-gcp-project/kisll-service:latest .
```
2. **Push the images to GCP Container Registry**:
```bash
docker push gcr.io/your-gcp-project/kisll-service:latest
```
3. **Configure Kubernetes Cluster**:
Ensure your kubectl context is set to the desired GKE cluster:
```bash
gcloud container clusters get-credentials your-cluster-name --zone your-zone --project your-gcp-project
```
4. **Deploy Services**:
Use Helm to deploy microservices:
```bash
helm install kisll-service ./helm/kisll-service
```

### Monitoring and Logging

KISLL incorporates monitoring and logging tools to ensure platform reliability:

* **Stackdriver (now Cloud Operations)**: For monitoring metrics and logs.
* **Prometheus & Grafana**: For advanced monitoring and dashboard visualization.
