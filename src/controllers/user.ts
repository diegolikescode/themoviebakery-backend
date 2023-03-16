import { Request, Response } from 'express'
import shortUUID from 'short-uuid'
import User from '../entities/user'
import { appDataSource } from '../server'

export default class UserController {
  create = async (req: Request, res: Response) => {
    const userRepository = appDataSource.getRepository(User) // passar por construtor depois
    const { email, displayName } = req.body

    // verificar email igual
    const userExist = await userRepository.findOneBy({ email })
    if (userExist) {
      return res.json({ message: 'user already exists' }).status(409)
    }
    const newUser = new User()
    newUser.userId = shortUUID.generate()
    newUser.email = email
    newUser.displayName = displayName

    const a = await userRepository.insert(newUser)

    const allUsers = await userRepository.find()

    res.json(allUsers).status(201)
  }

  // need it????? lets see frontend's response
  createGoogle = (req: Request, res: Response) => {
    res.json({ yet: 'to implement' }).status(200)
  }

  getByEmail = (req: Request, res: Response) => {
    res.json({ yet: 'to implement' }).status(200)
  }

  getById = (req: Request, res: Response) => {
    res.json({ yet: 'to implement' }).status(200)
  }
}
