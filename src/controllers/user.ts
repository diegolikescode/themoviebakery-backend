import { Request, Response } from 'express'
import shortUUID from 'short-uuid'
import User from '../entities/user'
import { appDataSource } from '../server'

type FindUser = {
  email?: string
  userId?: string
}

export default class UserController {
  create = async (req: Request, res: Response) => {
    const userRepository = appDataSource.getRepository(User) // passar por construtor depois
    const { email, displayName } = req.body

    const userExist = await userRepository.findOneBy({ email })
    if (userExist) {
      return res.json({ message: 'user already exists' }).status(409)
    }
    const newUser = new User()
    newUser.userId = shortUUID.generate()
    newUser.email = email
    newUser.displayName = displayName

    await userRepository.insert(newUser)

    res.json(newUser).status(201)
  }

  // need it????? lets see frontend's response
  createGoogle = (req: Request, res: Response) => {
    res.json({ yet: 'to implement' }).status(200)
  }

  getByEmail = async (req: Request, res: Response) => {
    const userRepository = appDataSource.getRepository(User) // passar por construtor depois
    const { email }: FindUser = req.query

    const user = await userRepository.findOneBy({ email })
    if (!user) {
      return res.json({ message: 'user not found' }).status(404)
    }

    return res.json(user).status(200)
  }

  getById = async (req: Request, res: Response) => {
    if (!req.query.userId)
      return res.json({ message: 'you need to passa userId' }).status(400)
    const userRepository = appDataSource.getRepository(User) // passar por construtor depois
    const { userId }: FindUser = req.query

    const user = await userRepository.findOneBy({ userId })
    if (!user) {
      return res.json({ message: 'user not found' }).status(404)
    }

    return res.json(user).status(200)
  }
}
