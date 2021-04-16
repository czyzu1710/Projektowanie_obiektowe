FROM kprzystalski/projobj2:latest

USER root

RUN apt install libz-dev

RUN mkdir /home/kprzystalski/data/

VOLUME /home/kprzystalski/data/

EXPOSE 8000 8080 5000 38351

USER kprzystalski
