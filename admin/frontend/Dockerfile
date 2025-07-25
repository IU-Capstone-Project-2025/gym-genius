FROM node:lts-alpine AS builder
WORKDIR /app

COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile

COPY . .
RUN yarn build

FROM node:lts-alpine AS runner
WORKDIR /app

COPY --from=builder /app/.output /app/.output
COPY --from=builder /app/package.json /app/

EXPOSE 3000
CMD ["node", ".output/server/index.mjs"]
