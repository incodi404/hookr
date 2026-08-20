FROM node:20-alpine as base

WORKDIR /app

# Service manifest
COPY package.json package-lock.json ./

# Root workspace configuration + lockfile
COPY apps/workers/delivery-worker/package.json ./apps/workers/delivery-worker/

# All shared workspace packages
COPY packages ./packages

# Install dependencies
RUN npm ci

# Service source
COPY apps/workers/delivery-worker ./apps/workers/delivery-worker/

# ============
# Development
# ============
FROM base As dev
CMD ["npm", "run", "dev", "--workspace=@hookr/delivery-worker"]