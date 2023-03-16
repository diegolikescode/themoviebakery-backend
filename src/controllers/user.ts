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

    if (!email || !displayName)
      return res
        .json({ message: 'you need to send email and displayName' })
        .status(400)

    const userExist = await userRepository.findOneBy({ email })
    if (userExist) {
      return res.status(409).json({ message: 'user already exists' })
    }
    const newUser = new User()
    newUser.userId = shortUUID.generate()
    newUser.email = email
    newUser.displayName = displayName

    await userRepository.insert(newUser)

    res.status(201).json(newUser)
  }

  getByEmail = async (req: Request, res: Response) => {
    const userRepository = appDataSource.getRepository(User) // passar por construtor depois
    const { email }: FindUser = req.query

    const user = await userRepository.findOneBy({ email })
    if (!user) {
      return res.status(404).json({ message: 'user not found' })
    }

    return res.status(200).json(user)
  }

  getById = async (req: Request, res: Response) => {
    if (!req.query.userId)
      return res.status(400).json({ message: 'you need to passa userId' })
    const userRepository = appDataSource.getRepository(User) // passar por construtor depois
    const { userId }: FindUser = req.query

    const user = await userRepository.findOneBy({ userId })
    if (!user) {
      return res.status(404).json({ message: 'user not found' })
    }

    return res.status(200).json(user)
  }
}
