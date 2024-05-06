# Execution : sudo docker image build -t forum /home/fabien/Zone01/Forum

FROM golang

LABEL Authors = "Alexandre Valin, Antoine Salaun, Fabien Olivier, Armand Auvray, Nicolas Legay"
LABEL Description : "This is for educational purpose only, a simple forum which handle database interactions"

WORKDIR /app
ADD . ./

# Best good practices to do for docker
RUN ./launch

CMD ["./forum"]
EXPOSE 8080

