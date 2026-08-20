FROM node:20-alpine as base

WORKDIR /app

# Service manifest
COPY package.json package-lock.json ./

# Root workspace configuration + lockfile
COPY apps/workers/email-worker/package.json ./apps/workers/email-worker/

# All shared workspace packages
COPY packages ./packages

# Install dependencies
RUN npm ci

# Service source
COPY apps/workers/email-worker ./apps/workers/email-worker/

# ============
# Development
# ============
FROM base As dev
CMD ["npm", "run", "dev", "--workspace=@hookr/email-worker"]