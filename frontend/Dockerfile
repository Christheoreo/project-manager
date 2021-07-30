FROM node:14

WORKDIR /usr/src/app/frontend

COPY package*.json ./
COPY pnpm-lock.yaml ./

RUN npm install pnpm -g

RUN pnpm install

EXPOSE 5000

CMD ["pnpm", "run", "dev"]