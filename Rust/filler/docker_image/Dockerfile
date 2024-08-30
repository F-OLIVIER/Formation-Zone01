FROM rust:1.63-buster

COPY ./maps			        /filler/maps
COPY ./linux_robots		    /filler/linux_robots
COPY ./m1_robots		    /filler/m1_robots
COPY ./linux_game_engine	/filler/linux_game_engine
COPY ./m1_game_engine	    /filler/m1_game_engine
COPY ./audit.sh	            /filler/audit.sh

WORKDIR /filler/

ENTRYPOINT /bin/bash
