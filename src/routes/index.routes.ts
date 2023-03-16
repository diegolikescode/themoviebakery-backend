import { Router } from 'express'
import ratingRouter from './rating.routes'
import userRouter from './user.routes'

const router = Router()

router.use('/user', userRouter)
router.use('/rating', ratingRouter)

export default router
