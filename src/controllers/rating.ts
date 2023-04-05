import { Request, Response } from 'express'
import Rating from '../entities/rating'
import { appDataSource } from '../server'

type CreateRating = {
  userId: number
  movieId: string
  ratingValue: number
}

type FindRating = {
  userId?: number
  movieId?: string
}

export default class RatingController {
  createAndUpdate = async (req: Request, res: Response) => {
    const ratingRepository = appDataSource.getRepository(Rating) // passar por construtor depois
    const { userId, movieId, ratingValue }: CreateRating = req.body
    const rating = await ratingRepository.findOneBy({ userId, movieId })
    if (rating) {
      await ratingRepository.update(
        { userId, movieId },
        { ...rating, ratingValue }
      )
      return res.status(200).json(rating)
    }

    const newRating = new Rating()
    newRating.movieId = movieId
    newRating.userId = userId
    newRating.ratingValue = ratingValue

    await ratingRepository.insert(newRating)

    res.status(201).json(newRating)
  }

  allRatingsAllUsers = async (_: Request, res: Response) => {
    const ratingRepository = appDataSource.getRepository(Rating) // passar por construtor depois
    const allRatings = await ratingRepository.find()

    res.status(200).json(allRatings)
  }

  findByUser = async (req: Request, res: Response) => {
    const { userId }: FindRating = req.query

    const ratingRepository = appDataSource.getRepository(Rating) // passar por construtor depois
    const allRatings = await ratingRepository.find({ where: { userId } })

    res.status(200).json(allRatings)
  }
}
