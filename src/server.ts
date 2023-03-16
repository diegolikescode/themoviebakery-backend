import 'reflect-metadata'

import express from 'express'
import cors from 'cors'
import router from './routes/index.routes'
import { DataSource } from 'typeorm'

import user from './entities/user'
import rating from './entities/rating'

const app = express()

export const appDataSource = new DataSource({
  type: 'postgres',
  host: 'localhost',
  port: 5432,
  username: 'admin',
  password: 'admin',
  database: 'TheMovieBakery',
  entities: [user, rating],
  synchronize: true,
})

appDataSource
  .initialize()
  .then(() => {
    app.use(express.json())
    app.use(cors())
    app.use('/api/v1', router)

    app.listen(8080, () => 'running')
  })
  .catch(() => console.log('error in the datasource'))
