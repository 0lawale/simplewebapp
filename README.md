# 🚀 Go Web App — CI/CD Pipeline with GitHub Actions & AWS EC2

A simple **Golang web application** deployed automatically to an **AWS EC2 instance** using **GitHub Actions**.  
This project demonstrates a full **CI/CD pipeline** — including building, testing, and deploying an app in a DevOps workflow.

---

## 🧩 Project Overview

This repository showcases an end-to-end DevOps pipeline:
- ✅ **Continuous Integration (CI)** — Build and test a Go web application using GitHub Actions.
- 🚀 **Continuous Deployment (CD)** — Automatically deploys to an AWS EC2 instance via SSH.
- 🧱 **Infrastructure-as-Code Ready** — Designed to integrate with Terraform or Ansible for provisioning in future steps.

---

## 📁 Project Structure
```bash
.
└── webapp
    ├── docker-compose.yml
    ├── Dockerfile
    ├── github
    │   └── workflows
    │       └── ci-cd.yml
    ├── go.mod
    ├── main.go
    ├── main_test.go
    └── README.md

```

---

## ⚙️ Prerequisites

Before running or deploying this project, ensure you have the following installed:

- [Go](https://golang.org/dl/) (v1.21 or later)
- [Git](https://git-scm.com/)
- [AWS CLI](https://aws.amazon.com/cli/)
- A configured **AWS EC2 instance**
- A valid **SSH key pair** (either created via AWS or imported from `ssh-keygen`)

---

## 🧠 Understanding the Stack

| Component | Purpose |
|------------|----------|
| **Go (Golang)** | Backend language for building the web app |
| **GitHub Actions** | Automates build, test, and deployment |
| **AWS EC2** | Hosts the web app as a Linux VM |
| **SSH Keys** | Secure communication for deployment |
| **Systemd or Nginx (optional)** | Process and reverse-proxy management for production setup |

---

## 🚧 How to Run Locally

1. **Clone this repository**
   ```bash
   git clone https://github.com/<your-username>/go-webapp.git
   cd go-webapp

2. **install dependencies**
    ```bash
    go mod tidy
    ```
3. **run the app**
    ```bash
    go run main.go
    ```
4. **Access locally**
open your browser and visit:
👉 http://127.0.0.1:8080

## 🤖 CI/CD Pipeline (GitHub Actions)

The workflow file is located at .github/workflows/ci-cd.yml.
### pipeline statges

| stage | description |
|------------|----------|
| **Build** | compiles the go app and ensures dependencies are valid |
| **test** | Runs unit tests |
| **Deploy** | SSHs into an AWS EC2 instance and deploys the latest code version |

### Environment Variables in GitHub Secrets


| secret name | Purpose |
|------------|----------|
| **DOCKERHUB_USERNAME** | The docker hub username |
| **DOCKERHUB_TOKEN** | a Docker Hub access token  |
| **SERVER_HOST** | The ip address of your server |
| **SERVER_USER** | The user on your server |
| **SSH_PRIVATE_KEY** | private key used by the GitHub action to authenticate via SSH |

## 🌍 Deployment to AWS EC2

1. Generate or import your SSH key
2. Launch an EC2 instance using that key pair
3. Add your private key as a GitHub Secret (SSH_PRIVATE_KEY)
4. Push code to the main branch — GitHub Actions will:
- Build and test the Go app
- SSH into EC2
- Pull and restart the app automatically 🎉

## 🔒 Security Notes
- Never commit your .env or private key files.
- Use GitHub Secrets for all sensitive data.
- Rotate keys periodically for production environments

## 🧠 Future Improvements

- Add Terraform for AWS provisioning

- Add Docker and ECR integration

- Set up monitoring with Prometheus & Grafana

- Add Nginx reverse proxy and HTTPS with Let’s Encrypt

## 👨‍💻 Author

Ayoola Philip Olawale
- DevOps Engineer & Petroleum Engineer
- 🌍 Passionate about automation, infrastructure, and scalable web systems.

📫 Connect: [Github](https://github.com/0lawale)


