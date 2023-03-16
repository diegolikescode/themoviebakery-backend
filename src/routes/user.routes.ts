import { Router } from 'express'
import UserController from '../controllers/user'

const userRouter = Router()
const userController = new UserController()

userRouter.post('/create', userController.create)
userRouter.post('/create-google', userController.createGoogle)
userRouter.get('/get-by-email', userController.getByEmail)
userRouter.get('/get-by-id', userController.getById)

export default userRouter
