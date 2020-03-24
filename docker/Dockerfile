FROM node:current-slim
WORKDIR /usr/src/app
COPY node-bulletin-board/bulletin-board-app/ .
RUN npm install
EXPOSE 8080
CMD ["npm", "start"]

