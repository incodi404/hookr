FROM node:20-alpine as base

WORKDIR /app

# Crucial: Ensure node execution finds hoisted monorepo binaries
# ENV PATH=/app/node_modules/.bin:$PATH

# Service manifest
COPY package.json package-lock.json ./

# Copy ALL workspace package.json files so npm can complete the graph
COPY packages ./packages
COPY apps/websocket/package.json ./apps/websocket/
# COPY apps/api/package.json ./apps/api/
# COPY apps/workers/delivery-worker/package.json ./apps/workers/delivery-worker/
# COPY apps/workers/email-worker/package.json ./apps/workers/email-worker/
# COPY apps/workers/schedule-worker/package.json ./apps/workers/schedule-worker/

# Install dependencies
RUN npm ci

# Service source
COPY apps/websocket ./apps/websocket/

# ============
# Development
# ============
FROM base As dev
CMD ["npm", "run", "dev", "--workspace=@hookr/websocket"]