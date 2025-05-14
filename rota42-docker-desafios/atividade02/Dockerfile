# Usando uma imagem base Python
FROM python:3.11-slim

# Definir o diretório de trabalho no container
WORKDIR /app

# Copiar os arquivos do diretório atual (atividade02) para o diretório de trabalho /app no container
COPY . .

# Instalar dependências
RUN pip install -r requirements.txt

# Definir o comando para rodar a aplicação
CMD ["python", "app.py"]
