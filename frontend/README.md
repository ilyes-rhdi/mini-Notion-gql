mini notion front (public)

what was fixed
- removed tailwind so the UI does not rely on px/rem utility sizes
- removed duplicated folder src - Copie
- fixed broken imports (Login/Verify file casing)
- aligned routes under /auth/* and /app/*
- auth forms now use fluid CSS (vh/vw/%) and consistent routes
- added minimal workspace/page/editor UI so the app runs end-to-end
- removed dead subscription code (backend has no websocket)

how to run
1) create .env (or .env.local)
   VITE_API_BASE=http://localhost:3000
   VITE_GQL_URL=http://localhost:3000/graphql

2) install and start
   npm install
   npm run dev

notes
- the editor uses simple blocks (paragraph) and stores block.data as JSON {text: "..."}
- if your backend schema uses different field names, update the graphql documents in src/graphql/*
