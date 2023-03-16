import { Router } from 'express'
import RatingController from '../controllers/rating'

const ratingRouter = Router()
const ratingController = new RatingController()

ratingRouter.post('/create', ratingController.createAndUpdate)
ratingRouter.get('/get-all', ratingController.allRatingsAllUsers)
ratingRouter.get('/get-by-user', ratingController.findByUser)

export default ratingRouter
