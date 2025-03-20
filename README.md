# 🚀 Go API Vagas

API RESTful para gerenciamento de vagas de emprego, desenvolvida em **Go** com o framework **Gin** e **GORM**.

---

## 📚 Tecnologias Utilizadas

- [Go](https://go.dev/)
- [Gin Framework](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Swagger](https://swagger.io/) para documentação da API

---

## 📂 Estrutura Docker Compose

Este projeto possui dois ambientes configurados com Docker Compose:

- `dev.yml` → **Ambiente de Desenvolvimento** com hot reload (Air).
- `prod.yml` → **Ambiente de Produção** com build otimizado e leve.

---

## ✅ Como Executar o Projeto

### Pré-requisitos

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)

---

### Ambiente de Desenvolvimento

Com suporte a hot reload usando [Air](https://github.com/cosmtrek/air).

```bash
docker compose -f dev.yml up

```

### Ambiente de Producao


```bash
docker compose -f prod.yml up
```

### Gerando swagger

```bash
swag init
```
Acessando swagger (http:://localhost:8080/swagger/index.html).