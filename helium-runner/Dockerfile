FROM node:8-alpine

WORKDIR /home/src
COPY . /home/src

RUN yarn install && yarn build

CMD yarn run initialize && yarn run start
