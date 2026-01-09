# Conversor de Moedas

Um simples utilitário de linha de comando (CLI) para converter valores entre moedas (BRL, USD, EUR) utilizando a API da [FreeCurrencyAPI](https://freecurrencyapi.com/).

Este projeto foi construído com foco no aprendizado e na aplicação de bons princípios de design de software em Go

## Instalação

Para compilar o projeto, você precisa ter o [Go](https://go.dev/) instalado.

1.  **Clone o repositório:**
    ```bash
    git clone https://github.com/emanuelfelicio/conversor-de-moedas.git
    cd conversor-de-moedas
    ```

2.  **Construa o executável:**
    ```bash
    go build -o conversor-de-moedas .
    ```

## Configuração

A aplicação requer uma chave de API da FreeCurrencyAPI para funcionar.

1.  **Obtenha uma chave de API:** Crie uma conta gratuita em [https://app.freecurrencyapi.com/register](https://app.freecurrencyapi.com/register).

2.  **Crie um arquivo `.env`:** Na raiz do projeto, crie um arquivo chamado `.env`.

3.  **Adicione sua chave:** Adicione a seguinte linha ao arquivo `.env`, substituindo `SUA_CHAVE_AQUI` pela chave que você obteve.
    ```
    EXCHANGE_API_KEY="SUA_CHAVE_AQUI"
    ```

## Uso

Após a [Instalação](#instalação) e [Configuração](#configuração), você pode executar a aplicação de duas formas:

*   **Executando o binário compilado:**
    ```bash
    ./conversor-de-moedas
    ```

*   **Executando diretamente com o Go (para desenvolvimento):**
    ```bash
    go run main.go
    ```

A aplicação irá apresentar um formulário interativo para que você insira a moeda de origem, a moeda de destino e o valor a ser convertido.

<img width="677" height="321" alt="image" src="https://github.com/user-attachments/assets/011dc6b0-1406-4e9c-b63a-39d9bf22ee84" />


