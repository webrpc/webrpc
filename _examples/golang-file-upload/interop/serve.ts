// Standalone entrypoint: hosts the TypeScript FileService on PORT (4243 by
// default). The Go client in ../main.go can be pointed at it with:
//
//   go run .. client http://localhost:4243
import { createFileServiceHttpServer } from './server'

const port = Number(process.env['PORT'] || 4243)
createFileServiceHttpServer().listen(port, () => {
  console.log(`FileService (TypeScript) serving on http://localhost:${port}`)
})
