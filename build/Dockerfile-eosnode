FROM eosio/eos
USER root

RUN apt-get update
RUN apt-get install git curl -y

#RUN curl "https://s3.wasabisys.com/eosnode.tools/blocks/blocks_2019-02-12-07-02.tar.gz" --create-dirs -o /opt/eosio/bin/blocks.tar.gz

RUN curl "https://raw.githubusercontent.com/CryptoLions/EOS-MainNet/master/genesis.json" -o /opt/eosio/bin/genesis.json --create-dirs
RUN curl "https://raw.githubusercontent.com/CryptoLions/EOS-MainNet/master/config.ini" -o /opt/eosio/bin/config.ini --create-dirs 

WORKDIR /opt/eosio/bin

#RUN tar xvzf blocks.tar.gz

RUN mv /opt/eosio/bin/nodeos /bin/nodeos

CMD ["/bin/bash", "-c", "nodeos --data-dir=/opt/eosio/bin --hard-replay"]
