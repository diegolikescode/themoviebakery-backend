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
  host: 'ls-b9321f739756e12f0d5637fc356b634e638b7821.cqybf2hflrg3.us-east-1.rds.amazonaws.com',
  port: 5432,
  username: 'dbmasteruser',
  password: ']=5WYWBh?nhA4L=0av*Ms<))6=L6Fkzr',
  database: 'my-postgres',
  entities: [user, rating],
  synchronize: true,
})

appDataSource
  .initialize()
  .then(() => {
    app.use(express.json())
    app.use(cors())
    app.use('/api/v1', router)

    app.listen(8080, () => console.log('running in 8080'))
  })
  .catch(() => console.log('was not able to connect to postgreSQL'))
