FROM debian:bullseye-slim

# Instale o ffmpeg
RUN apt-get update && \
    apt-get install -y ffmpeg && \
    apt-get clean

# Copie o script convert.sh
COPY convert.sh /usr/local/bin/convert.sh

# Torne o script executável
RUN chmod +x /usr/local/bin/convert.sh

# Configure o diretório de trabalho
WORKDIR /videos

# Configure entrypoint e comando padrão
ENTRYPOINT ["convert.sh"]
CMD ["input.mp4", "mp3"]
