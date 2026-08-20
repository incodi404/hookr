FROM node:20-alpine as base

WORKDIR /app

# Service manifest
COPY package.json package-lock.json ./

# Root workspace configuration + lockfile
COPY apps/workers/schedule-worker/package.json ./apps/workers/schedule-worker/

# All shared workspace packages
COPY packages ./packages

# Install dependencies
RUN npm ci

# Service source
COPY apps/workers/schedule-worker ./apps/workers/schedule-worker/

# ============
# Development
# ============
FROM base As dev
CMD ["npm", "run", "dev", "--workspace=@hookr/schedule-worker"]