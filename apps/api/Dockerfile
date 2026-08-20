FROM node:20-alpine as base

WORKDIR /app

# Service manifest
COPY package.json package-lock.json ./

# Root workspace configuration + lockfile
COPY apps/api/package.json ./apps/api/

# All shared workspace packages
COPY packages ./packages

# Install dependencies
RUN npm ci

# Service source
COPY apps/api ./apps/api/

# ============
# Development
# ============
FROM base As dev
CMD ["npm", "run", "dev", "--workspace=@hookr/api"]