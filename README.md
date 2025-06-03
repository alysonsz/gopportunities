📑 Gopportunities API
📌 Descrição do Projeto

O Gopportunities API é uma aplicação desenvolvida em Golang, com foco em fornecer um sistema robusto de notificações de vagas de emprego em tempo real. Utilizando boas práticas de arquitetura MVC, JWT para autenticação, e SSE (Server-Sent Events) para comunicação em tempo real, a API resolve o problema comum de usuários perderem oportunidades por atualizações tardias.

🚀 Objetivos do Projeto

  • Entregar atualizações instantâneas sobre novas oportunidades de emprego usando SSE.

  • Oferecer controle total ao usuário por meio de preferências personalizadas (localização, tipo de vaga, palavras-chave).

  • Seguir boas práticas de arquitetura em Golang (repository, service, controller).

  • Garantir segurança e escalabilidade com JWT e segregação de responsabilidades.

  • Demonstrar domínio em integração de tecnologias como GORM, Gin e Swagger.

🛠️ Tecnologias Utilizadas

Backend:

  • Golang

  • Gin (framework web)

  • GORM (ORM)

  • SQLite (para prototipagem)

  • JWT (autenticação)

  • Swagger (documentação)

  • SSE (Server-Sent Events)

Padrões e Conceitos:

  • Arquitetura MVC

  • Repository-Service-Controller

  • Validação de entrada (binding do Gin)

  • Comunicação real-time com SSE

📚 Funcionalidades Principais

  • CRUD de oportunidades de emprego: Criação, leitura, atualização e exclusão.

  • Autenticação segura com JWT: Login e registro de usuários.

  • Preferências de notificação: Configuração de alertas personalizados.

  • Notificações em tempo real (SSE): Receba alertas assim que novas oportunidades forem publicadas.

  • Documentação completa com Swagger.

🔗 Estrutura das Entidades

  • User: Usuários autenticados, com JWT.

  • Opportunity: Representa vagas de emprego.

  • UserPreference: Preferências configuráveis para notificações (localização, tipo, palavras-chave).

🧩 Organização do Projeto

gopportunities/
├── authentication/
│   └── authJwt.go
├── models/
│   ├── user.go
│   ├── opportunity.go
│   ├── userPreference.go
│   ├── loginRequest.go
│   ├── loginResponse.go
│   └── registerRequest.go
├── controllers/
│   ├── opportunityController.go
│   ├── authController.go
│   ├── notificationController.go
│   └── userPreferenceController.go
├── repositories/
│   ├── opportunityRepository.go
│   ├── userRepository.go
│   └── userPreferenceRepository.go
├── services/
│   ├── opportunityService.go
│   ├── authService.go
│   ├── notificationService.go
│   └── userPreferenceService.go
├── router/
│   ├── routes.go
│   └── router.go
├── utils/
│   └── jwt.go
├── docs/
│   └── swagger docs (gerados com swag init)
├── main.go
└── README.md

🔐 Autenticação JWT

  • /api/v1/register: Registra um novo usuário.

  • /api/v1/login: Autentica e retorna um JWT.

  • Use o token JWT como Bearer Token no Swagger para acessar rotas protegidas.

📡 Notificações em Tempo Real (SSE)

  • /api/v1/notifications [GET]

      • Conecta ao servidor via SSE.

      • Recebe alertas automáticos sempre que uma nova oportunidade for criada.

      • Exemplo de consumo:

  curl -H "Accept: text/event-stream" http://localhost:3030/api/v1/notifications

      const evtSource = new EventSource("http://localhost:3030/api/v1/notifications");
      evtSource.onmessage = (e) => console.log("Nova oportunidade:", e.data);

⚙️ Preferências do Usuário

  • /api/v1/preferences [POST]: Define preferências (localização, tipo, palavras-chave).

  • /api/v1/preferences [GET]: Recupera as preferências salvas.

📑 Documentação Swagger

  • Acesse http://localhost:3030/swagger/index.html

  • Explore e teste todas as rotas (incluindo JWT e SSE).

📈 Futuras Implementações

  • Migração do SQLite para um banco escalável (PostgreSQL, MySQL).

  • Notificações personalizadas com base nas preferências do usuário.

  • Sistema de papéis (admin, publisher) com RBAC.

  • Deploy com Docker e Kubernetes.

📌 Como Rodar o Projeto

  • Clone o repositório.

  • Rode go mod tidy para baixar dependências.

  • Execute swag init para gerar documentação Swagger.

  • Rode go run main.go ou go build -o gopportunities.git.exe e execute.

  • Acesse a API via http://localhost:3030 e o Swagger via /swagger/index.html.

👨‍💻 Autor

  • Alyson Souza Carregosa 👨‍💻 Back-end Developer

📝 Licença

Este projeto está disponível sob a licença MIT.