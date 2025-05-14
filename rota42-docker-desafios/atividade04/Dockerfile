# Etapa 1: Build (compilação)
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copie todo o código
COPY . .

# Compile seu app
RUN go build -o videoconverter .

# Etapa 2: Imagem final leve
FROM alpine:latest

# Copie o binário compilado na etapa anterior
COPY --from=builder /app/videoconverter /app/videoconverter

# Configure o diretório de trabalho
WORKDIR /app

# Configure o comando padrão
ENTRYPOINT ["/app/videoconverter"]
# Opcional: se desejar um comando padrão
# CMD ["input.mp4", "mp3"]
