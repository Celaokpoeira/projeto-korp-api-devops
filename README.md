# 🚀 Projeto Korp - API & Infraestrutura DevOps

Uma arquitetura robusta de API, conteinerizada e focada em alta disponibilidade, observabilidade ativa e automação de deploy (CI/CD). 

Este projeto não é apenas sobre escrever código, mas sobre como entregá-lo e mantê-lo em um ambiente de produção real.

## 🏗️ Arquitetura e Tecnologias

A infraestrutura foi desenhada para garantir resiliência e monitoramento proativo:

*   **Docker & Docker Compose:** Todo o ecossistema é conteinerizado, garantindo isolamento e facilidade de deploy.
*   **NGINX:** Atuando como proxy reverso, gerenciando o tráfego de forma eficiente e segura.
*   **Prometheus:** Responsável pela coleta de métricas de saúde do sistema e do servidor.
*   **Grafana:** Dashboards visuais para análise de dados em tempo real.
*   **Integração com Telegram:** Sistema de alertas autônomo (SRE proativo). O Grafana monitora o status da aplicação e dispara notificações instantâneas no celular em caso de instabilidade (`[FIRING]`) ou recuperação (`[RESOLVED]`).
*   **GitHub Actions (CI/CD):** Pipeline de entrega contínua. Qualquer atualização na *branch main* aciona um robô que conecta na VPS via SSH, puxa as atualizações e reconstrói os containers automaticamente, sem *downtime* perceptível ou intervenção manual.

## ⚙️ Como Executar o Projeto

Como o ambiente é totalmente baseado em Docker, rodar o projeto localmente ou em um novo servidor requer apenas um comando.

1. Clone o repositório:
   ```bash
   git clone [https://github.com/SEU_USUARIO/projeto-korp-api-devops.git](https://github.com/SEU_USUARIO/projeto-korp-api-devops.git)
