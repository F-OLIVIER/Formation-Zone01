# ----------------------------
# Utilisation container Docker
# 
# VERSION               0.0.1
# ----------------------------

# Definition de l'image 
# syntax=docker/dockerfile:1
FROM golang:1.20

# Information sur l'application
LABEL description="ascii art web dockerize"
LABEL authors="OLIVIER Fabien"
LABEL release-date="20/06/2023"

# Dossier de travail Docker
RUN mkdir /app
# Copie du code source 
# (add pour fichier et dossier)
# (copy que pour les fichiers)
ADD . /app
WORKDIR /app

# Si besoin de dépendance
RUN go mod download

# Compilation de l'appplication
# action option nom_executable emplacement_fichier_main
RUN go build -o /ascii-art-web cmd/web/ascii-art-web.go

# Run l'image
CMD ["/ascii-art-web"]
